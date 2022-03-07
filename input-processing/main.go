package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var resultChan = make(chan string)
var errChan = make(chan error)

func main() {
	go inputProcessing(os.Stdin)
	for result := range resultChan {
		fmt.Println(result)
	}

	for result := range errChan {
		fmt.Println(result)
	}
}

func inputProcessing(reader io.Reader) {
	r := bufio.NewReader(reader)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			errChan <- fmt.Errorf("%v resulted in partial data", err)
		}

		if strings.Contains(line, "error") {
			resultChan <- line
		}
	}
	defer close(resultChan)
	defer close(errChan)
}
