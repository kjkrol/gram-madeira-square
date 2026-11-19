package main

import (
	"image/color"
	"math"

	"github.com/kjkrol/gram/plugins/forest"
	"github.com/kjkrol/gram/plugins/world"
)

// foliage is what a tree is painted in: its crown, the lit leaves on it, its trunk, and what it
// bears, fruit or blossom, where it bears any
type foliage struct {
	crown  color.RGBA
	leaves color.RGBA
	trunk  color.RGBA
	bears  color.RGBA
}

// Madeira's trees, each a body of clumps of leaves and limbs of wood, sized in metres of life. No
// crown is wider than a cell, so each is as tall as a tree of its kind stands under such a crown,
// a young one beside a cottage of 4 m, not a veteran of 25. Of the laurisilva: the til, the
// laurel, the vinhático, the barbusano and the faia; of the heights: the tree heather and the
// Madeira cedar; brought in: the maritime pine, the eucalyptus, the acacia and the chestnut; of
// the coast and the terraces: the dragon tree, the palm, the banana and the vine.
var trees = map[string]forest.Spec{
	TilKind: broadleaf(10, 6.5, foliage{
		crown:  color.RGBA{R: 22, G: 70, B: 38, A: 255},
		leaves: color.RGBA{R: 50, G: 104, B: 52, A: 255},
		trunk:  color.RGBA{R: 62, G: 48, B: 36, A: 255},
	}),
	LaurelKind: broadleaf(7, 5.5, foliage{
		crown:  color.RGBA{R: 30, G: 84, B: 42, A: 255},
		leaves: color.RGBA{R: 64, G: 122, B: 58, A: 255},
		trunk:  color.RGBA{R: 72, G: 54, B: 38, A: 255},
		bears:  color.RGBA{R: 226, G: 218, B: 150, A: 255},
	}),
	VinhaticoKind: broadleaf(9, 6, foliage{
		crown:  color.RGBA{R: 44, G: 96, B: 48, A: 255},
		leaves: color.RGBA{R: 150, G: 96, B: 60, A: 255},
		trunk:  color.RGBA{R: 96, G: 64, B: 44, A: 255},
	}),
	BarbusanoKind: broadleaf(8, 5.5, foliage{
		crown:  color.RGBA{R: 40, G: 76, B: 36, A: 255},
		leaves: color.RGBA{R: 86, G: 112, B: 48, A: 255},
		trunk:  color.RGBA{R: 58, G: 46, B: 38, A: 255},
	}),
	FaiaKind: broadleaf(6, 5, foliage{
		crown:  color.RGBA{R: 66, G: 110, B: 46, A: 255},
		leaves: color.RGBA{R: 112, G: 150, B: 64, A: 255},
		trunk:  color.RGBA{R: 90, G: 74, B: 56, A: 255},
		bears:  color.RGBA{R: 128, G: 44, B: 62, A: 255},
	}),
	ChestnutKind: broadleaf(9, 6.5, foliage{
		crown:  color.RGBA{R: 58, G: 112, B: 44, A: 255},
		leaves: color.RGBA{R: 132, G: 160, B: 62, A: 255},
		trunk:  color.RGBA{R: 84, G: 62, B: 44, A: 255},
		bears:  color.RGBA{R: 198, G: 206, B: 104, A: 255},
	}),
	AcaciaKind: broadleaf(7, 5, foliage{
		crown:  color.RGBA{R: 70, G: 104, B: 62, A: 255},
		leaves: color.RGBA{R: 214, G: 190, B: 70, A: 255},
		trunk:  color.RGBA{R: 70, G: 58, B: 48, A: 255},
	}),
	HeatherKind: bush(4, 4.5, foliage{
		crown:  color.RGBA{R: 98, G: 112, B: 62, A: 255},
		leaves: color.RGBA{R: 156, G: 160, B: 100, A: 255},
		trunk:  color.RGBA{R: 92, G: 72, B: 52, A: 255},
		bears:  color.RGBA{R: 232, G: 220, B: 226, A: 255},
	}),
	VineKind: bush(1.8, 2.5, foliage{
		crown:  color.RGBA{R: 96, G: 132, B: 52, A: 255},
		leaves: color.RGBA{R: 148, G: 172, B: 70, A: 255},
		trunk:  color.RGBA{R: 86, G: 66, B: 48, A: 255},
		bears:  color.RGBA{R: 96, G: 62, B: 124, A: 255},
	}),
	PineKind: conifer(10, 3.75, foliage{
		crown:  color.RGBA{R: 38, G: 90, B: 70, A: 255},
		leaves: color.RGBA{R: 54, G: 108, B: 82, A: 255},
		trunk:  color.RGBA{R: 84, G: 58, B: 40, A: 255},
		bears:  color.RGBA{R: 124, G: 88, B: 56, A: 255},
	}),
	CedarKind: conifer(8, 5, foliage{
		crown:  color.RGBA{R: 44, G: 82, B: 84, A: 255},
		leaves: color.RGBA{R: 86, G: 124, B: 122, A: 255},
		trunk:  color.RGBA{R: 104, G: 72, B: 56, A: 255},
	}),
	EucalyptusKind: slender(12, 5, foliage{
		crown:  color.RGBA{R: 84, G: 116, B: 96, A: 255},
		leaves: color.RGBA{R: 134, G: 158, B: 136, A: 255},
		trunk:  color.RGBA{R: 196, G: 184, B: 164, A: 255},
	}),
	DragonKind: umbrella(5.5, 6, foliage{
		crown:  color.RGBA{R: 62, G: 98, B: 66, A: 255},
		leaves: color.RGBA{R: 110, G: 140, B: 92, A: 255},
		trunk:  color.RGBA{R: 150, G: 132, B: 112, A: 255},
		bears:  color.RGBA{R: 222, G: 132, B: 52, A: 255},
	}),
	PalmKind: fronded(8, 5, 0.6, 0.08, foliage{
		crown:  color.RGBA{R: 52, G: 110, B: 48, A: 255},
		leaves: color.RGBA{R: 98, G: 150, B: 60, A: 255},
		trunk:  color.RGBA{R: 120, G: 92, B: 64, A: 255},
		bears:  color.RGBA{R: 224, G: 150, B: 52, A: 255},
	}),
	BananaKind: fronded(3, 3.5, 0.25, 0.16, foliage{
		crown:  color.RGBA{R: 72, G: 136, B: 50, A: 255},
		leaves: color.RGBA{R: 136, G: 180, B: 66, A: 255},
		trunk:  color.RGBA{R: 112, G: 124, B: 62, A: 255},
		bears:  color.RGBA{R: 226, G: 208, B: 84, A: 255},
	}),
}

// colours are f as a tree's body is painted in it
func (f foliage) colours() forest.Colours {
	return forest.Colours{Crown: f.crown, Leaves: f.leaves, Trunk: f.trunk, Bears: f.bears}
}

// body is a tree's body given in crown widths off its foot, across, deep and up, laid in world
// units for a crown as wide as crown metres of life
type body struct {
	crown float32
	parts []world.Part
}

func bodyOf(crown float32) *body { return &body{crown: scale.Figure(crown)} }

// clump adds a clump of leaves r round at (x, y) off the trunk and z up
func (b *body) clump(x, y, z, r float32) { b.tuft(x, y, z, r, r) }

// tuft adds a clump of leaves r round across and half as high as half says, at (x, y) off the
// trunk and z up
func (b *body) tuft(x, y, z, r, half float32) {
	c := b.crown
	b.parts = append(b.parts, world.Part{
		Shape: world.Spheroid,
		From:  [3]float32{x * c, y * c, (z - half) * c},
		To:    [3]float32{x * c, y * c, (z + half) * c},
		Wide:  2 * r * c,
		Paint: forest.Leaves,
	})
}

// frond adds a long leaf from one point to another, wide across its middle
func (b *body) frond(from, to [3]float32, wide float32) {
	c := b.crown
	b.parts = append(b.parts, world.Part{
		Shape: world.Spheroid,
		From:  [3]float32{from[0] * c, from[1] * c, from[2] * c},
		To:    [3]float32{to[0] * c, to[1] * c, to[2] * c},
		Wide:  wide * c,
		Paint: forest.Leaves,
	})
}

// limb adds a limb of wood from one point to another, wide at its foot and tip at its end
func (b *body) limb(from, to [3]float32, wide, tip float32) {
	c := b.crown
	b.parts = append(b.parts, world.Part{
		Shape: world.Cone,
		From:  [3]float32{from[0] * c, from[1] * c, from[2] * c},
		To:    [3]float32{to[0] * c, to[1] * c, to[2] * c},
		Wide:  wide * c,
		Tip:   tip * c,
		Paint: forest.Wood,
	})
}

// tier adds a tier of boughs, a cone half wide to either side at base, its tip up at tip
func (b *body) tier(base, tip, half float32) {
	c := b.crown
	b.parts = append(b.parts, world.Part{
		Shape: world.Cone,
		From:  [3]float32{0, 0, base * c},
		To:    [3]float32{0, 0, tip * c},
		Wide:  2 * half * c,
		Paint: forest.Leaves,
	})
}

// spec is the kind of tree of the body b, height and crown metres of life, swaying as far as
// sway says
func (b *body) spec(height, crown, sway float32, f foliage) forest.Spec {
	return forest.Spec{
		Height:  scale.Figure(height),
		Crown:   scale.Figure(crown),
		Sway:    sway,
		Colours: f.colours(),
		Parts:   b.parts,
	}
}

// broadleaf is a tree of a round crown in clumps on a trunk forking under it
func broadleaf(height, crown float32, f foliage) forest.Spec {
	top := height / crown
	b := bodyOf(crown)
	b.limb([3]float32{}, [3]float32{0, 0, top - 0.8}, 0.15, 0.09)
	b.limb([3]float32{0, 0, top - 1}, [3]float32{-0.2, 0.08, top - 0.7}, 0.06, 0.04)
	b.limb([3]float32{0, 0, top - 1}, [3]float32{0.2, -0.08, top - 0.7}, 0.06, 0.04)
	// four clumps round the trunk under three round it higher, one over them all
	for _, at := range [][2]float32{{-0.19, -0.19}, {0.19, -0.19}, {0.19, 0.19}, {-0.19, 0.19}} {
		b.clump(at[0], at[1], top-0.66, 0.24)
	}
	for _, at := range [][2]float32{{0, -0.2}, {0.18, 0.1}, {-0.18, 0.1}} {
		b.clump(at[0], at[1], top-0.42, 0.27)
	}
	b.clump(0, 0, top-0.28, 0.28)
	return b.spec(height, crown, 0.35, f)
}

// conifer is a tree of boughs in tiers narrowing to its top, each hanging over the one under it
func conifer(height, crown float32, f foliage) forest.Spec {
	top := height / crown
	foot := top * 0.12
	tiers := min(max(int(top*2), 3), 7)
	rise := (top - foot) / float32(tiers)
	b := bodyOf(crown)
	// its trunk runs up to its tip under the boughs: no tier stands off it
	b.limb([3]float32{}, [3]float32{0, 0, top - 0.05}, 0.1, 0.02)
	for i := range tiers {
		base := foot + rise*float32(i)
		b.tier(base, min(base+rise*1.3, top), 0.5*(1-float32(i)/float32(tiers)))
	}
	return b.spec(height, crown, 0.3, f)
}

// bush is a low thing of three stems, wider than tall
func bush(height, crown float32, f foliage) forest.Spec {
	top := height / crown
	b := bodyOf(crown)
	b.limb([3]float32{-0.07, 0, 0}, [3]float32{-0.21, -0.04, top * 0.4}, 0.05, 0.04)
	b.limb([3]float32{0, 0.03, 0}, [3]float32{0, 0.05, top * 0.45}, 0.05, 0.04)
	b.limb([3]float32{0.07, 0, 0}, [3]float32{0.21, 0.04, top * 0.4}, 0.05, 0.04)
	// a low bush's clumps are smaller, so its top one stays under its height
	k := min(top/0.9, 1)
	b.clump(-0.29, -0.05, top*0.38, 0.2*k)
	b.clump(0.29, 0.05, top*0.38, 0.2*k)
	b.clump(0.13, -0.15, top*0.5, 0.24*k)
	b.clump(-0.13, 0.15, top*0.5, 0.25*k)
	b.clump(0, 0, top*0.7, 0.26*k)
	return b.spec(height, crown, 0.5, f)
}

// slender is a tall tree, a long pale trunk under a narrow crown in loose tufts, a pair a
// storey down its trunk, each storey turned a quarter off the one over it
func slender(height, crown float32, f foliage) forest.Spec {
	top := height / crown
	b := bodyOf(crown)
	b.limb([3]float32{}, [3]float32{0, 0, top - 0.4}, 0.13, 0.06)
	const storeys = 4
	for i := range storeys {
		z := top - 0.38 - float32(i)*0.33
		sin, cos := math.Sincos(float64(i) * math.Pi / 2)
		x, y := float32(cos)*0.18, float32(sin)*0.18
		b.clump(x, y, z-0.05, 0.2)
		b.clump(-x, -y, z+0.05, 0.22)
	}
	b.clump(0, 0, top-0.2, 0.19)
	return b.spec(height, crown, 0.6, f)
}

// umbrella is a tree of a thick bare trunk forking under a flat crown of tufts
func umbrella(height, crown float32, f foliage) forest.Spec {
	top := height / crown
	fork := top * 0.34
	b := bodyOf(crown)
	b.limb([3]float32{}, [3]float32{0, 0, fork}, 0.22, 0.16)
	for i := range 4 {
		sin, cos := math.Sincos(float64(i)*math.Pi/2 + math.Pi/4)
		b.limb([3]float32{0, 0, fork - 0.04},
			[3]float32{float32(cos) * 0.26, float32(sin) * 0.26, top - 0.3}, 0.11, 0.07)
	}
	// its tufts in a dome, each flatter than round
	for i := range 6 {
		sin, cos := math.Sincos(float64(i) * math.Pi / 3)
		b.tuft(float32(cos)*0.3, float32(sin)*0.3, top-0.24, 0.17, 0.1)
	}
	b.tuft(0, 0, top-0.14, 0.2, 0.1)
	return b.spec(height, crown, 0.15, f)
}

// fronded is a plant of long leaves arching out of the top of its stem, the stem the share stem
// of its height, each leaf broad crown widths broad
func fronded(height, crown, stem, broad float32, f foliage) forest.Spec {
	top := height / crown
	from := top * stem
	b := bodyOf(crown)
	b.limb([3]float32{}, [3]float32{0, 0, from}, 0.12, 0.09)
	// its leaves stand out all round, every other one higher, none past its crown's width
	reach := min(top-from, 0.5)
	out := 0.48 - broad
	const leaves = 8
	for i := range leaves {
		sin, cos := math.Sincos(float64(i) * 2 * math.Pi / leaves)
		up := reach * 0.3
		if i%2 == 0 {
			up = reach * 0.72
		}
		b.frond([3]float32{0, 0, from},
			[3]float32{float32(cos) * out, float32(sin) * out, from + up}, 2*broad)
	}
	b.clump(0, 0, from+min(reach*0.6, top-from-0.1), 0.1)
	return b.spec(height, crown, 0.65, f)
}
