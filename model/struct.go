package model

import "strings"

type (
	UnderlyingType string

	Schema struct {
		Name   string
		Fields []*Field
		Type   *Type
	}

	Field struct {
		Name    string
		Type    *Type
		IsSlice bool
		IsPtr   bool
		Tags    []*Tag
		Schema  *Schema
		Comment []string
	}

	Type struct {
		Underlying UnderlyingType // github.com/xxx/yy.ZZZ
		Package    string         // github.com/xxx/yy
		PkPlusName string         // yy.ZZZ
		TypeName   string         // ZZZ
		// Enum ? []T
	}

	Tag struct {
		Key      string
		Values   []string
		RawValue string
	}
)

func (u UnderlyingType) pk() (pack string, pkPlusName string) {
	arr := strings.Split(string(u), "/")
	if len(arr) == 1 {
		return
	}
	pkPlusName = arr[len(arr)-1]
	withouType := strings.Split(pkPlusName, ".")
	pack = strings.Join(arr[0:len(arr)-1], "/") + "/" + withouType[0]
	return
}

func (t *Type) SetPackage() {
	t.Package, t.PkPlusName = t.Underlying.pk()
}
