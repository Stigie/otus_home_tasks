package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	inputFilePath   = "./testdata/input.txt"
	outputFilePath  = "./output.txt"
	wrongOutputPath = "./output"
	wrongInputPath  = "./input"
)

func TestCopy(t *testing.T) {
	t.Run("offet more than file size", func(t *testing.T) {
		err := Copy(inputFilePath, outputFilePath, 7000, 0)
		defer os.RemoveAll(outputFilePath)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("limit more than file size", func(t *testing.T) {
		err := Copy(inputFilePath, outputFilePath, 0, 7000)
		defer os.RemoveAll(outputFilePath)
		require.Equal(t, nil, err)
	})

	t.Run("limit = 0", func(t *testing.T) {
		err := Copy(inputFilePath, outputFilePath, 0, 7000)
		defer os.RemoveAll(outputFilePath)
		require.Equal(t, nil, err)
	})
	t.Run("is dir", func(t *testing.T) {
		err := Copy(wrongInputPath, wrongOutputPath, 0, 7000)
		defer os.RemoveAll(outputFilePath)
		require.NotNil(t, err)
	})
	t.Run("test on unsupported files", func(t *testing.T) {
		err := Copy("/dev/random", outputFilePath, 0, 0)
		defer os.RemoveAll(outputFilePath)
		require.Equal(t, ErrUnsupportedFile, err)
	})
}
