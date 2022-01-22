// Content managed by Project Forge, see [projectforge.md] for details.
package history

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type HistoryHistory struct {
	ID        uuid.UUID     `json:"id"`
	HistoryID string        `json:"historyID"`
	Old       util.ValueMap `json:"o,omitempty"`
	New       util.ValueMap `json:"n,omitempty"`
	Changes   util.Diffs    `json:"c,omitempty"`
	Created   time.Time     `json:"created"`
}

func (h *HistoryHistory) ToData() []interface{} {
	return []interface{}{
		h.ID,
		h.HistoryID,
		util.ToJSONBytes(h.Old, true),
		util.ToJSONBytes(h.New, true),
		util.ToJSONBytes(h.Changes, true),
		h.Created,
	}
}

type historyDTO struct {
	ID        uuid.UUID       `db:"id"`
	HistoryID string          `db:"history_id"`
	Old       json.RawMessage `db:"o"`
	New       json.RawMessage `db:"n"`
	Changes   json.RawMessage `db:"c"`
	Created   time.Time       `db:"created"`
}

func (h *historyDTO) ToHistory() *HistoryHistory {
	o := util.ValueMap{}
	_ = util.FromJSON(h.Old, &o)
	n := util.ValueMap{}
	_ = util.FromJSON(h.New, &n)
	c := util.Diffs{}
	_ = util.FromJSON(h.Changes, &c)
	return &HistoryHistory{ID: h.ID, HistoryID: h.HistoryID, Old: o, New: n, Changes: c, Created: h.Created}
}
