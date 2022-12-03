// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"golang.org/x/exp/slices"

	"github.com/kyleu/pftest/app/util"
)

type Troubles []*Trouble

func (t Troubles) Get(from string, where []string) *Trouble {
	for _, x := range t {
		if x.From == from && slices.Equal(x.Where, where) {
			return x
		}
	}
	return nil
}

func (t Troubles) GetByFroms(froms ...string) Troubles {
	var ret Troubles
	for _, x := range t {
		if slices.Contains(froms, x.From) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (t Troubles) Froms() []string {
	ret := make([]string, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.From)
	}
	return ret
}

func (t Troubles) FromStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.From)
	}
	return ret
}

func (t Troubles) Wheres() [][]string {
	ret := make([][]string, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.Where)
	}
	return ret
}

func (t Troubles) WhereStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, util.ToJSON(&x.Where))
	}
	return ret
}

func (t Troubles) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range t {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (t Troubles) Clone() Troubles {
	return slices.Clone(t)
}
