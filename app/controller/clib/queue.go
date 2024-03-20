// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"context"
	"net/http"
	"time"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/queue"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vadmin"
)

var QueueInstance *queue.Queue

func initQueueInstance(ctx context.Context, logger util.Logger) (*queue.Status, error) {
	if QueueInstance == nil {
		db, err := database.OpenSQLiteDatabase(ctx, "queue", &database.SQLiteParams{File: "tmp/queue.sqlite"}, logger)
		if err != nil {
			return nil, err
		}
		QueueInstance, err = queue.New(ctx, "queue", 0, 10*time.Second, "", db, logger)
		if err != nil {
			return nil, err
		}
	}
	return QueueInstance.Status(), nil
}

func QueueIndex(w http.ResponseWriter, r *http.Request) {
	controller.Act("queue.index", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		st, err := initQueueInstance(ps.Context, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Queue", st)
		return controller.Render(w, r, as, &vadmin.Queue{Status: st}, ps, "admin", "Queue")
	})
}

func QueueSend(w http.ResponseWriter, r *http.Request) {
	controller.Act("queue.send", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		st, err := initQueueInstance(ps.Context, ps.Logger)
		if err != nil {
			return "", err
		}
		msg := queue.NewMessage("foo", "OK")
		err = QueueInstance.Send(ps.Context, nil, msg, ps.Logger)
		if err != nil {
			return "", err
		}
		msg, err = QueueInstance.Receive(ps.Context, nil, "foo", ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Queue", util.ValueMap{"message": msg, "status": st})
		return controller.Render(w, r, as, &vadmin.Queue{Status: st, Message: msg}, ps, "admin", "Queue")
	})
}
