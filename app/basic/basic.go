package basic

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/basic"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Basic)(nil)

type Basic struct {
	ID      uuid.UUID `json:"id,omitzero"`
	Name    string    `json:"name,omitzero"`
	Status  string    `json:"status,omitzero"`
	Created time.Time `json:"created,omitzero"`
}

func NewBasic(id uuid.UUID) *Basic {
	return &Basic{ID: id}
}

func (b *Basic) Clone() *Basic {
	return &Basic{ID: b.ID, Name: b.Name, Status: b.Status, Created: b.Created}
}

func (b *Basic) String() string {
	return b.ID.String()
}

func (b *Basic) TitleString() string {
	if xx := b.Name; xx != "" {
		return xx
	}
	return b.String()
}

func RandomBasic() *Basic {
	return &Basic{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Status:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (b *Basic) Strings() []string {
	return []string{b.ID.String(), b.Name, b.Status, util.TimeToFull(&b.Created)}
}

func (b *Basic) ToCSV() ([]string, [][]string) {
	return BasicFieldDescs.Keys(), [][]string{b.Strings()}
}

func (b *Basic) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(b.ID.String()))...)
}

func (b *Basic) Breadcrumb(extra ...string) string {
	return b.TitleString() + "||" + b.WebPath(extra...) + "**star"
}

func (b *Basic) ToData() []any {
	return []any{b.ID, b.Name, b.Status, b.Created}
}

var BasicFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
