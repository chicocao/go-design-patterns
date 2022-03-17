package bridge

import "testing"

func TestBridge(t *testing.T) {
	sR := NewShapeStraightImpl(&RedLine{})
	sR.Draw(1, 2)

	sG := NewShapeStraightImpl(&GreenLine{})
	sG.Draw(3, 4)

	sCR := NewShapeCurveImpl(&RedLine{})
	sCR.Draw(5, 6)

	sCG := NewShapeCurveImpl(&GreenLine{})
	sCG.Draw(7, 8)
}
