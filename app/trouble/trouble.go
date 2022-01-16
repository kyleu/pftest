package trouble

import (
	"fmt"
	"time"

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
		Delete:    util.NowPointer(),
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

func (t *Trouble) String() string {
	return fmt.Sprintf("%s::%s", t.From, fmt.Sprint(t.Where))
}

func (t *Trouble) WebPath() string {
	return "/trouble" + "/" + t.From + "/" + fmt.Sprint(t.Where)
}

func (t *Trouble) ToData() []interface{} {
	return []interface{}{t.From, t.Where, t.Selectcol, t.Limit, t.Group, t.Delete}
}

func (t *Trouble) ToDataCore() []interface{} {
	return []interface{}{t.From, t.Where, t.Selectcol, t.Limit, t.Delete}
}

func (t *Trouble) ToDataSelectcol() []interface{} {
	return []interface{}{t.From, t.Where, t.Selectcol, t.Group}
}

type Troubles []*Trouble
