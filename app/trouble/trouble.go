package trouble

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

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

func New(from string, where []string) *Trouble {
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

func Random() *Trouble {
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
	return FieldDescs.Keys(), [][]string{t.Strings()}
}

func (t *Trouble) WebPath() string {
	return "/troub/le/" + url.QueryEscape(t.From) + "/" + strings.Join(t.Where, ",")
}

func (t *Trouble) ToData() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}

var FieldDescs = util.FieldDescs{
	{Key: "from", Title: "From", Description: "", Type: "string"},
	{Key: "where", Title: "Where", Description: "", Type: "[]string"},
	{Key: "selectcol", Title: "Selectcol", Description: "", Type: "int"},
	{Key: "limit", Title: "Limit", Description: "", Type: "string"},
	{Key: "group", Title: "Group", Description: "", Type: "string"},
	{Key: "delete", Title: "Delete", Description: "", Type: "timestamp"},
}
