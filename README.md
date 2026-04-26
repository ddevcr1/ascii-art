# ascii-art

A Go program that takes a string as an argument and outputs it as ASCII graphics using banner files.

## Description

The project reads a banner file with a specific graphical template and generates a graphical representation of the input string. Each character has a height of 8 lines and is separated by a newline `\n`.

Supported:
- Latin letters (upper and lower case)
- Digits
- Spaces
- Special characters (`!`, `@`, `#`, `{`, `}`, etc.)
- Escape sequence `\n` inside the argument (line break)

## Project structure

```
ascii-art/
├── main.go          # main code
├── main_test.go     # unit tests
├── standard.txt     # standard banner
├── shadow.txt       # shadow banner
├── thinkertoy.txt   # thinkertoy banner
└── README.md
```

## Usage

```bash
go run . "Hello"
go run . "hello"
go run . "HELLO"
go run . "Hello\nThere"
go run . "1Hello 2There"
go run . "{Hello & There #}"
```

An empty string outputs nothing:
```bash
go run . ""
```

## Changing the font

Currently uses `standard.txt`. To switch the banner to `shadow` or `thinkertoy`, edit the line in `main.go`:

```go
lines, err := loadBanner("standard.txt")
```

## Tests

Run unit tests:

```bash
go test -v ./...
```
