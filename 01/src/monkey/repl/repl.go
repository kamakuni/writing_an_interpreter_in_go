package repl

import (
	"bufio"
	"fmt"
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/lexer"
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/parser"
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			continue
		}
	}
}
