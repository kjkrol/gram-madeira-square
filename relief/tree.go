package main

import (
	"image/color"
	"math"

	"github.com/kjkrol/gram/plugins/forest"
	"github.com/kjkrol/gram/render"
)

// foliage is what a tree is painted in: its crown, the lit leaves on it, its trunk, and what it
// bears, fruit or blossom, where it bears any
type foliage struct {
	crown  color.RGBA
	leaves color.RGBA
	trunk  color.RGBA
	bears  color.RGBA
}

// Madeira's trees, each drawn from above and from the side, its side spanning its crown across and
// its height down, sized in metres of life. No crown is wider than a cell, so each is as tall as a
// tree of its kind stands under such a crown, a young one beside a cottage of 4 m, not a veteran of
// 25. Of the laurisilva: the til, the laurel, the vinhático, the barbusano and the faia; of the
// heights: the tree heather and the Madeira cedar; brought in: the maritime pine, the eucalyptus,
// the acacia and the chestnut; of the coast and the terraces: the dragon tree, the palm, the banana
// and the vine.
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

// side draws a tree from the side on a square stretched to its crown across and its height
// down: w is the crown's width in pixels, and a point is given across in crown widths from the
// left and up in crown widths from the ground, so a disc stays round on the tree
type side struct {
	dst    *render.Canvas
	w      float32
	aspect float32 // the tree's height over its crown's width
	f      foliage
}

func sideOf(dst *render.Canvas, size int, height, crown float32, f foliage) side {
	return side{dst: dst, w: float32(size), aspect: height / crown, f: f}
}

// how much of its light a colour keeps in the shade
const shade = 0.68

// dim is c with the share k of its light, lighter past 1
func dim(c color.RGBA, k float32) color.RGBA {
	lit := func(v uint8) uint8 { return uint8(min(float32(v)*k, 255)) }
	return color.RGBA{R: lit(c.R), G: lit(c.G), B: lit(c.B), A: c.A}
}

// wobble is the i-th of the shares between -1 and 1 that are the tree's own, by what it is
// painted in: two kinds of one form grow apart by them
func (s side) wobble(i int) float32 {
	c := s.f.crown
	h := (uint32(c.R)<<16|uint32(c.G)<<8|uint32(c.B))*2654435761 + uint32(i)*40503
	h ^= h >> 15
	h *= 2246822519
	h ^= h >> 13
	return float32(h%2001)/1000 - 1
}

// disc fills a disc r crown widths round at x across and y up
func (s side) disc(x, y, r float32, col color.RGBA) {
	cy, ry := s.w*(1-y/s.aspect), s.w*r/s.aspect
	for row := float32(math.Floor(float64(cy - ry))); row <= cy+ry; row++ {
		d := (row + 0.5 - cy) / ry
		if d*d >= 1 {
			continue
		}
		half := s.w * r * float32(math.Sqrt(float64(1-d*d)))
		s.dst.FillRect(s.w*x-half, row, 2*half, 1, col)
	}
}

// rows fills every row of pixels between the heights y0 and y1 from left to right, as across
// says for the row's height
func (s side) rows(y0, y1 float32, across func(y float32) (left, right float32), col color.RGBA) {
	for row := float32(0); row < s.w; row++ {
		y := (1 - (row+0.5)/s.w) * s.aspect
		if y < y0 || y > y1 {
			continue
		}
		if left, right := across(y); right > left {
			s.dst.FillRect(s.w*left, row, s.w*(right-left), 1, col)
		}
	}
}

// trunk fills a trunk at x across from the ground up to y, foot wide at the ground and neck
// wide at its top, its right side in the shade
func (s side) trunk(x, y, foot, neck float32) {
	wide := func(at float32) float32 { return foot + (neck-foot)*at/y }
	s.rows(0, y, func(at float32) (float32, float32) {
		return x - wide(at)/2, x + wide(at)/2
	}, s.f.trunk)
	s.rows(0, y, func(at float32) (float32, float32) {
		return x + wide(at)/6, x + wide(at)/2
	}, dim(s.f.trunk, shade))
}

// bough strokes a line from one point to another, wide crown widths wide
func (s side) bough(x0, y0, x1, y1, wide float32, col color.RGBA) {
	s.dst.StrokeLine(s.w*x0, s.w*(1-y0/s.aspect), s.w*x1, s.w*(1-y1/s.aspect),
		max(s.w*wide/s.aspect, 1), col)
}

// tuft fills a tuft of leaves r round at x across and y up, its lower right in the shade
func (s side) tuft(x, y, r float32) {
	s.disc(x, y, r, dim(s.f.crown, shade))
	s.disc(x-r*0.1, y+r*0.1, r*0.9, s.f.crown)
}

// clump fills a tuft of leaves with the light on its upper left
func (s side) clump(x, y, r float32) {
	s.tuft(x, y, r)
	s.disc(x-r*0.3, y+r*0.3, r*0.36, s.f.leaves)
}

// crown fills the clumps given, each across, up and round, from the hindmost to the one in
// front, each a little off its place and its size as the tree's own wobble has it, none past the
// crown's width or over the height top
func (s side) crown(top float32, clumps ...[3]float32) {
	for i, c := range clumps {
		r := c[2] * (1 + 0.14*s.wobble(3*i))
		x := min(max(c[0]+0.05*s.wobble(3*i+1), r), 1-r)
		y := min(c[1]+0.05*s.wobble(3*i+2), top-r)
		s.clump(x, y, r)
	}
}

// bear dots what the tree bears at the points given, each across and up, where the tree is
// drawn already: nothing hangs in the air, and a tree that bears nothing shows nothing
func (s side) bear(r float32, at ...[2]float32) {
	if s.f.bears == (color.RGBA{}) {
		return
	}
	for _, p := range at {
		if s.dst.RGBAAt(int(s.w*p[0]), int(s.w*(1-p[1]/s.aspect))).A == 255 {
			s.disc(p[0], p[1], r, s.f.bears)
		}
	}
}

// broadleaf is a tree of a round crown in clumps on a trunk forking under it
func broadleaf(height, crown float32, f foliage) forest.Spec {
	return forest.Spec{
		Height: scale.Figure(height),
		Crown:  scale.Figure(crown),
		Sway:   0.35,
		Top: func(dst *render.Canvas, size int) {
			s := float32(size)
			dst.FillCircle(s/2, s/2, s/2, f.crown)
			dst.FillCircle(s*0.36, s*0.36, s*0.18, f.leaves)
			dst.FillCircle(s*0.64, s*0.32, s*0.12, f.leaves)
			dst.FillCircle(s*0.32, s*0.64, s*0.1, f.leaves)
		},
		Side: func(dst *render.Canvas, size int) {
			s := sideOf(dst, size, height, crown, f)
			top := s.aspect
			s.trunk(0.5, top-0.8, 0.15, 0.09)
			s.bough(0.5, top-1, 0.3, top-0.7, 0.06, f.trunk)
			s.bough(0.5, top-1, 0.7, top-0.7, 0.06, dim(f.trunk, shade))
			s.crown(top,
				[3]float32{0.24, top - 0.66, 0.24},
				[3]float32{0.76, top - 0.66, 0.24},
				[3]float32{0.5, top - 0.7, 0.22},
				[3]float32{0.72, top - 0.42, 0.27},
				[3]float32{0.3, top - 0.42, 0.28},
				[3]float32{0.52, top - 0.28, 0.28},
				[3]float32{0.5, top - 0.52, 0.26},
			)
			s.bear(0.035,
				[2]float32{0.2, top - 0.6},
				[2]float32{0.4, top - 0.74},
				[2]float32{0.62, top - 0.3},
				[2]float32{0.82, top - 0.5},
				[2]float32{0.66, top - 0.64},
				[2]float32{0.4, top - 0.2},
			)
		},
	}
}

// conifer is a tree of boughs in tiers narrowing to its top, each hanging over the one under it
func conifer(height, crown float32, f foliage) forest.Spec {
	return forest.Spec{
		Height: scale.Figure(height),
		Crown:  scale.Figure(crown),
		Sway:   0.3,
		Top: func(dst *render.Canvas, size int) {
			s := float32(size)
			dst.FillCircle(s/2, s/2, s/2, f.crown)
			for i := range 8 {
				sin, cos := math.Sincos(float64(i) * math.Pi / 4)
				dst.StrokeLine(s/2, s/2, s/2+float32(cos)*s*0.42, s/2+float32(sin)*s*0.42, s/16,
					f.leaves)
			}
			dst.FillCircle(s/2, s/2, s*0.1, f.leaves)
		},
		Side: func(dst *render.Canvas, size int) {
			s := sideOf(dst, size, height, crown, f)
			top := s.aspect
			foot := top * 0.12
			// its trunk runs up to its tip under the boughs: no tier stands off it
			s.trunk(0.5, top-0.05, 0.1, 0.02)
			tiers := min(max(int(top*2), 3), 7)
			rise := (top - foot) / float32(tiers)
			for i := range tiers {
				base := foot + rise*float32(i)
				half := 0.5 * (1 - float32(i)/float32(tiers))
				tip := min(base+rise*1.7, top)
				if i > 0 {
					s.tier(base-rise*0.14, tip, half, dim(f.crown, shade))
				}
				s.tier(base, tip, half, f.crown)
				s.bough(0.5-half*0.74, base+rise*0.08, 0.5-half*0.36, base+rise*0.8, 0.05,
					f.leaves)
			}
			s.bear(0.03,
				[2]float32{0.4, foot + rise*0.5},
				[2]float32{0.64, foot + rise*0.4},
				[2]float32{0.56, foot + rise*1.5},
			)
		},
	}
}

// tier fills a tier of boughs from base up to its tip, half crown widths to either side at its
// base, its hem hanging in scallops
func (s side) tier(base, tip, half float32, col color.RGBA) {
	s.rows(base, tip, func(y float32) (float32, float32) {
		reach := half * (tip - y) / (tip - base)
		return 0.5 - reach, 0.5 + reach
	}, col)
	// a scallop a tenth of the crown wide, three at the least
	n := max(int(half/0.1+0.5), 3)
	hem := half / float32(n)
	for i := range n {
		s.disc(0.5-half+hem*float32(2*i+1), base+hem*0.2, hem, col)
	}
}

// bush is a low thing of many stems, wider than tall
func bush(height, crown float32, f foliage) forest.Spec {
	return forest.Spec{
		Height: scale.Figure(height),
		Crown:  scale.Figure(crown),
		Sway:   0.5,
		Top: func(dst *render.Canvas, size int) {
			s := float32(size)
			dst.FillCircle(s/2, s/2, s*0.4, f.crown)
			for i := range 6 {
				sin, cos := math.Sincos(float64(i) * math.Pi / 3)
				dst.FillCircle(s/2+float32(cos)*s*0.3, s/2+float32(sin)*s*0.3, s*0.18, f.crown)
			}
			dst.FillCircle(s*0.38, s*0.36, s*0.12, f.leaves)
			dst.FillCircle(s*0.62, s*0.6, s*0.08, f.leaves)
		},
		Side: func(dst *render.Canvas, size int) {
			s := sideOf(dst, size, height, crown, f)
			top := s.aspect
			// up into its clumps, whichever way they wobble
			s.bough(0.29, top*0.4, 0.43, 0, 0.05, f.trunk)
			s.bough(0.5, top*0.45, 0.5, 0, 0.05, f.trunk)
			s.bough(0.71, top*0.4, 0.57, 0, 0.05, dim(f.trunk, shade))
			// a low bush's clumps are smaller, so its top one stays under its height
			k := min(top/0.9, 1)
			s.crown(top,
				[3]float32{0.2, top * 0.38, 0.2 * k},
				[3]float32{0.8, top * 0.38, 0.2 * k},
				[3]float32{0.64, top * 0.5, 0.24 * k},
				[3]float32{0.36, top * 0.5, 0.25 * k},
				[3]float32{0.5, top * 0.7, 0.26 * k},
			)
			s.bear(0.03,
				[2]float32{0.16, top * 0.44},
				[2]float32{0.3, top * 0.66},
				[2]float32{0.44, top * 0.4},
				[2]float32{0.56, top * 0.84},
				[2]float32{0.62, top * 0.56},
				[2]float32{0.78, top * 0.6},
				[2]float32{0.86, top * 0.36},
			)
		},
	}
}

// slender is a tall tree, a long pale trunk shedding its bark under a narrow crown in loose tufts
func slender(height, crown float32, f foliage) forest.Spec {
	return forest.Spec{
		Height: scale.Figure(height),
		Crown:  scale.Figure(crown),
		Sway:   0.6,
		Top: func(dst *render.Canvas, size int) {
			s := float32(size)
			dst.FillCircle(s*0.42, s*0.4, s*0.34, f.crown)
			dst.FillCircle(s*0.62, s*0.6, s*0.3, f.crown)
			dst.FillCircle(s*0.36, s*0.34, s*0.14, f.leaves)
			dst.FillCircle(s*0.64, s*0.62, s*0.1, f.leaves)
		},
		Side: func(dst *render.Canvas, size int) {
			s := sideOf(dst, size, height, crown, f)
			top := s.aspect
			s.trunk(0.5, top-0.4, 0.13, 0.06)
			// the bark it sheds, in strips down its trunk
			for i := range 4 {
				y := top * (0.08 + 0.1*float32(i))
				s.bough(0.47+0.03*float32(i%2), y, 0.47+0.03*float32(i%2), y+top*0.05, 0.025,
					dim(f.trunk, 0.8))
			}
			// its crown in storeys down its trunk, each a clump to either side, one higher than
			// the other in turn, and one between them, on boughs of their own, the lowest's
			// leaves hanging
			const storeys = 4
			var clumps [][3]float32
			for i := storeys - 1; i >= 0; i-- {
				y := top - 0.3 - float32(i)*0.33
				lean := float32(1 - 2*(i%2))
				s.bough(0.5, y-0.34, 0.5-0.18*lean, y-0.02, 0.04, f.trunk)
				s.bough(0.5, y-0.34, 0.5+0.18*lean, y-0.12, 0.04, dim(f.trunk, shade))
				if i == storeys-1 {
					for _, x := range []float32{0.22, 0.32, 0.44, 0.58, 0.7, 0.8} {
						s.bough(x, y-0.1, x, y-0.4, 0.03, dim(f.crown, shade))
					}
				}
				clumps = append(clumps,
					[3]float32{0.5 + 0.18*lean, y - 0.12, 0.2},
					[3]float32{0.5 - 0.18*lean, y - 0.02, 0.22},
					[3]float32{0.5, y + 0.1, 0.19},
				)
			}
			s.crown(top, clumps...)
		},
	}
}

// umbrella is a tree of a thick bare trunk forking twice under a flat crown of tufts
func umbrella(height, crown float32, f foliage) forest.Spec {
	return forest.Spec{
		Height: scale.Figure(height),
		Crown:  scale.Figure(crown),
		Sway:   0.15,
		Top: func(dst *render.Canvas, size int) {
			s := float32(size)
			dst.FillCircle(s/2, s/2, s/2, f.crown)
			for i := range 10 {
				sin, cos := math.Sincos(float64(i) * math.Pi / 5)
				dst.FillCircle(s/2+float32(cos)*s*0.3, s/2+float32(sin)*s*0.3, s*0.07, f.leaves)
			}
			dst.FillCircle(s/2, s/2, s*0.08, f.leaves)
		},
		Side: func(dst *render.Canvas, size int) {
			s := sideOf(dst, size, height, crown, f)
			top := s.aspect
			fork, limbs := top*0.3, top*0.52
			s.trunk(0.5, fork+0.04, 0.22, 0.16)
			dark := dim(f.trunk, shade)
			s.bough(0.5, fork, 0.27, limbs, 0.11, f.trunk)
			s.bough(0.5, fork, 0.73, limbs, 0.11, dark)
			s.bough(0.5, fork, 0.5, top-0.3, 0.09, f.trunk)
			s.bough(0.27, limbs, 0.13, top-0.3, 0.07, f.trunk)
			s.bough(0.27, limbs, 0.33, top-0.3, 0.07, f.trunk)
			s.bough(0.73, limbs, 0.67, top-0.3, 0.07, dark)
			s.bough(0.73, limbs, 0.87, top-0.3, 0.07, dark)
			// its tufts in a dome, each a rosette of blades standing out of it
			for _, c := range [][2]float32{
				{0.3, 0.3},
				{0.7, 0.3},
				{0.5, 0.29},
				{0.87, 0.3},
				{0.13, 0.3},
				{0.7, 0.2},
				{0.3, 0.2},
				{0.5, 0.16},
			} {
				x, y := c[0], top-c[1]
				s.clump(x, y, 0.14)
				out, up := (x-0.5)*0.34, 0.11-(x-0.5)*(x-0.5)*0.2
				for _, lean := range []float32{-0.07, 0, 0.07} {
					s.bough(x, y+0.04, x+out+lean, y+up+0.04, 0.022, f.leaves)
				}
			}
			s.bear(0.03,
				[2]float32{0.2, top - 0.4},
				[2]float32{0.4, top - 0.41},
				[2]float32{0.6, top - 0.4},
				[2]float32{0.8, top - 0.41},
			)
		},
	}
}

// fronded is a plant of long leaves arching out of the top of its stem, the stem the share stem
// of its height, each leaf broad crown widths broad
func fronded(height, crown, stem, broad float32, f foliage) forest.Spec {
	return forest.Spec{
		Height: scale.Figure(height),
		Crown:  scale.Figure(crown),
		Sway:   0.65,
		Top: func(dst *render.Canvas, size int) {
			s := float32(size)
			for i := range 9 {
				sin, cos := math.Sincos(float64(i) * 2 * math.Pi / 9)
				col := f.crown
				if i%2 == 0 {
					col = f.leaves
				}
				dst.StrokeLine(s/2, s/2, s/2+float32(cos)*s*0.46, s/2+float32(sin)*s*0.46, s/7,
					col)
			}
			dst.FillCircle(s/2, s/2, s*0.1, f.trunk)
		},
		Side: func(dst *render.Canvas, size int) {
			s := sideOf(dst, size, height, crown, f)
			top := s.aspect
			from := top * stem
			s.trunk(0.5, from, 0.12, 0.09)
			// the scars its fallen leaves left round its stem
			for y := float32(0.1); y < from-0.06; y += 0.11 {
				s.bough(0.455, y, 0.545, y, 0.02, dim(f.trunk, shade))
			}
			s.bear(0.045,
				[2]float32{0.42, from - 0.05},
				[2]float32{0.58, from - 0.05},
				[2]float32{0.5, from - 0.11},
			)
			// its leaves arch out of the top of the stem and droop at their ends, the hindmost
			// in the shade
			reach := min(top-from, 0.5)
			leaves := 9
			for i := range leaves {
				at := math.Pi * (float64(i)/float64(leaves-1)*1.24 - 0.12)
				sin, cos := math.Sincos(at)
				col := f.crown
				switch {
				case i%2 == 0:
					col = f.leaves
				case i%4 == 1:
					col = dim(f.crown, shade)
				}
				bx, by := 0.5+float32(cos)*0.28, from+float32(sin)*reach*0.72+0.04
				ex, ey := 0.5+float32(cos)*0.48, from+float32(sin)*reach-0.04
				s.bough(0.5, from, bx, by, broad, col)
				s.bough(bx, by, ex, ey, broad*0.7, col)
				if broad > 0.1 { // a rib down the middle of a broad leaf
					s.bough(0.5, from, bx, by, 0.025, dim(col, 1.18))
					s.bough(bx, by, ex, ey, 0.025, dim(col, 1.18))
				}
			}
			s.disc(0.5, from, 0.07, dim(f.crown, shade))
		},
	}
}
