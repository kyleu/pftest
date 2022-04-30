// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"fmt"
	"time"

	"golang.org/x/exp/slices"

	"github.com/kyleu/pftest/app/util"
)

type Group struct {
	ID      string        `json:"id"`
	Group   string        `json:"group"`
	Data    util.ValueMap `json:"data"`
	Created time.Time     `json:"created"`
	Updated *time.Time    `json:"updated,omitempty"`
	Deleted *time.Time    `json:"deleted,omitempty"`
}

func New(id string) *Group {
	return &Group{ID: id}
}

func Random() *Group {
	return &Group{
		ID:      util.RandomString(12),
		Group:   util.RandomString(12),
		Data:    util.RandomValueMap(4),
		Created: time.Now(),
		Updated: util.NowPointer(),
		Deleted: nil,
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Group, error) {
	ret := &Group{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Group, err = m.ParseString("group", true, true)
	if err != nil {
		return nil, err
	}
	ret.Data, err = m.ParseMap("data", true, true)
	if err != nil {
		return nil, err
	}
	ret.Deleted, err = m.ParseTime("deleted", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (g *Group) Clone() *Group {
	return &Group{
		ID:      g.ID,
		Group:   g.Group,
		Data:    g.Data.Clone(),
		Created: g.Created,
		Updated: g.Updated,
		Deleted: g.Deleted,
	}
}

func (g *Group) String() string {
	return g.ID
}

func (g *Group) WebPath() string {
	return "/group" + "/" + g.ID
}

func (g *Group) Diff(gx *Group) util.Diffs {
	var diffs util.Diffs
	if g.ID != gx.ID {
		diffs = append(diffs, util.NewDiff("id", g.ID, gx.ID))
	}
	if g.Group != gx.Group {
		diffs = append(diffs, util.NewDiff("group", g.Group, gx.Group))
	}
	diffs = append(diffs, util.DiffObjects(g.Data, gx.Data, "data")...)
	if g.Created != gx.Created {
		diffs = append(diffs, util.NewDiff("created", g.Created.String(), gx.Created.String()))
	}
	if (g.Deleted == nil && gx.Deleted != nil) || (g.Deleted != nil && gx.Deleted == nil) || (g.Deleted != nil && gx.Deleted != nil && *g.Deleted != *gx.Deleted) {
		diffs = append(diffs, util.NewDiff("deleted", fmt.Sprint(g.Deleted), fmt.Sprint(gx.Deleted))) // nolint:gocritic // it's nullable
	}
	return diffs
}

func (g *Group) ToData() []any {
	return []any{g.ID, g.Group, g.Data, g.Created, g.Updated, g.Deleted}
}

type Groups []*Group

func (g Groups) Clone() Groups {
	return slices.Clone(g)
}
