package softdel

import (
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/softdel"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Softdel)(nil)

type Softdel struct {
	ID      string     `json:"id,omitzero"`
	Created time.Time  `json:"created,omitzero"`
	Updated *time.Time `json:"updated,omitzero"`
	Deleted *time.Time `json:"deleted,omitzero"`
}

func NewSoftdel(id string) *Softdel {
	return &Softdel{ID: id}
}

func (s *Softdel) Clone() *Softdel {
	return &Softdel{ID: s.ID, Created: s.Created, Updated: s.Updated, Deleted: s.Deleted}
}

func (s *Softdel) String() string {
	return s.ID
}

func (s *Softdel) TitleString() string {
	return s.String()
}

func RandomSoftdel() *Softdel {
	return &Softdel{
		ID:      util.RandomString(12),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
		Deleted: nil,
	}
}

func (s *Softdel) Strings() []string {
	return []string{s.ID, util.TimeToFull(&s.Created), util.TimeToFull(s.Updated), util.TimeToFull(s.Deleted)}
}

func (s *Softdel) ToCSV() ([]string, [][]string) {
	return SoftdelFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Softdel) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.ID))...)
}

func (s *Softdel) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**star"
}

func (s *Softdel) ToData() []any {
	return []any{s.ID, s.Created, s.Updated, s.Deleted}
}

var SoftdelFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "string"},
	{Key: "created", Title: "Created", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Type: "timestamp"},
	{Key: "deleted", Title: "Deleted", Type: "timestamp"},
}
