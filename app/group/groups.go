// Content managed by Project Forge, see [projectforge.md] for details.
package group

import "golang.org/x/exp/slices"

type Groups []*Group

func (g Groups) Get(id string) *Group {
	for _, x := range g {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (g Groups) IDs() []string {
	ret := make([]string, 0, len(g)+1)
	for _, x := range g {
		ret = append(ret, x.ID)
	}
	return ret
}

func (g Groups) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(g)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range g {
		ret = append(ret, x.ID)
	}
	return ret
}

func (g Groups) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(g)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range g {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (g Groups) Clone() Groups {
	return slices.Clone(g)
}
