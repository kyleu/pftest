package reference

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/reference"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Reference)(nil)

type Reference struct {
	ID      uuid.UUID   `json:"id,omitzero"`
	Custom  *foo.Custom `json:"custom,omitzero"`
	Self    *SelfCustom `json:"self,omitzero"`
	Created time.Time   `json:"created,omitzero"`
}

func NewReference(id uuid.UUID) *Reference {
	return &Reference{ID: id}
}

func (r *Reference) Clone() *Reference {
	return &Reference{r.ID, r.Custom.Clone(), r.Self.Clone(), r.Created}
}

func (r *Reference) String() string {
	return r.ID.String()
}

func (r *Reference) TitleString() string {
	return r.String()
}

func RandomReference() *Reference {
	return &Reference{
		ID:      util.UUID(),
		Custom:  nil,
		Self:    nil,
		Created: util.TimeCurrent(),
	}
}

func (r *Reference) Strings() []string {
	return []string{r.ID.String(), util.ToJSONCompact(r.Custom), util.ToJSONCompact(r.Self), util.TimeToFull(&r.Created)}
}

func (r *Reference) ToCSV() ([]string, [][]string) {
	return ReferenceFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *Reference) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(r.ID.String()))...)
}

func (r *Reference) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**star"
}

func (r *Reference) ToData() []any {
	return []any{r.ID, r.Custom, r.Self, r.Created}
}

var ReferenceFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "custom", Title: "Custom", Description: "", Type: "Custom"},
	{Key: "self", Title: "Self", Description: "", Type: "SelfCustom"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
