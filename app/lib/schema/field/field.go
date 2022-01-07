package field

import (
	"time"

	"github.com/kyleu/pftest/app/lib/schema/types"
	"github.com/kyleu/pftest/app/util"
)

type Field struct {
	Key      string         `json:"key"`
	Type     *types.Wrapped `json:"type"`
	Title    string         `json:"-"` // override only
	Plural   string         `json:"-"` // override only
	Hidden   bool           `json:"-"` // override only
	Default  interface{}    `json:"default,omitempty"`
	ReadOnly bool           `json:"readOnly,omitempty"`
	Metadata *Metadata      `json:"metadata,omitempty"`
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

func (f *Field) DefaultClean() interface{} {
	switch f.Default {
	case nil:
		return f.Type.Default(f.Key)
	case "now()":
		return time.Now()
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
	ret := make([]string, 0, len(s))
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}
