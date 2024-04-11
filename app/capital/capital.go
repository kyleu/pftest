// Package capital - Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

var _ svc.Model = (*Capital)(nil)

type Capital struct {
	ID       string     `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Birthday time.Time  `json:"birthday,omitempty"`
	Deathday *time.Time `json:"deathday,omitempty"`
}

func New(id string) *Capital {
	return &Capital{ID: id}
}

func (c *Capital) Clone() *Capital {
	return &Capital{c.ID, c.Name, c.Birthday, c.Deathday}
}

func (c *Capital) String() string {
	return c.ID
}

func (c *Capital) TitleString() string {
	return c.Name
}

func Random() *Capital {
	return &Capital{
		ID:       util.RandomString(12),
		Name:     util.RandomString(12),
		Birthday: util.TimeCurrent(),
		Deathday: util.TimeCurrentP(),
	}
}

func (c *Capital) Strings() []string {
	return []string{c.ID, c.Name, util.TimeToFull(&c.Birthday), util.TimeToFull(c.Deathday)}
}

func (c *Capital) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{c.Strings()}
}

func (c *Capital) WebPath() string {
	return "/capital/" + url.QueryEscape(c.ID)
}

func (c *Capital) ToData() []any {
	return []any{c.ID, c.Name, c.Birthday, c.Deathday}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "string"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "birthday", Title: "Birthday", Description: "", Type: "timestamp"},
	{Key: "deathday", Title: "Deathday", Description: "", Type: "timestamp"},
}
