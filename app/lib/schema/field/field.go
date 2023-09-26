// Package field - Content managed by Project Forge, see [projectforge.md] for details.
package field

import (
	"reflect"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/types"
	"github.com/kyleu/pftest/app/util"
)

type Field struct {
	Key      string         `json:"key"`
	Type     *types.Wrapped `json:"type"`
	Title    string         `json:"-"` // override only
	Plural   string         `json:"-"` // override only
	Hidden   bool           `json:"-"` // override only
	Default  any            `json:"default,omitempty"`
	ReadOnly bool           `json:"readOnly,omitempty"`
	Metadata *Metadata      `json:"metadata,omitempty"`
}

func NewFieldByType(key string, t reflect.Type, ro bool, md *Metadata) *Field {
	return &Field{Key: key, Type: types.FromReflect(t), ReadOnly: ro, Metadata: md}
}

func (f *Field) Name() string {
	if f.Title == "" {
		return util.StringToSingular(util.StringToTitle(f.Key))
	}
	return f.Title
}

func (f *Field) PluralName() string {
	if f.Plural == "" {
		ret := f.Name()
		return util.StringToPlural(ret)
	}
	return f.Plural
}

func (f *Field) String() string {
	return f.Key + " " + f.Type.String()
}

func (f *Field) Description() string {
	if f.Metadata == nil {
		return ""
	}
	return f.Metadata.Description
}

func (f *Field) DefaultClean() any {
	switch f.Default {
	case nil:
		return f.Type.Default(f.Key)
	case "now()":
		return util.TimeCurrent()
	default:
		return f.Default
	}
}

type Fields []*Field

func (s Fields) Get(key string) (int, *Field) {
	for idx, x := range s {
		if x.Key == key {
			return idx, x
		}
	}
	return -1, nil
}

func (s Fields) Names() []string {
	return lo.Map(s, func(x *Field, _ int) string {
		return x.Key
	})
}
