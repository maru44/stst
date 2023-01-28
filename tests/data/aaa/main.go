package aaa

type (
	Intf interface {
		Hello() string
	}

	IntSample int

	Sample struct {
		Str string `tag0:"xxx" tag1:"yyy,zzz"` // comment
	}

	prefixes []*[][][]*struct{}
)
