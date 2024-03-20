package main

// Token represents a lexical token.
type Token struct {
	Type    TokenType
	Literal string
}

// TokenType represents the type of tokens.
type TokenType string

const (
	// Define token types here
	INT    TokenType = "INT"
	PLUS   TokenType = "PLUS"
	MINUS  TokenType = "MINUS"
	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
	EOF    TokenType = "EOF"
)

// Lexer breaks the input text into tokens.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// NewLexer returns a new instance of Lexer.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character from the input and advances the lexer's positions.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token in the input.
func (l *Lexer) NextToken() Token {
	var tok Token

	// Implement tokenization logic here

	return tok
}

// Parser represents a parser.
type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
}

// NewParser returns a new instance of Parser.
func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken reads the next token from the lexer and updates the parser's tokens.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// ParseProgram parses the input and returns an abstract syntax tree.
func (p *Parser) ParseProgram() *Program {
	// Implement the parsing logic here

	return &Program{}
}

// Program represents the root node of the abstract syntax tree.
type Program struct {
	// Define the AST nodes here
}

// Compiler represents a compiler.
type Compiler struct {
	parser *Parser
}

// NewCompiler returns a new instance of Compiler.
func NewCompiler(parser *Parser) *Compiler {
	return &Compiler{parser: parser}
}

// Compile compiles the source code into target code.
func (c *Compiler) Compile() {
	// Implement the compilation logic here
}

// func main() {
// 	input := "2 + 2" // Example input
// 	lexer := NewLexer(input)
// 	parser := NewParser(lexer)
// 	compiler := NewCompiler(parser)

// 	compiler.Compile()

// 	fmt.Println("Compilation complete.")
// }
