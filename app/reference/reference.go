// Package reference - Content managed by Project Forge, see [projectforge.md] for details.
package reference

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/util"
)

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

func (r *Reference) WebPath() string {
	return "/reference/" + r.ID.String()
}

func (r *Reference) ToData() []any {
	return []any{r.ID, r.Custom, r.Self, r.Created}
}
