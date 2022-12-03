// Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Paths []*Path

func (p Paths) Get(id uuid.UUID) *Path {
	for _, x := range p {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (p Paths) GetByIDs(ids ...uuid.UUID) Paths {
	var ret Paths
	for _, x := range p {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (p Paths) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(p)+1)
	for _, x := range p {
		ret = append(ret, x.ID)
	}
	return ret
}

func (p Paths) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(p)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range p {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (p Paths) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(p)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range p {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (p Paths) Clone() Paths {
	return slices.Clone(p)
}
