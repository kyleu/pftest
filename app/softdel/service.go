// Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

type Service struct {
	db     *database.Service
	logger util.Logger
}

func NewService(db *database.Service, logger util.Logger) *Service {
	logger = logger.With("svc", "softdel")
	filter.AllowedColumns["softdel"] = columns
	return &Service{db: db, logger: logger}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("softdel", &filter.Ordering{Column: "created"})
}
