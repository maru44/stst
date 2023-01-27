package data

//go:generate go run github.com/maru44/stst/tests/gen

import (
	"time"

	"github.com/maru44/stst/tests/data/aaa"
)

type SampleString string

type Person struct {
	Name  string `bigquery:"name"` // comment
	Age   int    `bigquery:"age"`
	Sex   string `bigquery:"-"`
	Hobby string `bigquery:"hobby,nullable"`
	Good         // no name // no name
}

type Animal struct {
	ID      string `bigquery:"id"`
	Kind    string `bigquery:"kind"`
	Goods   []*Good
	GoodPtr *Good `bigquery:"good_ptr"`
	GoodStr Good  `bigquery:"good_str" aaa:"bbb,nullable"`
	strPtr  *string
	tim     time.Time
	timPtr  *time.Time
}

type Good struct {
	Name      string
	Sample    SampleString
	SamplePtr *SampleString
}

type intf struct {
	error
	str string
	aaa.Intf
	*Good
	intef aaa.Intf
	aaa.StrSample
}
