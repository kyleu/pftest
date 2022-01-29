// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) RecordsForAudit(ctx context.Context, tx *sqlx.Tx, auditID uuid.UUID, params *filter.Params) (Records, error) {
	params = filters(params)
	wc := `"audit_id" = $1`
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := recordDTOs{}
	err := s.db.Select(ctx, &ret, q, tx, auditID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get audit_records by auditID [%v]", auditID)
	}
	return ret.ToRecords(), nil
}
