package patch

import (
	"go/ast"
	"go/types"
)

func renameType(f *ast.File, info *types.Info, old, new string) {
	renameInFunc := func(fn *ast.FuncDecl) {
		rename := func(decl *ast.AssignStmt) {
			for _, v := range decl.Rhs {
				switch v := v.(type) {
				case *ast.TypeAssertExpr:
					if i, ok := v.Type.(*ast.Ident); ok && i.Name == old {
						i.Name = new
					}
					s, ok := v.Type.(*ast.StarExpr)
					if !ok {
						continue
					}
					if i, ok := s.X.(*ast.Ident); ok && i.Name == old {
						i.Name = new
					}
				case *ast.CompositeLit:
					if t, ok := v.Type.(*ast.Ident); ok && t.Name == old {
						t.Name = new
					}
				}
			}
			return
		}
		for _, v := range fn.Body.List {
			switch v := v.(type) {
			case *ast.DeclStmt:
				d, ok := v.Decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, s := range d.Specs {
					v, ok := s.(*ast.ValueSpec)
					if !ok {
						continue
					}
					if t, ok := v.Type.(*ast.Ident); ok && t.Name == old {
						t.Name = new
					}
				}
			case *ast.IfStmt:
				i, ok := v.Cond.(*ast.Ident)
				if !ok {
					continue
				}
				if v, ok := i.Obj.Decl.(*ast.AssignStmt); ok {
					rename(v)
				}
			case *ast.TypeSwitchStmt:
				for _, v := range v.Body.List {
					c, ok := v.(*ast.CaseClause)
					if !ok {
						continue
					}
					for _, v := range c.List {
						if i, ok := v.(*ast.Ident); ok && i.Name == old {
							i.Name = new
						}
					}
				}
			case *ast.AssignStmt:
				rename(v)
			}
		}
	}
	for _, v := range f.Decls {
		fn, ok := v.(*ast.FuncDecl)
		if !ok {
			continue
		}
		renameInFunc(fn)
	}
	for id, v := range info.Defs {
		if v == nil {
			continue
		}
		if v.Name() == old {
			id.Name = new
		}
	}
	for id, v := range info.Uses {
		if v == nil {
			continue
		}
		if v.Name() == old {
			id.Name = new
		}
	}
}
