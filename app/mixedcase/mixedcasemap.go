// Package mixedcase - Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import "github.com/kyleu/pftest/app/util"

func FromMap(m util.ValueMap, setPK bool) (*MixedCase, util.ValueMap, error) {
	ret := &MixedCase{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				ret.ID, err = m.ParseString(k, true, true)
			}
		case "testField":
			ret.TestField, err = m.ParseString(k, true, true)
		case "anotherField":
			ret.AnotherField, err = m.ParseString(k, true, true)
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
