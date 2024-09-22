package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: lox [script]")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		parseFile(os.Args[1])
	} else {
		parsePrompt()
	}
}

func parseFile(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Could not open file")
		os.Exit(1)
	}

	scanner := NewScanner()
	scanner.Parse(string(file))
}

func parsePrompt() {
	reader := bufio.NewReader(os.Stdin)
	scanner := NewScanner()

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			return
		}

		scanner.Parse(text)
	}
}
