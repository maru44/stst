package data

import (
	"time"
)

type (
	SampleString string

	MapSimple map[any]struct{ a SampleString }
	MapS      map[any]struct{}

	SamplePrefixMap map[string][]*[][]*[]any

	Person struct {
		Name  string `bigquery:"name"` // comment
		Age   int    `bigquery:"age"`
		Sex   string `bigquery:"-"`
		Hobby string `bigquery:"hobby,nullable"`
		Good         // no name // no name
	}

	Animal struct {
		ID      string `bigquery:"id"`
		Kind    string `bigquery:"kind"`
		Goods   []*Good
		GoodPtr *Good `bigquery:"good_ptr"`
		GoodStr Good  `bigquery:"good_str" aaa:"bbb,nullable"`
		strPtr  *string
		tim     time.Time
		timPtr  *time.Time
		strs    []string
		funcs   []func(v any, ints []int, ptrInts *[]*int) bool
	}

	Gene[T any] struct {
		One T
	}

	Good struct {
		Name      string
		Sample    SampleString
		SamplePtr *SampleString
	}
)
