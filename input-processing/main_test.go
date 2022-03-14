package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func readTestFile() io.Reader {
	f, err := os.Open("stdin.txt")
	if err != nil {
		return nil
	}
	return f
}

func TestInputProcessing(t *testing.T) {
	testInput := readTestFile()

	if file, ok := testInput.(*os.File); ok {
		defer file.Close()
	}

	go inputProcessing(testInput)

	for {
		select {
		case result := <-resultChan:
			if !strings.Contains(result, "error") {
				t.Fail()
			}
		case err := <-errChan:
			if err != nil {
				t.Fail()
			}
		case <-endChan:
			return
		}
	}
}
