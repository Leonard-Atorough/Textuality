package main

import (
	"unicode"
)

// Token types
const (
	TokenTypeKeyword  = iota
	TokenTypeDigit    = "TokenTypeDigit"
	TokenTypeCharater = "TokenTypeCharater"
	TokenTypeSymbol   = "TokenTypeSymbol"
	TokenTypeMath     = "TokenTypeMath"
	TokenTypePunct    = "TokenTypePunct"
	TokenTypeOther    = "TokenTypeOther"
	// Add more token types as needed.
)

// Token represents a lexical token.
type Token struct {
	Type  string
	Value rune
}

func tokenize(text [][]rune) [][]Token {
	var tokens [][]Token

	for _, l := range text {
		var tokenLine []Token
		for _, r := range l {
			switch {
			// case unicode.IsLetter(r):
			// 	tokenLine = append(tokenLine, Token{TokenTypeCharater, r})
			case unicode.IsDigit(r):
				tokenLine = append(tokenLine, Token{TokenTypeDigit, r})
			case unicode.IsPunct(r):
				tokenLine = append(tokenLine, Token{TokenTypePunct, r})
			case unicode.Is(unicode.Sm, r):
				tokenLine = append(tokenLine, Token{TokenTypeMath, r})
			case unicode.IsSymbol(r):
				tokenLine = append(tokenLine, Token{TokenTypeSymbol, r})
			default:
				tokenLine = append(tokenLine, Token{TokenTypeOther, r})
			}
		}
		tokens = append(tokens, tokenLine)
	}

	return tokens
}
