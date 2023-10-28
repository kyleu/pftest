// Package audited - Content managed by Project Forge, see [projectforge.md] for details.
package audited

import "github.com/kyleu/pftest/app/util"

func (a *Audited) Diff(ax *Audited) util.Diffs {
	var diffs util.Diffs
	if a.ID != ax.ID {
		diffs = append(diffs, util.NewDiff("id", a.ID.String(), ax.ID.String()))
	}
	if a.Name != ax.Name {
		diffs = append(diffs, util.NewDiff("name", a.Name, ax.Name))
	}
	return diffs
}
