package mixedcase

import "github.com/kyleu/pftest/app/util"

func (m *MixedCase) Diff(mx *MixedCase) util.Diffs {
	var diffs util.Diffs
	if m.ID != mx.ID {
		diffs = append(diffs, util.NewDiff("id", m.ID, mx.ID))
	}
	if m.TestField != mx.TestField {
		diffs = append(diffs, util.NewDiff("testField", m.TestField, mx.TestField))
	}
	if m.AnotherField != mx.AnotherField {
		diffs = append(diffs, util.NewDiff("anotherField", m.AnotherField, mx.AnotherField))
	}
	return diffs
}
