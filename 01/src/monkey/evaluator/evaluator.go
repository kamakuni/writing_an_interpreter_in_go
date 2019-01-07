package evaluator

import (
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/ast"
	"github.com/kamakuni/writing_an_interpreter_in_go/01/src/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
