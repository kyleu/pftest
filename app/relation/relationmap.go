// Package relation - Content managed by Project Forge, see [projectforge.md] for details.
package relation

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Relation, error) {
	ret := &Relation{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retBasicID, e := m.ParseUUID("basicID", true, true)
	if e != nil {
		return nil, e
	}
	if retBasicID != nil {
		ret.BasicID = *retBasicID
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
