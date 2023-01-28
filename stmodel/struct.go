package stmodel

import "strings"

type (
	UnderlyingType string
	TypePrefix     string

	// Schema is information for defined as type
	Schema struct {
		Name         string
		Fields       []*Field
		Type         *Type
		Func         *Func
		Map          *Map
		IsInterface  bool
		TypePrefixes []TypePrefix
	}

	// Func has information of args and results
	Func struct {
		Args    []*Field
		Results []*Field
	}

	Field struct {
		Name                string
		Type                *Type
		IsUntitledStruct    bool
		IsUntitledInterface bool
		Tags                []*Tag
		Comment             []string
		Func                *Func
		Map                 *Map
		TypePrefixes        []TypePrefix
		// Schema is only for untitled struct or untitled interface
		Schema *Schema
	}

	// Type is type information.
	Type struct {
		Underlying UnderlyingType // xxx/yy.ZZZ
		// Package is package id.
		Package    string // xxx/yy
		PkPlusName string // yy.ZZZ
		TypeName   string // ZZZ
	}

	Map struct {
		Key   *Field
		Value *Field
	}

	Tag struct {
		Key      string
		Values   []string
		RawValue string
	}
)

const (
	TypePrefixPtr   = TypePrefix("*")
	TypePrefixSlice = TypePrefix("[]")
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

func (f *Field) IsFunc() bool {
	return f.Func != nil
}

func (s *Schema) IsMap() bool {
	return s.Map != nil
}

func (f *Field) IsMap() bool {
	return f.Map != nil
}
