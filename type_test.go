package path

import "testing"

func TestPathClean(t *testing.T) {
	a := Path("/a/b/..").Clean()
	isEqual(t, a, Path("/a"), "")

	b := Path("/a//./b/..").Clean()
	isEqual(t, b, Path("/a"), "")

	c := Path("///").Clean()
	isEqual(t, c, Path("/"), "")
}

func TestPathSplit(t *testing.T) {
	a, b := Path("/a/b/zz.png").Split()
	isEqual(t, a, Path("/a/b/"), "")
	isEqual(t, b, "zz.png", "")
}

func TestPathOf(t *testing.T) {
	isEqual(t, Of("", "cc", "d", "", "/e/", "x.png"), Path("cc/d/e/x.png"), "")
	isEqual(t, Of("", "/cc", "d", "/e/", ""), Path("/cc/d/e"), "")
	isEqual(t, Of(), Path(""), "")
}

func TestPathOfAny(t *testing.T) {
	isEqual(t, OfAny("", "cc", "d", "", 1, Path("/e/f"), "x.png"), Path("cc/d/1/e/f/x.png"), "")
	isEqual(t, OfAny(), Path(""), "")
}

func TestPathPrepend(t *testing.T) {
	isEqual(t, Path("a/b/xx.png").Prepend("", "/cc", "d", ""), Path("/cc/d/a/b/xx.png"), "")
	isEqual(t, Path("//a/b/xx.png").Prepend("", "cc", "d/"), Path("cc/d/a/b/xx.png"), "")
}

func TestPathAppend(t *testing.T) {
	isEqual(t, Path("/a/b").Append("", "cc", "d", "xx.png", ""), Path("/a/b/cc/d/xx.png"), "")
	isEqual(t, Path("/a/b/").Append("", "cc", "d", "xx.png", ""), Path("/a/b/cc/d/xx.png"), "")
	isEqual(t, Path("/a/b/").Append("", "/cc", "d", "xx.png", ""), Path("/a/b/cc/d/xx.png"), "")
	isEqual(t, Path("/a/b").Append("", "/cc", "d", "xx.png", ""), Path("/a/b/cc/d/xx.png"), "")
}

func TestPathExt(t *testing.T) {
	isEqual(t, Path("/a/b/zz.png").Ext(), ".png", "")
	isEqual(t, Path("/a/b/zz").Ext(), "", "")
}

func TestPathBase(t *testing.T) {
	isEqual(t, Path("/a/b/zz.png").Base(), "zz.png", "")
}

func TestPathIsAbs(t *testing.T) {
	isEqual(t, Path("/a/b/zz.png").IsAbs(), true, "")
	isEqual(t, Path("a/b/zz.png").IsAbs(), false, "")
}

func TestPathHasPrefix(t *testing.T) {
	isEqual(t, Path("/a/b/zz.png").HasPrefix("/a/b"), true, "")
	isEqual(t, Path("/a/b/zz.png").HasPrefix("/a/b/"), true, "")
	isEqual(t, Path("/a/b/zz.png").HasPrefix("/a/c/"), false, "")
}

func TestPathHasSuffix(t *testing.T) {
	isEqual(t, Path("/a/b/zz.png").HasSuffix("/zz.png"), true, "")
	isEqual(t, Path("/a/b/zz.png").HasSuffix("b/zz.png"), true, "")
	isEqual(t, Path("/a/b/zz.png").HasSuffix("b/aa.png"), false, "")
}

func TestPathDir(t *testing.T) {
	isEqual(t, Path("/a/b/zz.png").Dir(), Path("/a/b"), "")
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
	isEqual(t, Path("/a/b/c/zz.png").Take(2), Path("/a/b"), "")
}

func TestPathDrop(t *testing.T) {
	isEqual(t, Path("/a/b/c/zz.png").Drop(2), Path("/c/zz.png"), "")
}

func TestPathIsEmpty(t *testing.T) {
	isEqual(t, Path("/a/b/c/zz.png").IsEmpty(), false, "")
	isEqual(t, Path("").IsEmpty(), true, "")
}

func TestPathSegments(t *testing.T) {
	isEqual(t, Path("/a/b/c/zz.png").Segments(), []string{"a", "b", "c", "zz.png"}, "")
	isEqual(t, Path("a/b/c/zz.png").Segments(), []string{"a", "b", "c", "zz.png"}, "")
	isEqual(t, Path("/a/b/c/").Segments(), []string{"a", "b", "c"}, "")
	isEqual(t, Path("/").Segments(), []string(nil), "")
	isEqual(t, Path("").Segments(), []string(nil), "")
}

func TestPathString(t *testing.T) {
	isEqual(t, Path("/a/b/c/zz.png").String(), "/a/b/c/zz.png", "")
}

func TestPathScan(t *testing.T) {
	a := new(Path)

	err := a.Scan(nil)
	isNil(t, err, "")
	isEqual(t, *a, Path(""), "")

	err = a.Scan("/a/b/c/zz.png")
	isNil(t, err, "")
	isEqual(t, *a, Path("/a/b/c/zz.png"), "")

	err = a.Scan([]byte("/a/b/c/zz.png"))
	isNil(t, err, "")
	isEqual(t, *a, Path("/a/b/c/zz.png"), "")

	err = a.Scan(123)
	isEqual(t, err.Error(), "Path.Scan(123)", "")
}

func TestPathValue(t *testing.T) {
	a, err := Path("/a/b/c/zz.png").Value()
	isNil(t, err, "")
	isEqual(t, a, "/a/b/c/zz.png", "")
}
