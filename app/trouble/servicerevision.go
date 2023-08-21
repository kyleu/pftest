// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

type IDRev struct {
	From             string   `db:"from"`
	Where            []string `db:"where"`
	CurrentSelectcol int      `db:"current_selectcol"`
}

//nolint:lll
func (s *Service) GetAllSelectcols(ctx context.Context, tx *sqlx.Tx, from string, where []string, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := "\"from\" = $1 and \"where\" = $2"
	wc = addDeletedClause(wc, includeDeleted)
	tablesJoinedParam := fmt.Sprintf("%q t join %q tr on t.\"from\" = tr.\"trouble_from\" and t.\"where\" = tr.\"trouble_where\"", table, tableSelectcol)
	q := database.SQLSelect(columnsString, tablesJoinedParam, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, from, where)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Troubles")
	}
	return ret.ToTroubles(), nil
}

func (s *Service) GetSelectcol(ctx context.Context, tx *sqlx.Tx, from string, where []string, selectcol int, logger util.Logger) (*Trouble, error) {
	wc := "\"from\" = $1 and \"where\" = $2 and \"selectcol\" = $3"
	ret := &row{}
	tablesJoinedParam := fmt.Sprintf("%q t join %q tr on t.\"from\" = tr.\"trouble_from\" and t.\"where\" = tr.\"trouble_where\"", table, tableSelectcol)
	q := database.SQLSelectSimple(columnsString, tablesJoinedParam, s.db.Placeholder(), wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, from, where, selectcol)
	if err != nil {
		return nil, err
	}
	return ret.ToTrouble(), nil
}

func (s *Service) getCurrentSelectcols(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) (map[string]int, error) {
	stmts := lo.Map(models, func(_ *Trouble, i int) string {
		return fmt.Sprintf(`"from" = $%d and "where" = $%d`, (i*2)+1, (i*2)+2)
	})
	q := database.SQLSelectSimple(`"from", "where", "current_selectcol"`, tableQuoted, s.db.Placeholder(), strings.Join(stmts, " or "))
	vals := lo.FlatMap(models, func(model *Trouble, _ int) []any {
		return []any{model.From, model.Where}
	})
	var results []*IDRev
	err := s.dbRead.Select(ctx, &results, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Troubles")
	}

	ret := make(map[string]int, len(models))
	lo.ForEach(models, func(model *Trouble, _ int) {
		curr := 0
		lo.ForEach(results, func(x *IDRev, _ int) {
			if x.From == model.From && slices.Equal(x.Where, model.Where) {
				curr = x.CurrentSelectcol
			}
		})
		ret[model.String()] = curr
	})
	return ret, nil
}
