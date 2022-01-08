package group

import (
	"time"

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

func FromMap(m util.ValueMap, setPK bool) (*Group, error) {
	ret := &Group{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
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
	return ret, nil
}

func (g *Group) String() string {
	return g.ID
}

func (g *Group) WebPath() string {
	return "/group" + "/" + g.ID
}

func (g *Group) ToData() []interface{} {
	return []interface{}{g.ID, g.Group, g.Data, g.Created, g.Updated, g.Deleted}
}

type Groups []*Group
