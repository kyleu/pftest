package seed

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

var _ svc.Model = (*Seed)(nil)

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
	if xx := s.Name; xx != "" {
		return xx
	}
	return s.String()
}

func Random() *Seed {
	return &Seed{
		ID:   util.UUID(),
		Name: util.RandomString(12),
		Size: util.RandomInt(10000),
		Obj:  util.RandomValueMap(4),
	}
}

func (s *Seed) Strings() []string {
	return []string{s.ID.String(), s.Name, fmt.Sprint(s.Size), util.ToJSON(s.Obj)}
}

func (s *Seed) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Seed) WebPath() string {
	return "/seed/" + s.ID.String()
}

func (s *Seed) ToData() []any {
	return []any{s.ID, s.Name, s.Size, s.Obj}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "size", Title: "Size", Description: "", Type: "int"},
	{Key: "obj", Title: "Obj", Description: "", Type: "map"},
}
