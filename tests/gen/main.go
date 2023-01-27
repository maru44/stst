package main

import (
	"fmt"

	"github.com/maru44/stst"
	"github.com/maru44/stst/model"
	"golang.org/x/tools/go/packages"
)

func main() {
	ps, _ := loadPackages("github.com/maru44/stst/tests/data")

	var schemas []*model.Schema

	for _, pk := range ps {
		p := stst.NewParser(pk)
		s := p.Parse()
		schemas = append(schemas, s...)
	}
}

func loadPackages(ps ...string) ([]*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, ps...)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}
	return pkgs, err
}
