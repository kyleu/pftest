// Package trouble - Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

type PK struct {
	From  string   `json:"from,omitempty"`
	Where []string `json:"where,omitempty"`
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

func FromMap(m util.ValueMap, setPK bool) (*Trouble, error) {
	ret := &Trouble{}
	var err error
	if setPK {
		ret.From, err = m.ParseString("from", true, true)
		if err != nil {
			return nil, err
		}
		ret.Where, err = m.ParseArrayString("where", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Selectcol, err = m.ParseInt("selectcol", true, true)
	if err != nil {
		return nil, err
	}
	ret.Limit, err = m.ParseString("limit", true, true)
	if err != nil {
		return nil, err
	}
	ret.Group, err = m.ParseString("group", true, true)
	if err != nil {
		return nil, err
	}
	ret.Delete, err = m.ParseTime("delete", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
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

func (t *Trouble) WebPath() string {
	return "/troub/le/" + url.QueryEscape(t.From) + "/" + strings.Join(t.Where, ",")
}

func (t *Trouble) Diff(tx *Trouble) util.Diffs {
	var diffs util.Diffs
	if t.From != tx.From {
		diffs = append(diffs, util.NewDiff("from", t.From, tx.From))
	}
	diffs = append(diffs, util.DiffObjects(t.Where, tx.Where, "where")...)
	if t.Selectcol != tx.Selectcol {
		diffs = append(diffs, util.NewDiff("selectcol", fmt.Sprint(t.Selectcol), fmt.Sprint(tx.Selectcol)))
	}
	if t.Limit != tx.Limit {
		diffs = append(diffs, util.NewDiff("limit", t.Limit, tx.Limit))
	}
	if t.Group != tx.Group {
		diffs = append(diffs, util.NewDiff("group", t.Group, tx.Group))
	}
	if (t.Delete == nil && tx.Delete != nil) || (t.Delete != nil && tx.Delete == nil) || (t.Delete != nil && tx.Delete != nil && *t.Delete != *tx.Delete) {
		diffs = append(diffs, util.NewDiff("delete", fmt.Sprint(t.Delete), fmt.Sprint(tx.Delete))) //nolint:gocritic // it's nullable
	}
	return diffs
}

func (t *Trouble) ToData() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}
