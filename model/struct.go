package model

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
