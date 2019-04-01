package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestIfStringIsWrittenToFile(t *testing.T) {
	s := "Creating Test"
	f := File{FileName: "TestFile.txt"}
	f.Save(s)

	f1, err := os.Open("TestFile.txt")
	if err != nil {
		t.Error("Could not open file", err)
	}

	scanner := bufio.NewScanner(f1)

	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "Creating Test") {
			t.Errorf("File doesn't contain the expected string")
		}
	}
}
