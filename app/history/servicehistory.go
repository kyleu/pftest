// Content managed by Project Forge, see [projectforge.md] for details.
package history

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
	historyColumns       = []string{"id", "history_id", "o", "n", "c", "created"}
	historyColumnsQuoted = util.StringArrayQuoted(historyColumns)
	historyColumnsString = strings.Join(historyColumnsQuoted, ", ")

	historyTable       = table + "_history"
	historyTableQuoted = fmt.Sprintf("%q", historyTable)
)

func (s *Service) GetHistory(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*HistoryHistory, error) {
	q := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, "id = $1")
	ret := historyDTO{}
	err := s.db.Get(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get history history [%s]", id.String())
	}
	return ret.ToHistory(), nil
}

func (s *Service) GetHistories(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) (HistoryHistories, error) {
	q := database.SQLSelectSimple(historyColumnsString, historyTableQuoted, "history_id = $1")
	ret := historyDTOs{}
	err := s.db.Select(ctx, &ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get histories by id [%v]", id)
	}
	return ret.ToHistories(), nil
}

func (s *Service) SaveHistory(ctx context.Context, tx *sqlx.Tx, o *History, n *History, logger util.Logger) (*HistoryHistory, error) {
	q := database.SQLInsert(historyTableQuoted, historyColumns, 1, "")
	h := &historyDTO{
		ID:        util.UUID(),
		HistoryID: o.ID,
		Old:       util.ToJSONBytes(o, true),
		New:       util.ToJSONBytes(n, true),
		Changes:   util.ToJSONBytes(o.Diff(n), true),
		Created:   time.Now(),
	}
	hist := h.ToHistory()
	err := s.db.Insert(ctx, q, tx, logger, hist.ToData()...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to insert history")
	}
	return hist, nil
}
