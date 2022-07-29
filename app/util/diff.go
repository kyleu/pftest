// Content managed by Project Forge, see [projectforge.md] for details.
package util

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Diff struct {
	Path string `json:"path"`
	Old  string `json:"o,omitempty"`
	New  string `json:"n"`
}

func NewDiff(p string, o string, n string) *Diff {
	return &Diff{Path: p, Old: o, New: n}
}

func (d Diff) String() string {
	return d.Path
}

func (d Diff) StringVerbose() string {
	return fmt.Sprintf("%s (%q != %q)", d.Path, d.Old, d.New)
}

type Diffs []*Diff

func (d Diffs) String() string {
	sb := make([]string, 0, len(d))
	for _, x := range d {
		sb = append(sb, x.String())
	}
	return strings.Join(sb, "; ")
}

func (d Diffs) StringVerbose() string {
	sb := make([]string, 0, len(d))
	for _, x := range d {
		sb = append(sb, x.StringVerbose())
	}
	return strings.Join(sb, "; ")
}

func DiffObjects(l any, r any, path ...string) Diffs {
	return DiffObjectsIgnoring(l, r, nil, path...)
}

func DiffObjectsIgnoring(l any, r any, ignored []string, path ...string) Diffs {
	var ret Diffs
	if len(path) > 0 && slices.Contains(ignored, path[len(path)-1]) {
		return ret
	}
	if l == nil {
		return append(ret, NewDiff(strings.Join(path, "."), "", fmt.Sprint(r)))
	}
	if r == nil {
		return append(ret, NewDiff(strings.Join(path, "."), fmt.Sprint(l), ""))
	}
	if lt, rt := fmt.Sprintf("%T", l), fmt.Sprintf("%T", r); lt != rt {
		return append(ret, NewDiff(strings.Join(path, "."), ToJSONCompact(l), ToJSONCompact(r)))
	}

	switch t := l.(type) {
	case ValueMap:
		ret = append(ret, diffMaps(t, r, ignored, path...)...)
	case map[string]any:
		ret = append(ret, diffMaps(t, r, ignored, path...)...)
	case map[string]int:
		ret = append(ret, diffIntMaps(t, r, ignored, path...)...)
	case []any:
		ret = append(ret, diffArrays(t, r, ignored, path...)...)
	case Diffs:
		rm, _ := r.(Diffs)
		for idx, v := range t {
			rv := rm[idx]
			ret = append(ret, DiffObjectsIgnoring(v, rv, ignored, append([]string{}, path...)...)...)
		}
	case int:
		i, _ := r.(int)
		if t != i {
			ret = append(ret, NewDiff(strings.Join(path, "."), fmt.Sprint(t), fmt.Sprint(i)))
		}
	case string:
		s, _ := r.(string)
		if t != s {
			ret = append(ret, NewDiff(strings.Join(path, "."), t, s))
		}
	default:
		if lj, rj := ToJSONCompact(l), ToJSONCompact(r); lj != rj {
			ret = append(ret, NewDiff(strings.Join(path, "."), lj, rj))
		}
	}

	return ret
}

func diffArrays(l []any, r any, ignored []string, path ...string) Diffs {
	var ret Diffs
	rm, _ := r.([]any)
	for idx, v := range l {
		if len(rm) > idx {
			rv := rm[idx]
			ret = append(ret, DiffObjectsIgnoring(v, rv, ignored, append(append([]string{}, path...), fmt.Sprint(idx))...)...)
		} else {
			ret = append(ret, DiffObjectsIgnoring(v, nil, ignored, append(append([]string{}, path...), fmt.Sprint(idx))...)...)
		}
	}
	if len(rm) > len(l) {
		for i := len(l); i < len(rm); i++ {
			ret = append(ret, DiffObjectsIgnoring(nil, rm[i], ignored, append(append([]string{}, path...), fmt.Sprint(i))...)...)
		}
	}
	return ret
}

func diffMaps(l map[string]any, r any, ignored []string, path ...string) Diffs {
	var ret Diffs
	rm, ok := r.(map[string]any)
	if !ok {
		rm, _ = r.(ValueMap)
	}
	for k, v := range l {
		if slices.Contains(ignored, k) {
			continue
		}
		rv := rm[k]
		ret = append(ret, DiffObjectsIgnoring(v, rv, ignored, append(append([]string{}, path...), k)...)...)
	}
	for k, v := range rm {
		if slices.Contains(ignored, k) {
			continue
		}
		if _, exists := l[k]; !exists {
			ret = append(ret, DiffObjectsIgnoring(nil, v, ignored, append(append([]string{}, path...), k)...)...)
		}
	}
	return ret
}

func diffIntMaps(l map[string]int, r any, ignored []string, path ...string) Diffs {
	var ret Diffs
	rm, _ := r.(map[string]int)
	for k, v := range l {
		if slices.Contains(ignored, k) {
			continue
		}
		rv := rm[k]
		ret = append(ret, DiffObjectsIgnoring(v, rv, ignored, append(append([]string{}, path...), k)...)...)
	}
	for k, v := range rm {
		if slices.Contains(ignored, k) {
			continue
		}
		if _, exists := l[k]; !exists {
			ret = append(ret, DiffObjectsIgnoring(nil, v, ignored, append(append([]string{}, path...), k)...)...)
		}
	}
	return ret
}
