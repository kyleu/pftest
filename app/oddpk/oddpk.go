package oddpk

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/oddpk"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*OddPK)(nil)

type PK struct {
	Project uuid.UUID `json:"project,omitzero"`
	Path    string    `json:"path,omitzero"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%v • %s", p.Project, p.Path)
}

type OddPK struct {
	Project uuid.UUID `json:"project,omitzero"`
	Path    string    `json:"path,omitzero"`
	Name    string    `json:"name,omitzero"`
}

func NewOddPK(project uuid.UUID, path string) *OddPK {
	return &OddPK{Project: project, Path: path}
}

func (o *OddPK) Clone() *OddPK {
	return &OddPK{Project: o.Project, Path: o.Path, Name: o.Name}
}

func (o *OddPK) String() string {
	return fmt.Sprintf("%s • %s", o.Project.String(), o.Path)
}

func (o *OddPK) TitleString() string {
	if xx := o.Name; xx != "" {
		return xx
	}
	return o.String()
}

func (o *OddPK) ToPK() *PK {
	return &PK{
		Project: o.Project,
		Path:    o.Path,
	}
}

func RandomOddPK() *OddPK {
	return &OddPK{
		Project: util.UUID(),
		Path:    util.RandomString(12),
		Name:    util.RandomString(12),
	}
}

func (o *OddPK) Strings() []string {
	return []string{o.Project.String(), o.Path, o.Name}
}

func (o *OddPK) ToCSV() ([]string, [][]string) {
	return OddPKFieldDescs.Keys(), [][]string{o.Strings()}
}

func (o *OddPK) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(o.Project.String()), url.QueryEscape(strings.ReplaceAll(o.Path, "/", "||")))...)
}

func (o *OddPK) Breadcrumb(extra ...string) string {
	return o.TitleString() + "||" + o.WebPath(extra...) + "**star"
}

func (o *OddPK) ToData() []any {
	return []any{o.Project, o.Path, o.Name}
}

var OddPKFieldDescs = util.FieldDescs{
	{Key: "project", Title: "Project", Type: "uuid"},
	{Key: "path", Title: "Path", Type: "string"},
	{Key: "name", Title: "Name", Type: "string"},
}
