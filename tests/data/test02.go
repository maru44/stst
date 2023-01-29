package data

import "github.com/maru44/stst/tests/data/aaa"

type (
	withIntf struct {
		error
		str string
		aaa.Intf
		*Good
		intef aaa.Intf
		aaa.IntSample
		intf
		notEmbeddedIntf intf
		fn              func(v any) error
		ma              map[string]any
	}

	intf interface {
		AAA(in string, good Good, sample aaa.Sample) (string, error)
		BBB()
		aaa.Intf
		childIntf
	}

	childIntf interface {
		CCC()
	}
)
