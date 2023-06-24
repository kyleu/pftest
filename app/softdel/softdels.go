// Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Softdels []*Softdel

func (s Softdels) Get(id string) *Softdel {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Softdels) GetByIDs(ids ...string) Softdels {
	var ret Softdels
	for _, x := range s {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s Softdels) IDs() []string {
	ret := make([]string, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.ID)
	}
	return ret
}

func (s Softdels) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.ID)
	}
	return ret
}

func (s Softdels) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s Softdels) Clone() Softdels {
	return slices.Clone(s)
}
