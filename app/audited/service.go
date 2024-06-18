package audited

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/svc"
)

var (
	_ svc.ServiceID[*Audited, Auditeds, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Auditeds]                  = (*Service)(nil)
)

type Service struct {
	db     *database.Service
	dbRead *database.Service
	audit  *audit.Service
}

func NewService(db *database.Service, aud *audit.Service) *Service {
	filter.AllowedColumns["audited"] = columns
	return &Service{db: db, audit: aud}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("audited", &filter.Ordering{Column: "name"})
}
