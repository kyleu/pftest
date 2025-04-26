package app

import (
	"context"

	"github.com/kyleu/pftest/app/audited"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/capital"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/mixedcase"
	"github.com/kyleu/pftest/app/oddpk"
	"github.com/kyleu/pftest/app/oddpk/oddrel"
	"github.com/kyleu/pftest/app/reference"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/seed"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/app/util"
)

type GeneratedServices struct {
	Capital   *capital.Service
	Audited   *audited.Service
	Basic     *basic.Service
	MixedCase *mixedcase.Service
	OddPK     *oddpk.Service
	Oddrel    *oddrel.Service
	Path      *path.Service
	Reference *reference.Service
	Relation  *relation.Service
	Seed      *seed.Service
	Softdel   *softdel.Service
	Timestamp *timestamp.Service
	Trouble   *trouble.Service
}

func initGeneratedServices(ctx context.Context, st *State, audSvc *audit.Service, logger util.Logger) GeneratedServices {
	return GeneratedServices{
		Capital:   capital.NewService(st.DB, st.DBRead),
		Audited:   audited.NewService(st.DB, audSvc),
		Basic:     basic.NewService(st.DB, st.DBRead),
		MixedCase: mixedcase.NewService(st.DB, st.DBRead),
		OddPK:     oddpk.NewService(st.DB, st.DBRead),
		Oddrel:    oddrel.NewService(st.DB, st.DBRead),
		Path:      path.NewService(st.DB, st.DBRead),
		Reference: reference.NewService(st.DB, st.DBRead),
		Relation:  relation.NewService(st.DB, st.DBRead),
		Seed:      seed.NewService(st.DB, st.DBRead),
		Softdel:   softdel.NewService(st.DB, st.DBRead),
		Timestamp: timestamp.NewService(st.DB, st.DBRead),
		Trouble:   trouble.NewService(st.DB, st.DBRead),
	}
}
