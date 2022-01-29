// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Audit struct {
	ID        uuid.UUID     `json:"id"`
	App       string        `json:"app"`
	Act       string        `json:"act"`
	Client    string        `json:"client"`
	Server    string        `json:"server"`
	User      string        `json:"user"`
	Metadata  util.ValueMap `json:"metadata"`
	Message   string        `json:"message"`
	Started   time.Time     `json:"started"`
	Completed time.Time     `json:"completed"`
}

func New(id uuid.UUID) *Audit {
	return &Audit{ID: id}
}

func Random() *Audit {
	return &Audit{
		ID:        util.UUID(),
		App:       util.RandomString(12),
		Act:       util.RandomString(12),
		Client:    util.RandomString(12),
		Server:    util.RandomString(12),
		User:      util.RandomString(12),
		Metadata:  util.RandomValueMap(4),
		Message:   util.RandomString(12),
		Started:   time.Now(),
		Completed: time.Now(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Audit, error) {
	ret := &Audit{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.App, err = m.ParseString("app", true, true)
	if err != nil {
		return nil, err
	}
	ret.Act, err = m.ParseString("act", true, true)
	if err != nil {
		return nil, err
	}
	ret.Client, err = m.ParseString("client", true, true)
	if err != nil {
		return nil, err
	}
	ret.Server, err = m.ParseString("server", true, true)
	if err != nil {
		return nil, err
	}
	ret.User, err = m.ParseString("user", true, true)
	if err != nil {
		return nil, err
	}
	ret.Metadata, err = m.ParseMap("metadata", true, true)
	if err != nil {
		return nil, err
	}
	ret.Message, err = m.ParseString("message", true, true)
	if err != nil {
		return nil, err
	}
	retStarted, e := m.ParseTime("started", true, true)
	if e != nil {
		return nil, e
	}
	if retStarted != nil {
		ret.Started = *retStarted
	}
	retCompleted, e := m.ParseTime("completed", true, true)
	if e != nil {
		return nil, e
	}
	if retCompleted != nil {
		ret.Completed = *retCompleted
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (a *Audit) Clone() *Audit {
	return &Audit{
		ID:        a.ID,
		App:       a.App,
		Act:       a.Act,
		Client:    a.Client,
		Server:    a.Server,
		User:      a.User,
		Metadata:  a.Metadata,
		Message:   a.Message,
		Started:   a.Started,
		Completed: a.Completed,
	}
}

func (a *Audit) String() string {
	return a.ID.String()
}

func (a *Audit) WebPath() string {
	return "/admin/audit" + "/" + a.ID.String()
}

func (a *Audit) Diff(ax *Audit) util.Diffs {
	var diffs util.Diffs
	if a.ID != ax.ID {
		diffs = append(diffs, util.NewDiff("id", a.ID.String(), ax.ID.String()))
	}
	if a.App != ax.App {
		diffs = append(diffs, util.NewDiff("app", a.App, ax.App))
	}
	if a.Act != ax.Act {
		diffs = append(diffs, util.NewDiff("act", a.Act, ax.Act))
	}
	if a.Client != ax.Client {
		diffs = append(diffs, util.NewDiff("client", a.Client, ax.Client))
	}
	if a.Server != ax.Server {
		diffs = append(diffs, util.NewDiff("server", a.Server, ax.Server))
	}
	if a.User != ax.User {
		diffs = append(diffs, util.NewDiff("user", a.User, ax.User))
	}
	diffs = append(diffs, util.DiffObjects(a.Metadata, ax.Metadata, "metadata")...)
	if a.Message != ax.Message {
		diffs = append(diffs, util.NewDiff("message", a.Message, ax.Message))
	}
	if a.Started != ax.Started {
		diffs = append(diffs, util.NewDiff("started", fmt.Sprint(a.Started), fmt.Sprint(ax.Started)))
	}
	if a.Completed != ax.Completed {
		diffs = append(diffs, util.NewDiff("completed", fmt.Sprint(a.Completed), fmt.Sprint(ax.Completed)))
	}
	return diffs
}

func (a *Audit) ToData() []interface{} {
	return []interface{}{a.ID, a.App, a.Act, a.Client, a.Server, a.User, a.Metadata, a.Message, a.Started, a.Completed}
}

type Audits []*Audit
