// Content managed by Project Forge, see [projectforge.md] for details.
package capital

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

func (s *Service) GetAllVersions(ctx context.Context, tx *sqlx.Tx, id string, params *filter.Params, logger util.Logger) (Capitals, error) {
	params = filters(params)
	wc := "\"ID\" = $1"
	tablesJoinedParam := fmt.Sprintf("%q c join %q cr on c.\"ID\" = cr.\"Capital_ID\"", table, tableVersion)
	q := database.SQLSelect(columnsString, tablesJoinedParam, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Capitals")
	}
	return ret.ToCapitals(), nil
}

func (s *Service) GetVersion(ctx context.Context, tx *sqlx.Tx, id string, version int, logger util.Logger) (*Capital, error) {
	wc := "\"ID\" = $1 and \"Version\" = $2"
	ret := &dto{}
	tablesJoinedParam := fmt.Sprintf("%q c join %q cr on c.\"ID\" = cr.\"Capital_ID\"", table, tableVersion)
	q := database.SQLSelectSimple(columnsString, tablesJoinedParam, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id, version)
	if err != nil {
		return nil, err
	}
	return ret.ToCapital(), nil
}

func (s *Service) getCurrentVersions(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Capital) (map[string]int, error) {
	stmts := make([]string, 0, len(models))
	for i := range models {
		stmts = append(stmts, fmt.Sprintf(`"ID" = $%d`, i+1))
	}
	q := database.SQLSelectSimple(`"ID", "current_Version"`, tableQuoted, strings.Join(stmts, " or "))
	vals := make([]any, 0, len(models))
	for _, model := range models {
		vals = append(vals, model.ID)
	}
	var results []*struct {
		ID             string `db:"ID"`
		CurrentVersion int    `db:"current_Version"`
	}
	err := s.dbRead.Select(ctx, &results, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Capitals")
	}

	ret := make(map[string]int, len(models))
	for _, model := range models {
		curr := 0
		for _, x := range results {
			if x.ID == model.ID {
				curr = x.CurrentVersion
			}
		}
		ret[model.String()] = curr
	}
	return ret, nil
}
