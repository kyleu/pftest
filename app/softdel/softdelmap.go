// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Softdel, error) {
	ret := &Softdel{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Deleted, err = m.ParseTime("deleted", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
