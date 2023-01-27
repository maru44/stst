package model

import "strings"

type (
	UnderlyingType string

	Schema struct {
		Name        string
		Fields      []*Field
		Type        *Type
		Func        *Func
		IsInterface bool
	}

	Func struct {
		Args    []*Field
		Results []*Field
	}

	Field struct {
		Name              string
		Type              *Type
		IsSlice           bool
		IsPtr             bool
		Tags              []*Tag
		Schema            *Schema
		Comment           []string
		IsInterface       bool
		PossibleInterface bool
	}

	Type struct {
		Underlying UnderlyingType // github.com/xxx/yy.ZZZ
		Package    string         // github.com/xxx/yy
		PkPlusName string         // yy.ZZZ
		TypeName   string         // ZZZ
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
		withoutType := strings.Split(string(u), ".")
		if len(withoutType) != 1 {
			pkPlusName = string(u)
			pack = withoutType[0]
		}
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

func (s *Schema) IsFunc() bool {
	return s.Func != nil
}
