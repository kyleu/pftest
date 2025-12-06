package oddrel

import (
	"net/url"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/oddpk/oddrel"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Oddrel)(nil)

type Oddrel struct {
	ID      uuid.UUID `json:"id,omitzero"`
	Project uuid.UUID `json:"project,omitzero"`
	Path    string    `json:"path,omitzero"`
}

func NewOddrel(id uuid.UUID) *Oddrel {
	return &Oddrel{ID: id}
}

func (o *Oddrel) Clone() *Oddrel {
	if o == nil {
		return nil
	}
	return &Oddrel{ID: o.ID, Project: o.Project, Path: o.Path}
}

func (o *Oddrel) String() string {
	return o.ID.String()
}

func (o *Oddrel) TitleString() string {
	return o.String()
}

func RandomOddrel() *Oddrel {
	return &Oddrel{
		ID:      util.UUID(),
		Project: util.UUID(),
		Path:    util.RandomString(12),
	}
}

func (o *Oddrel) Strings() []string {
	return []string{o.ID.String(), o.Project.String(), o.Path}
}

func (o *Oddrel) ToCSV() ([]string, [][]string) {
	return OddrelFieldDescs.Keys(), [][]string{o.Strings()}
}

func (o *Oddrel) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(o.ID.String()))...)
}

func (o *Oddrel) Breadcrumb(extra ...string) string {
	return o.TitleString() + "||" + o.WebPath(extra...) + "**star"
}

func (o *Oddrel) ToData() []any {
	return []any{o.ID, o.Project, o.Path}
}

var OddrelFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Type: "uuid"},
	{Key: "project", Title: "Project", Type: "uuid"},
	{Key: "path", Title: "Path", Type: "string"},
}
