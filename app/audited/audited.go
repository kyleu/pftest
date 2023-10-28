// Package audited - Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Audited struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

func New(id uuid.UUID) *Audited {
	return &Audited{ID: id}
}

func (a *Audited) Clone() *Audited {
	return &Audited{a.ID, a.Name}
}

func (a *Audited) String() string {
	return a.ID.String()
}

func (a *Audited) TitleString() string {
	return a.Name
}

func Random() *Audited {
	return &Audited{
		ID:   util.UUID(),
		Name: util.RandomString(12),
	}
}

func (a *Audited) WebPath() string {
	return "/audited/" + a.ID.String()
}

func (a *Audited) ToData() []any {
	return []any{a.ID, a.Name}
}
