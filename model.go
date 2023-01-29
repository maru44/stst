package stst

import (
	"regexp"
	"strconv"
	"strings"
)

type (
	UnderlyingType string
	TypePrefix     string
	TypePrefixKind string

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
		// PkgID is package id.
		PkgID       string // xxx/yy
		PkgPlusName string // yy.ZZZ
		TypeName    string // ZZZ
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

	TypePrefixKindPtr     = TypePrefixKind("pointer")
	TypePrefixKindSlice   = TypePrefixKind("slice")
	TypePrefixKindArray   = TypePrefixKind("array")
	TypePrefixKindUnknown = TypePrefixKind("unknown")
)

var intReg = regexp.MustCompile("[0-9]+")

// SetPackage is method to set PkgID and PkgPlusName by UnderlyingType.
func (t *Type) SetPackage() {
	t.PkgID, t.PkgPlusName = t.Underlying.pk()
}

// IsFunc returns whether the Schema is function or not.
func (s *Schema) IsFunc() bool {
	return s.Func != nil
}

// IsFunc returns whether the Schema is map or not.
func (s *Schema) IsMap() bool {
	return s.Map != nil
}

// IsFunc returns whether the Field is function or not.
func (f *Field) IsFunc() bool {
	return f.Func != nil
}

// IsFunc returns whether the Field is map or not.
func (f *Field) IsMap() bool {
	return f.Map != nil
}

func (t TypePrefix) Kind() TypePrefixKind {
	if t == TypePrefixPtr {
		return TypePrefixKindPtr
	}
	if t == TypePrefixSlice {
		return TypePrefixKindSlice
	}
	if _, ok := t.ArrayLength(); ok {
		return TypePrefixKindArray
	}
	return TypePrefixKindUnknown
}

func (t TypePrefix) ArrayLength() (int, bool) {
	b := intReg.Find([]byte(t))
	if len(b) == 0 {
		return 0, false
	}
	out, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, false
	}
	return out, true
}

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
