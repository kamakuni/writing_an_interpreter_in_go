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

func (p *Parser) NextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}
