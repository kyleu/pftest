package basic

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Basics, error) {
	params = filters(params)
	wc := ""
	sql := database.SQLSelect(columnsString, table, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Basics")
	}
	return ret.ToBasics(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (*Basic, error) {
	wc := "id = $1"
	ret := &dto{}
	sql := database.SQLSelectSimple(columnsString, table, wc)
	err := s.db.Get(ctx, ret, sql, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get basic by id [%s]", id)
	}
	return ret.ToBasic(), nil
}

const searchClause = "(lower(id::text) like $1 or lower(name) like $1)"

func (s *Service) Search(ctx context.Context, q string, tx *sqlx.Tx, params *filter.Params) (Basics, error) {
	params = filters(params)
	wc := searchClause
	sql := database.SQLSelect(columnsString, table, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, "%"+strings.ToLower(q)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToBasics(), nil
}
