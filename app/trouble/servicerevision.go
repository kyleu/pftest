// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) GetAllSelectcols(ctx context.Context, tx *sqlx.Tx, from string, where []string, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := "\"from\" = $1 and \"where\" = $2"
	wc = addDeletedClause(wc, includeDeleted)
	tablesJoinedParam := fmt.Sprintf("%q t join %q tr on t.\"from\" = tr.\"trouble_from\" and t.\"where\" = tr.\"trouble_where\"", table, tableSelectcol)
	q := database.SQLSelect(columnsString, tablesJoinedParam, wc, params.OrderByString(), params.Limit, params.Offset)
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
	q := database.SQLSelectSimple(columnsString, tablesJoinedParam, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, from, where, selectcol)
	if err != nil {
		return nil, err
	}
	return ret.ToTrouble(), nil
}

func (s *Service) getCurrentSelectcols(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Trouble) (map[string]int, error) {
	stmts := make([]string, 0, len(models))
	for i := range models {
		stmts = append(stmts, fmt.Sprintf(`"from" = $%d and "where" = $%d`, (i*2)+1, (i*2)+2))
	}
	q := database.SQLSelectSimple(`"from", "where", "current_selectcol"`, tableQuoted, strings.Join(stmts, " or "))
	vals := make([]any, 0, len(models))
	for _, model := range models {
		vals = append(vals, model.From, model.Where)
	}
	var results []*struct {
		From             string   `db:"from"`
		Where            []string `db:"where"`
		CurrentSelectcol int      `db:"current_selectcol"`
	}
	err := s.dbRead.Select(ctx, &results, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Troubles")
	}

	ret := make(map[string]int, len(models))
	for _, model := range models {
		curr := 0
		for _, x := range results {
			if x.From == model.From && slices.Equal(x.Where, model.Where) {
				curr = x.CurrentSelectcol
			}
		}
		ret[model.String()] = curr
	}
	return ret, nil
}
