package relation

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/relation"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Relation)(nil)

type Relation struct {
	ID      uuid.UUID `json:"id,omitempty"`
	BasicID uuid.UUID `json:"basicID,omitempty"`
	Name    string    `json:"name,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func NewRelation(id uuid.UUID) *Relation {
	return &Relation{ID: id}
}

func (r *Relation) Clone() *Relation {
	return &Relation{r.ID, r.BasicID, r.Name, r.Created}
}

func (r *Relation) String() string {
	return r.ID.String()
}

func (r *Relation) TitleString() string {
	if xx := r.Name; xx != "" {
		return xx
	}
	return r.String()
}

func RandomRelation() *Relation {
	return &Relation{
		ID:      util.UUID(),
		BasicID: util.UUID(),
		Name:    util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (r *Relation) Strings() []string {
	return []string{r.ID.String(), r.BasicID.String(), r.Name, util.TimeToFull(&r.Created)}
}

func (r *Relation) ToCSV() ([]string, [][]string) {
	return RelationFieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *Relation) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(r.ID.String()))...)
}

func (r *Relation) Breadcrumb(extra ...string) string {
	return r.TitleString() + "||" + r.WebPath(extra...) + "**star"
}

func (r *Relation) ToData() []any {
	return []any{r.ID, r.BasicID, r.Name, r.Created}
}

var RelationFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "basicID", Title: "Basic ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
