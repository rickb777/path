package path

import (
	std "path"
	"strings"
)

// Path is just a string with specialised methods.
type Path string

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
//	1. Replace multiple slashes with a single slash.
//	2. Eliminate each . path name element (the current directory).
//	3. Eliminate each inner .. path name element (the parent directory)
//	   along with the non-.. element that precedes it.
//	4. Eliminate .. elements that begin a rooted path:
//	   that is, replace "/.." by "/" at the beginning of a path.
//
// The returned path ends in a slash only if it is the root "/".
//
// If the result of this process is an empty string, Clean
// returns the string ".".
//
// See also Rob Pike, ``Lexical File Names in Plan 9 or
// Getting Dot-Dot Right,''
// https://9p.io/sys/doc/lexnames.html
func (p Path) Clean() Path {
	return Path(std.Clean(string(p)))
}

// Split splits path immediately following the final slash,
// separating it into a directory and file name component.
// If there is no slash in path, Split returns an empty dir and
// file set to path.
// The returned values have the property that path = dir+file.
func (p Path) Split() (dir Path, file string) {
	d, f := std.Split(string(p))
	return Path(d), f
}

// Append joins any number of path elements into a single path, adding a
// separating slash if necessary. The result is Cleaned; in particular,
// all empty strings are ignored.
func (p Path) Append(elem ...string) Path {
	if !strings.HasSuffix(string(p), "/") {
		p = p + "/"
	}
	return p + Path(std.Join(elem...))
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
func (p Path) Ext() string {
	return std.Ext(string(p))
}

// Base returns the last element of path.
// Trailing slashes are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of slashes, Base returns "/".
func (p Path) Base() string {
	return std.Base(string(p))
}

// IsAbs reports whether the path is absolute.
func (p Path) IsAbs() bool {
	return std.IsAbs(string(p))
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element using Split, the path is Cleaned and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of slashes followed by non-slash bytes, Dir
// returns a single slash. In any other case, the returned path does not end in a
// slash.
func (p Path) Dir() Path {
	return Path(std.Dir(string(p)))
}

// Divide divides a path at the nth slash, not counting the leading slash
// if there is one.
//
// The resulting pair (head, tail) always satisfy
//
//   head + tail = path
func (p Path) Divide(nth int) (Path, Path) {
	head, tail := Divide(string(p), nth)
	return Path(head), Path(tail)
}

// Drop is a helper for Divide that returns the tail part only.
func (p Path) Drop(unwanted int) Path {
	_, tail := p.Divide(unwanted)
	return tail
}

// Take is a helper for Divide that returns the head part only.
func (p Path) Take(wanted int) Path {
	head, _ := p.Divide(wanted)
	return head
}

// Next returns the first segment (without any leading '/') and the rest. It can
// be used for iterating through the path segments; the end has been reached when
// the tail is empty (see IsEmpty).
func (p Path) Next() (string, Path) {
	head, tail := p.Divide(1)
	next := string(head)
	if strings.HasPrefix(next, "/") {
		next = next[1:]
	}
	return next, tail
}

// IsEmpty returns true if the path is empty.
func (p Path) IsEmpty() bool {
	return len(p) == 0
}

// Segments returns the path split into the parts between slashes. Any leading or
// trailing slash on the path is removed before the path is split, so there is no
// leading or trailing blank string in the result.
func (p Path) Segments() []string {
	if p == "/" {
		return nil
	}
	if strings.HasPrefix(string(p), "/") {
		p = p[1:]
	}
	if strings.HasSuffix(string(p), "/") {
		p = p[:len(p)-1]
	}
	return strings.Split(string(p), "/")
}

// String simply converts the type to a string.
func (p Path) String() string {
	return string(p)
}

