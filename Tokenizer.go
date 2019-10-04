package main

import (
	"unicode"
)

const (
	OPENING_BRACKET = iota
	CLOSING_BRACKET = iota
	LITERAL         = iota
)

type TokenType int

type Token struct {
	Type     TokenType
	Text     string
	Position int
}

func Tokenize(text string) <-chan *Token {
	rr := []rune(text)
	position := 0
	result := make(chan *Token)

	skipSpaces := func() {
		for position < len(text) && unicode.IsSpace(rr[position]) {
			position++
		}
	}

	go func() {
		defer close(result)

		skipSpaces()
		for position < len(text) {
			if text[position] == '[' {
				result <- &Token{Type: OPENING_BRACKET, Position: position}
				position++
				skipSpaces()
			} else if text[position] == ']' {
				result <- &Token{Type: CLOSING_BRACKET, Position: position}
				position++
				skipSpaces()
			} else {
				start := position
				for position < len(text) && !unicode.IsSpace(rr[position]) {
					position++
				}
				result <- &Token{Type: LITERAL, Text: string(rr[start:position]), Position: start}
				skipSpaces()
			}
		}
	}()

	return result
}
