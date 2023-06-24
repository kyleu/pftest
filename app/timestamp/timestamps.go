// Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Timestamps []*Timestamp

func (t Timestamps) Get(id string) *Timestamp {
	for _, x := range t {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (t Timestamps) GetByIDs(ids ...string) Timestamps {
	var ret Timestamps
	for _, x := range t {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (t Timestamps) IDs() []string {
	ret := make([]string, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.ID)
	}
	return ret
}

func (t Timestamps) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.ID)
	}
	return ret
}

func (t Timestamps) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range t {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (t Timestamps) Clone() Timestamps {
	return slices.Clone(t)
}
