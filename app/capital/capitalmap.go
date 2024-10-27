package capital

import "github.com/kyleu/pftest/app/util"

func (c *Capital) ToMap() util.ValueMap {
	return util.ValueMap{"id": c.ID, "name": c.Name, "birthday": c.Birthday, "deathday": c.Deathday}
}

func CapitalFromMap(m util.ValueMap, setPK bool) (*Capital, util.ValueMap, error) {
	ret := &Capital{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				ret.ID, err = m.ParseString(k, true, true)
			}
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		case "birthday":
			retBirthday, e := m.ParseTime(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retBirthday != nil {
				ret.Birthday = *retBirthday
			}
		case "deathday":
			ret.Deathday, err = m.ParseTime(k, true, true)
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
