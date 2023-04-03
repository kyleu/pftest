// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := ""
	if !includeDeleted {
		wc = "\"delete\" is null"
	}
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get troubles")
	}
	return ret.ToTroubles(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, includeDeleted bool, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	if !includeDeleted {
		if whereClause == "" {
			whereClause = "\"delete\" is null"
		} else {
			whereClause += " and " + "\"delete\" is null"
		}
	}
	q := database.SQLSelectSimple("count(*) as x", tablesJoined, s.db.Placeholder(), whereClause)
	ret, err := s.dbRead.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of troubles")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, from string, where []string, includeDeleted bool, logger util.Logger) (*Trouble, error) {
	wc := defaultWC(0)
	wc = addDeletedClause(wc, includeDeleted)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, s.db.Placeholder(), wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, from, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get trouble by from [%v], where [%v]", from, where)
	}
	return ret.ToTrouble(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, includeDeleted bool, logger util.Logger, pks ...*PK) (Troubles, error) {
	if len(pks) == 0 {
		return Troubles{}, nil
	}
	wc := "("
	for idx := range pks {
		if idx > 0 {
			wc += " or "
		}
		wc += fmt.Sprintf("(from = $%d and where = $%d)", (idx*2)+1, (idx*2)+2)
	}
	wc += ")"
	wc = addDeletedClause(wc, includeDeleted)
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, s.db.Placeholder(), wc)
	vals := make([]any, 0, len(pks)*2)
	for _, x := range pks {
		vals = append(vals, x.From, x.Where)
	}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Troubles for [%d] pks", len(pks))
	}
	return ret.ToTroubles(), nil
}

func (s *Service) GetByFrom(ctx context.Context, tx *sqlx.Tx, from string, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := "\"from\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, from)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by from [%v]", from)
	}
	return ret.ToTroubles(), nil
}

func (s *Service) GetByWhere(ctx context.Context, tx *sqlx.Tx, where []string, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) { //nolint:lll
	params = filters(params)
	wc := "\"where\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by where [%v]", where)
	}
	return ret.ToTroubles(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Troubles, error) {
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get troubles using custom SQL")
	}
	return ret.ToTroubles(), nil
}
