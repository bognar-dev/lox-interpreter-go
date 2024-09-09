package parsing

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/tokens"
)

type Expr interface {
	Accept(visitor Visitor) any
}

type LiteralExpr struct {
	Value any
}

func (l *LiteralExpr) Accept(visitor Visitor) any {
	return visitor.VisitLiteralExpr(l)
}

type BinaryExpr struct {
	Left     Expr
	Operator tokens.Token
	Right    Expr
}

func (b *BinaryExpr) Accept(visitor Visitor) any {
	return visitor.VisitBinaryExpr(b)
}

type UnaryExpr struct {
	Operator tokens.Token
	Right    Expr
}

func (u *UnaryExpr) Accept(visitor Visitor) any {
	return visitor.VisitUnaryExpr(u)
}

type GroupingExpr struct {
	Expression Expr
}

func (g *GroupingExpr) Accept(visitor Visitor) any {
	return visitor.VisitGroupingExpr(g)
}
