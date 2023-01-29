package stst_test

import (
	"testing"

	"github.com/maru44/stst"
	"github.com/stretchr/testify/assert"
)

func TestSetPackage(t *testing.T) {
	tests := []struct {
		name string
		typ  *stst.Type
		want *stst.Type
	}{
		{
			name: "ok: set",
			typ: &stst.Type{
				Underlying: "aaa/bbb/ccc/ddd/eee.Fff",
				TypeName:   "Fff",
			},
			want: &stst.Type{
				Underlying:  "aaa/bbb/ccc/ddd/eee.Fff",
				PkgID:       "aaa/bbb/ccc/ddd/eee",
				PkgPlusName: "eee.Fff",
				TypeName:    "Fff",
			},
		},
		{
			name: "ok: without set",
			typ: &stst.Type{
				Underlying: "string",
				TypeName:   "Fff",
			},
			want: &stst.Type{
				Underlying: "string",
				TypeName:   "Fff",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.typ.SetPackage()
			assert.Equal(t, tt.typ, tt.want)
		})
	}
}

func TestArrayLength(t *testing.T) {
	tests := []struct {
		name   string
		prefix stst.TypePrefix
		array  bool
		length int
	}{
		{
			name:   "ok: array",
			prefix: stst.TypePrefix("[888]"),
			array:  true,
			length: 888,
		},
		{
			name:   "ok: not array pattern is not correct",
			prefix: stst.TypePrefix("[aaa]"),
		},
		{
			name:   "ok: not array slice",
			prefix: stst.TypePrefixSlice,
		},
	}

	for _, tt := range tests {
		tt := tt
		len, array := tt.prefix.ArrayLength()
		assert.Equal(t, tt.length, len)
		assert.Equal(t, tt.array, array)
	}
}

func TestTypePrefixKind(t *testing.T) {
	tests := []struct {
		name   string
		prefix stst.TypePrefix
		kind   stst.TypePrefixKind
	}{
		{
			name:   "ok: array",
			prefix: stst.TypePrefix("[888]"),
			kind:   stst.TypePrefixKindArray,
		},
		{
			name:   "ok: not array slice",
			prefix: stst.TypePrefixSlice,
			kind:   stst.TypePrefixKindSlice,
		},
		{
			name:   "ok: ptr",
			prefix: stst.TypePrefixPtr,
			kind:   stst.TypePrefixKindPtr,
		},
		{
			name:   "ok: not array pattern is not correct",
			prefix: stst.TypePrefix("[aaa]"),
			kind:   stst.TypePrefixKindUnknown,
		},
		{
			name:   "ok: unknown",
			prefix: stst.TypePrefix("uk"),
			kind:   stst.TypePrefixKindUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		got := tt.prefix.Kind()
		assert.Equal(t, tt.kind, got)
	}
}
