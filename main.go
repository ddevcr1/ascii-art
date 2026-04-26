package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	text := strings.ReplaceAll(string(content), "\r\n", "\n")
	lines := strings.Split(text, "\n")
	return lines, nil
}

func getChar(lines []string, ch rune)[]string {
	index := (int(ch) - 32) * 9
	return lines[index+1 : index+9]
}

func printAscii(input string, lines[]string) {
	if input == strings.Repeat(`\n`, len(input)/2) {
		for i := 0; i < len(input)/2; i++ {
			fmt.Println()
		}
		return
	}

	parts := strings.Split(input, `\n`)
	for _, part := range parts {
		if part == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range part {
				if ch < 32 || ch > 126 {
					fmt.Printf("Error: input %q contains unsupported characters\n", input)
					os.Exit(1)
				}
				charLines := getChar(lines, ch)
				fmt.Print(charLines[row])
			}
			fmt.Println()
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <string>")
		return
	}
	input := os.Args[1]
	if input == "" {
		return
	}

	lines, err := loadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	printAscii(input, lines)
}
