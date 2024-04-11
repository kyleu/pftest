// Package basic - Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/svc"
)

var (
	_ svc.ServiceID[*Basic, Basics, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Basics]                = (*Service)(nil)
)

type Service struct {
	db     *database.Service
	dbRead *database.Service
}

func NewService(db *database.Service, dbRead *database.Service) *Service {
	filter.AllowedColumns["basic"] = columns
	return &Service{db: db, dbRead: dbRead}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("basic", &filter.Ordering{Column: "created"})
}
