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
const MONKEY_FACE = ""

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
			printParserErrors(out, p.Errors())
			continue
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkye business here!\n")
	io.WriteString(out, "  parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
