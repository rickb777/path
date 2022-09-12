package path

import (
	"strings"
)

// Divide divides a path at the nth slash, not counting the leading slash
// if there is one.
//
// The resulting pair (head, tail) always satisfy
//
//	head + tail = path
func Divide(path string, nth int) (string, string) {
	if path == "" {
		return "", ""
	}

	if nth == 0 {
		return "", path
	}

	pivot := 0
	if path[0] == '/' {
		pivot++
	}

	for i := nth; i > 0; i-- {
		slash := strings.IndexByte(path[pivot:], '/') + pivot
		if slash <= pivot {
			return path, ""
		}
		pivot = slash + 1
	}

	pivot--
	return path[:pivot], path[pivot:]
}

// Drop is a helper for Divide that returns the tail part only.
func Drop(path string, unwanted int) string {
	_, tail := Divide(path, unwanted)
	return tail
}

// Take is a helper for Divide that returns the head part only.
func Take(path string, wanted int) string {
	head, _ := Divide(path, wanted)
	return head
}

// SplitExt splits the file name from its extension.
// The extension is the suffix beginning at the final dot
// in the final slash-separated element of path;
// it is empty if there is no dot.
// The dot is included in the extension.
//
// Everything prior to the last dot is returned as the first result.
func SplitExt(path string) (string, string) {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return path, ""
}
