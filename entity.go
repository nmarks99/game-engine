package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
    "math"
)

type Entity interface {
	Update()
	Draw()
	Id() uint64
	addToGame(game *Game, args ...any)
}

type EntityBase struct {
	position       Vector2
	angle          float64
	color          rl.Color
	// updateCallback func(*Entity)
	// drawCallback   func(*Entity)
	id             uint64
	physical       bool
	velocity       Vector2
	velocityMax    float64
	mass           float64
	elasticity     float64
	friction       float64
	cpBody         *cp.Body
	cpShape        *cp.Shape
}

func (e EntityBase) Id() uint64{
    return e.id
}

func (e *EntityBase) SetColor(color rl.Color) {
    e.color = color
}

func (e EntityBase) Color() rl.Color {
	return e.color
}

func (e EntityBase) IsPhysical() bool {
	return e.physical
}

func (e *EntityBase) SetMass(m float64) {
	e.mass = m
	if e.cpBody != nil {
		e.cpBody.SetMass(m)
	}
}

func (e *EntityBase) Mass() float64 {
	if e.cpBody != nil {
		e.mass = e.cpBody.Mass()
	}
	return e.mass
}

func (e *EntityBase) SetElasticity(elasticity float64) {
	e.elasticity = elasticity
	if e.cpShape != nil {
		e.cpShape.SetElasticity(elasticity)
	}
}

func (e *EntityBase) Elasticity() float64 {
	if e.cpShape != nil {
		e.elasticity = e.cpShape.Elasticity()
	}
	return e.elasticity
}

func (e *EntityBase) SetFriction(f float64) {
	e.friction = f
	if e.cpShape != nil {
		e.cpShape.SetFriction(f)
	}
}

func (e *EntityBase) Friction() float64 {
	if e.cpShape != nil {
		e.friction = e.cpShape.Friction()
	}
	return e.friction
}

func (e *EntityBase) SetAngle(a float64) {
	e.angle = a
	if e.cpBody != nil {
		e.cpBody.SetAngle(a)
	}
}

func (e *EntityBase) Angle() float64 {
	if e.cpBody != nil {
		e.angle = e.cpBody.Angle()
	}
	return e.angle
}

func (e *EntityBase) Fix() {
	if e.cpBody != nil {
		e.SetVelocity(0, 0)
		e.cpBody.SetMass(math.Inf(1))
	}
}

func (e *EntityBase) Unfix() {
	if e.cpBody != nil {
		e.cpBody.SetMass(e.mass)
	}
}

func (e *EntityBase) SetPosition(x float64, y float64) {
	e.position.X = x
	e.position.Y = y
	if e.cpBody != nil {
		e.cpBody.SetPosition(e.Position().ToChipmunk())
	}
}

func (e *EntityBase) Position() Vector2 {
	if e.cpBody != nil {
		e.position = Vector2(e.cpBody.Position())
	}
	return e.position
}

func (e *EntityBase) SetVelocity(x float64, y float64) {
	e.velocity.X = x
	e.velocity.Y = y
	if e.cpBody != nil {
		e.cpBody.SetVelocity(x, y)
	}
}

func (e *EntityBase) Velocity() Vector2 {
	if e.cpBody != nil {
		e.velocity = Vector2(e.cpBody.Velocity())
	}
	return e.velocity
}

func (e *EntityBase) SetVelocityMax(v float64) {
	e.velocityMax = v
}

func (e *EntityBase) VelocityMax() float64 {
	return e.velocityMax
}
