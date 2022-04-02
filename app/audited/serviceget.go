// Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Auditeds, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get auditeds")
	}
	return ret.ToAuditeds(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (*Audited, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, s.logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get audited by id [%v]", id)
	}
	return ret.ToAudited(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, ids ...uuid.UUID) (Auditeds, error) {
	if len(ids) == 0 {
		return Auditeds{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Auditeds for [%d] ids", len(ids))
	}
	return ret.ToAuditeds(), nil
}

const searchClause = "(lower(id::text) like $1 or lower(name) like $1)"

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params) (Auditeds, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToAuditeds(), nil
}
