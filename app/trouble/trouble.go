package trouble

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/troub/le"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Trouble)(nil)

type PK struct {
	From  string   `json:"from,omitempty"`
	Where []string `json:"where,omitempty"`
}

func (p *PK) String() string {
	return fmt.Sprintf("%s::%v", p.From, p.Where)
}

type Trouble struct {
	From      string     `json:"from,omitempty"`
	Where     []string   `json:"where,omitempty"`
	Selectcol int        `json:"selectcol,omitempty"`
	Limit     string     `json:"limit,omitempty"`
	Group     string     `json:"group,omitempty"`
	Delete    *time.Time `json:"delete,omitempty"`
}

func NewTrouble(from string, where []string) *Trouble {
	return &Trouble{From: from, Where: where}
}

func (t *Trouble) Clone() *Trouble {
	return &Trouble{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}

func (t *Trouble) String() string {
	return fmt.Sprintf("%s::%s", t.From, t.Where)
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
	return []string{t.From, util.ToJSON(t.Where), fmt.Sprint(t.Selectcol), t.Limit, t.Group, util.TimeToFull(t.Delete)}
}

func (t *Trouble) ToCSV() ([]string, [][]string) {
	return TroubleFieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *Trouble) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(t.From), url.QueryEscape(strings.Join(t.Where, ",")))...)
}

func (t *Trouble) ToData() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}

var TroubleFieldDescs = util.FieldDescs{
	{Key: "from", Title: "From", Description: "", Type: "string"},
	{Key: "where", Title: "Where", Description: "", Type: "[]string"},
	{Key: "selectcol", Title: "Selectcol", Description: "", Type: "int"},
	{Key: "limit", Title: "Limit", Description: "", Type: "string"},
	{Key: "group", Title: "Group", Description: "", Type: "string"},
	{Key: "delete", Title: "Delete", Description: "", Type: "timestamp"},
}
