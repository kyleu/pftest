package capital

import (
	"net/url"
	"time"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/capital"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Capital)(nil)

type Capital struct {
	ID       string     `json:"id,omitzero"`
	Name     string     `json:"name,omitzero"`
	Birthday time.Time  `json:"birthday,omitzero"`
	Deathday *time.Time `json:"deathday,omitzero"`
}

func NewCapital(id string) *Capital {
	return &Capital{ID: id}
}

func (c *Capital) Clone() *Capital {
	return &Capital{ID: c.ID, Name: c.Name, Birthday: c.Birthday, Deathday: c.Deathday}
}

func (c *Capital) String() string {
	return c.ID
}

func (c *Capital) TitleString() string {
	if xx := c.Name; xx != "" {
		return xx
	}
	return c.String()
}

func RandomCapital() *Capital {
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
	return CapitalFieldDescs.Keys(), [][]string{c.Strings()}
}

func (c *Capital) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(c.ID))...)
}

func (c *Capital) Breadcrumb(extra ...string) string {
	return c.TitleString() + "||" + c.WebPath(extra...) + "**star"
}

func (c *Capital) ToData() []any {
	return []any{c.ID, c.Name, c.Birthday, c.Deathday}
}

var CapitalFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "string"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "birthday", Title: "Birthday", Description: "", Type: "timestamp"},
	{Key: "deathday", Title: "Deathday", Description: "", Type: "timestamp"},
}
