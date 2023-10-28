// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"fmt"

	"github.com/kyleu/pftest/app/util"
)

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
