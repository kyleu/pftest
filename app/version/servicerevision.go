// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) GetAllRevisions(ctx context.Context, tx *sqlx.Tx, id string, params *filter.Params, includeDeleted bool) (Versions, error) {
	params = filters(params)
	wc := "\"id\" = $1"
	tablesJoinedParam := fmt.Sprintf("%q v join %q vr on v.\"id\" = vr.\"version_id\"", table, tableRevision)
	q := database.SQLSelect(columnsString, tablesJoinedParam, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Versions")
	}
	return ret.ToVersions(), nil
}

func (s *Service) GetRevision(ctx context.Context, tx *sqlx.Tx, id string, revision int) (*Version, error) {
	wc := "\"id\" = $1 and \"revision\" = $2"
	ret := &dto{}
	tablesJoinedParam := fmt.Sprintf("%q v join %q vr on v.\"id\" = vr.\"version_id\"", table, tableRevision)
	q := database.SQLSelectSimple(columnsString, tablesJoinedParam, wc)
	err := s.db.Get(ctx, ret, q, tx, s.logger, id, revision)
	if err != nil {
		return nil, err
	}
	return ret.ToVersion(), nil
}

func (s *Service) getCurrentRevisions(ctx context.Context, tx *sqlx.Tx, models ...*Version) (map[string]int, error) {
	stmts := make([]string, 0, len(models))
	for i := range models {
		stmts = append(stmts, fmt.Sprintf(`"id" = $%d`, i+1))
	}
	q := database.SQLSelectSimple(`"id", "current_revision"`, tableQuoted, strings.Join(stmts, " or "))
	vals := make([]any, 0, len(models))
	for _, model := range models {
		vals = append(vals, model.ID)
	}
	var results []*struct {
		ID              string `db:"id"`
		CurrentRevision int    `db:"current_revision"`
	}
	err := s.db.Select(ctx, &results, q, tx, s.logger, vals...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Versions")
	}

	ret := make(map[string]int, len(models))
	for _, model := range models {
		curr := 0
		for _, x := range results {
			if x.ID == model.ID {
				curr = x.CurrentRevision
			}
		}
		ret[model.String()] = curr
	}
	return ret, nil
}
