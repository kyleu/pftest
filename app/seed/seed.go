// Package seed - Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Seed struct {
	ID   uuid.UUID     `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
	Size int           `json:"size,omitempty"`
	Obj  util.ValueMap `json:"obj,omitempty"`
}

func New(id uuid.UUID) *Seed {
	return &Seed{ID: id}
}

func (s *Seed) Clone() *Seed {
	return &Seed{s.ID, s.Name, s.Size, s.Obj.Clone()}
}

func (s *Seed) String() string {
	return s.ID.String()
}

func (s *Seed) TitleString() string {
	return s.Name
}

func Random() *Seed {
	return &Seed{
		ID:   util.UUID(),
		Name: util.RandomString(12),
		Size: util.RandomInt(10000),
		Obj:  util.RandomValueMap(4),
	}
}

func (s *Seed) WebPath() string {
	return "/seed/" + s.ID.String()
}

func (s *Seed) ToData() []any {
	return []any{s.ID, s.Name, s.Size, s.Obj}
}
