package parser

import "go/ast"

// GetErrors will collect all of the errors from the construction
// of the AST. If there were no errors, this will return nil. If the
// limit is set to positive number, it will limit the number of errors
// to the given number.
func GetErrors(root ast.Node, limit int) []error {
	v := errorVisitor{limit: limit}
	ast.Walk(&v, root)
	return v.errs
}

type errorVisitor struct {
	errs  []error
	limit int
}

func (ev *errorVisitor) Visit(n ast.Node) ast.Visitor {
	if ev.limit > 0 && len(ev.errs) >= ev.limit {
		return nil
	} else if n == nil {
		return ev
	}

	if err, ok := n.(error); ok {
		ev.errs = append(ev.errs, err)
	}
	return ev
}
