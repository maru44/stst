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

func (p *Parser) Parse() []*model.Schema {
	var schemas []*model.Schema
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

func (p *Parser) ParseFile(f *ast.File) []*model.Schema {
	var schemas []*model.Schema
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

func (p *Parser) parseTypeSpec(spec *ast.TypeSpec) *model.Schema {
	sc := &model.Schema{
		Name: spec.Name.Name,
	}

	switch typ := spec.Type.(type) {
	case *ast.StructType:
		sc.Type = p.parseIdent(spec.Name)
		sc.Type.SetPackage()
		p.parseStruct(typ, sc)
	case *ast.Ident:
		sc.Type = p.parseIdent(typ)
		sc.Type.SetPackage()
	case *ast.InterfaceType:
		sc.Type = p.parseIdent(spec.Name)
		sc.Type.SetPackage()
		sc.IsInterface = true
		p.parseInterface(typ, sc)
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

func (p *Parser) parseInterface(in *ast.InterfaceType, sc *model.Schema) {
	for _, m := range in.Methods.List {
		ff, ok := p.parseField(m)
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
		out.Type = p.parseIdent(typ)

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
				out.Type = p.parseIdent(x)
			case *ast.SelectorExpr:
				out.Type = &model.Type{
					TypeName:   x.Sel.Name,
					Underlying: model.UnderlyingType(strings.TrimLeft(p.Pkg.TypesInfo.TypeOf(typ).String(), "*")),
				}
			}
		case *ast.Ident:
			out.Type = p.parseIdent(elt)
		case *ast.FuncType:
			out.Func = p.parseFunc(elt)
		}
	case *ast.StarExpr:
		out.IsPtr = true
		switch x := typ.X.(type) {
		case *ast.Ident:
			// set name to embeded pointer
			if name == "" {
				name = x.Name
			}
			out.Type = p.parseIdent(x)
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
	case *ast.FuncType:
		out.Func = p.parseFunc(typ)
	}

	out.Name = name
	if out.Type != nil {
		out.Type.SetPackage()
	}
	return out, true
}

func (p *Parser) parseFunc(fn *ast.FuncType) *model.Func {
	var args, results []*model.Field
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
	return &model.Func{
		Args:    args,
		Results: results,
	}
}

func (p *Parser) parseIdent(ide *ast.Ident) *model.Type {
	if ide.Obj == nil {
		return &model.Type{
			TypeName:   ide.Name,
			Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
		}
	}
	if ide.Obj.Decl != nil {
		if spec, ok := ide.Obj.Decl.(*ast.TypeSpec); ok {
			switch typ := spec.Type.(type) {
			case *ast.Ident:
				// enum like
				return &model.Type{
					TypeName:   ide.Name,
					Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(typ).String()),
				}
			case *ast.StructType:
				// sc := p.parseTypeSpec(spec)
				return &model.Type{
					TypeName:   ide.Name,
					Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
				}
			}
		}
	}
	return &model.Type{
		TypeName:   ide.Obj.Name,
		Underlying: model.UnderlyingType(p.Pkg.TypesInfo.TypeOf(ide).String()),
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

func (p *Parser) samePackage(pkgID string) bool {
	return p.Pkg.ID == pkgID
}
