package basic

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	table         = "basic"
	columns       = []string{"id", "name", "created"}
	columnsString = strings.Join(columns, ", ")
)

type dto struct {
	ID      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}

func (d *dto) ToBasic() *Basic {
	if d == nil {
		return nil
	}
	return &Basic{ID: d.ID, Name: d.Name, Created: d.Created}
}

type dtos []*dto

func (x dtos) ToBasics() Basics {
	ret := make(Basics, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToBasic())
	}
	return ret
}
