package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hasError = false

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

	run(string(file))

	if hasError {
		os.Exit(1)
	}
}

func parsePrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			return
		}

		run(text)

		hasError = false
	}
}

func run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}

// fuck is just the error function, but collides with a builtin so whatever
func fuck(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, msg)
	hasError = true
}
