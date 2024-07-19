package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestReadInput(t *testing.T) {
	input := "test input\n"
	inputChan := make(chan string)
	stopChan := make(chan struct{})
	isTesting = true

	go readInput(inputChan, stopChan, strings.NewReader(input))

	select {
	case result := <-inputChan:
		if result != "test input" {
			t.Errorf("TestReadInput error, got '%s'", result)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("TestReadInput timeout error")
	}

	close(stopChan)
}

func TestWriteToFile(t *testing.T) {
	inputChan := make(chan string)
	stopChan := make(chan struct{})
	var buf bytes.Buffer

	go writeToFile(inputChan, stopChan, &buf)

	inputChan <- "input test text"
	close(inputChan)

	time.Sleep(1 * time.Second)

	if buf.String() != "input test text\n" {
		t.Errorf("TestWriteToFile err, got '%s'", buf.String())
	}

	close(stopChan)
}
