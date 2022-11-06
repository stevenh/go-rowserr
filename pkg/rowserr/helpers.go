package rowserr

import (
	"go/ast"
	"go/types"
)

// isRowsPtrType returns true if typ is contained in rowTypes, false otherwise.
func isRowsPtrType(typ types.Type, rowTypes map[string]struct{}) bool {
	if typ == nil {
		return false
	}

	_, ok := rowTypes[typ.String()]
	return ok
}

// rootIdent finds the root identifier x in a chain of selections x.y.z, or nil if not found.
func rootIdent(n ast.Node) *ast.Ident {
	switch n := n.(type) {
	case *ast.SelectorExpr:
		return rootIdent(n.X)
	case *ast.Ident:
		return n
	default:
		return nil
	}
}

// restOfBlock, given a traversal stack, finds the innermost containing
// block and returns the suffix of its statements starting with the current
// node.
func restOfBlock(stack []ast.Node) []ast.Stmt {
	for i := len(stack) - 1; i >= 0; i-- {
		if b, ok := stack[i].(*ast.BlockStmt); ok {
			for j, v := range b.List {
				if v == stack[i+1] {
					return b.List[j:]
				}
			}
			return nil
		}
	}

	return nil
}
