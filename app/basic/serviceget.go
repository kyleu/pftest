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
	wc := ""
	sql := database.SQLSelect(columnsString, table, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Basics")
	}
	return ret.ToBasics(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, name string) (*Basic, error) {
	wc := "id = $1 and name = $2"
	ret := &dto{}
	sql := database.SQLSelectSimple(columnsString, table, wc)
	err := s.db.Get(ctx, ret, sql, tx, id, name)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get basic by id [%s], name [%s]", id, name)
	}
	return ret.ToBasic(), nil
}

func (s *Service) GetByID(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, params *filter.Params) (Basics, error) {
	wc := "id = $1"
	sql := database.SQLSelect(columnsString, table, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get basics by id [%s]", id)
	}
	return ret.ToBasics(), nil
}

func (s *Service) GetByName(ctx context.Context, tx *sqlx.Tx, name string, params *filter.Params) (Basics, error) {
	wc := "name = $1"
	sql := database.SQLSelect(columnsString, table, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, name)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get basics by name [%s]", name)
	}
	return ret.ToBasics(), nil
}

func (s *Service) Search(ctx context.Context, q string, tx *sqlx.Tx, params *filter.Params) (Basics, error) {
	wc := "(lower(id::text) like $1 or lower(name) like $1)"
	sql := database.SQLSelect(columnsString, table, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, "%"+strings.ToLower(q)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToBasics(), nil
}
