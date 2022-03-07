package main

import (
	"os"
	"strings"
	"testing"
)

func readTestFile() *os.File {
	f, err := os.Open("stdin.txt")
	if err != nil {
		return nil
	}
	return f
}

func TestInputProcessing(t *testing.T) {
	testInput := readTestFile()
	defer testInput.Close()
	go inputProcessing(testInput)
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
