package path

import (
	std "path"
	"strings"
	"fmt"
	"database/sql/driver"
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
	return Path(std.Join(elem...))
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
func (p Path) Prepend(elem ...string) Path {
	if !strings.HasPrefix(string(p), "/") {
		p = "/" + p
	}
	q := Path(std.Join(elem...)) + p
	return q.Clean()
}

// Append joins any number of path elements to the end of the path, adding a
// separating slashes as necessary. The result is Cleaned; in particular,
// all empty strings are ignored.
func (p Path) Append(elem ...string) Path {
	if !strings.HasSuffix(string(p), "/") {
		p = p + "/"
	}
	q := p + Path(std.Join(elem...))
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

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
//
// Unlike ExtOnly, the dot is included in the result.
func (p Path) Ext() string {
	return std.Ext(string(p))
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
func (p Path) Base() string {
	return std.Base(string(p))
}

// IsAbs reports whether the path is absolute.
func (p Path) IsAbs() bool {
	return std.IsAbs(string(p))
}

// HasPrefix reports whether the path starts with a particular prefix.
func (p Path) HasPrefix(other Path) bool {
	return strings.HasPrefix(string(p), string(other))
}

// HasSuffix reports whether the path ends with a particular suffix.
func (p Path) HasSuffix(other Path) bool {
	return strings.HasSuffix(string(p), string(other))
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
//
// The root path "/" will return nil. A blank path will also return nil; this ensures
// that the Segments of the zero value of Path is a zero value of []string.
func (p Path) Segments() []string {
	if p == "/" || p == "" {
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

//-------------------------------------------------------------------------------------------------

// Scan parses some value. It implements sql.Scanner,
// https://golang.org/pkg/database/sql/#Scanner
func (p *Path) Scan(value interface{}) error {
	if value == nil {
		*p = Path("")
		return nil
	}

	switch value.(type) {
	case string:
		*p = Path(value.(string))
	case []byte:
		*p = Path(string(value.([]byte)))
	case nil:
	default:
		return fmt.Errorf("Path.Scan(%#v)", value)
	}
	return nil
}

// Value converts the value to a string. It implements driver.Valuer,
// https://golang.org/pkg/database/sql/driver/#Valuer
func (p Path) Value() (driver.Value, error) {
	return string(p), nil
}
