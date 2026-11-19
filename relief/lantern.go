package main

import (
	"github.com/kjkrol/ceremony/comp"
	"github.com/kjkrol/gram/entity/kind"
	"github.com/kjkrol/gram/plugins/board"
	"github.com/kjkrol/gram/plugins/board/cell"
	"github.com/kjkrol/gram/plugins/light"
	"github.com/kjkrol/gram/plugins/world"
)

type lanternRow struct {
	At cell.ID `yaml:"at"`
}

// lantern is a lamp standing at a bridge's head: it shines warm as far as it reaches and casts
// no shadow, its glass shining whatever the light.
type lantern struct{}

var _ board.Figure = lantern{}

func (lantern) Define(b *board.Plugin) {
	lanterns := board.NewEntities(b, board.Shape{Size: scale.Figure(1), Height: scale.Figure(4)},
		func(r lanternRow) cell.ID { return r.At })
	lanterns.Define(LanternKind,
		comp.Const(light.Source{Color: lanternColor, Strength: 1.5,
			Reach: float32(scale.Figure(10))}),
		comp.Const(light.Shadow{Cast: light.NoShadow}))
}

func (lantern) Look(a *world.Atlas, w *world.Plugin) {
	a.AddBody(kind.Named[lanternRow](w.Kinds(), LanternKind), 16, lampBody())
}

// lampBody is a post on a foot under a globe of glass, which shines, and a cap over it, a metre
// across and four tall: a lantern's and a settlement's lamp's
func lampBody() world.Body {
	m := scale.Figure
	return world.Body{
		Wide: m(1),
		Tall: m(4),
		Parts: []world.Part{
			{Shape: world.Box, To: [3]float32{0, 0, m(0.3)}, Wide: m(0.4), Tip: m(0.4)},
			{Shape: world.Box, To: [3]float32{0, 0, m(3.2)}, Wide: m(0.12), Tip: m(0.12)},
			{Shape: world.Spheroid, From: [3]float32{0, 0, m(3.1)}, To: [3]float32{0, 0, m(3.7)},
				Wide: m(0.6), Paint: 1},
			{Shape: world.Cone, From: [3]float32{0, 0, m(3.62)}, To: [3]float32{0, 0, m(3.92)},
				Wide: m(0.45), Tip: m(0.08)},
			{Shape: world.Spheroid, From: [3]float32{0, 0, m(3.9)}, To: [3]float32{0, 0, m(4)},
				Wide: m(0.1)},
		},
		Paints: []world.Paint{
			{Colour: pierColor},
			{Colour: lanternGlow, Shines: true},
		},
	}
}
