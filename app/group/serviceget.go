// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Groups, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get groups")
	}
	return ret.ToGroups(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*Group, error) {
	wc := defaultWC
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get group by id [%v]", id)
	}
	return ret.ToGroup(), nil
}

func (s *Service) GetGroups(ctx context.Context, tx *sqlx.Tx) ([]*util.KeyValInt, error) {
	wc := ""
	q := database.SQLSelectGrouped("\"group\" as key, count(*) as val", tableQuoted, wc, "\"group\"", "\"group\"", 0, 0)
	var ret []*util.KeyValInt
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get groups by group")
	}
	return ret, nil
}

func (s *Service) GetByGroup(ctx context.Context, tx *sqlx.Tx, group string, params *filter.Params) (Groups, error) {
	params = filters(params)
	wc := defaultWC
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, group)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get groups by group [%v]", group)
	}
	return ret.ToGroups(), nil
}
