// Package reference - Content managed by Project Forge, see [projectforge.md] for details.
package reference

import (
	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/util"
)

func FromMap(m util.ValueMap, setPK bool) (*Reference, error) {
	ret := &Reference{}
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
	tmpCustom, err := m.ParseString("custom", true, true)
	if err != nil {
		return nil, err
	}
	customArg := &foo.Custom{}
	err = util.FromJSON([]byte(tmpCustom), customArg)
	if err != nil {
		return nil, err
	}
	ret.Custom = customArg
	tmpSelf, err := m.ParseString("self", true, true)
	if err != nil {
		return nil, err
	}
	selfArg := &SelfCustom{}
	err = util.FromJSON([]byte(tmpSelf), selfArg)
	if err != nil {
		return nil, err
	}
	ret.Self = selfArg
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}
