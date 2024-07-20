package path

import (
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/g1/g2/path"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Path)(nil)

type Path struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Status  string    `json:"status,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(id uuid.UUID) *Path {
	return &Path{ID: id}
}

func (p *Path) Clone() *Path {
	return &Path{p.ID, p.Name, p.Status, p.Created}
}

func (p *Path) String() string {
	return p.ID.String()
}

func (p *Path) TitleString() string {
	if xx := p.Name; xx != "" {
		return xx
	}
	return p.String()
}

func Random() *Path {
	return &Path{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Status:  util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (p *Path) Strings() []string {
	return []string{p.ID.String(), p.Name, p.Status, util.TimeToFull(&p.Created)}
}

func (p *Path) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{p.Strings()}
}

func (p *Path) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(p.ID.String()))...)
}

func (p *Path) ToData() []any {
	return []any{p.ID, p.Name, p.Status, p.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "status", Title: "Status", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
