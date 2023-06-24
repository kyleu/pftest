// Content managed by Project Forge, see [projectforge.md] for details.
package reference

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type References []*Reference

func (r References) Get(id uuid.UUID) *Reference {
	for _, x := range r {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (r References) GetByIDs(ids ...uuid.UUID) References {
	var ret References
	for _, x := range r {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (r References) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.ID)
	}
	return ret
}

func (r References) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (r References) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r References) Clone() References {
	return slices.Clone(r)
}
