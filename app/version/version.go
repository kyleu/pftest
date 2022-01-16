package version

import (
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Version struct {
	ID       string        `json:"id"`
	Revision int           `json:"revision"`
	Constcol string        `json:"constcol"`
	Varcol   util.ValueMap `json:"varcol"`
	Created  time.Time     `json:"created"`
	Updated  *time.Time    `json:"updated,omitempty"`
	Deleted  *time.Time    `json:"deleted,omitempty"`
}

func New(id string) *Version {
	return &Version{ID: id}
}

func Random() *Version {
	return &Version{
		ID:       util.RandomString(12),
		Revision: util.RandomInt(10000),
		Constcol: util.RandomString(12),
		Varcol:   util.RandomValueMap(4),
		Created:  time.Now(),
		Updated:  util.NowPointer(),
		Deleted:  util.NowPointer(),
	}
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
	ret.Constcol, err = m.ParseString("constcol", true, true)
	if err != nil {
		return nil, err
	}
	ret.Varcol, err = m.ParseMap("varcol", true, true)
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
	return []interface{}{v.ID, v.Revision, v.Constcol, v.Varcol, v.Created, v.Updated, v.Deleted}
}

func (v *Version) ToDataCore() []interface{} {
	return []interface{}{v.ID, v.Revision, v.Constcol, v.Updated, v.Deleted}
}

func (v *Version) ToDataRevision() []interface{} {
	return []interface{}{v.ID, v.Revision, v.Varcol, v.Created}
}

type Versions []*Version
