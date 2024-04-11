// Package relation - Content managed by Project Forge, see [projectforge.md] for details.
package relation

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/svc"
)

var (
	_ svc.ServiceID[*Relation, Relations, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Relations]                   = (*Service)(nil)
)

type Service struct {
	db     *database.Service
	dbRead *database.Service
}

func NewService(db *database.Service, dbRead *database.Service) *Service {
	filter.AllowedColumns["relation"] = columns
	return &Service{db: db, dbRead: dbRead}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("relation", &filter.Ordering{Column: "created"})
}
