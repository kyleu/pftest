// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

type Service struct {
	db     *database.Service
	logger *zap.SugaredLogger
}

func NewService(db *database.Service, logger *zap.SugaredLogger) *Service {
	logger = logger.With("svc", "audit")
	filter.AllowedColumns["audit"] = columns
	filter.AllowedColumns["audit_record"] = recordColumns
	return &Service{db: db, logger: logger}
}

func Apply(a *Audit, r Records) (*Audit, Records, error) {
	ret := &Audit{ID: util.UUID()}
	records := Records{}
	return ret, records, nil
}
