package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestRunCmd(t *testing.T) {
	t.Run("Check set env", func(t *testing.T) {
		env := Environment{
			"BAR": "bar",
		}
		re := captureOutput(func() {
			RunCmd([]string{"env"}, env)
		})
		require.Contains(t, re, "BAR=bar")
	})

	t.Run("Check unset empty env", func(t *testing.T) {
		env := Environment{
			"UNSET": "",
		}
		re := captureOutput(func() {
			RunCmd([]string{"env"}, env)
		})
		require.NotContains(t, re, "UNSET")
	})

	t.Run("check deleting empty env", func(t *testing.T) {
		env := Environment{"UNSET": ""}
		os.Setenv("UNSET", "VALUE")
		re := captureOutput(func() {
			RunCmd([]string{"env"}, env)
		})
		require.NotContains(t, re, "UNSET")
	})
}


