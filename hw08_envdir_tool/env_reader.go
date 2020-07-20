package main

import (
	"bufio"
	"bytes"
	"errors"
	io2 "io"
	io "io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]string

var ErrIncorrectEnvName = errors.New("incorrect env name")

func ReadLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	if err != nil {
		return "", err
	}

	line = bytes.ReplaceAll(line, []byte("\x00"), []byte("\n"))
	return string(line), nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := io.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envMap := make(Environment)

	for _, file := range files {
		if strings.Contains(file.Name(), "=") {
			return nil, ErrIncorrectEnvName
		}

		env, err := ReadLine(path.Join(dir, file.Name()))
		if err != nil && err != io2.EOF {
			return nil, err
		}
		env = strings.TrimRight(env, " \t\n")
		envMap[file.Name()] = env
	}

	return envMap, nil
}
