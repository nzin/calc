package parser

import "fmt"

type Node struct {
	Left     *Node
	Operator Token // PLUS, MINUS, MULTIPLY, DIVIDE
	Number   *float64
	Right    *Node
}

func (n *Node) Compute() (float64, error) {
	if n.Operator == EOF {
		return *n.Number, nil
	}
	a, err := n.Left.Compute()
	if err != nil {
		return 0, err
	}
	b, err := n.Right.Compute()
	if err != nil {
		return 0, err
	}
	switch n.Operator {
	case PLUS:
		return a + b, nil
	case MINUS:
		return a - b, nil
	case MULTIPLY:
		return a * b, nil
	case DIVIDE:
		if b == 0 {
			return 0, fmt.Errorf("division by 0")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("missing operator")
}
