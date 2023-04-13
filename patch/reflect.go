package patch

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"strings"
)

func (p *Patcher) patchNoReflect(id *ast.Ident, obj types.Object) {
	var found bool
	for _, v := range p.noReflect {
		if found = v == obj; found {
			break
		}
	}
	if !found {
		return
	}
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return
	}
	pos := make(map[string]token.Pos)

	// remove all reflect methods
	for i := 0; i < named.NumMethods(); i++ {
		m := named.Method(i)
		switch m.Name() {
		case "String":
		case "Reset":
		case "ProtoMessage":
		case "ProtoReflect":
		case "Descriptor":
		case "Type":
		case "Number":
		case "EnumDescriptor":
		default:
			continue
		}
		pos[m.Name()] = m.Pos()
	}
	// remove fields
	if s, ok := id.Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType); ok {
		var keep []*ast.Field
		for _, v := range s.Fields.List {
			if len(v.Names) == 0 {
				keep = append(keep, v)
				continue
			}
			for _, vv := range v.Names {
				switch vv.Name {
				case "state":
				case "sizeCache":
				// we need to preserve the unknown fields
				case "unknownFields":
					fallthrough
				default:
					keep = append(keep, v)
				}
			}
		}
		s.Fields.List = keep
	}
	var comments []*ast.Comment
	addComments := func(g *ast.CommentGroup) {
		if g == nil {
			return
		}
		comments = append(comments, g.List...)
	}
	file := p.fileOf(id)
	if file == nil {
		log.Printf("Warning: file not found: %v", obj)
		return
	}
	for k, v := range pos {
		found := false
		for i, vv := range file.Decls {
			vv, ok := vv.(*ast.FuncDecl)
			if !ok || vv.Recv == nil {
				continue
			}
			if found = v == vv.Name.NamePos; found {
				file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)
				addComments(vv.Doc)
				break
			}
		}
		if !found {
			log.Printf("Warning: getter not found: %v `%s`", obj, k)
		}
	}
	filterSpec := func(s ast.Spec) bool {
		switch v := s.(type) {
		case *ast.ImportSpec:
			switch v.Name.Name {
			case "protoreflect":
			case "reflect":
			case "sync":
			default:
				return false
			}
			return true
		case *ast.ValueSpec:
			if len(v.Names) == 0 {
				return false
			}
			switch {
			case v.Names[0].Name == "_":
			case strings.ToLower(v.Names[0].Name)[0] == v.Names[0].Name[0]:
			case strings.HasPrefix(v.Names[0].Name, "File_"):
			default:
				return false
			}
			addComments(v.Doc)
			addComments(v.Comment)
			return true
		}
		return false
	}
	filterDecl := func(d ast.Decl) bool {
		if d == nil {
			return true
		}
		switch v := d.(type) {
		case *ast.GenDecl:
			var specs []ast.Spec
			for _, vv := range v.Specs {
				if !filterSpec(vv) {
					specs = append(specs, vv)
				}
			}
			v.Specs = specs
			if len(v.Specs) == 0 {
				addComments(v.Doc)
				return true
			}
		case *ast.FuncDecl:
			if v.Recv != nil {
				return false
			}
			addComments(v.Doc)
			return true
		}
		return false
	}
	// remove file level declarations, e.g. all unexported vars and methods
	var decls []ast.Decl
	for _, v := range file.Decls {
		if !filterDecl(v) {
			decls = append(decls, v)
		}
	}
	file.Decls = decls
	// remove comments
	for _, v := range comments {
		for _, vv := range file.Comments {
			var keep []*ast.Comment
			for _, vvv := range vv.List {
				if vvv == v {
					continue
				}
				keep = append(keep, vvv)
			}
			vv.List = keep
		}
	}
}
