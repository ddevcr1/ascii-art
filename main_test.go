package main

import (
	"strings"
	"testing"
)

var bannerLines []string

func init() {
	var err error
	bannerLines, err = loadBanner("standard.txt")
	if err != nil {
		panic("could not load standard.txt: " + err.Error())
	}
}

func captureOutput(input string) string {
	parts := strings.Split(input, `\n`)
	var result strings.Builder
	for _, part := range parts {
		if part == "" {
			result.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range part {
				charLines := getChar(bannerLines, ch)
				result.WriteString(charLines[row])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}

func TestHello(t *testing.T) {
	out := captureOutput("hello")
	if !strings.Contains(out, "| |") {
		t.Error("output for 'hello' looks wrong")
	}
}

func TestNewline(t *testing.T) {
	out := captureOutput(`Hello\nThere`)
	lines := strings.Split(out, "\n")
	if len(lines) < 17 {
		t.Errorf("expected at least 17 lines for 'Hello\\nThere', got %d", len(lines))
	}
}

func TestDoubleNewline(t *testing.T) {
	out := captureOutput(`Hello\n\nThere`)
	lines := strings.Split(out, "\n")
	if len(lines) < 18 {
		t.Errorf("expected at least 18 lines for 'Hello\\n\\nThere', got %d", len(lines))
	}
}

func TestNumbers(t *testing.T) {
	out := captureOutput("123")
	if !strings.Contains(out, "_") {
		t.Error("output for '123' looks wrong")
	}
}

func TestSpecialChars(t *testing.T) {
	out := captureOutput("{Hello}")
	if !strings.Contains(out, "__") {
		t.Error("output for '{Hello}' looks wrong")
	}
}

func TestSpace(t *testing.T) {
	out := captureOutput("Hello World")
	if !strings.Contains(out, "  ") {
		t.Error("output for 'Hello World' should contain spaces")
	}
}
