// Content managed by Project Forge, see [projectforge.md] for details.
package version

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
	ID              string `db:"id"`
	CurrentRevision int    `db:"current_revision"`
}

func (s *Service) GetAllRevisions(ctx context.Context, tx *sqlx.Tx, id string, params *filter.Params, logger util.Logger) (Versions, error) {
	params = filters(params)
	wc := "\"id\" = $1"
	tablesJoinedParam := fmt.Sprintf("%q v join %q vr on v.\"id\" = vr.\"version_id\"", table, tableRevision)
	q := database.SQLSelect(columnsString, tablesJoinedParam, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Versions")
	}
	return ret.ToVersions(), nil
}

func (s *Service) GetRevision(ctx context.Context, tx *sqlx.Tx, id string, revision int, logger util.Logger) (*Version, error) {
	wc := "\"id\" = $1 and \"revision\" = $2"
	ret := &row{}
	tablesJoinedParam := fmt.Sprintf("%q v join %q vr on v.\"id\" = vr.\"version_id\"", table, tableRevision)
	q := database.SQLSelectSimple(columnsString, tablesJoinedParam, s.db.Placeholder(), wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id, revision)
	if err != nil {
		return nil, err
	}
	return ret.ToVersion(), nil
}

func (s *Service) getCurrentRevisions(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Version) (map[string]int, error) {
	stmts := lo.Map(models, func(_ *Version, i int) string {
		return fmt.Sprintf(`"id" = $%d`, i+1)
	})
	q := database.SQLSelectSimple(`"id", "current_revision"`, tableQuoted, s.db.Placeholder(), strings.Join(stmts, " or "))
	vals := lo.Map(models, func(model *Version, _ int) any {
		return model.ID
	})
	var results []*IDRev
	err := s.dbRead.Select(ctx, &results, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get Versions")
	}

	ret := make(map[string]int, len(models))
	lo.ForEach(models, func(model *Version, _ int) {
		curr := 0
		lo.ForEach(results, func(x *IDRev, _ int) {
			if x.ID == model.ID {
				curr = x.CurrentRevision
			}
		})
		ret[model.String()] = curr
	})
	return ret, nil
}
