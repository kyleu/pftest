// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"slices"

	"github.com/samber/lo"
)

type Versions []*Version

func (v Versions) Get(id string) *Version {
	return lo.FindOrElse(v, nil, func(x *Version) bool {
		return x.ID == id
	})
}

func (v Versions) GetByIDs(ids ...string) Versions {
	return lo.Filter(v, func(x *Version, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (v Versions) IDs() []string {
	return lo.Map(v, func(x *Version, _ int) string {
		return x.ID
	})
}

func (v Versions) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(v)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(v, func(x *Version, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (v Versions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(v)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(v, func(x *Version, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (v Versions) Clone() Versions {
	return slices.Clone(v)
}
