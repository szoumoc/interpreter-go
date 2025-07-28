package lexer

import "go-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

/* The ones that might cause some confusion
right now are position and readPosition. Both will be used to access characters in input by
using them as an index, e.g.: l.input[l.readPosition]. The reason for these two “pointers”
pointing into our input string is the fact that we will need to be able to “peek” further into
the input and look after the current character to see what comes up next. readPosition always
points to the “next” character in the input. position points to the character in the input that
corresponds to the ch byte.
*/

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // l.ch = NUL

	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Some language implementations do create tokens for newline characters for example and throw parsing errors if they are
// not at the correct place in the stream of tokens. We skip over newline characters to make the
// parsing step later on a little easier.

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// I don’t know if you noticed, but we simplified things a lot in readNumber. We only read in
// integers. What about floats? Or numbers in hex notation? Octal notation? We ignore them
// and just say that Monkey doesn’t support this. Of course, the reason for this is again the
// educational aim and limited scope of this book.

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
