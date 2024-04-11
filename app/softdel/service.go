// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/svc"
)

var _ svc.ServiceSoftDeleteID[*Softdel, Softdels, string] = (*Service)(nil)

type Service struct {
	db     *database.Service
	dbRead *database.Service
}

func NewService(db *database.Service, dbRead *database.Service) *Service {
	filter.AllowedColumns["softdel"] = columns
	return &Service{db: db, dbRead: dbRead}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("softdel", &filter.Ordering{Column: "created"})
}
