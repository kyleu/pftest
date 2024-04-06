// Package audited - Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Auditeds, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get auditeds")
	}
	return ret.ToAuditeds(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Type, whereClause)
	ret, err := s.dbRead.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of auditeds")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Audited, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get audited by id [%v]", id)
	}
	return ret.ToAudited(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Auditeds, error) {
	if len(ids) == 0 {
		return Auditeds{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Auditeds for [%d] ids", len(ids))
	}
	return ret.ToAuditeds(), nil
}

const searchClause = "(lower(id::text) like $1 or lower(name) like $1)"

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Auditeds, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToAuditeds(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Auditeds, error) {
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get auditeds using custom SQL")
	}
	return ret.ToAuditeds(), nil
}

func (s *Service) ListWhere(ctx context.Context, tx *sqlx.Tx, where string, params *filter.Params, logger util.Logger, values ...any) (Auditeds, error) {
	params = filters(params)
	sql := database.SQLSelect(columnsString, tableQuoted, where, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	return s.ListSQL(ctx, tx, sql, logger, values...)
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Audited, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random auditeds")
	}
	return ret.ToAudited(), nil
}
