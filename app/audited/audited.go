package audited

import (
	"net/url"
	"path"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/audited"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Audited)(nil)

type Audited struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

func NewAudited(id uuid.UUID) *Audited {
	return &Audited{ID: id}
}

func (a *Audited) Clone() *Audited {
	return &Audited{a.ID, a.Name}
}

func (a *Audited) String() string {
	return a.ID.String()
}

func (a *Audited) TitleString() string {
	if xx := a.Name; xx != "" {
		return xx
	}
	return a.String()
}

func RandomAudited() *Audited {
	return &Audited{
		ID:   util.UUID(),
		Name: util.RandomString(12),
	}
}

func (a *Audited) Strings() []string {
	return []string{a.ID.String(), a.Name}
}

func (a *Audited) ToCSV() ([]string, [][]string) {
	return AuditedFieldDescs.Keys(), [][]string{a.Strings()}
}

func (a *Audited) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(a.ID.String()))...)
}

func (a *Audited) ToData() []any {
	return []any{a.ID, a.Name}
}

var AuditedFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
}
