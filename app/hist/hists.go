// Content managed by Project Forge, see [projectforge.md] for details.
package hist

import "golang.org/x/exp/slices"

type Hists []*Hist

func (h Hists) Get(id string) *Hist {
	for _, x := range h {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (h Hists) GetByIDs(ids ...string) Hists {
	var ret Hists
	for _, x := range h {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (h Hists) IDs() []string {
	ret := make([]string, 0, len(h)+1)
	for _, x := range h {
		ret = append(ret, x.ID)
	}
	return ret
}

func (h Hists) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(h)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range h {
		ret = append(ret, x.ID)
	}
	return ret
}

func (h Hists) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(h)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range h {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (h Hists) Clone() Hists {
	return slices.Clone(h)
}
