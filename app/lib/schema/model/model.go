package model

import (
	"fmt"
	"sort"

	"github.com/kyleu/pftest/app/lib/schema/field"
	"github.com/kyleu/pftest/app/util"
)

type Model struct {
	Key           string        `json:"key"`
	Pkg           util.Pkg      `json:"pkg,omitempty"`
	Type          Type          `json:"type"`
	Title         string        `json:"-"` // override only
	Plural        string        `json:"-"` // override only
	Interfaces    []string      `json:"interfaces,omitempty"`
	Fields        field.Fields  `json:"fields,omitempty"`
	Indexes       Indexes       `json:"indexes,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	References    References    `json:"-"` // internal cache
	Metadata      *Metadata     `json:"metadata,omitempty"`
	pk            []string      // internal cache
}

func NewModel(pkg util.Pkg, key string) *Model {
	return &Model{Key: key, Pkg: pkg}
}

func (m *Model) String() string {
	if len(m.Pkg) == 0 {
		return m.Key
	}
	return m.Pkg.ToPath(m.Key)
}

func (m *Model) Name() string {
	if m.Title == "" {
		return util.StringToSingular(util.StringToTitle(m.Key))
	}
	return m.Title
}

func (m *Model) PluralName() string {
	if m.Plural == "" {
		ret := m.Name()
		return util.StringToPlural(ret)
	}
	return m.Plural
}

func (m *Model) Description() string {
	return fmt.Sprintf("%s model [%s]", m.Type.String(), m.Key)
}

func (m *Model) Path() util.Pkg {
	return m.Pkg.With(m.Key)
}

func (m *Model) PathString() string {
	return m.Pkg.ToPath(m.Key)
}

type Models []*Model

func (m Models) Get(pkg util.Pkg, key string) *Model {
	for _, x := range m {
		if x.Pkg.Equals(pkg) && x.Key == key {
			return x
		}
	}
	return nil
}

func (m Models) Sort() {
	sort.Slice(m, func(l int, r int) bool {
		return m[l].Key < m[r].Key
	})
}

func (m Models) Names() []string {
	ret := make([]string, 0, len(m))
	for _, md := range m {
		ret = append(ret, md.Key)
	}
	return ret
}
