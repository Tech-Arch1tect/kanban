// ast.go - Defines the AST (Abstract Syntax Tree) structures for query expressions.
package taskquery

type Expression interface{}

type ComparisonExpr struct {
	Field    string
	Operator string
	Value    interface{}
}

type LogicalExpr struct {
	Left     Expression
	Operator string
	Right    Expression
}
