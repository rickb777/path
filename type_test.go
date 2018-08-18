package path

import "testing"

func TestPathClean(t *testing.T) {
	a := Path("/a/b/..").Clean()
	isEqual(t, a, Path("/a"), "")
}

func TestPathSplit(t *testing.T) {
	a, b := Path("/a/b/zz.png").Split()
	isEqual(t, a, Path("/a/b/"), "")
	isEqual(t, b, "zz.png", "")
}

func TestPathAppend(t *testing.T) {
	a := Path("/a/b").Append("", "cc", "d", "xx.png")
	isEqual(t, a, Path("/a/b/cc/d/xx.png"), "")
}

func TestPathExt(t *testing.T) {
	a := Path("/a/b/zz.png").Ext()
	isEqual(t, a, ".png", "")
}

func TestPathBase(t *testing.T) {
	a := Path("/a/b/zz.png").Base()
	isEqual(t, a, "zz.png", "")
}

func TestPathIsAbs(t *testing.T) {
	a := Path("/a/b/zz.png").IsAbs()
	isEqual(t, a, true, "")
}

func TestPathDir(t *testing.T) {
	a := Path("/a/b/zz.png").Dir()
	isEqual(t, a, Path("/a/b"), "")
}

func TestPathNext(t *testing.T) {
	a, b := Path("/a/b/c/zz.png").Next()
	isEqual(t, a, "a", "")
	isEqual(t, b, Path("/b/c/zz.png"), "")
	_, c := b.Next()
	isEqual(t, c, Path("/c/zz.png"), "")
}

func TestPathDivide(t *testing.T) {
	a, b := Path("/a/b/c/zz.png").Divide(2)
	isEqual(t, a, Path("/a/b"), "")
	isEqual(t, b, Path("/c/zz.png"), "")
}

func TestPathTake(t *testing.T) {
	a := Path("/a/b/c/zz.png").Take(2)
	isEqual(t, a, Path("/a/b"), "")
}

func TestPathDrop(t *testing.T) {
	b := Path("/a/b/c/zz.png").Drop(2)
	isEqual(t, b, Path("/c/zz.png"), "")
}

func TestPathIsEmpty(t *testing.T) {
	isEqual(t, Path("/a/b/c/zz.png").IsEmpty(), false, "")
	isEqual(t, Path("").IsEmpty(), true, "")
}

func TestPathSegments(t *testing.T) {
	a := Path("/a/b/c/zz.png").Segments()
	isEqual(t, a, []string{"", "a", "b", "c", "zz.png"}, "")

	b := Path("a/b/c/zz.png").Segments()
	isEqual(t, b, []string{"a", "b", "c", "zz.png"}, "")
}

func TestPathString(t *testing.T) {
	a := Path("/a/b/c/zz.png").String()
	isEqual(t, a, "/a/b/c/zz.png", "")
}
