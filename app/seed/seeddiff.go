// Package seed - Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"fmt"

	"github.com/kyleu/pftest/app/util"
)

func (s *Seed) Diff(sx *Seed) util.Diffs {
	var diffs util.Diffs
	if s.ID != sx.ID {
		diffs = append(diffs, util.NewDiff("id", s.ID.String(), sx.ID.String()))
	}
	if s.Name != sx.Name {
		diffs = append(diffs, util.NewDiff("name", s.Name, sx.Name))
	}
	if s.Size != sx.Size {
		diffs = append(diffs, util.NewDiff("size", fmt.Sprint(s.Size), fmt.Sprint(sx.Size)))
	}
	diffs = append(diffs, util.DiffObjects(s.Obj, sx.Obj, "obj")...)
	return diffs
}
