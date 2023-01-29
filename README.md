# stst

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/maru44/stst/blob/master/LICENSE)
![ActionsCI](https://github.com/maru44/stst/workflows/lint-test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/maru44/stst)](https://goreportcard.com/report/github.com/maru44/stst)

The `stst` is package to get information defined as type in go package by static analysis.

## How to Use

**Install**

```shell
go install github.com/maru44/stst@latest
```

1. Load package
2. Create `stst.Parser` by `stst.NewParser` with loaded package
3. Execute `Parse` or `ParseFile` method of created `stst.Parser`

sample code:

```go
package main

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/maru44/stst"
	"golang.org/x/tools/go/packages"
)

func main() {
	ps, _ := loadPackages("github.com/maru44/stst/tests/data/aaa")

	var schemas []*stst.Schema
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
[]*stst.Schema{
  &stst.Schema{
    Name:   "Intf",
    Fields: []*stst.Field{
      &stst.Field{
        Name:                "Hello",
        Func:                &stst.Func{
          Args:    []*stst.Field{},
          Results: []*stst.Field{
            &stst.Field{
              Name: "string",
              Type: &stst.Type{
                Underlying: "string",
                TypeName:   "string",
              },
            },
          },
        },
      },
    },
    Type: &stst.Type{
      Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
      PkgID:    "github.com/maru44/stst/tests/data/aaa",
      PkgPlusName: "aaa.Intf",
      TypeName:   "Intf",
    },
    IsInterface:  true,
  },
  &stst.Schema{
    Name:   "IntSample",
    Type:   &stst.Type{
      Underlying: "int",
      TypeName:   "int",
    },
  },
  &stst.Schema{
    Name:   "Sample",
    Fields: []*stst.Field{
      &stst.Field{
        Name: "Str",
        Type: &stst.Type{
          Underlying: "string",
          TypeName:   "string",
        },
        Tags:                []*stst.Tag{
          &stst.Tag{
            Key:    "tag0",
            Values: []string{
              "xxx",
            },
            RawValue: "xxx",
          },
          &stst.Tag{
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
    Type: &stst.Type{
      Underlying: "github.com/maru44/stst/tests/data/aaa.Sample",
      PkgID:    "github.com/maru44/stst/tests/data/aaa",
      PkgPlusName: "aaa.Sample",
      TypeName:   "Sample",
    },
  },
  &stst.Schema{
    Name:   "prefixes",
    Type:   &stst.Type{
      Underlying: "github.com/maru44/stst/tests/data/aaa.prefixes",
      PkgID:    "github.com/maru44/stst/tests/data/aaa",
      PkgPlusName: "aaa.prefixes",
      TypeName:   "prefixes",
    },
    TypePrefixes: []stst.TypePrefix{
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
