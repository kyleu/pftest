// Package capital - Content managed by Project Forge, see [projectforge.md] for details.
package capital

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*Capital, error) {
	ret := &Capital{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	retBirthday, e := m.ParseTime("birthday", true, true)
	if e != nil {
		return nil, e
	}
	if retBirthday != nil {
		ret.Birthday = *retBirthday
	}
	ret.Deathday, err = m.ParseTime("deathday", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
