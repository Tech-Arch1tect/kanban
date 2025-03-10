// parser.go - Implements a parser for converting query strings into an AST.
package taskquery

import (
	"errors"
	"fmt"
)

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(input string) *Parser {
	l := NewLexer(input)
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseQuery() (Expression, error) {
	return p.parseExpression()
}

func (p *Parser) parseExpression() (Expression, error) {
	left, err := p.parseComparison()
	if err != nil {
		return nil, err
	}
	for p.curToken.Type == TokenLogical {
		op := p.curToken.Value
		p.nextToken()
		right, err := p.parseComparison()
		if err != nil {
			return nil, err
		}
		left = &LogicalExpr{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}
	return left, nil
}

func (p *Parser) parseComparison() (Expression, error) {
	if p.curToken.Type == TokenLParen {
		p.nextToken()
		expr, err := p.ParseQuery()
		if err != nil {
			return nil, err
		}
		if p.curToken.Type != TokenRParen {
			return nil, errors.New("expected ')'")
		}
		p.nextToken()
		return expr, nil
	}
	if p.curToken.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected identifier, got %s", p.curToken.Value)
	}
	field := p.curToken.Value
	p.nextToken()

	if p.curToken.Type != TokenOperator {
		return nil, fmt.Errorf("expected operator, got %s", p.curToken.Value)
	}
	operator := p.curToken.Value
	p.nextToken()

	var value interface{}
	switch p.curToken.Type {
	case TokenString, TokenNumber, TokenIdentifier:
		value = p.curToken.Value
	default:
		return nil, fmt.Errorf("expected literal, got %s", p.curToken.Value)
	}
	p.nextToken()

	return &ComparisonExpr{
		Field:    field,
		Operator: operator,
		Value:    value,
	}, nil
}
