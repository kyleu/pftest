package seed

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/seed"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Seed)(nil)

type Seed struct {
	ID   uuid.UUID     `json:"id,omitzero"`
	Name string        `json:"name,omitzero"`
	Size int           `json:"size,omitzero"`
	Obj  util.ValueMap `json:"obj,omitzero"`
}

func NewSeed(id uuid.UUID) *Seed {
	return &Seed{ID: id}
}

func (s *Seed) Clone() *Seed {
	return &Seed{ID: s.ID, Name: s.Name, Size: s.Size, Obj: s.Obj.Clone()}
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

func RandomSeed() *Seed {
	return &Seed{
		ID:   util.UUID(),
		Name: util.RandomString(12),
		Size: util.RandomInt(10000),
		Obj:  util.RandomValueMap(4),
	}
}

func (s *Seed) Strings() []string {
	return []string{s.ID.String(), s.Name, fmt.Sprint(s.Size), util.ToJSONCompact(s.Obj)}
}

func (s *Seed) ToCSV() ([]string, [][]string) {
	return SeedFieldDescs.Keys(), [][]string{s.Strings()}
}

func (s *Seed) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(s.ID.String()))...)
}

func (s *Seed) Breadcrumb(extra ...string) string {
	return s.TitleString() + "||" + s.WebPath(extra...) + "**star"
}

func (s *Seed) ToData() []any {
	return []any{s.ID, s.Name, s.Size, s.Obj}
}

var SeedFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "size", Title: "Size", Description: "", Type: "int"},
	{Key: "obj", Title: "Obj", Description: "", Type: "map"},
}
