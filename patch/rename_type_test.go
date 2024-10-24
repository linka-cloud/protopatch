package patch

import (
	"bytes"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenameType(t *testing.T) {
	src := `package main

type T struct {
	s Something
}

type Something struct {}

func (s Something) Do(v any) {
	s, ok := v.(Something)
	if !ok {
		return
	}
	_ = s
}

func (s *Something) Do2() {}

type SomethingElse = Something

func do(v any) {
	s, ok := v.(Something)
	if !ok {
		return
	}
	_ = s
}

func do2(v any) {
	if s, ok := v.(Something); ok {
		_ = s
	}
	if s, ok := v.(*Something); ok {
		_ = s
	}
}

func do3(v any) {
	var s Something
	var s2 Something = s
	_ = s2
}

func do4(v any) {
	switch v.(type) {
	case Something:
	}
}

func do5(v SomethingElse) {
	s := Something(v)
	_ = s
}
`
	want := `package main

type T struct {
	s Other
}

type Other struct{}

func (s Other) Do(v any) {
	s, ok := v.(Other)
	if !ok {
		return
	}
	_ = s
}

func (s *Other) Do2()	{}

type SomethingElse = Other

func do(v any) {
	s, ok := v.(Other)
	if !ok {
		return
	}
	_ = s
}

func do2(v any) {
	if s, ok := v.(Other); ok {
		_ = s
	}
	if s, ok := v.(*Other); ok {
		_ = s
	}
}

func do3(v any) {
	var s Other
	var s2 Other = s
	_ = s2
}

func do4(v any) {
	switch v.(type) {
	case Other:
	}
}

func do5(v SomethingElse) {
	s := Other(v)
	_ = s
}
`
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	files := []*ast.File{f}
	conf := types.Config{Importer: importer.Default(), Error: func(err error) {}}
	pkg, err := conf.Check("", fset, files, info)
	if err != nil {
		t.Fatal(err)
	}
	_ = pkg

	renameType(f, info, "Something", "Other")

	b := &bytes.Buffer{}
	if err := printer.Fprint(b, fset, f); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, b.String())
}
