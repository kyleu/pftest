// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"fmt"
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

func (c *Capital) Diff(cx *Capital) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID, cx.ID))
	}
	if c.Name != cx.Name {
		diffs = append(diffs, util.NewDiff("name", c.Name, cx.Name))
	}
	if c.Birthday != cx.Birthday {
		diffs = append(diffs, util.NewDiff("birthday", c.Birthday.String(), cx.Birthday.String()))
	}
	if c.Version != cx.Version {
		diffs = append(diffs, util.NewDiff("version", fmt.Sprint(c.Version), fmt.Sprint(cx.Version)))
	}
	if (c.Deathday == nil && cx.Deathday != nil) || (c.Deathday != nil && cx.Deathday == nil) || (c.Deathday != nil && cx.Deathday != nil && *c.Deathday != *cx.Deathday) {
		diffs = append(diffs, util.NewDiff("deathday", fmt.Sprint(c.Deathday), fmt.Sprint(cx.Deathday))) // nolint:gocritic // it's nullable
	}
	return diffs
}

func (c *Capital) ToData() []any {
	return []any{c.ID, c.Name, c.Birthday, c.Version, c.Deathday}
}

func (c *Capital) ToDataCore() []any {
	return []any{c.ID, c.Version}
}

func (c *Capital) ToDataVersion() []any {
	return []any{c.ID, c.Version, c.Name, c.Birthday, c.Deathday}
}

type Capitals []*Capital
