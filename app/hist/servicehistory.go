// Content managed by Project Forge, see [projectforge.md] for details.
package hist

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
)

var (
	historyColumns       = []string{"id", "hist_id", "o", "n", "c", "created"}
	historyColumnsQuoted = util.StringArrayQuoted(historyColumns)
	historyColumnsString = strings.Join(historyColumnsQuoted, ", ")

	historyTable       = table + "_history"
	historyTableQuoted = fmt.Sprintf("%q", historyTable)
)

func (s *Service) GetHistory(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*History, error) {
	q := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, "id = $1")
	ret := historyDTO{}
	err := s.dbRead.Get(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get hist history [%s]", id.String())
	}
	return ret.ToHistory(), nil
}

func (s *Service) GetHistories(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) (Histories, error) {
	q := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, "hist_id = $1")
	ret := historyDTOs{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get hists by id [%v]", id)
	}
	return ret.ToHistories(), nil
}

func (s *Service) SaveHistory(ctx context.Context, tx *sqlx.Tx, o *Hist, n *Hist, logger util.Logger) (*History, error) {
	diffs := o.Diff(n)
	if len(diffs) == 0 {
		return nil, nil
	}
	q := database.SQLInsert(historyTableQuoted, historyColumns, 1, "")
	h := &historyDTO{
		ID:     util.UUID(),
		HistID: o.ID,
		Old:    util.ToJSONBytes(o, true),
		New:    util.ToJSONBytes(n, true),
		Changes: util.ToJSONBytes(diffs, true),
		Created: time.Now(),
	}
	hist := h.ToHistory()
	err := s.db.Insert(ctx, q, tx, logger, hist.ToData()...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert hist")
	}
	return hist, nil
}
