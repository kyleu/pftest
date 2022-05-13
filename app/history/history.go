// Content managed by Project Forge, see [projectforge.md] for details.
package history

import (
	"time"

	"golang.org/x/exp/slices"

	"github.com/kyleu/pftest/app/util"
)

type History struct {
	ID      string        `json:"id"`
	Data    util.ValueMap `json:"data"`
	Created time.Time     `json:"created"`
	Updated *time.Time    `json:"updated,omitempty"`
}

func New(id string) *History {
	return &History{ID: id}
}

func Random() *History {
	return &History{
		ID:      util.RandomString(12),
		Data:    util.RandomValueMap(4),
		Created: time.Now(),
		Updated: util.NowPointer(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*History, error) {
	ret := &History{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Data, err = m.ParseMap("data", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (h *History) Clone() *History {
	return &History{
		ID:      h.ID,
		Data:    h.Data.Clone(),
		Created: h.Created,
		Updated: h.Updated,
	}
}

func (h *History) String() string {
	return h.ID
}

func (h *History) TitleString() string {
	return h.String()
}

func (h *History) WebPath() string {
	return "/history" + "/" + h.ID
}

func (h *History) Diff(hx *History) util.Diffs {
	var diffs util.Diffs
	if h.ID != hx.ID {
		diffs = append(diffs, util.NewDiff("id", h.ID, hx.ID))
	}
	diffs = append(diffs, util.DiffObjects(h.Data, hx.Data, "data")...)
	if h.Created != hx.Created {
		diffs = append(diffs, util.NewDiff("created", h.Created.String(), hx.Created.String()))
	}
	return diffs
}

func (h *History) ToData() []any {
	return []any{h.ID, h.Data, h.Created, h.Updated}
}

type Histories []*History

func (h Histories) Get(id string) *History {
	for _, x := range h {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (h Histories) Clone() Histories {
	return slices.Clone(h)
}
