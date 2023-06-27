// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

type IDRev struct {
	ID             string `db:"ID"`
	CurrentVersion int    `db:"current_Version"`
}

func (s *Service) GetAllVersions(ctx context.Context, tx *sqlx.Tx, id string, params *filter.Params, logger util.Logger) (Capitals, error) {
	params = filters(params)
	wc := "\"ID\" = $1"
	tablesJoinedParam := fmt.Sprintf("%q c join %q cr on c.\"ID\" = cr.\"Capital_ID\"", table, tableVersion)
	q := database.SQLSelect(columnsString, tablesJoinedParam, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Capitals")
	}
	return ret.ToCapitals(), nil
}

func (s *Service) GetVersion(ctx context.Context, tx *sqlx.Tx, id string, version int, logger util.Logger) (*Capital, error) {
	wc := "\"ID\" = $1 and \"Version\" = $2"
	ret := &row{}
	tablesJoinedParam := fmt.Sprintf("%q c join %q cr on c.\"ID\" = cr.\"Capital_ID\"", table, tableVersion)
	q := database.SQLSelectSimple(columnsString, tablesJoinedParam, s.db.Placeholder(), wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id, version)
	if err != nil {
		return nil, err
	}
	return ret.ToCapital(), nil
}

func (s *Service) getCurrentVersions(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Capital) (map[string]int, error) {
	stmts := lo.Map(models, func(_ *Capital, i int) string {
		return fmt.Sprintf(`"ID" = $%d`, i+1)
	})
	q := database.SQLSelectSimple(`"ID", "current_Version"`, tableQuoted, s.db.Placeholder(), strings.Join(stmts, " or "))
	vals := lo.FlatMap(models, func(model *Capital, _ int) []any {
		return []any{model.ID}
	})
	var results []*IDRev
	err := s.dbRead.Select(ctx, &results, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Capitals")
	}

	ret := make(map[string]int, len(models))
	lo.ForEach(models, func(model *Capital, _ int) {
		curr := 0
		lo.ForEach(results, func(x *IDRev, _ int) {
			if x.ID == model.ID {
				curr = x.CurrentVersion
			}
		})
		ret[model.String()] = curr
	})
	return ret, nil
}
