// Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Seeds []*Seed

func (s Seeds) Get(id uuid.UUID) *Seed {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Seeds) GetByIDs(ids ...uuid.UUID) Seeds {
	var ret Seeds
	for _, x := range s {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s Seeds) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.ID)
	}
	return ret
}

func (s Seeds) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (s Seeds) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s Seeds) Clone() Seeds {
	return slices.Clone(s)
}
