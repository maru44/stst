//go:build exclude
// +build exclude

package main

import (
	"fmt"

	"github.com/maru44/stst"
	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	ps, err := packages.Load(cfg, "github.com/maru44/stst/tests/data/aaa")
	if err != nil {
		panic(err)
	}

	for _, pk := range ps {
		p := stst.NewParser(pk)
		s := p.Parse()
		for _, ss := range s {
			fmt.Println(ss)
		}
	}
}
