// Package timestamp - Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Timestamp struct {
	ID      string     `json:"id,omitempty"`
	Created time.Time  `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Deleted *time.Time `json:"deleted,omitempty"`
}

func New(id string) *Timestamp {
	return &Timestamp{ID: id}
}

func (t *Timestamp) Clone() *Timestamp {
	return &Timestamp{t.ID, t.Created, t.Updated, t.Deleted}
}

func (t *Timestamp) String() string {
	return t.ID
}

func (t *Timestamp) TitleString() string {
	return t.String()
}

func Random() *Timestamp {
	return &Timestamp{
		ID:      util.RandomString(12),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
		Deleted: nil,
	}
}

func (t *Timestamp) WebPath() string {
	return "/timestamp/" + url.QueryEscape(t.ID)
}

func (t *Timestamp) ToData() []any {
	return []any{t.ID, t.Created, t.Updated, t.Deleted}
}
