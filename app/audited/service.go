// Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

type Service struct {
	db     *database.Service
	audit  *audit.Service
	logger *zap.SugaredLogger
}

func NewService(db *database.Service, aud *audit.Service, logger *zap.SugaredLogger) *Service {
	logger = logger.With("svc", "audited")
	filter.AllowedColumns["audited"] = columns
	return &Service{db: db, audit: aud, logger: logger}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("audited", &filter.Ordering{Column: "created"})
}
