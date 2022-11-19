// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import "golang.org/x/exp/slices"

type Capitals []*Capital

func (c Capitals) Get(id string) *Capital {
	for _, x := range c {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (c Capitals) IDs() []string {
	ret := make([]string, 0, len(c)+1)
	for _, x := range c {
		ret = append(ret, x.ID)
	}
	return ret
}

func (c Capitals) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(c)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range c {
		ret = append(ret, x.ID)
	}
	return ret
}

func (c Capitals) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(c)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range c {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (c Capitals) Clone() Capitals {
	return slices.Clone(c)
}
