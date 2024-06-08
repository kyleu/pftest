// Package trouble - Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Trouble, util.ValueMap, error) {
	ret := &Trouble{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "from":
			if setPK {
				ret.From, err = m.ParseString(k, true, true)
			}
		case "where":
			if setPK {
				ret.Where, err = m.ParseArrayString(k, true, true)
				if err != nil {
					return nil, nil, err
				}
			}
		case "selectcol":
			ret.Selectcol, err = m.ParseInt(k, true, true)
		case "limit":
			ret.Limit, err = m.ParseString(k, true, true)
		case "group":
			ret.Group, err = m.ParseString(k, true, true)
		case "delete":
			ret.Delete, err = m.ParseTime(k, true, true)
		default:
			extra[k] = v
		}
		if err != nil {
			return nil, nil, err
		}
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, extra, nil
}
