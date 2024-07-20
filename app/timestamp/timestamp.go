package timestamp

import (
	"net/url"
	"path"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/timestamp"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Timestamp)(nil)

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

func (t *Timestamp) Strings() []string {
	return []string{t.ID, util.TimeToFull(&t.Created), util.TimeToFull(t.Updated), util.TimeToFull(t.Deleted)}
}

func (t *Timestamp) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *Timestamp) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(t.ID))...)
}

func (t *Timestamp) ToData() []any {
	return []any{t.ID, t.Created, t.Updated, t.Deleted}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
	{Key: "updated", Title: "Updated", Description: "", Type: "timestamp"},
	{Key: "deleted", Title: "Deleted", Description: "", Type: "timestamp"},
}
