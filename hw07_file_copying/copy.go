package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	file, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if size := fileInfo.Size(); size < limit || limit == 0 {
		limit = fileInfo.Size()
	}
	if limit+offset > fileInfo.Size() {
		limit = fileInfo.Size() - offset
	}
	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	_, err = file.Seek(offset, 0)
	if err != nil {
		return err
	}

	reader := io.LimitReader(file, limit)
	writer, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer writer.Close()

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)
	_, err = io.CopyN(writer, barReader, limit)
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}
