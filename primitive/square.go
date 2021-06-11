package primitive

import (
	"fmt"
	"github.com/fogleman/gg"
)

type Square struct {
	Worker *Worker
	X, Y   int
	Size   int
}

func NewRandomSquare(worker *Worker) *Square {
	rnd := worker.Rnd
	x := rnd.Intn(worker.W)
	y := rnd.Intn(worker.H)
	size := rnd.Intn(32) + 1
	if x+size > worker.W-1 {
		x = x - size
	}
	if y+size > worker.H-1 {
		y = y - size
	}
	return &Square{worker, x, y, size}
}

func (r *Square) bounds() (x1, y1, size int) {
	x1, y1, size = r.X, r.Y, r.Size

	return
}

func (r *Square) Draw(dc *gg.Context, scale float64) {
	x, y, size := r.bounds()
	dc.DrawRectangle(float64(x), float64(y), float64(size), float64(size))
	dc.Fill()
}

func (r *Square) SVG(attrs string) string {
	x, y, size := r.bounds()
	w := size
	h := size
	return fmt.Sprintf(
		"<rect %s x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" />",
		attrs, x, y, w, h)
}

func (r *Square) Copy() Shape {
	a := *r
	return &a
}

func (r *Square) Mutate() {
	w := r.Worker.W
	h := r.Worker.H
	rnd := r.Worker.Rnd
	switch rnd.Intn(2) {
	case 0:
		r.X = clampInt(r.X+int(rnd.NormFloat64()*16), 0, w-1)
		r.Y = clampInt(r.Y+int(rnd.NormFloat64()*16), 0, h-1)
		if r.X+r.Size > w-1 {
			r.X = r.X - r.Size
		}
		if r.Y+r.Size > h-1 {
			r.Y = r.Y - r.Size
		}
	case 1:
		r.Size = clampInt(r.Size+int(rnd.NormFloat64()*16), 1, w-1)
		if r.X+r.Size > w-1 {
			r.X = r.X - r.Size
		}
		if r.Y+r.Size > h-1 {
			r.Y = r.Y - r.Size
		}

	}

}

func (r *Square) Rasterize() []Scanline {
	X, Y, size := r.bounds()
	lines := r.Worker.Lines[:0]
	for y := Y; y <= Y+size; y++ {
		lines = append(lines, Scanline{y, X, X + size, 0xffff})
	}
	return lines
}
