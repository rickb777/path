package path

import (
	"database/sql/driver"
	"fmt"
	std "path"
	"strings"
)

// Path is just a string with specialised methods.
type Path string

// Of joins any number of path elements to build a new path, adding a
// separating slashes as necessary. The result is Cleaned; in particular,
// all empty strings are ignored.
//
// If the first non-blank elem has a leading slash, the path will also
// have a leading slash. It will not have a trailing slash.
func Of(elem ...string) Path {
	return Path(std.Join(elem...)) // Join includes Clean
}

// OfAny joins any number of path elements to build a new path, adding a
// separating slashes as necessary. Each element is treated as either a
// string, or a Path, or some other type; in the latter case, it is formatted
// using fmt.Sprintf so the "%v" rules apply.
//
// The result is Cleaned; in particular, all empty strings are ignored.
//
// If the first non-blank elem has a leading slash, the path will also
// have a leading slash. It will not have a trailing slash.
func OfAny(elem ...interface{}) Path {
	ss := make([]string, len(elem))
	for i, v := range elem {
		switch x := v.(type) {
		case string:
			ss[i] = x
		case Path:
			ss[i] = string(x)
		default:
			ss[i] = fmt.Sprintf("%v", v)
		}
	}
	return Of(ss...)
}

// Prepend joins any number of path elements to the beginning of the path, adding a
// separating slashes as necessary. The result is Cleaned; in particular,
// all empty strings are ignored.
func (path Path) Prepend(elem ...string) Path {
	if !strings.HasPrefix(string(path), "/") {
		path = "/" + path
	}
	q := Path(std.Join(elem...)) + path
	return q.Clean()
}

// Append joins any number of path elements to the end of the path, adding a
// separating slashes as necessary. The result is Cleaned; in particular,
// all empty strings are ignored.
func (path Path) Append(elem ...string) Path {
	if !strings.HasSuffix(string(path), "/") {
		path = path + "/"
	}
	q := path + Path(std.Join(elem...))
	return q.Clean()
}

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
func (path Path) Clean() Path {
	return Path(std.Clean(string(path)))
}

// Split splits path immediately following the final slash,
// separating it into a directory and file name component.
// If there is no slash in path, Split returns an empty dir and
// file set to path.
// The returned values have the property that path = dir+file.
func (path Path) Split() (dir Path, file string) {
	d, f := std.Split(string(path))
	return Path(d), f
}

// SplitExt splits the file name from its extension.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
// The dot is included in the extension.
//
// Everything prior to the last dot is returned as the first result.
func (path Path) SplitExt() (Path, string) {
	p, e := SplitExt(string(path))
	return Path(p), e
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
//
// Unlike ExtOnly, the dot is included in the result.
func (path Path) Ext() string {
	return std.Ext(string(path))
}

// ExtOnly returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
//
// Unlike Ext, the dot is not included in the result.
func (p Path) ExtOnly() string {
	ext := std.Ext(string(p))
	if ext == "" {
		return ""
	}
	return ext[1:]
}

// Base returns the last element of path.
// Trailing slashes are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of slashes, Base returns "/".
func (path Path) Base() string {
	return std.Base(string(path))
}

// IsAbs reports whether the path is absolute.
func (path Path) IsAbs() bool {
	return std.IsAbs(string(path))
}

// HasPrefix reports whether the path starts with a particular prefix.
func (path Path) HasPrefix(other Path) bool {
	return strings.HasPrefix(string(path), string(other))
}

// HasSuffix reports whether the path ends with a particular suffix.
func (path Path) HasSuffix(other Path) bool {
	return strings.HasSuffix(string(path), string(other))
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element using Split, the path is Cleaned and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of slashes followed by non-slash bytes, Dir
// returns a single slash. In any other case, the returned path does not end in a
// slash.
func (path Path) Dir() Path {
	return Path(std.Dir(string(path)))
}

// Divide divides a path at the nth slash, not counting the leading slash
// if there is one.
//
// The resulting pair (head, tail) always satisfy
//
//   head + tail = path
func (path Path) Divide(nth int) (Path, Path) {
	head, tail := Divide(string(path), nth)
	return Path(head), Path(tail)
}

// Drop is a helper for Divide that returns the tail part only.
func (path Path) Drop(unwanted int) Path {
	_, tail := path.Divide(unwanted)
	return tail
}

// Take is a helper for Divide that returns the head part only.
func (path Path) Take(wanted int) Path {
	head, _ := path.Divide(wanted)
	return head
}

// Next returns the first segment (without any leading '/') and the rest. It can
// be used for iterating through the path segments; the end has been reached when
// the tail is empty (see IsEmpty).
func (path Path) Next() (string, Path) {
	head, tail := path.Divide(1)
	next := string(head)
	if strings.HasPrefix(next, "/") {
		next = next[1:]
	}
	return next, tail
}

// IsEmpty returns true if the path is empty.
func (path Path) IsEmpty() bool {
	return len(path) == 0
}

// Segments returns the path split into the parts between slashes. Any leading or
// trailing slash on the path is removed before the path is split, so there is no
// leading or trailing blank string in the result.
//
// The root path "/" will return nil. A blank path will also return nil; this ensures
// that the Segments of the zero value of Path is a zero value of []string.
func (path Path) Segments() []string {
	if path == "/" || path == "" {
		return nil
	}
	if strings.HasPrefix(string(path), "/") {
		path = path[1:]
	}
	if strings.HasSuffix(string(path), "/") {
		path = path[:len(path)-1]
	}
	return strings.Split(string(path), "/")
}

// String simply converts the type to a string.
func (path Path) String() string {
	return string(path)
}

//-------------------------------------------------------------------------------------------------

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func (path *Path) Scan(value interface{}) error {
	if value == nil {
		*path = Path("")
		return nil
	}

	switch value.(type) {
	case string:
		*path = Path(value.(string))
	case []byte:
		*path = Path(string(value.([]byte)))
	case nil:
	default:
		return fmt.Errorf("Path.Scan(%#v)", value)
	}
	return nil
}

// Value converts the value to a string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func (path Path) Value() (driver.Value, error) {
	return string(path), nil
}
