package reference

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/svc"
)

var (
	_ svc.ServiceID[*Reference, References, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[References]                    = (*Service)(nil)
)

type Service struct {
	db     *database.Service
	dbRead *database.Service
}

func NewService(db *database.Service, dbRead *database.Service) *Service {
	filter.AllowedColumns["reference"] = columns
	return &Service{db: db, dbRead: dbRead}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("reference", &filter.Ordering{Column: "created"})
}
