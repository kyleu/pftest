// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"fmt"
	"time"

	"golang.org/x/exp/slices"

	"github.com/kyleu/pftest/app/util"
)

type Trouble struct {
	From      string     `json:"from"`
	Where     int        `json:"where"`
	Selectcol int        `json:"selectcol"`
	Limit     string     `json:"limit"`
	Group     string     `json:"group"`
	Delete    *time.Time `json:"delete,omitempty"`
}

func New(from string, where int) *Trouble {
	return &Trouble{From: from, Where: where}
}

func Random() *Trouble {
	return &Trouble{
		From:      util.RandomString(12),
		Where:     util.RandomInt(10000),
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
		ret.Where, err = m.ParseInt("where", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
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
	return &Trouble{
		From:      t.From,
		Where:     t.Where,
		Selectcol: t.Selectcol,
		Limit:     t.Limit,
		Group:     t.Group,
		Delete:    t.Delete,
	}
}

func (t *Trouble) String() string {
	return fmt.Sprintf("%s::%s", t.From, fmt.Sprint(t.Where))
}

func (t *Trouble) TitleString() string {
	return t.String()
}

func (t *Trouble) WebPath() string {
	return "/troub/le" + "/" + t.From + "/" + fmt.Sprint(t.Where)
}

func (t *Trouble) Diff(tx *Trouble) util.Diffs {
	var diffs util.Diffs
	if t.From != tx.From {
		diffs = append(diffs, util.NewDiff("from", t.From, tx.From))
	}
	if t.Where != tx.Where {
		diffs = append(diffs, util.NewDiff("where", fmt.Sprint(t.Where), fmt.Sprint(tx.Where)))
	}
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
		diffs = append(diffs, util.NewDiff("delete", fmt.Sprint(t.Delete), fmt.Sprint(tx.Delete))) // nolint:gocritic // it's nullable
	}
	return diffs
}

func (t *Trouble) ToData() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}

func (t *Trouble) ToDataCore() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Limit, t.Delete}
}

func (t *Trouble) ToDataSelectcol() []any {
	return []any{t.From, t.Where, t.Selectcol, t.Group}
}

type Troubles []*Trouble

func (t Troubles) Get(from string, where int) *Trouble {
	for _, x := range t {
		if x.From == from && x.Where == where {
			return x
		}
	}
	return nil
}

func (t Troubles) Clone() Troubles {
	return slices.Clone(t)
}
