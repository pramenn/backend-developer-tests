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
	reader := readTestFile()
	go inputProcessing(reader)
	for result := range resultChan {
		if !strings.Contains(result, "error") {
			t.Fail()
		}
	}
	for result := range errChan {
		if result != nil {
			t.Fail()
		}
	}
}
