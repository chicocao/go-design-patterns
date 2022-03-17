package bridge

import "fmt"

type DrawInterface interface {
	DrawLine(int, int)
}
type RedLine struct {
}

func (r *RedLine) DrawLine(start, end int) {
	fmt.Printf("red line start: %d end: %d\n", start, end)
}

type GreenLine struct {
}

func (g *GreenLine) DrawLine(start, end int) {
	fmt.Printf("green line start: %d end: %d\n", start, end)
}

type ShapeInterface interface {
	Draw(int, int)
}

type ShapeStraightImpl struct {
	d DrawInterface
}

func (s *ShapeStraightImpl) Draw(start, end int) {
	fmt.Printf("straight draw: ")
	s.d.DrawLine(start, end)
}

func NewShapeStraightImpl(d DrawInterface) *ShapeStraightImpl {
	return &ShapeStraightImpl{d: d}
}

type ShapeCurveImpl struct {
	d DrawInterface
}

func (s *ShapeCurveImpl) Draw(start, end int) {
	fmt.Printf("curve draw: ")
	s.d.DrawLine(start, end)
}

func NewShapeCurveImpl(d DrawInterface) *ShapeCurveImpl {
	return &ShapeCurveImpl{d: d}
}
