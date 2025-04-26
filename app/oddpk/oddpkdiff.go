package oddpk

import "github.com/kyleu/pftest/app/util"

func (o *OddPK) Diff(ox *OddPK) util.Diffs {
	var diffs util.Diffs
	if o.Project != ox.Project {
		diffs = append(diffs, util.NewDiff("project", o.Project.String(), ox.Project.String()))
	}
	if o.Path != ox.Path {
		diffs = append(diffs, util.NewDiff("path", o.Path, ox.Path))
	}
	if o.Name != ox.Name {
		diffs = append(diffs, util.NewDiff("name", o.Name, ox.Name))
	}
	return diffs
}
