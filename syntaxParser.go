package main

import (
	"unicode"
)

type TokenType string

const (
	TokenTypeKeyword   = iota
	TokenTypeDigit     = "TokenTypeDigit"
	TokenTypeCharacter = "TokenTypeCharacter"
	TokenTypeSymbol    = "TokenTypeSymbol"
	TokenTypeMath      = "TokenTypeMath"
	TokenTypePunct     = "TokenTypePunct"
	TokenTypeOther     = "TokenTypeOther"
	// Add more token types as needed.
)

// Token represents a lexical token.
type Token struct {
	Type  TokenType
	Value string
}

func tokenize(text [][]rune) [][]Token {
	var tokens [][]Token
	for _, line := range text {
		var currentToken string
		var tokenLine []Token
		for _, r := range line {
			switch {
			case unicode.IsLetter(r):
				currentToken += string(r)
			case unicode.IsDigit(r), unicode.IsPunct(r), unicode.Is(unicode.Sm, r), unicode.IsSymbol(r):
				var tokenType TokenType
				switch {
				case unicode.IsDigit(r):
					tokenType = TokenTypeDigit
				case unicode.IsPunct(r):
					tokenType = TokenTypePunct
				case unicode.Is(unicode.Sm, r):
					tokenType = TokenTypeMath
				case unicode.IsSymbol(r):
					tokenType = TokenTypeSymbol
				}
				tokenLine = append_token(tokenLine, tokenType, string(r), &currentToken)
			default:
				tokenLine = append_token(tokenLine, TokenTypeOther, string(r), &currentToken)
			}
		}
		if currentToken != "" {
			tokenLine = append(tokenLine, Token{TokenTypeOther, currentToken})
		}
		tokens = append(tokens, tokenLine)
	}

	return tokens
}

func is_not_empty_token(currentToken string) bool {
	return currentToken != ""
}

func append_token(tokenLine []Token, tokenType TokenType, value string, currentToken *string) []Token {
	if is_not_empty_token(*currentToken) {
		tokenLine = append(tokenLine, Token{TokenTypeCharacter, *currentToken})
		*currentToken = ""
	}
	tokenLine = append(tokenLine, Token{tokenType, value})
	return tokenLine
}
