// Package relation - Content managed by Project Forge, see [projectforge.md] for details.
package relation

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Relation struct {
	ID      uuid.UUID `json:"id,omitempty"`
	BasicID uuid.UUID `json:"basicID,omitempty"`
	Name    string    `json:"name,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(id uuid.UUID) *Relation {
	return &Relation{ID: id}
}

func (r *Relation) Clone() *Relation {
	return &Relation{r.ID, r.BasicID, r.Name, r.Created}
}

func (r *Relation) String() string {
	return r.ID.String()
}

func (r *Relation) TitleString() string {
	return r.Name
}

func Random() *Relation {
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
	return FieldDescs.Keys(), [][]string{r.Strings()}
}

func (r *Relation) WebPath() string {
	return "/relation/" + r.ID.String()
}

func (r *Relation) ToData() []any {
	return []any{r.ID, r.BasicID, r.Name, r.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "basicID", Title: "Basic ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
