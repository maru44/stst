package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/maru44/stst"
	"github.com/maru44/stst/model"
	"golang.org/x/tools/go/packages"
)

func main() {
	ps, _ := loadPackages()

	var schemas []*model.Schema

	for _, pk := range ps {
		p := stst.NewParser(pk)
		for _, f := range p.Pkg.Syntax {
			structs := p.ParseFile(f)
			schemas = append(schemas, structs...)
		}
	}

	pp.Println("result: ")
	pp.Println(schemas)
}

func loadPackages() ([]*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}
	return pkgs, err
}
