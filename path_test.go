package path

import (
	"reflect"
	"testing"
)

func TestDivideAndDropLeading(t *testing.T) {
	cases := []struct {
		n                 int
		input, head, tail string
	}{
		{0, "", "", ""},

		{0, "a/b/c/x.png", "", "a/b/c/x.png"},
		{1, "a/b/c/x.png", "a", "/b/c/x.png"},
		{2, "a/b/c/x.png", "a/b", "/c/x.png"},
		{3, "a/b/c/x.png", "a/b/c", "/x.png"},
		{4, "a/b/c/x.png", "a/b/c/x.png", ""},
		{5, "a/b/c/x.png", "a/b/c/x.png", ""},

		{0, "/a/b/c/x.png", "", "/a/b/c/x.png"},
		{1, "/a/b/c/x.png", "/a", "/b/c/x.png"},
		{2, "/a/b/c/x.png", "/a/b", "/c/x.png"},
		{3, "/a/b/c/x.png", "/a/b/c", "/x.png"},
		{4, "/a/b/c/x.png", "/a/b/c/x.png", ""},
		{5, "/a/b/c/x.png", "/a/b/c/x.png", ""},

		{0, "/a/b/c/", "", "/a/b/c/"},
		{1, "/a/b/c/", "/a", "/b/c/"},
		{2, "/a/b/c/", "/a/b", "/c/"},
		{3, "/a/b/c/", "/a/b/c", "/"},
		{4, "/a/b/c/", "/a/b/c/", ""},
		{5, "/a/b/c/", "/a/b/c/", ""},
	}

	for i, test := range cases {
		p1, p2 := Divide(test.input, test.n)
		isEqual(t, p1, test.head, i)
		isEqual(t, p2, test.tail, i)

		head := Take(test.input, test.n)
		isEqual(t, head, test.head, i)

		tail := Drop(test.input, test.n)
		isEqual(t, tail, test.tail, i)
	}
}

//-------------------------------------------------------------------------------------------------

func isNil(t *testing.T, a, hint interface{}) {
	t.Helper()
	if a != nil {
		t.Errorf("Got %#v; expected nil - for %v\n", a, hint)
	}
}

func isEqual(t *testing.T, a, b, hint interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Got %#v; expected %#v - for %v\n", a, b, hint)
	}
}

//func isNotEqual(t *testing.T, a, b, hint interface{}) {
//	t.Helper()
//	if reflect.DeepEqual(a, b) {
//		t.Errorf("Got %#v; expected something else - for %v\n", a, hint)
//	}
//}

//func isGte(t *testing.T, a, b int, hint interface{}) {
//	t.Helper()
//	if a < b {
//		t.Errorf("Got %d; expected at least %d - for %v\n", a, b, hint)
//	}
//}
