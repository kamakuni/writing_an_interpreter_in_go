package parser

import (
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/ast"
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/lexer"
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
