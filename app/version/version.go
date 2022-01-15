package version

import (
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Version struct {
	ID       string        `json:"id"`
	Revision int           `json:"revision"`
	Const    string        `json:"const"`
	Var      util.ValueMap `json:"var"`
	Created  time.Time     `json:"created"`
	Updated  *time.Time    `json:"updated,omitempty"`
	Deleted  *time.Time    `json:"deleted,omitempty"`
}

func New(id string) *Version {
	return &Version{ID: id}
}

func FromMap(m util.ValueMap, setPK bool) (*Version, error) {
	ret := &Version{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Const, err = m.ParseString("const", true, true)
	if err != nil {
		return nil, err
	}
	ret.Var, err = m.ParseMap("var", true, true)
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

func (v *Version) String() string {
	return v.ID
}

func (v *Version) WebPath() string {
	return "/version" + "/" + v.ID
}

func (v *Version) ToData() []interface{} {
	return []interface{}{v.ID, v.Revision, v.Const, v.Var, v.Created, v.Updated, v.Deleted}
}

func (v *Version) ToDataCore() []interface{} {
	return []interface{}{v.ID, v.Revision, v.Const, v.Updated, v.Deleted}
}

func (v *Version) ToDataRevision() []interface{} {
	return []interface{}{v.ID, v.Revision, v.Var, v.Created}
}

type Versions []*Version
