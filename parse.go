package stst

import (
	"go/ast"
	"strings"

	"github.com/maru44/stst/model"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	Pkg *packages.Package
}

func NewParser(pkg *packages.Package) *Parser {
	return &Parser{
		Pkg: pkg,
	}
}

func (p *Parser) ParseFile(file *ast.File) []*model.Schema {
	var schemas []*model.Schema
	for _, decl := range file.Decls {
		if it, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range it.Specs {
				switch ts := spec.(type) {
				case *ast.TypeSpec:
					sc := p.parseTypeSpec(ts)
					schemas = append(schemas, sc)
				}
			}
		}
	}
	return schemas
}

func (p *Parser) parseTypeSpec(spec *ast.TypeSpec) *model.Schema {
	sc := &model.Schema{
		Name: spec.Name.Name,
	}

	switch typ := spec.Type.(type) {
	case *ast.StructType:
		p.parseStruct(typ, sc)
	case *ast.Ident:
		// need?
		sc.Type, _ = p.parseIdent(typ)
	}
	return sc
}

func (p *Parser) parseStruct(st *ast.StructType, sc *model.Schema) {
	for _, f := range st.Fields.List {
		ff, ok := p.parseField(f)
		if !ok {
			continue
		}

		sc.Fields = append(sc.Fields, ff)
	}
}

func (p *Parser) parseField(f *ast.Field) (*model.Field, bool) {
	var name string
	if len(f.Names) != 0 {
		name = f.Names[0].Name
	}

	out := &model.Field{
		Tags: p.parseTag(f.Tag),
	}

	if f.Comment != nil {
		coms := make([]string, len(f.Comment.List))
		for i, c := range f.Comment.List {
			coms[i] = c.Text
		}
		out.Comment = coms
	}

	switch typ := f.Type.(type) {
	case *ast.Ident:
		out.Type, out.Schema = p.parseIdent(typ)

		// set name for embeded struct
		if len(f.Names) == 0 {
			name = typ.Name
		}
		if name == "" {
			return nil, false
		}
	case *ast.ArrayType:
		out.IsSlice = true
		switch elt := typ.Elt.(type) {
		case *ast.StarExpr:
			out.IsPtr = true
			switch x := elt.X.(type) {
			case *ast.Ident:
				out.Type, out.Schema = p.parseIdent(x)
			case *ast.SelectorExpr:
				out.Type = &model.Type{
					TypeName:   x.Sel.Name,
					Underlying: model.UnderlyingType(strings.TrimLeft(p.Pkg.TypesInfo.TypeOf(typ).String(), "*")),
				}
			}
		case *ast.Ident:
			out.Type, out.Schema = p.parseIdent(elt)
		}
	case *ast.StarExpr:
		out.IsPtr = true
		switch x := typ.X.(type) {
		case *ast.Ident:
			out.Type, out.Schema = p.parseIdent(x)
		case *ast.SelectorExpr:
			// imported pointer type (like *time.Time)
			out.Type = &model.Type{
				TypeName:   x.Sel.Name,
				Underlying: model.UnderlyingType(strings.TrimLeft(p.Pkg.TypesInfo.TypeOf(typ).String(), "*")),
			}
		}
	case *ast.SelectorExpr:
		// interface, something imported struct (like time.Time)

		// set name for embeded interface
		if name == "" {
			name = typ.Sel.Name
		}
		out.Type = &model.Type{
			TypeName:   typ.Sel.Name,
			Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(typ).String()),
		}
	}
	out.Name = name
	if out.Type != nil {
		out.Type.SetPackage()
	}
	return out, true
}

func (p *Parser) parseIdent(ide *ast.Ident) (*model.Type, *model.Schema) {
	if ide.Obj == nil {
		return &model.Type{
			TypeName:   ide.Name,
			Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
		}, nil
	}
	if ide.Obj.Decl != nil {
		if spec, ok := ide.Obj.Decl.(*ast.TypeSpec); ok {
			switch typ := spec.Type.(type) {
			case *ast.Ident:
				// enum like
				return &model.Type{
					TypeName:   ide.Name,
					Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(typ).String()),
				}, nil
			case *ast.StructType:
				sc := p.parseTypeSpec(spec)
				return &model.Type{
					TypeName:   ide.Name,
					Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
				}, sc
			}
		}
	}
	return &model.Type{
		TypeName:   ide.Obj.Name,
		Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
	}, nil
}

func (p *Parser) parseTag(tag *ast.BasicLit) []*model.Tag {
	if tag == nil {
		return nil
	}

	var out []*model.Tag
	tags := strings.Split(strings.Trim(tag.Value, "`"), " ")
	for _, t := range tags {
		kv := strings.Split(t, ":")
		if len(kv) != 2 {
			continue
		}

		v := strings.Trim(kv[1], `"`)
		tag := &model.Tag{
			Key:      kv[0],
			Values:   strings.Split(v, ","),
			RawValue: v,
		}
		out = append(out, tag)
	}
	return out
}
