// Content managed by Project Forge, see [projectforge.md] for details.
package relation

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Relations []*Relation

func (r Relations) Get(id uuid.UUID) *Relation {
	for _, x := range r {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (r Relations) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.ID)
	}
	return ret
}

func (r Relations) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (r Relations) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r Relations) Clone() Relations {
	return slices.Clone(r)
}
