// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Softdel struct {
	ID      string     `json:"id,omitempty"`
	Created time.Time  `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Deleted *time.Time `json:"deleted,omitempty"`
}

func New(id string) *Softdel {
	return &Softdel{ID: id}
}

func (s *Softdel) Clone() *Softdel {
	return &Softdel{s.ID, s.Created, s.Updated, s.Deleted}
}

func (s *Softdel) String() string {
	return s.ID
}

func (s *Softdel) TitleString() string {
	return s.String()
}

func Random() *Softdel {
	return &Softdel{
		ID:      util.RandomString(12),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
		Deleted: nil,
	}
}

func (s *Softdel) WebPath() string {
	return "/softdel/" + url.QueryEscape(s.ID)
}

func (s *Softdel) ToData() []any {
	return []any{s.ID, s.Created, s.Updated, s.Deleted}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
	{Key: "deleted", Title: "Deleted", Description: "", Type: "timestamp"},
}
