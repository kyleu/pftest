// Content managed by Project Forge, see [projectforge.md] for details.
package hist

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type History struct {
	ID     uuid.UUID     `json:"id"`
	HistID string        `json:"histID"`
	Old    util.ValueMap `json:"o,omitempty"`
	New    util.ValueMap `json:"n,omitempty"`
	Changes util.Diffs    `json:"c,omitempty"`
	Created time.Time     `json:"created"`
}

func (h *History) ToData() []any {
	return []any{
		h.ID,
		h.HistID,
		util.ToJSONBytes(h.Old, true),
		util.ToJSONBytes(h.New, true),
		util.ToJSONBytes(h.Changes, true),
		h.Created,
	}
}

type Histories []*History

type historyDTO struct {
	ID     uuid.UUID       `db:"id"`
	HistID string          `db:"hist_id"`
	Old    json.RawMessage `db:"o"`
	New    json.RawMessage `db:"n"`
	Changes json.RawMessage `db:"c"`
	Created time.Time       `db:"created"`
}

func (h *historyDTO) ToHistory() *History {
	o := util.ValueMap{}
	_ = util.FromJSON(h.Old, &o)
	n := util.ValueMap{}
	_ = util.FromJSON(h.New, &n)
	c := util.Diffs{}
	_ = util.FromJSON(h.Changes, &c)
	return &History{ID: h.ID, HistID: h.HistID, Old: o, New: n, Changes: c, Created: h.Created}
}

type historyDTOs []*historyDTO

func (h historyDTOs) ToHistories() Histories {
	ret := make(Histories, 0, len(h))
	for _, x := range h {
		ret = append(ret, x.ToHistory())
	}
	return ret
}
