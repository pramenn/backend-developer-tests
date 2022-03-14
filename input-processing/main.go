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
var endChan = make(chan struct{})

func main() {
	go inputProcessing(os.Stdin)
	for {
		select {
		case result := <-resultChan:
			if !strings.Contains(result, "error") {
				fmt.Println(result)
			}
		case err := <-errChan:
			fmt.Println(err)
		case <-endChan:
			return
		}
	}
}

func inputProcessing(reader io.Reader) {
	r := bufio.NewReader(reader)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				endChan <- struct{}{}
				break
			}
			errChan <- fmt.Errorf("%v resulted in partial data", err)
		}

		if strings.Contains(line, "error") {
			resultChan <- line
		}
	}
}
