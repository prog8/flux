package ast

// A Visitor's Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(Node) Visitor
}

func Walk(v Visitor, n Node) {
	if v = v.Visit(n); v == nil {
		return
	}

	switch n := n.(type) {
	case *ArrayExpression:
		for _, element := range n.Elements {
			v.Visit(element)
		}

	case *ArrowFunctionExpression:
		for _, param := range n.Params {
			v.Visit(param)
		}
		v.Visit(n.Body)

	case *BinaryExpression:
		v.Visit(n.Left)
		v.Visit(n.Right)

	case *BlockStatement:
		for _, stmt := range n.Body {
			v.Visit(stmt)
		}

	case *BooleanLiteral:
		// nothing to do

	case *CallExpression:
		v.Visit(n.Callee)
		for _, arg := range n.Arguments {
			v.Visit(arg)
		}

	case *ConditionalExpression:
		// todo(jsternberg): this is probably the wrong order and this
		// expression isn't defined in the parser grammar.
		v.Visit(n.Test)
		v.Visit(n.Consequent)
		v.Visit(n.Alternate)

	case *DateTimeLiteral:
		// nothing to do

	case *DurationLiteral:
		// nothing to do

	case *ExpressionStatement:
		v.Visit(n.Expression)

	case *FloatLiteral:
		// nothing to do

	case *Identifier:
		// nothing to do

	case *IndexExpression:
		v.Visit(n.Array)
		v.Visit(n.Index)

	case *IntegerLiteral:
		// nothing to do

	case *LogicalExpression:
		v.Visit(n.Left)
		v.Visit(n.Right)

	case *MemberExpression:
		v.Visit(n.Object)
		v.Visit(n.Property)

	case *ObjectExpression:
		for _, property := range n.Properties {
			v.Visit(property)
		}

	case *OptionStatement:
		v.Visit(n.Declaration)

	case *PipeExpression:
		v.Visit(n.Argument)
		v.Visit(n.Call)

	case *PipeLiteral:
		// nothing to do

	case *Program:
		for _, stmt := range n.Body {
			v.Visit(stmt)
		}

	case *Property:
		v.Visit(n.Key)
		v.Visit(n.Value)

	case *RegexpLiteral:
		// nothing to do

	case *ReturnStatement:
		v.Visit(n.Argument)

	case *StringLiteral:
		// nothing to do

	case *UnaryExpression:
		v.Visit(n.Argument)

	case *UnsignedIntegerLiteral:
		// nothing to do

	case *VariableDeclaration:
		for _, decl := range n.Declarations {
			v.Visit(decl)
		}

	case *VariableDeclarator:
		v.Visit(n.ID)
		v.Visit(n.Init)
	}

	v.Visit(nil)
}
