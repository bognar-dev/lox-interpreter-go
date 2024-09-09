package parsing

import (
	"github.com/codecrafters-io/interpreter-starter-go/cmd/myinterpreter/scanning"
)

type Evaluator struct {
}

func (a *Evaluator) VisitLiteralExpr(expr *LiteralExpr) any {
	return expr.Value
}

func (a *Evaluator) VisitBinaryExpr(expr *BinaryExpr) any {
	left := expr.Left.Accept(a)
	right := expr.Right.Accept(a)

	switch expr.Operator.TokenType {
	case scanning.MINUS:
		return left.(float64) - right.(float64)
	case scanning.SLASH:
		return left.(float64) / right.(float64)
	case scanning.STAR:
		return left.(float64) * right.(float64)
	case scanning.PLUS:
		if leftFloat, ok := left.(float64); ok {
			if rightFloat, ok := right.(float64); ok {
				return leftFloat + rightFloat
			}
		}
		if leftStr, ok := left.(string); ok {
			if rightStr, ok := right.(string); ok {
				return leftStr + rightStr
			}
		}
	case scanning.GREATER:
		return left.(float64) > right.(float64)
	case scanning.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case scanning.LESS:
		return left.(float64) < right.(float64)
	case scanning.LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case scanning.BANG_EQUAL:
		return !isEqual(left, right)
	case scanning.EQUAL_EQUAL:
		return isEqual(left, right)
	}

	return nil
}

func (a *Evaluator) VisitUnaryExpr(expr *UnaryExpr) any {
	right := expr.Right.Accept(a)

	switch expr.Operator.TokenType {
	case scanning.MINUS:
		return -(right.(float64))
	case scanning.BANG:
		return !isTruthy(right)
	}

	return nil
}

func (a *Evaluator) VisitGroupingExpr(expr *GroupingExpr) any {
	return a.evaluate(expr.Expression)
}

func (a *Evaluator) evaluate(expr Expr) any {
	return expr.Accept(a)
}

func isTruthy(value any) bool {
	if value == nil {
		return false
	}
	if b, ok := value.(bool); ok {
		return b
	}
	return true
}

func isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}
