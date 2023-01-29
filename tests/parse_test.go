package tests_test

import (
	"testing"

	"github.com/maru44/stst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	ps, err := loadPackages("github.com/maru44/stst/tests/data")
	require.NoError(t, err)
	require.Len(t, ps, 1)

	var schemas []*stst.Schema
	for _, pk := range ps {
		p := stst.NewParser(pk)
		s := p.Parse()
		schemas = append(schemas, s...)
	}

	want := []*stst.Schema{
		{
			Name: "SampleString",
			Type: &stst.Type{
				Underlying: "string",
				TypeName:   "string",
			},
			Comment: []string{"// comment"},
		},
		{
			Name: "MapSimple",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.MapSimple",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.MapSimple",
				TypeName:    "MapSimple",
			},
			Map: &stst.Map{
				Key: &stst.Field{
					Name: "any",
					Type: &stst.Type{
						Underlying: "any",
						TypeName:   "any",
					},
				},
				Value: &stst.Field{
					Name:             "",
					IsUntitledStruct: true,
					Schema: &stst.Schema{
						Fields: []*stst.Field{
							{
								Name: "a",
								Type: &stst.Type{
									Underlying: "string",
									TypeName:   "SampleString",
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "MapS",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.MapS",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.MapS",
				TypeName:    "MapS",
			},
			Map: &stst.Map{
				Key: &stst.Field{
					Name:                "",
					IsUntitledInterface: true,
					Schema: &stst.Schema{
						Name: "",
						Fields: []*stst.Field{
							{
								Name: "AAA",
								Func: &stst.Func{},
							},
						},
					},
				},
				Value: &stst.Field{
					Name:             "",
					IsUntitledStruct: true,
				},
			},
		},
		{
			Name: "SamplePrefixMap",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.SamplePrefixMap",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.SamplePrefixMap",
				TypeName:    "SamplePrefixMap",
			},
			Map: &stst.Map{
				Key: &stst.Field{
					Name: "string",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				Value: &stst.Field{
					Name: "any",
					Type: &stst.Type{
						TypeName:   "any",
						Underlying: "any",
					},
					TypePrefixes: []stst.TypePrefix{
						"[]", "*", "[]", "[]", "*", "[5]",
					},
				},
			},
			Comment: []string{"// pref comment"},
		},
		{
			Name: "Person",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.Person",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.Person",
				TypeName:    "Person",
			},
			Fields: []*stst.Field{
				{
					Name: "Name",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"name",
							},
							RawValue: "name",
						},
					},
					Comment: []string{
						"// comment",
					},
				},
				{
					Name: "Age",
					Type: &stst.Type{
						Underlying: "int",
						TypeName:   "int",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"age",
							},
							RawValue: "age",
						},
					},
				},
				{
					Name: "Sex",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"-",
							},
							RawValue: "-",
						},
					},
				},
				{
					Name: "Hobby",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"hobby",
								"nullable",
							},
							RawValue: "hobby,nullable",
						},
					},
				},
				{
					Name: "Good",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.Good",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.Good",
						TypeName:    "Good",
					},
					Comment: []string{
						"// no name // no name",
					},
				},
			},
		},
		{
			Name: "Animal",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.Animal",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.Animal",
				TypeName:    "Animal",
			},
			Fields: []*stst.Field{
				{
					Name: "ID",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"id",
							},
							RawValue: "id",
						},
					},
				},
				{
					Name: "Kind",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"kind",
							},
							RawValue: "kind",
						},
					},
				},
				{
					Name: "Goods",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.Good",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.Good",
						TypeName:    "Good",
					},
					TypePrefixes: []stst.TypePrefix{"[]", "*"},
				},
				{
					Name: "GoodPtr",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.Good",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.Good",
						TypeName:    "Good",
					},
					TypePrefixes: []stst.TypePrefix{"*"},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"good_ptr",
							},
							RawValue: "good_ptr",
						},
					},
				},
				{
					Name: "GoodStr",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.Good",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.Good",
						TypeName:    "Good",
					},
					Tags: []*stst.Tag{
						{
							Key: "bigquery",
							Values: []string{
								"good_str",
							},
							RawValue: "good_str",
						},
						{
							Key: "aaa",
							Values: []string{
								"bbb",
								"nullable",
							},
							RawValue: "bbb,nullable",
						},
					},
				},
				{
					Name: "strPtr",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					TypePrefixes: []stst.TypePrefix{"*"},
				},
				{
					Name: "tim",
					Type: &stst.Type{
						Underlying:  "time.Time",
						PkgID:       "time",
						PkgPlusName: "time.Time",
						TypeName:    "Time",
					},
				},
				{
					Name: "timPtr",
					Type: &stst.Type{
						Underlying:  "time.Time",
						PkgID:       "time",
						PkgPlusName: "time.Time",
						TypeName:    "Time",
					},
					TypePrefixes: []stst.TypePrefix{"*"},
				},
				{
					Name: "strs",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					TypePrefixes: []stst.TypePrefix{"[]"},
				},
				{
					Name:         "funcs",
					TypePrefixes: []stst.TypePrefix{"[]"},
					Func: &stst.Func{
						Args: []*stst.Field{
							{
								Name: "v",
								Type: &stst.Type{
									Underlying: "any",
									TypeName:   "any",
								},
							},
							{
								Name:         "ints",
								TypePrefixes: []stst.TypePrefix{"[]"},
								Type: &stst.Type{
									Underlying: "int",
									TypeName:   "int",
								},
							},
							{
								Name: "ptrInts",
								Type: &stst.Type{
									Underlying: "int",
									TypeName:   "int",
								},
								TypePrefixes: []stst.TypePrefix{"*", "[]", "*"},
							},
						},
						Results: []*stst.Field{
							{
								Name: "bool",
								Type: &stst.Type{
									Underlying:  "bool",
									PkgID:       "",
									PkgPlusName: "",
									TypeName:    "bool",
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "Gene",
			Fields: []*stst.Field{
				{
					Name: "One",
					Type: &stst.Type{
						Underlying: "T",
						TypeName:   "T",
					},
				},
			},
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.Gene[T any]",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.Gene[T any]",
				TypeName:    "Gene",
			},
		},
		{
			Name: "Good",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.Good",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.Good",
				TypeName:    "Good",
			},
			Fields: []*stst.Field{
				{
					Name: "Name",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				{
					Name: "Sample",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "SampleString",
					},
				},
				{
					Name: "SamplePtr",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "SampleString",
					},
					TypePrefixes: []stst.TypePrefix{"*"},
				},
			},
		},
		{
			Name: "withIntf",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.withIntf",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.withIntf",
				TypeName:    "withIntf",
			},
			Fields: []*stst.Field{
				{
					Name: "error",
					Type: &stst.Type{
						Underlying: "error",
						TypeName:   "error",
					},
				},
				{
					Name: "str",
					Type: &stst.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				{
					Name: "Intf",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data/aaa.Intf",
						PkgID:       "github.com/maru44/stst/tests/data/aaa",
						PkgPlusName: "aaa.Intf",
						TypeName:    "Intf",
					},
				},
				{
					Name: "Good",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.Good",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.Good",
						TypeName:    "Good",
					},
					TypePrefixes: []stst.TypePrefix{"*"},
				},
				{
					Name: "intef",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data/aaa.Intf",
						PkgID:       "github.com/maru44/stst/tests/data/aaa",
						PkgPlusName: "aaa.Intf",
						TypeName:    "Intf",
					},
				},
				{
					Name: "IntSample",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data/aaa.IntSample",
						PkgID:       "github.com/maru44/stst/tests/data/aaa",
						PkgPlusName: "aaa.IntSample",
						TypeName:    "IntSample",
					},
				},
				{
					Name: "intf",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.intf",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.intf",
						TypeName:    "intf",
					},
				},
				{
					Name: "notEmbededIntf",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.intf",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.intf",
						TypeName:    "intf",
					},
				},
				{
					Name: "fn",
					Func: &stst.Func{
						Args: []*stst.Field{
							{
								Name: "v",
								Type: &stst.Type{
									Underlying: "any",
									TypeName:   "any",
								},
							},
						},
						Results: []*stst.Field{
							{
								Name: "error",
								Type: &stst.Type{
									Underlying: "error",
									TypeName:   "error",
								},
							},
						},
					},
				},
				{
					Name: "ma",
					Map: &stst.Map{
						Key: &stst.Field{
							Name: "string",
							Type: &stst.Type{
								Underlying: "string",
								TypeName:   "string",
							},
						},
						Value: &stst.Field{
							Name: "any",
							Type: &stst.Type{
								Underlying: "any",
								TypeName:   "any",
							},
						},
					},
				},
			},
		},
		{
			Name: "intf",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.intf",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.intf",
				TypeName:    "intf",
			},
			Fields: []*stst.Field{
				{
					Name: "AAA",
					Func: &stst.Func{
						Args: []*stst.Field{
							{
								Name: "in",
								Type: &stst.Type{
									Underlying: "string",
									TypeName:   "string",
								},
							},
							{
								Name: "good",
								Type: &stst.Type{
									Underlying:  "github.com/maru44/stst/tests/data.Good",
									PkgID:       "github.com/maru44/stst/tests/data",
									PkgPlusName: "data.Good",
									TypeName:    "Good",
								},
							},
							{
								Name: "sample",
								Type: &stst.Type{
									Underlying:  "github.com/maru44/stst/tests/data/aaa.Sample",
									PkgID:       "github.com/maru44/stst/tests/data/aaa",
									PkgPlusName: "aaa.Sample",
									TypeName:    "Sample",
								},
							},
						},
						Results: []*stst.Field{
							{
								Name: "string",
								Type: &stst.Type{
									Underlying: "string",
									TypeName:   "string",
								},
							},
							{
								Name: "error",
								Type: &stst.Type{
									Underlying: "error",
									TypeName:   "error",
								},
							},
						},
					},
				},
				{
					Name: "BBB",
					Func: &stst.Func{},
				},
				{
					Name: "Intf",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data/aaa.Intf",
						PkgID:       "github.com/maru44/stst/tests/data/aaa",
						PkgPlusName: "aaa.Intf",
						TypeName:    "Intf",
					},
				},
				{
					Name: "childIntf",
					Type: &stst.Type{
						Underlying:  "github.com/maru44/stst/tests/data.childIntf",
						TypeName:    "childIntf",
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.childIntf",
					},
				},
			},
			IsInterface: true,
		},
		{
			Name: "childIntf",
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.childIntf",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.childIntf",
				TypeName:    "childIntf",
			},
			Fields: []*stst.Field{
				{
					Name: "CCC",
					Func: &stst.Func{},
				},
			},
			IsInterface: true,
		},
	}

	assert.Equal(t, want, schemas)
}
