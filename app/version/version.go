// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"fmt"
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
		Created:  util.TimeCurrent(),
		Updated:  util.TimeCurrentP(),
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
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (v *Version) Clone() *Version {
	return &Version{v.ID, v.Revision, v.Constcol, v.Varcol.Clone(), v.Created, v.Updated}
}

func (v *Version) String() string {
	return v.ID
}

func (v *Version) TitleString() string {
	return v.String()
}

func (v *Version) WebPath() string {
	return "/version/" + v.ID
}

func (v *Version) Diff(vx *Version) util.Diffs {
	var diffs util.Diffs
	if v.ID != vx.ID {
		diffs = append(diffs, util.NewDiff("id", v.ID, vx.ID))
	}
	if v.Revision != vx.Revision {
		diffs = append(diffs, util.NewDiff("revision", fmt.Sprint(v.Revision), fmt.Sprint(vx.Revision)))
	}
	if v.Constcol != vx.Constcol {
		diffs = append(diffs, util.NewDiff("constcol", v.Constcol, vx.Constcol))
	}
	diffs = append(diffs, util.DiffObjects(v.Varcol, vx.Varcol, "varcol")...)
	if v.Created != vx.Created {
		diffs = append(diffs, util.NewDiff("created", v.Created.String(), vx.Created.String()))
	}
	return diffs
}

func (v *Version) ToData() []any {
	return []any{v.ID, v.Revision, v.Constcol, v.Varcol, v.Created, v.Updated}
}

func (v *Version) ToDataCore() []any {
	return []any{v.ID, v.Revision, v.Constcol, v.Updated}
}

func (v *Version) ToDataRevision() []any {
	return []any{v.ID, v.Revision, v.Varcol, v.Created}
}
