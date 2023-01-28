package stst

import (
	"go/ast"
	"strings"

	"github.com/maru44/stst/stmodel"
	"golang.org/x/tools/go/packages"
)

type Parser struct {
	Pkg *packages.Package
}

func NewParser(pkg *packages.Package) *Parser {
	out := &Parser{
		Pkg: pkg,
	}
	return out
}

func (p *Parser) Parse() []*stmodel.Schema {
	var schemas []*stmodel.Schema
	for _, f := range p.Pkg.Syntax {
		for _, decl := range f.Decls {
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
	}
	return schemas
}

func (p *Parser) ParseFile(f *ast.File) []*stmodel.Schema {
	var schemas []*stmodel.Schema
	for _, decl := range f.Decls {
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

func (p *Parser) parseTypeSpec(spec *ast.TypeSpec) *stmodel.Schema {
	sc := &stmodel.Schema{
		Name: spec.Name.Name,
	}

	var fin bool
	var prefixes []stmodel.TypePrefix
	ex := spec.Type
	for !fin {
		var pref stmodel.TypePrefix
		ex, pref, fin = p.purgePointerOrSlice(ex)
		if fin {
			break
		}
		prefixes = append(prefixes, pref)
	}
	sc.TypePrefixes = prefixes

	switch typ := ex.(type) {
	case *ast.StructType:
		sc.Type = p.parseIdent(spec.Name)
		sc.Type.SetPackage()

		for _, f := range typ.Fields.List {
			ff, ok := p.parseField(f)
			if !ok {
				continue
			}

			sc.Fields = append(sc.Fields, ff)
		}
	case *ast.Ident:
		sc.Type = p.parseIdent(typ)
		sc.Type.SetPackage()
	case *ast.InterfaceType:
		sc.Type = p.parseIdent(spec.Name)
		sc.Type.SetPackage()
		sc.IsInterface = true

		for _, m := range typ.Methods.List {
			ff, ok := p.parseField(m)
			if !ok {
				continue
			}
			sc.Fields = append(sc.Fields, ff)
		}
	case *ast.MapType:
		sc.Type = p.parseIdent(spec.Name)
		sc.Type.SetPackage()
		sc.Map = p.parseMap(typ)
	}
	return sc
}

func (p *Parser) parseField(f *ast.Field) (*stmodel.Field, bool) {
	var name string
	if len(f.Names) != 0 {
		name = f.Names[0].Name
	}

	out := &stmodel.Field{
		Tags: p.parseTag(f.Tag),
	}

	if f.Comment != nil {
		coms := make([]string, len(f.Comment.List))
		for i, c := range f.Comment.List {
			coms[i] = c.Text
		}
		out.Comment = coms
	}

	var fin bool
	var prefixes []stmodel.TypePrefix
	ex := f.Type
	for !fin {
		var pref stmodel.TypePrefix
		ex, pref, fin = p.purgePointerOrSlice(ex)
		if fin {
			break
		}
		prefixes = append(prefixes, pref)
	}

	switch typ := ex.(type) {
	case *ast.Ident:
		out.Type = p.parseIdent(typ)

		// set name for embeded struct
		if len(f.Names) == 0 {
			name = typ.Name
		}
		if name == "" {
			return nil, false
		}
	case *ast.SelectorExpr:
		// interface, something imported struct (like time.Time)

		// set name for embeded interface
		if name == "" {
			name = typ.Sel.Name
		}
		out.Type = &stmodel.Type{
			TypeName:   typ.Sel.Name,
			Underlying: stmodel.UnderlyingType(p.Pkg.TypesInfo.TypeOf(typ).String()),
		}
	case *ast.FuncType:
		out.Func = p.parseFunc(typ)
	case *ast.MapType:
		out.Map = p.parseMap(typ)
	case *ast.StructType:
		out.IsUntitledStruct = true
		if len(typ.Fields.List) > 0 {
			sc := &stmodel.Schema{}
			for _, f := range typ.Fields.List {
				ff, ok := p.parseField(f)
				if !ok {
					continue
				}

				sc.Fields = append(sc.Fields, ff)
			}
			out.Schema = sc
		}
	}

	out.Name = name
	if out.Type != nil {
		out.Type.SetPackage()
	}
	out.TypePrefixes = prefixes
	return out, true
}

func (p *Parser) purgePointerOrSlice(ex ast.Expr) (ast.Expr, stmodel.TypePrefix, bool) {
	switch typ := ex.(type) {
	case *ast.StarExpr:
		return typ.X, stmodel.TypePrefixPtr, false
	case *ast.ArrayType:
		return typ.Elt, stmodel.TypePrefixSlice, false
	}
	return ex, "", true
}

func (p *Parser) parseMap(m *ast.MapType) *stmodel.Map {
	key, ok := p.parseField(&ast.Field{
		Type: m.Key,
	})
	if !ok {
		return nil
	}
	value, ok := p.parseField(&ast.Field{
		Type: m.Value,
	})
	if !ok {
		return nil
	}
	return &stmodel.Map{
		Key:   key,
		Value: value,
	}
}

func (p *Parser) parseFunc(fn *ast.FuncType) *stmodel.Func {
	var args, results []*stmodel.Field
	if fn.Params != nil {
		for _, param := range fn.Params.List {
			ff, ok := p.parseField(param)
			if ok {
				args = append(args, ff)
			}
		}
	}
	if fn.Results != nil {
		for _, res := range fn.Results.List {
			ff, ok := p.parseField(res)
			if ok {
				results = append(results, ff)
			}
		}
	}
	return &stmodel.Func{
		Args:    args,
		Results: results,
	}
}

func (p *Parser) parseIdent(ide *ast.Ident) *stmodel.Type {
	if ide.Obj == nil {
		return &stmodel.Type{
			TypeName:   ide.Name,
			Underlying: stmodel.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
		}
	}
	if ide.Obj.Decl != nil {
		if spec, ok := ide.Obj.Decl.(*ast.TypeSpec); ok {
			switch typ := spec.Type.(type) {
			case *ast.Ident:
				// like stringLike type
				return &stmodel.Type{
					TypeName:   ide.Name,
					Underlying: stmodel.UnderlyingType(p.Pkg.TypesInfo.TypeOf(typ).String()),
				}
			case *ast.StructType:
				return &stmodel.Type{
					TypeName:   ide.Name,
					Underlying: stmodel.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
				}
			}
		}
	}
	return &stmodel.Type{
		TypeName:   ide.Obj.Name,
		Underlying: stmodel.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
	}
}

func (p *Parser) parseTag(tag *ast.BasicLit) []*stmodel.Tag {
	if tag == nil {
		return nil
	}

	var out []*stmodel.Tag
	tags := strings.Split(strings.Trim(tag.Value, "`"), " ")
	for _, t := range tags {
		kv := strings.Split(t, ":")
		if len(kv) != 2 {
			continue
		}

		v := strings.Trim(kv[1], `"`)
		tag := &stmodel.Tag{
			Key:      kv[0],
			Values:   strings.Split(v, ","),
			RawValue: v,
		}
		out = append(out, tag)
	}
	return out
}

// func (p *Parser) samePackage(pkgID string) bool {
// 	return p.Pkg.ID == pkgID
// }
