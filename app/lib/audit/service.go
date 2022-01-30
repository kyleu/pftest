// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

type Service struct {
	db     *database.Service
	logger *zap.SugaredLogger
}

func NewService(db *database.Service, logger *zap.SugaredLogger) *Service {
	logger = logger.With("svc", "audit")
	filter.AllowedColumns["audit"] = columns
	return &Service{db: db, logger: logger}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("audit", &filter.Ordering{Column: "started"})
}

func recordFilters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("audit_record", &filter.Ordering{Column: "occurred"})
}
