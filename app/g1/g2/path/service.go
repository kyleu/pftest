// Package path - Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/svc"
)

var (
	_ svc.ServiceID[*Path, Paths, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Paths]               = (*Service)(nil)
)

type Service struct {
	db     *database.Service
	dbRead *database.Service
}

func NewService(db *database.Service, dbRead *database.Service) *Service {
	filter.AllowedColumns["path"] = columns
	return &Service{db: db, dbRead: dbRead}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("path", &filter.Ordering{Column: "created"})
}
