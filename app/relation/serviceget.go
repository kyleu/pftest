package relation

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Relation, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get relation by id [%v]", id)
	}
	return ret.ToRelation(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Relations, error) {
	if len(ids) == 0 {
		return Relations{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Relations for [%d] ids", len(ids))
	}
	return ret.ToRelations(), nil
}

func (s *Service) GetByBasicID(ctx context.Context, tx *sqlx.Tx, basicID uuid.UUID, params *filter.Params, logger util.Logger) (Relations, error) {
	params = filters(params)
	wc := "\"basic_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, basicID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Relations by basicID [%v]", basicID)
	}
	return ret.ToRelations(), nil
}

func (s *Service) GetByBasicIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, basicIDs ...uuid.UUID) (Relations, error) {
	if len(basicIDs) == 0 {
		return Relations{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("basic_id", len(basicIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(basicIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Relations for [%d] basicIDs", len(basicIDs))
	}
	return ret.ToRelations(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Relation, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random relations")
	}
	return ret.ToRelation(), nil
}
