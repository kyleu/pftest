// Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import "golang.org/x/exp/slices"

type MixedCases []*MixedCase

func (m MixedCases) Get(id string) *MixedCase {
	for _, x := range m {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (m MixedCases) IDs() []string {
	ret := make([]string, 0, len(m)+1)
	for _, x := range m {
		ret = append(ret, x.ID)
	}
	return ret
}

func (m MixedCases) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(m)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range m {
		ret = append(ret, x.ID)
	}
	return ret
}

func (m MixedCases) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(m)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range m {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (m MixedCases) Clone() MixedCases {
	return slices.Clone(m)
}
