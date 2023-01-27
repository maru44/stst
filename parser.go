package stst

import (
	"go/ast"
	"strings"

	"github.com/maru44/stst/model"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	pkg *packages.Package
}

func NewParser(pkg *packages.Package) *Parser {
	return &Parser{
		pkg: pkg,
	}
}

func (p *Parser) ParseFile(file *ast.File) []*model.Schema {
	var schemas []*model.Schema
	for _, decl := range file.Decls {
		if it, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range it.Specs {
				switch ts := spec.(type) {
				case *ast.TypeSpec:
					sc := &model.Schema{
						Name: ts.Name.Name,
					}

					switch typ := ts.Type.(type) {
					case *ast.StructType:
						p.parseStruct(typ, sc)
						// case
					}
					schemas = append(schemas)
				}
			}
		}
	}
	return schemas
}

func (p *Parser) parseStruct(st *ast.StructType, sc *model.Schema) {
	for _, f := range st.Fields.List {
		ff, ok := p.parseFields(f)
		if !ok {
			continue
		}

		sc.Fields = append(sc.Fields, ff)
	}
}

// TODO set schema and Type's pk
func (p *Parser) parseFields(f *ast.Field) (*model.Field, bool) {
	if len(f.Names) == 0 {
		return nil, false
	}

	out := &model.Field{
		Name: f.Names[0].Name,
		Tags: p.parseTag(f.Tag),
	}

	switch typ := f.Type.(type) {
	case *ast.Ident:
		out.Type = p.parseIdent(typ)
	case *ast.ArrayType:
		out.IsSlice = true
		switch elt := typ.Elt.(type) {
		case *ast.StarExpr:
			out.IsPtr = true
			switch x := elt.X.(type) {
			case *ast.Ident:
				out.Type = p.parseIdent(x)
			case *ast.SelectorExpr:
				out.Type = &model.Type{
					TypeName:   x.Sel.Name,
					Underlying: model.UnderlyingType(p.pkg.TypesInfo.TypeOf(typ).String()),
				}
			}
		case *ast.Ident:
			out.Type = p.parseIdent(elt)
		}
	case *ast.StarExpr:
		out.IsPtr = true
		if x, ok := typ.X.(*ast.Ident); ok {
			out.Type = p.parseIdent(x)
		}
		switch x := typ.X.(type) {
		case *ast.Ident:
			out.Type = p.parseIdent(x)
		case *ast.SelectorExpr:
			out.Type = &model.Type{
				TypeName:   x.Sel.Name,
				Underlying: model.UnderlyingType(p.pkg.TypesInfo.TypeOf(typ).String()),
			}
		}
	case *ast.SelectorExpr:
		// if x, ok := typ.X.(*ast.Ident); ok {
		// 	if x.Obj == nil {
		// 		out.typ = x.Name
		// 	}
		// }
		// out.typ += "." + typ.Sel.Name
		out.Type = &model.Type{
			TypeName:   typ.Sel.Name,
			Underlying: model.UnderlyingType(p.pkg.TypesInfo.TypeOf(typ).String()),
		}
	}
	return out, true
}

func (p *Parser) parseIdent(ide *ast.Ident) *model.Type {
	if ide.Obj == nil {
		return &model.Type{
			TypeName:   ide.Name,
			Underlying: model.UnderlyingType(p.pkg.TypesInfo.TypeOf(ide).String()),
		}
	}
	if ide.Obj.Decl != nil {
		if spec, ok := ide.Obj.Decl.(*ast.TypeSpec); ok {
			switch typ := spec.Type.(type) {
			case *ast.Ident:
				// enum like
				return &model.Type{
					TypeName:   ide.Name,
					Underlying: model.UnderlyingType(p.pkg.TypesInfo.TypeOf(typ).String()),
				}
			case *ast.StructType:
				// pass
			}
		}
	}
	return &model.Type{
		TypeName:   ide.Obj.Name,
		Underlying: model.UnderlyingType(p.pkg.TypesInfo.TypeOf(ide).String()),
	}
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

		tag := &model.Tag{
			Key:      kv[0],
			Values:   strings.Split(kv[1], ","),
			RawValue: kv[1],
		}
		out = append(out, tag)
	}
	return out
}
