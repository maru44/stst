package tests_test

import (
	"testing"

	"github.com/maru44/stst"
	"github.com/maru44/stst/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	ps, err := loadPackages("github.com/maru44/stst/tests/data")
	require.NoError(t, err)
	require.Len(t, ps, 1)

	var schemas []*model.Schema
	for _, pk := range ps {
		p := stst.NewParser(pk)
		s := p.Parse()
		schemas = append(schemas, s...)
	}

	want := []*model.Schema{
		{
			Name: "SampleString",
			Type: &model.Type{
				Underlying: "string",
				TypeName:   "string",
			},
		},
		{
			Name: "Person",
			Type: &model.Type{
				Underlying: "github.com/maru44/stst/tests/data.Person",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Person",
				TypeName:   "Person",
			},
			Fields: []*model.Field{
				{
					Name: "Name",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "int",
						TypeName:   "int",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
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
			Type: &model.Type{
				Underlying: "github.com/maru44/stst/tests/data.Animal",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Animal",
				TypeName:   "Animal",
			},
			Fields: []*model.Field{
				{
					Name: "ID",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					IsSlice: true,
					IsPtr:   true,
				},
				{
					Name: "GoodPtr",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					IsPtr: true,
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					Tags: []*model.Tag{
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
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					IsPtr: true,
				},
				{
					Name: "tim",
					Type: &model.Type{
						Underlying: "time.Time",
						Package:    "time",
						PkPlusName: "time.Time",
						TypeName:   "Time",
					},
				},
				{
					Name: "timPtr",
					Type: &model.Type{
						Underlying: "time.Time",
						Package:    "time",
						PkPlusName: "time.Time",
						TypeName:   "Time",
					},
					IsPtr: true,
				},
				{
					Name: "strs",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
					IsSlice: true,
				},
				{
					Name:    "funcs",
					IsSlice: true,
					Func: &model.Func{
						Args: []*model.Field{
							{
								Name: "v",
								Type: &model.Type{
									Underlying: "any",
									TypeName:   "any",
								},
							},
							{
								Name:    "ints",
								IsSlice: true,
								Type: &model.Type{
									Underlying: "int",
									TypeName:   "int",
								},
							},
						},
						Results: []*model.Field{
							{
								Name: "bool",
								Type: &model.Type{
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
			Name: "Good",
			Type: &model.Type{
				Underlying: "github.com/maru44/stst/tests/data.Good",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.Good",
				TypeName:   "Good",
			},
			Fields: []*model.Field{
				{
					Name: "Name",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				{
					Name: "Sample",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "SampleString",
					},
				},
				{
					Name: "SamplePtr",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "SampleString",
					},
					IsPtr: true,
				},
			},
		},
		{
			Name: "withIntf",
			Type: &model.Type{
				Underlying: "github.com/maru44/stst/tests/data.withIntf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.withIntf",
				TypeName:   "withIntf",
			},
			Fields: []*model.Field{
				{
					Name: "error",
					Type: &model.Type{
						Underlying: "error",
						TypeName:   "error",
					},
				},
				{
					Name: "str",
					Type: &model.Type{
						Underlying: "string",
						TypeName:   "string",
					},
				},
				{
					Name: "Intf",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.Intf",
						TypeName:   "Intf",
					},
				},
				{
					Name: "Good",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data.Good",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.Good",
						TypeName:   "Good",
					},
					IsPtr: true,
				},
				{
					Name: "intef",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.Intf",
						TypeName:   "Intf",
					},
				},
				{
					Name: "IntSample",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.IntSample",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.IntSample",
						TypeName:   "IntSample",
					},
				},
				{
					Name: "intf",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data.intf",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.intf",
						TypeName:   "intf",
					},
				},
				{
					Name: "notEmbededIntf",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data.intf",
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.intf",
						TypeName:   "intf",
					},
				},
				{
					Name: "fn",
					Func: &model.Func{
						Args: []*model.Field{
							{
								Name: "v",
								Type: &model.Type{
									Underlying: "any",
									TypeName:   "any",
								},
							},
						},
						Results: []*model.Field{
							{
								Name: "error",
								Type: &model.Type{
									Underlying: "error",
									TypeName:   "error",
								},
							},
						},
					},
				},
			},
		},
		{
			Name: "intf",
			Type: &model.Type{
				Underlying: "github.com/maru44/stst/tests/data.intf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.intf",
				TypeName:   "intf",
			},
			Fields: []*model.Field{
				{
					Name: "AAA",
					Func: &model.Func{
						Args: []*model.Field{
							{
								Name: "in",
								Type: &model.Type{
									Underlying: "string",
									TypeName:   "string",
								},
							},
							{
								Name: "good",
								Type: &model.Type{
									Underlying: "github.com/maru44/stst/tests/data.Good",
									Package:    "github.com/maru44/stst/tests/data",
									PkPlusName: "data.Good",
									TypeName:   "Good",
								},
							},
							{
								Name: "sample",
								Type: &model.Type{
									Underlying: "github.com/maru44/stst/tests/data/aaa.Sample",
									Package:    "github.com/maru44/stst/tests/data/aaa",
									PkPlusName: "aaa.Sample",
									TypeName:   "Sample",
								},
							},
						},
						Results: []*model.Field{
							{
								Name: "string",
								Type: &model.Type{
									Underlying: "string",
									TypeName:   "string",
								},
							},
							{
								Name: "error",
								Type: &model.Type{
									Underlying: "error",
									TypeName:   "error",
								},
							},
						},
					},
				},
				{
					Name: "BBB",
					Func: &model.Func{},
				},
				{
					Name: "Intf",
					Type: &model.Type{
						Underlying: "github.com/maru44/stst/tests/data/aaa.Intf",
						Package:    "github.com/maru44/stst/tests/data/aaa",
						PkPlusName: "aaa.Intf",
						TypeName:   "Intf",
					},
				},
				{
					Name: "childIntf",
					Type: &model.Type{
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
			Type: &model.Type{
				Underlying: "github.com/maru44/stst/tests/data.childIntf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.childIntf",
				TypeName:   "childIntf",
			},
			Fields: []*model.Field{
				{
					Name: "CCC",
					Func: &model.Func{},
				},
			},
			IsInterface: true,
		},
	}

	assert.EqualValues(t, want, schemas)
}
