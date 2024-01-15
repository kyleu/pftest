// Package basic - Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Basic struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Status  string    `json:"status,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(id uuid.UUID) *Basic {
	return &Basic{ID: id}
}

func (b *Basic) Clone() *Basic {
	return &Basic{b.ID, b.Name, b.Status, b.Created}
}

func (b *Basic) String() string {
	return b.ID.String()
}

func (b *Basic) TitleString() string {
	return b.Name
}

func Random() *Basic {
	return &Basic{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Status:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (b *Basic) WebPath() string {
	return "/basic/" + b.ID.String()
}

func (b *Basic) ToData() []any {
	return []any{b.ID, b.Name, b.Status, b.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
