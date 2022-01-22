// Content managed by Project Forge, see [projectforge.md] for details.
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
	retBirthday, e := m.ParseTime("birthday", true, true)
	if e != nil {
		return nil, e
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

func (c *Capital) Clone() *Capital {
	return &Capital{
		ID:       c.ID,
		Name:     c.Name,
		Birthday: c.Birthday,
		Version:  c.Version,
		Deathday: c.Deathday,
	}
}

func (c *Capital) String() string {
	return c.ID
}

func (c *Capital) WebPath() string {
	return "/capital" + "/" + c.ID
}

func (c *Capital) ToData() []interface{} {
	return []interface{}{c.ID, c.Name, c.Birthday, c.Version, c.Deathday}
}

func (c *Capital) ToDataCore() []interface{} {
	return []interface{}{c.ID, c.Version}
}

func (c *Capital) ToDataVersion() []interface{} {
	return []interface{}{c.ID, c.Version, c.Name, c.Birthday, c.Deathday}
}

type Capitals []*Capital
