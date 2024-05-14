package hg

import (
	"context"
	"database/sql"

	_ "github.com/it512/loong/hg/internal/ent/runtime"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/it512/loong/hg/internal/ent"
	"github.com/it512/loong/hg/internal/ent/hg"
)

func open(db *sql.DB) *ent.Client {
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

type HumanGroup interface {
	LoadGroup(ctx context.Context, code string) ([]*Hg, error)
	LoadHuman(ctx context.Context, code string) ([]*Hg, error)
}

type PGHumanGroup struct {
	client *ent.Client
}

func New(db *sql.DB) *PGHumanGroup {
	return &PGHumanGroup{client: open(db)}
}

func (h *PGHumanGroup) LoadHuman(ctx context.Context, gcode string) (rr []*Hg, err error) {
	var r []*ent.Hg
	q := h.client.Hg.Query()
	r, err = q.Where(hg.GroupCodeEQ(gcode)).All(ctx)
	if ent.IsNotFound(err) {
		return rr, nil
	}
	for _, h := range r {
		rr = append(rr, &Hg{
			HumanCode: h.HumanCode,
			HumanName: h.HumanName,
			GroupCode: h.GroupCode,
			GroupName: h.GroupName,
		})
	}

	return
}

func (h *PGHumanGroup) LoadGroup(ctx context.Context, hcode string) (rr []*Hg, err error) {
	var r []*ent.Hg
	q := h.client.Hg.Query()
	r, err = q.Where(hg.HumanCodeEQ(hcode)).All(ctx)
	if ent.IsNotFound(err) {
		return rr, nil
	}

	for _, h := range r {
		rr = append(rr, &Hg{
			HumanCode: h.HumanCode,
			HumanName: h.HumanName,
			GroupCode: h.GroupCode,
			GroupName: h.GroupName,
		})
	}
	rr = append(rr, &Hg{HumanCode: hcode, GroupCode: hcode})
	return
}
