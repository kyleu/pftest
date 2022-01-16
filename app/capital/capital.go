package capital

import (
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Capital struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Birthday time.Time  `json:"birthday"`
	Version  int        `json:"version"`
	Deathday *time.Time `json:"deathday,omitempty"`
}

func New(id string) *Capital {
	return &Capital{ID: id}
}

func Random() *Capital {
	return &Capital{
		ID:       util.RandomString(12),
		Name:     util.RandomString(12),
		Birthday: time.Now(),
		Version:  util.RandomInt(10000),
		Deathday: util.NowPointer(),
	}
}

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
	retBirthday, err := m.ParseTime("birthday", true, true)
	if err != nil {
		return nil, err
	}
	ret.Birthday = *retBirthday
	ret.Deathday, err = m.ParseTime("deathday", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (C *Capital) String() string {
	return C.ID
}

func (C *Capital) WebPath() string {
	return "/capital" + "/" + C.ID
}

func (C *Capital) ToData() []interface{} {
	return []interface{}{C.ID, C.Name, C.Birthday, C.Version, C.Deathday}
}

func (C *Capital) ToDataCore() []interface{} {
	return []interface{}{C.ID, C.Version}
}

func (C *Capital) ToDataVersion() []interface{} {
	return []interface{}{C.ID, C.Version, C.Name, C.Birthday, C.Deathday}
}

type Capitals []*Capital
