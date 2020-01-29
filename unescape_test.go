package unescape

import (
	"testing"
)

func TestUnescape(t *testing.T) {
	list := []struct {
		I string
		W string
	}{
		{I: `a%u1234%aa+%aa%u1234`, W: "a\u1234\xaa+\xaa\u1234"},
		{I: `a%u3042%u3042%09`, W: "a\u3042あ\t"},
		{I: `a%09+%20`, W: "a\t+ "},
		{I: `a%%+ab`, W: "a%%+ab"},
	}
	for _, i := range list {
		in := i.I
		re := Unescape(in)
		wa := i.W
		if re != wa {
			t.Errorf("%+q -> %+q, want %+q", in, re, wa)
		}
	}
}
