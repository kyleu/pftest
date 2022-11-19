// Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Basics []*Basic

func (b Basics) Get(id uuid.UUID) *Basic {
	for _, x := range b {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (b Basics) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(b)+1)
	for _, x := range b {
		ret = append(ret, x.ID)
	}
	return ret
}

func (b Basics) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(b)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range b {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (b Basics) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(b)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range b {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (b Basics) Clone() Basics {
	return slices.Clone(b)
}
