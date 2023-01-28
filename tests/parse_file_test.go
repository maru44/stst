package tests_test

import (
	"testing"

	"github.com/maru44/stst"
	"github.com/maru44/stst/stmodel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	ps, err := loadPackages("github.com/maru44/stst/tests/data")
	require.NoError(t, err)
	require.Len(t, ps, 1)

	var schemas []*stmodel.Schema
	for _, pk := range ps {
		require.NoError(t, err)
		p := stst.NewParser(pk)

		for i, f := range pk.Syntax {
			if i == 1 {
				s := p.ParseFile(f)
				schemas = append(schemas, s...)
			}
		}
	}

	want := []*stmodel.Schema{
		{
			Name: "withIntf",
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
					IsPtr: true,
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
			},
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.withIntf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.withIntf",
				TypeName:   "withIntf",
			},
		},
		{
			Name: "intf",
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
						Package:    "github.com/maru44/stst/tests/data",
						PkPlusName: "data.childIntf",
						TypeName:   "childIntf",
					},
				},
			},
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.intf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.intf",
				TypeName:   "intf",
			},
			IsInterface: true,
		},
		{
			Name: "childIntf",
			Fields: []*stmodel.Field{
				{
					Name: "CCC",
					Func: &stmodel.Func{},
				},
			},
			Type: &stmodel.Type{
				Underlying: "github.com/maru44/stst/tests/data.childIntf",
				Package:    "github.com/maru44/stst/tests/data",
				PkPlusName: "data.childIntf",
				TypeName:   "childIntf",
			},
			IsInterface: true,
		},
	}
	assert.Equal(t, want, schemas)
}
