// Package trouble - Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Trouble, error) {
	ret := &Trouble{}
	var err error
	if setPK {
		ret.From, err = m.ParseString("from", true, true)
		if err != nil {
			return nil, err
		}
		ret.Where, err = m.ParseArrayString("where", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Selectcol, err = m.ParseInt("selectcol", true, true)
	if err != nil {
		return nil, err
	}
	ret.Limit, err = m.ParseString("limit", true, true)
	if err != nil {
		return nil, err
	}
	ret.Group, err = m.ParseString("group", true, true)
	if err != nil {
		return nil, err
	}
	ret.Delete, err = m.ParseTime("delete", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
