package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(x float64, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}

func (v Vector2) ToRaylib() rl.Vector2 {
	return rl.NewVector2(float32(v.X), float32(v.Y))
}

func (v Vector2) ToChipmunk() cp.Vector {
	return cp.Vector{X: float64(v.X), Y: float64(v.Y)}
}

func Vector2FromRaylib(v rl.Vector2) Vector2 {
	return NewVector2(float64(v.X), float64(v.Y))
}

func Vector2FromChipmunk(v cp.Vector) Vector2 {
	return NewVector2(v.X, v.Y)
}
