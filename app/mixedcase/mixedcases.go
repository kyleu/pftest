// Package mixedcase - Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import (
	"slices"

	"github.com/samber/lo"
)

type MixedCases []*MixedCase

func (m MixedCases) Get(id string) *MixedCase {
	return lo.FindOrElse(m, nil, func(x *MixedCase) bool {
		return x.ID == id
	})
}

func (m MixedCases) GetByIDs(ids ...string) MixedCases {
	return lo.Filter(m, func(xx *MixedCase, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (m MixedCases) GetByID(id string) MixedCases {
	return lo.Filter(m, func(xx *MixedCase, _ int) bool {
		return xx.ID == id
	})
}

func (m MixedCases) IDs() []string {
	return lo.Map(m, func(x *MixedCase, _ int) string {
		return x.ID
	})
}

func (m MixedCases) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(m)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(m, func(x *MixedCase, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (m MixedCases) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(m)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(m, func(x *MixedCase, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (m MixedCases) Clone() MixedCases {
	return slices.Clone(m)
}
