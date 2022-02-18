package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type intCustomType int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const SidesTriangle intCustomType = 3
const SidesSquare intCustomType = 4
const SidesCircle intCustomType = 0

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (math.Sqrt(3) / 4) * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	default:
		return 0
	}
}
