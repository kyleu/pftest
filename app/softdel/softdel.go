// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/util"
)

type Softdel struct {
	ID      string     `json:"id"`
	Created time.Time  `json:"created"`
	Updated *time.Time `json:"updated,omitempty"`
	Deleted *time.Time `json:"deleted,omitempty"`
}

func New(id string) *Softdel {
	return &Softdel{ID: id}
}

func Random() *Softdel {
	return &Softdel{
		ID:      util.RandomString(12),
		Created: util.TimeCurrent(),
		Updated: util.TimeCurrentP(),
		Deleted: nil,
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Softdel, error) {
	ret := &Softdel{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.Deleted, err = m.ParseTime("deleted", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (s *Softdel) Clone() *Softdel {
	return &Softdel{s.ID, s.Created, s.Updated, s.Deleted}
}

func (s *Softdel) String() string {
	return s.ID
}

func (s *Softdel) TitleString() string {
	return s.String()
}

func (s *Softdel) WebPath() string {
	return "/softdel/" + url.QueryEscape(s.ID)
}

func (s *Softdel) Diff(sx *Softdel) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID, sx.ID))
	}
	if s.Created != sx.Created {
		diffs = append(diffs, util.NewDiff("created", s.Created.String(), sx.Created.String()))
	}
	if (s.Deleted == nil && sx.Deleted != nil) || (s.Deleted != nil && sx.Deleted == nil) || (s.Deleted != nil && sx.Deleted != nil && *s.Deleted != *sx.Deleted) {
		diffs = append(diffs, util.NewDiff("deleted", fmt.Sprint(s.Deleted), fmt.Sprint(sx.Deleted))) //nolint:gocritic // it's nullable
	}
	return diffs
}

func (s *Softdel) ToData() []any {
	return []any{s.ID, s.Created, s.Updated, s.Deleted}
}
