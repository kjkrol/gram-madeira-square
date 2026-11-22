package main

import (
	"github.com/kjkrol/aabbworld/geom"
	"github.com/kjkrol/aabbworld/plane"
	"github.com/kjkrol/ceremony/comp"
	"github.com/kjkrol/gram/entity/kind"
	"github.com/kjkrol/gram/plugins/board"
	"github.com/kjkrol/gram/plugins/board/cell"
	"github.com/kjkrol/gram/plugins/world"
	"github.com/kjkrol/gram/render"
)

type fieldRow struct {
	From  cell.ID `yaml:"from"`
	To    cell.ID `yaml:"to"`
	Cells int     `yaml:"cells"`
}

// field is a forest or a settlement of the map: the rectangle round its cells, which the simple
// map draws as they are, so nothing unfolds in it.
type field struct{}

var _ board.Figure = field{}

func (field) Define(b *board.Plugin) {
	w := b.Read(0).World()
	corners := func(r fieldRow) (geom.Vec, geom.Vec) {
		bd := b.Read(0).Board()
		return bd.CellCenter(r.From), bd.CellCenter(r.To)
	}
	for _, name := range []string{ForestKind, SettlementKind} {
		kind.Define[fieldRow](w.Kinds(), name, w.Roster().Spec(world.EntityTemplate,
			comp.Load(func(r fieldRow) world.Position {
				from, to := corners(r)
				corner := geom.NewVec((from.X+to.X-CellSize)/2, (from.Y+to.Y-CellSize)/2)
				return world.Position{AABB: plane.NewAABB(corner, CellSize, CellSize)}
			}),
			comp.Load(func(r fieldRow) world.Field {
				from, to := corners(r)
				return world.Field{Size: geom.NewVec(to.X-from.X+CellSize, to.Y-from.Y+CellSize)}
			}),
			comp.Const(world.Aggregate{})))
	}
}

func (field) Look(a *world.Atlas, w *world.Plugin) {
	for _, name := range []string{ForestKind, SettlementKind} {
		a.Add(kind.Named[fieldRow](w.Kinds(), name), CellSize, unseen)
	}
}

func unseen(*render.Canvas, int) {}
