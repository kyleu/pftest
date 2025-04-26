package oddpk

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, project uuid.UUID, path string, logger util.Logger) (*OddPK, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, project, strings.ReplaceAll(path, "||", "/"))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get oddPK by project [%v], path [%v]", project, path)
	}
	return ret.ToOddPK(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, pks ...*PK) (OddPKs, error) {
	if len(pks) == 0 {
		return OddPKs{}, nil
	}
	wc := "("
	lo.ForEach(pks, func(_ *PK, idx int) {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(project = $%d and path = $%d)", (idx*2)+1, (idx*2)+2)
	})
	wc += ")"
	ret := rows{}
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	vals := lo.FlatMap(pks, func(x *PK, _ int) []any {
		return []any{x.Project, x.Path}
	})
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get OddPKs for [%d] pks", len(pks))
	}
	return ret.ToOddPKs(), nil
}

func (s *Service) GetByPath(ctx context.Context, tx *sqlx.Tx, path string, params *filter.Params, logger util.Logger) (OddPKs, error) {
	params = filters(params)
	wc := "\"path\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Odd PKs by path [%v]", path)
	}
	return ret.ToOddPKs(), nil
}

func (s *Service) GetByPaths(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, paths ...string) (OddPKs, error) {
	if len(paths) == 0 {
		return OddPKs{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("path", len(paths), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(paths)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get OddPKs for [%d] paths", len(paths))
	}
	return ret.ToOddPKs(), nil
}

func (s *Service) GetByProject(ctx context.Context, tx *sqlx.Tx, project uuid.UUID, params *filter.Params, logger util.Logger) (OddPKs, error) {
	params = filters(params)
	wc := "\"project\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, project)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Odd PKs by project [%v]", project)
	}
	return ret.ToOddPKs(), nil
}

func (s *Service) GetByProjects(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, projects ...uuid.UUID) (OddPKs, error) {
	if len(projects) == 0 {
		return OddPKs{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("project", len(projects), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(projects)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get OddPKs for [%d] projects", len(projects))
	}
	return ret.ToOddPKs(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*OddPK, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random odd pks")
	}
	return ret.ToOddPK(), nil
}
