// Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["path"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("path", &filter.Ordering{Column: "created"})
}
