package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Parser represents a parser, including a scanner and the underlying raw input.
// It also contains a small buffer to allow for two unscans.
type Parser struct {
	s   *Lexer
	raw string
	buf TokenStack
}

// NewParser returns a new instance of Parser.
func NewParser(s string) *Parser {
	return &Parser{s: NewLexer(strings.NewReader(s)), raw: s}
}

// Parse takes the raw string and returns the root node of the AST.
func (p *Parser) Parse() (*Node, error) {
	operation, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	tok, lit := p.scanIgnoreWhitespace()
	if tok != EOF {
		return nil, fmt.Errorf("token %s after parsing not expected", lit)
	}
	return operation, nil
}

/*
4*8-6/2
func (p *Parser) parseExpression() (*Node, error) {
	// simple value or (something)
	node := findValue()

	for {
		operator := findOperator()
		nextValue := findValue()
	}
}
*/

func (p *Parser) parseExpression() (*Node, error) {
	// simple value or (something)
	node, err := p.findValue()
	if err != nil {
		return nil, err
	}

	for {
		// find operator (+,-,*,/)
		tok, lit := p.scanIgnoreWhitespace()
		if tok == EOF {
			return node, nil
		}
		if tok == CLOSED_PARENTHESIS {
			p.unscan(TokenInfo{Token: tok, Literal: lit})
			return node, nil
		}
		if tok != PLUS && tok != MINUS && tok != MULTIPLY && tok != DIVIDE {
			return nil, fmt.Errorf("missing operator, we found %s instead", lit)
		}

		if tok == PLUS || tok == MINUS {
			nextValue, err := p.parseExpression()
			if err != nil {
				return nil, err
			}

			newNode := &Node{
				Left:     node,
				Operator: tok,
				Right:    nextValue,
			}
			node = newNode
		} else {
			nextValue, err := p.findValue()
			if err != nil {
				return nil, err
			}
			newNode := &Node{
				Left:     node,
				Operator: tok,
				Right:    nextValue,
			}
			node = newNode
		}
	}
	return node, nil
}

func (p *Parser) findValue() (*Node, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok == OPEN_PARENTHESIS {
		node, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		tok, lit = p.scanIgnoreWhitespace()
		if tok != CLOSED_PARENTHESIS {
			return nil, fmt.Errorf("unexpected token %s", lit)
		}
		return node, nil
	} else if tok == NUMBER {
		node := &Node{}
		n, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			return nil, err
		}
		node.Number = &n
		node.Operator = EOF
		return node, nil
	} else {
		return nil, fmt.Errorf("unexpected token %s", lit)
	}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.Len() != 0 {
		// Can ignore the error since it's not empty.
		tokenInf, _ := p.buf.Pop()
		return tokenInf.Token, tokenInf.Literal
	}

	// Otherwise read the next token from the scanner.
	tokenInf := p.s.Scan()
	tok, lit = tokenInf.Token, tokenInf.Literal
	return tok, lit
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return tok, lit
}

// unscan pushes the previously read tokens back onto the buffer.
func (p *Parser) unscan(tok TokenInfo) {
	p.buf.Push(tok)
}
