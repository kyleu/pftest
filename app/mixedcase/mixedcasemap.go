// Package mixedcase - Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*MixedCase, error) {
	ret := &MixedCase{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.TestField, err = m.ParseString("testField", true, true)
	if err != nil {
		return nil, err
	}
	ret.AnotherField, err = m.ParseString("anotherField", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
