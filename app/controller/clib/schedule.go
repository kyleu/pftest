// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/http"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/views/vadmin"
)

const scheduleBC = "Schedule||/admin/schedule**stopwatch"

func ScheduleList(w http.ResponseWriter, r *http.Request) {
	controller.Act("schedule.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		jobs := as.Services.Schedule.ListJobs()
		ps.SetTitleAndData("Schedules", jobs)
		return controller.Render(r, as, &vadmin.Schedule{Jobs: jobs}, ps, "admin", scheduleBC)
	})
}

func ScheduleDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("schedule.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		id, err := cutil.PathUUID(r, "id")
		if err != nil {
			return "", err
		}

		job := as.Services.Schedule.GetJob(*id)
		if job == nil {
			return controller.ERsp("no scheduled job with id [%s]", id)
		}
		res := as.Services.Schedule.Results[*id]

		ps.SetTitleAndData(job.ID.String(), job)
		return controller.Render(r, as, &vadmin.ScheduleDetail{Job: job, Result: res}, ps, "admin", scheduleBC, job.ID.String())
	})
}
