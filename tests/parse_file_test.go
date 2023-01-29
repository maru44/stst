package tests_test

import (
	"testing"

	"github.com/maru44/stst"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFile(t *testing.T) {
	ps, err := loadPackages("github.com/maru44/stst/tests/data")
	require.NoError(t, err)
	require.Len(t, ps, 1)

	var schemas []*stst.Schema
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

	want := []*stst.Schema{
		{
			Name: "withIntf",
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
					Name: "notEmbeddedIntf",
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
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.withIntf",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.withIntf",
				TypeName:    "withIntf",
			},
		},
		{
			Name: "intf",
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
						PkgID:       "github.com/maru44/stst/tests/data",
						PkgPlusName: "data.childIntf",
						TypeName:    "childIntf",
					},
				},
			},
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.intf",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.intf",
				TypeName:    "intf",
			},
			IsInterface: true,
		},
		{
			Name: "childIntf",
			Fields: []*stst.Field{
				{
					Name: "CCC",
					Func: &stst.Func{},
				},
			},
			Type: &stst.Type{
				Underlying:  "github.com/maru44/stst/tests/data.childIntf",
				PkgID:       "github.com/maru44/stst/tests/data",
				PkgPlusName: "data.childIntf",
				TypeName:    "childIntf",
			},
			IsInterface: true,
		},
	}
	assert.Equal(t, want, schemas)
}
