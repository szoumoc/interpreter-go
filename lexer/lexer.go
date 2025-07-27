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
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
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

// We added a default branch to our switch statement, so we can check for identifiers whenever
// the l.ch is not one of the recognized characters. We also added the generation of token.ILLEGAL
// tokens. If we end up there, we truly don’t know how to handle the current character and declare
// it as token.ILLEGAL.
// The isLetter helper function just checks whether the given argument is a letter. That sounds
// easy enough, but what’s noteworthy about isLetter is that changing this function has a larger
// impact on the language our interpreter will be able to parse than one would expect from such
// a small function. As you can see, in our case it contains the check ch == '_', which means that
// we’ll treat _ as a letter and allow it in identifiers and keywords. That means we can use variable
// names like foo_bar. Other programming languages even allow ! and ? in identifiers. If you
// want to allow that too, this is the place to sneak it in.
// readIdentifier() does exactly what its name suggests: it reads in an identifier and advances
// our lexer’s positions until it encounters a non-letter-character.
// In the default: branch of the switch statement we use readIdentifier() to set the Literal field
// of our current token. But what about its Type? Now that we have read identifiers like let, fn
// or foobar, we need to be able to tell user-defined identifiers apart from language keywords. We
// need a function that returns the correct TokenType for the token literal we have. What better
// place than the token package to add such a function?
