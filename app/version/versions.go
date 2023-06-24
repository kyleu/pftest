// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Versions []*Version

func (v Versions) Get(id string) *Version {
	for _, x := range v {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (v Versions) GetByIDs(ids ...string) Versions {
	var ret Versions
	for _, x := range v {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (v Versions) IDs() []string {
	ret := make([]string, 0, len(v)+1)
	for _, x := range v {
		ret = append(ret, x.ID)
	}
	return ret
}

func (v Versions) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(v)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range v {
		ret = append(ret, x.ID)
	}
	return ret
}

func (v Versions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(v)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range v {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (v Versions) Clone() Versions {
	return slices.Clone(v)
}
