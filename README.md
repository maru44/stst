# stst

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/maru44/stst/blob/master/LICENSE)
![ActionsCI](https://github.com/maru44/stst/workflows/lint-test/badge.svg)

The `stst` is package to get information defined as type in go package by static analysis.

## How to Use

sample code:

```go
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/maru44/stst"
	"github.com/maru44/stst/stmodel"
	"golang.org/x/tools/go/packages"
)

func main() {
	ps, _ := loadPackages("github.com/maru44/stst/tests/data/aaa")

	var schemas []*stmodel.Schema
	for _, pk := range ps {
		p := stst.NewParser(pk)
		s := p.Parse()
		schemas = append(schemas, s...)
	}
	pp.Println(schemas)
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

```

result:

```
[]*stmodel.Schema{
  &stmodel.Schema{
    Name:   "Intf",
    Fields: []*stmodel.Field{
      &stmodel.Field{
        Name:                "Hello",
        Func:                &stmodel.Func{
          Args:    []*stmodel.Field{},
          Results: []*stmodel.Field{
            &stmodel.Field{
              Name: "string",
              Type: &stmodel.Type{
                Underlying: "string",
                TypeName:   "string",
              },
            },
          },
        },
      },
    },
    Type: &stmodel.Type{
      Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
      Package:    "github.com/maru44/stst/tests/data/aaa",
      PkPlusName: "aaa.Intf",
      TypeName:   "Intf",
    },
    IsInterface:  true,
  },
  &stmodel.Schema{
    Name:   "IntSample",
    Type:   &stmodel.Type{
      Underlying: "int",
      TypeName:   "int",
    },
  },
  &stmodel.Schema{
    Name:   "Sample",
    Fields: []*stmodel.Field{
      &stmodel.Field{
        Name: "Str",
        Type: &stmodel.Type{
          Underlying: "string",
          TypeName:   "string",
        },
        Tags:                []*stmodel.Tag{
          &stmodel.Tag{
            Key:    "tag0",
            Values: []string{
              "xxx",
            },
            RawValue: "xxx",
          },
          &stmodel.Tag{
            Key:    "tag1",
            Values: []string{
              "yyy",
              "zzz",
            },
            RawValue: "yyy,zzz",
          },
        },
        Comment: []string{
          "// comment",
        },
      },
    },
    Type: &stmodel.Type{
      Underlying: "github.com/maru44/stst/tests/data/aaa.Sample",
      Package:    "github.com/maru44/stst/tests/data/aaa",
      PkPlusName: "aaa.Sample",
      TypeName:   "Sample",
    },
  },
  &stmodel.Schema{
    Name:   "prefixes",
    Type:   &stmodel.Type{
      Underlying: "github.com/maru44/stst/tests/data/aaa.prefixes",
      Package:    "github.com/maru44/stst/tests/data/aaa",
      PkPlusName: "aaa.prefixes",
      TypeName:   "prefixes",
    },
    TypePrefixes: []stmodel.TypePrefix{
      "[]",
      "*",
      "[]",
      "[]",
      "[]",
      "*",
    },
  },
}
```
