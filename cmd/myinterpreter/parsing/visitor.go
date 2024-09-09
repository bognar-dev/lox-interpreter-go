package parsing

type Visitor interface {
	VisitLiteralExpr(expr *LiteralExpr) any
	VisitBinaryExpr(expr *BinaryExpr) any
	VisitUnaryExpr(expr *UnaryExpr) any
	VisitGroupingExpr(expr *GroupingExpr) any
}
