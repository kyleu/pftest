// Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Auditeds []*Audited

func (a Auditeds) Get(id uuid.UUID) *Audited {
	for _, x := range a {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (a Auditeds) GetByIDs(ids ...uuid.UUID) Auditeds {
	var ret Auditeds
	for _, x := range a {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (a Auditeds) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(a)+1)
	for _, x := range a {
		ret = append(ret, x.ID)
	}
	return ret
}

func (a Auditeds) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(a)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range a {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (a Auditeds) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(a)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range a {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (a Auditeds) Clone() Auditeds {
	return slices.Clone(a)
}
