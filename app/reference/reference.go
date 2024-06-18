package reference

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

var _ svc.Model = (*Reference)(nil)

type Reference struct {
	ID      uuid.UUID   `json:"id,omitempty"`
	Custom  *foo.Custom `json:"custom,omitempty"`
	Self    *SelfCustom `json:"self,omitempty"`
	Created time.Time   `json:"created,omitempty"`
}

func New(id uuid.UUID) *Reference {
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

func Random() *Reference {
	return &Reference{
		ID:      util.UUID(),
		Custom:  nil,
		Self:    nil,
		Created: util.TimeCurrent(),
	}
}

func (r *Reference) Strings() []string {
	return []string{r.ID.String(), r.Custom.String(), r.Self.String(), util.TimeToFull(&r.Created)}
}

func (r *Reference) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *Reference) WebPath() string {
	return "/reference/" + r.ID.String()
}

func (r *Reference) ToData() []any {
	return []any{r.ID, r.Custom, r.Self, r.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "custom", Title: "Custom", Description: "", Type: "ref:github.com/kyleu/pftest/app/foo/Custom"},
	{Key: "self", Title: "Self", Description: "", Type: "ref:github.com/kyleu/pftest/app/reference/SelfCustom"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
