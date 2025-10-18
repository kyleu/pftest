package trouble

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/troub/le"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Trouble)(nil)

type PK struct {
	From  string   `json:"from,omitzero"`
	Where []string `json:"where,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%s • %v", p.From, p.Where)
}

type Trouble struct {
	From      string     `json:"from,omitzero"`
	Where     []string   `json:"where,omitempty"`
	Selectcol int        `json:"selectcol,omitzero"`
	Limit     string     `json:"limit,omitzero"`
	Group     string     `json:"group,omitzero"`
	Delete    *time.Time `json:"delete,omitzero"`
}

func NewTrouble(from string, where []string) *Trouble {
	return &Trouble{From: from, Where: where}
}

func (t *Trouble) Clone() *Trouble {
	return &Trouble{
		From: t.From, Where: util.ArrayCopy(t.Where), Selectcol: t.Selectcol, Limit: t.Limit, Group: t.Group, Delete: t.Delete,
	}
}

func (t *Trouble) String() string {
	return fmt.Sprintf("%s • %s", t.From, t.Where)
}

func (t *Trouble) TitleString() string {
	return t.String()
}

func (t *Trouble) ToPK() *PK {
	return &PK{
		From:  t.From,
		Where: t.Where,
	}
}

func RandomTrouble() *Trouble {
	return &Trouble{
		From:      util.RandomString(12),
		Where:     []string{util.RandomString(12), util.RandomString(12)},
		Selectcol: util.RandomInt(10000),
		Limit:     util.RandomString(12),
		Group:     util.RandomString(12),
		Delete:    nil,
	}
}

func (t *Trouble) Strings() []string {
	return []string{t.From, util.ToJSONCompact(t.Where), fmt.Sprint(t.Selectcol), t.Limit, t.Group, util.TimeToFull(t.Delete)}
}

func (t *Trouble) ToCSV() ([]string, [][]string) {
	return TroubleFieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *Trouble) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(t.From), url.QueryEscape(util.StringJoin(t.Where, ",")))...)
}

func (t *Trouble) Breadcrumb(extra ...string) string {
	return t.TitleString() + "||" + t.WebPath(extra...) + "**star"
}

func (t *Trouble) ToData() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}

var TroubleFieldDescs = util.FieldDescs{
	{Key: "from", Title: "From", Type: "string"},
	{Key: "where", Title: "Where", Type: "[]string"},
	{Key: "selectcol", Title: "Selectcol", Type: "int"},
	{Key: "limit", Title: "Limit", Type: "string"},
	{Key: "group", Title: "Group", Type: "string"},
	{Key: "delete", Title: "Delete", Type: "timestamp"},
}
