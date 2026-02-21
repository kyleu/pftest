package app

import (
	"context"

	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/exec"
	"github.com/kyleu/pftest/app/lib/har"
	"github.com/kyleu/pftest/app/lib/help"
	"github.com/kyleu/pftest/app/lib/notebook"
	"github.com/kyleu/pftest/app/lib/schedule"
	"github.com/kyleu/pftest/app/lib/scripting"
	"github.com/kyleu/pftest/app/lib/settings"
	"github.com/kyleu/pftest/app/lib/task"
	"github.com/kyleu/pftest/app/lib/websocket"
	"github.com/kyleu/pftest/app/user"
	"github.com/kyleu/pftest/app/util"
)

type CoreServices struct {
	Audit    *audit.Service
	User     *user.Service
	Har      *har.Service
	Exec     *exec.Service
	Notebook *notebook.Service
	Schedule *schedule.Service
	Script   *scripting.Service
	Settings *settings.Service
	Socket   *websocket.Service
	Task     *task.Service
	Help     *help.Service
}

func initCoreServices(ctx context.Context, st *State, auditSvc *audit.Service, logger util.Logger) CoreServices {
	return CoreServices{
		Audit:    auditSvc,
		User:     user.NewService(st.Files, logger),
		Har:      har.NewService(st.Files),
		Exec:     exec.NewService(),
		Notebook: notebook.NewService(),
		Schedule: schedule.NewService(),
		Script:   scripting.NewService(st.Files, "scripts"),
		Settings: settings.NewService(st.Files),
		Socket:   websocket.NewService(nil, nil),
		Task:     task.NewService(st.Files, "task_history"),
		Help:     help.NewService(logger),
	}
}
