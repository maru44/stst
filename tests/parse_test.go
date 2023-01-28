package tests_test

import (
	"testing"

	"github.com/maru44/stst"
	"github.com/maru44/stst/stmodel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	ps, err := loadPackages("github.com/maru44/stst/tests/data")
	require.NoError(t, err)
	require.Len(t, ps, 1)

	var schemas []*stmodel.Schema
	for _, pk := range ps {
		p := stst.NewParser(pk)
		s := p.Parse()
		schemas = append(schemas, s...)
	}

	want := []*stmodel.Schema{
		{
			Name: "SampleString",
			Type: &stmodel.Type{
				Underlying: "string",
				TypeName:   "string",
			},
		},
		{
			Name: "MapSimple",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.MapSimple",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.MapSimple",
				TypeName:   "MapSimple",
			},
			Map: &stmodel.Map{
				Key: &stmodel.Field{
					Name: "any",
					Type: &stmodel.Type{
						Underlying: "any",
						TypeName:   "any",
					},
				},
				Value: &stmodel.Field{
					Name:             "",
					IsUntitledStruct: true,
					Schema: &stmodel.Schema{
						Fields: []*stmodel.Field{
							{
								Name: "a",
								Type: &stmodel.Type{
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
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.MapS",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.MapS",
				TypeName:   "MapS",
			},
			Map: &stmodel.Map{
				Key: &stmodel.Field{
					Name:                "",
					IsUntitledInterface: true,
					Schema: &stmodel.Schema{
						Name: "",
						Fields: []*stmodel.Field{
							{
								Name: "AAA",
								Func: &stmodel.Func{},
							},
						},
					},
				},
				Value: &stmodel.Field{
					Name:             "",
					IsUntitledStruct: true,
				},
			},
		},
		{
			Name: "SamplePrefixMap",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.SamplePrefixMap",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.SamplePrefixMap",
				TypeName:   "SamplePrefixMap",
			},
			Map: &stmodel.Map{
				Key: &stmodel.Field{
					Name: "string",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				Value: &stmodel.Field{
					Name: "any",
					Type: &stmodel.Type{
						TypeName:   "any",
						Underlying: "any",
					},
					TypePrefixes: []stmodel.TypePrefix{
						"[]", "*", "[]", "[]", "*", "[]",
					},
				},
			},
		},
		{
			Name: "Person",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.Person",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Person",
				TypeName:   "Person",
			},
			Fields: []*stmodel.Field{
				{
					Name: "Name",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "int",
						TypeName:   "int",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					Comment: []string{
						"// no name // no name",
					},
				},
			},
		},
		{
			Name: "Animal",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.Animal",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Animal",
				TypeName:   "Animal",
			},
			Fields: []*stmodel.Field{
				{
					Name: "ID",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					TypePrefixes: []stmodel.TypePrefix{"[]", "*"},
				},
				{
					Name: "GoodPtr",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					TypePrefixes: []stmodel.TypePrefix{"*"},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					Tags: []*stmodel.Tag{
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
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					TypePrefixes: []stmodel.TypePrefix{"*"},
				},
				{
					Name: "tim",
					Type: &stmodel.Type{
						Underlying: "time.Time",
						Package:    "time",
						PkPlusName: "time.Time",
						TypeName:   "Time",
					},
				},
				{
					Name: "timPtr",
					Type: &stmodel.Type{
						Underlying: "time.Time",
						Package:    "time",
						PkPlusName: "time.Time",
						TypeName:   "Time",
					},
					TypePrefixes: []stmodel.TypePrefix{"*"},
				},
				{
					Name: "strs",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					TypePrefixes: []stmodel.TypePrefix{"[]"},
				},
				{
					Name:         "funcs",
					TypePrefixes: []stmodel.TypePrefix{"[]"},
					Func: &stmodel.Func{
						Args: []*stmodel.Field{
							{
								Name: "v",
								Type: &stmodel.Type{
									Underlying: "any",
									TypeName:   "any",
								},
							},
							{
								Name:         "ints",
								TypePrefixes: []stmodel.TypePrefix{"[]"},
								Type: &stmodel.Type{
									Underlying: "int",
									TypeName:   "int",
								},
							},
							{
								Name: "ptrInts",
								Type: &stmodel.Type{
									Underlying: "int",
									TypeName:   "int",
								},
								TypePrefixes: []stmodel.TypePrefix{"*", "[]", "*"},
							},
						},
						Results: []*stmodel.Field{
							{
								Name: "bool",
								Type: &stmodel.Type{
									Underlying: "bool",
									Package:    "",
									PkPlusName: "",
									TypeName:   "bool",
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "Gene",
			Fields: []*stmodel.Field{
				{
					Name: "One",
					Type: &stmodel.Type{
						Underlying: "T",
						TypeName:   "T",
					},
				},
			},
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.Gene[T any]",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Gene[T any]",
				TypeName:   "Gene",
			},
		},
		{
			Name: "Good",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.Good",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Good",
				TypeName:   "Good",
			},
			Fields: []*stmodel.Field{
				{
					Name: "Name",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				{
					Name: "Sample",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "SampleString",
					},
				},
				{
					Name: "SamplePtr",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "SampleString",
					},
					TypePrefixes: []stmodel.TypePrefix{"*"},
				},
			},
		},
		{
			Name: "withIntf",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.withIntf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.withIntf",
				TypeName:   "withIntf",
			},
			Fields: []*stmodel.Field{
				{
					Name: "error",
					Type: &stmodel.Type{
						Underlying: "error",
						TypeName:   "error",
					},
				},
				{
					Name: "str",
					Type: &stmodel.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				{
					Name: "Intf",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.Intf",
						TypeName:   "Intf",
					},
				},
				{
					Name: "Good",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					TypePrefixes: []stmodel.TypePrefix{"*"},
				},
				{
					Name: "intef",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.Intf",
						TypeName:   "Intf",
					},
				},
				{
					Name: "IntSample",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.IntSample",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.IntSample",
						TypeName:   "IntSample",
					},
				},
				{
					Name: "intf",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.intf",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.intf",
						TypeName:   "intf",
					},
				},
				{
					Name: "notEmbededIntf",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.intf",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.intf",
						TypeName:   "intf",
					},
				},
				{
					Name: "fn",
					Func: &stmodel.Func{
						Args: []*stmodel.Field{
							{
								Name: "v",
								Type: &stmodel.Type{
									Underlying: "any",
									TypeName:   "any",
								},
							},
						},
						Results: []*stmodel.Field{
							{
								Name: "error",
								Type: &stmodel.Type{
									Underlying: "error",
									TypeName:   "error",
								},
							},
						},
					},
				},
				{
					Name: "ma",
					Map: &stmodel.Map{
						Key: &stmodel.Field{
							Name: "string",
							Type: &stmodel.Type{
								Underlying: "string",
								TypeName:   "string",
							},
						},
						Value: &stmodel.Field{
							Name: "any",
							Type: &stmodel.Type{
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
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.intf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.intf",
				TypeName:   "intf",
			},
			Fields: []*stmodel.Field{
				{
					Name: "AAA",
					Func: &stmodel.Func{
						Args: []*stmodel.Field{
							{
								Name: "in",
								Type: &stmodel.Type{
									Underlying: "string",
									TypeName:   "string",
								},
							},
							{
								Name: "good",
								Type: &stmodel.Type{
									Underlying: "github.com/maru44/stst/tests/data.Good",
									Package:    "github.com/maru44/stst/tests/data",
									PkPlusName: "data.Good",
									TypeName:   "Good",
								},
							},
							{
								Name: "sample",
								Type: &stmodel.Type{
									Underlying: "github.com/maru44/stst/tests/data/aaa.Sample",
									Package:    "github.com/maru44/stst/tests/data/aaa",
									PkPlusName: "aaa.Sample",
									TypeName:   "Sample",
								},
							},
						},
						Results: []*stmodel.Field{
							{
								Name: "string",
								Type: &stmodel.Type{
									Underlying: "string",
									TypeName:   "string",
								},
							},
							{
								Name: "error",
								Type: &stmodel.Type{
									Underlying: "error",
									TypeName:   "error",
								},
							},
						},
					},
				},
				{
					Name: "BBB",
					Func: &stmodel.Func{},
				},
				{
					Name: "Intf",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.Intf",
						TypeName:   "Intf",
					},
				},
				{
					Name: "childIntf",
					Type: &stmodel.Type{
						Underlying: "github.com/maru44/stst/tests/data.childIntf",
						TypeName:   "childIntf",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.childIntf",
					},
				},
			},
			IsInterface: true,
		},
		{
			Name: "childIntf",
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.childIntf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.childIntf",
				TypeName:   "childIntf",
			},
			Fields: []*stmodel.Field{
				{
					Name: "CCC",
					Func: &stmodel.Func{},
				},
			},
			IsInterface: true,
		},
	}

	assert.Equal(t, want, schemas)
}
