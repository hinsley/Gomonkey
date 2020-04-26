package ast

import "github.com/hinsley/Gomonkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // This is not actually useful.
}

type Expression interface {
	Node
	expressionNode() // This is not actually useful.
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
