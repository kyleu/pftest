// Package enum - Content managed by Project Forge, see [projectforge.md] for details.
package enum

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/types"
	"github.com/kyleu/pftest/app/util"
)

const defaultIcon = "hammer"

type Enum struct {
	Name           string        `json:"name"`
	Package        string        `json:"package"`
	Group          []string      `json:"group,omitempty"`
	Description    string        `json:"description,omitempty"`
	Icon           string        `json:"icon,omitempty"`
	Values         Values        `json:"values,omitempty"`
	Tags           []string      `json:"tags,omitempty"`
	TitleOverride  string        `json:"title,omitempty"`
	ProperOverride string        `json:"proper,omitempty"`
	RouteOverride  string        `json:"route,omitempty"`
	Config         util.ValueMap `json:"config,omitempty"`
}

func (e *Enum) Title() string {
	if e.TitleOverride == "" {
		return util.StringToTitle(e.Name)
	}
	return e.TitleOverride
}

func (e *Enum) TitleLower() string {
	return strings.ToLower(e.Title())
}

func (e *Enum) Proper() string {
	if e.ProperOverride == "" {
		return util.StringToCamel(e.Name)
	}
	return util.StringToCamel(e.ProperOverride)
}

func (e *Enum) ProperPlural() string {
	ret := util.StringToPlural(e.Proper())
	if ret == e.Proper() {
		return ret + "Set"
	}
	return ret
}

func (e *Enum) FirstLetter() any {
	return strings.ToLower(e.Name[0:1])
}

func (e *Enum) IconSafe() string {
	if _, ok := util.SVGLibrary[e.Icon]; ok {
		return e.Icon
	}
	return defaultIcon
}

func (e *Enum) Camel() string {
	return util.StringToLowerCamel(e.Name)
}

func (e *Enum) ExtraFields() *util.OrderedMap[string] {
	ret := util.NewOrderedMap[string](false, 0)
	for _, v := range e.Values {
		if v.Extra == nil {
			continue
		}
		for _, k := range v.Extra.Order {
			x := v.Extra.GetSimple(k)
			if _, exists := ret.Get(k); exists {
				continue
			}
			typ := ""
			switch x.(type) {
			case string:
				typ = types.KeyString
			case float64:
				typ = types.KeyFloat
			case int, int32, int64:
				typ = types.KeyInt
			case bool:
				typ = types.KeyBool
			}
			if x := e.Config.GetStringOpt("type:" + k); x != "" {
				switch x {
				case types.KeyString:
					typ = types.KeyString
				case types.KeyInt:
					typ = types.KeyInt
				case types.KeyFloat:
					typ = types.KeyFloat
				case types.KeyBool:
					typ = types.KeyBool
				case types.KeyTimestamp:
					typ = types.KeyTimestamp
				default:
					typ = "unknown config type [" + x + "]"
				}
			}
			ret.Append(k, typ)
		}
	}
	return ret
}

func (e *Enum) ExtraFieldValues(k string) ([]any, bool) {
	ret := make([]any, 0, len(e.Values))
	for _, v := range e.Values {
		if v.Extra == nil {
			continue
		}
		if x, ok := v.Extra.Get(k); ok && x != nil {
			ret = append(ret, x)
		}
	}
	return ret, len(lo.Uniq(ret)) == len(ret)
}

func (e *Enum) PackageWithGroup(prefix string) string {
	if len(e.Group) == 0 {
		if prefix != "" && !strings.HasSuffix(prefix, "/") {
			prefix += "/"
		}
		return prefix + e.Package
	}
	var x []string
	if prefix != "" {
		x = append(x, prefix)
	}
	x = append(x, e.Group...)
	x = append(x, e.Package)
	return strings.Join(x, "/")
}

func (e *Enum) HasTag(t string) bool {
	return lo.Contains(e.Tags, t)
}

func (e *Enum) Breadcrumbs() string {
	ret := lo.Map(e.Group, func(g string, _ int) string {
		return fmt.Sprintf("%q", g)
	})
	ret = append(ret, fmt.Sprintf("%q", e.Package))
	return strings.Join(ret, ", ")
}

func (e *Enum) ValuesCamel() []string {
	return lo.Map(e.Values, func(x *Value, _ int) string {
		return util.StringToCamel(x.Key)
	})
}

func (e *Enum) Simple() bool {
	return e.Values.AllSimple()
}

type Enums []*Enum

func (e Enums) Get(key string) *Enum {
	return lo.FindOrElse(e, nil, func(x *Enum) bool {
		return x.Name == key
	})
}
