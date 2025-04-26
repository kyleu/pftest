package oddrel

import "github.com/kyleu/pftest/app/util"

func (o *Oddrel) Diff(ox *Oddrel) util.Diffs {
	var diffs util.Diffs
	if o.ID != ox.ID {
		diffs = append(diffs, util.NewDiff("id", o.ID.String(), ox.ID.String()))
	}
	if o.Project != ox.Project {
		diffs = append(diffs, util.NewDiff("project", o.Project.String(), ox.Project.String()))
	}
	if o.Path != ox.Path {
		diffs = append(diffs, util.NewDiff("path", o.Path, ox.Path))
	}
	return diffs
}
