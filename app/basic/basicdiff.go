// Package basic - Content managed by Project Forge, see [projectforge.md] for details.
package basic

import "github.com/kyleu/pftest/app/util"

func (b *Basic) Diff(bx *Basic) util.Diffs {
	var diffs util.Diffs
	if b.ID != bx.ID {
		diffs = append(diffs, util.NewDiff("id", b.ID.String(), bx.ID.String()))
	}
	if b.Name != bx.Name {
		diffs = append(diffs, util.NewDiff("name", b.Name, bx.Name))
	}
	if b.Status != bx.Status {
		diffs = append(diffs, util.NewDiff("status", b.Status, bx.Status))
	}
	if b.Created != bx.Created {
		diffs = append(diffs, util.NewDiff("created", b.Created.String(), bx.Created.String()))
	}
	return diffs
}
