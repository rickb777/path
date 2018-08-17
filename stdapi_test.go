package path

import "testing"

func TestClean(t *testing.T) {
	a := Clean("/a/b/..")
	isEqual(t, a, "/a", "")
}

func TestSplit(t *testing.T) {
	a, b := Split("/a/b/zz.png")
	isEqual(t, a, "/a/b/", "")
	isEqual(t, b, "zz.png", "")
}

func TestJoin(t *testing.T) {
	a := Join("/a", "b", "cc")
	isEqual(t, a, "/a/b/cc", "")
}

func TestExt(t *testing.T) {
	a := Ext("/a/b/zz.png")
	isEqual(t, a, ".png", "")
}

func TestBase(t *testing.T) {
	a := Base("/a/b/zz.png")
	isEqual(t, a, "zz.png", "")
}

func TestIsAbs(t *testing.T) {
	a := IsAbs("/a/b/zz.png")
	isEqual(t, a, true, "")
}

func TestDir(t *testing.T) {
	a := Dir("/a/b/zz.png")
	isEqual(t, a, "/a/b", "")
}

func TestMatch(t *testing.T) {
	a, e := Match("/a/b/*.png", "/a/b/zz.png")
	isEqual(t, a, true, "")
	isEqual(t, e, nil, "")
}
