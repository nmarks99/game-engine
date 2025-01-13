package engine

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	// "fmt"
)

type Entity interface{}

type Particle struct {
	position       Vector2
	velocity       Vector2
	velocityMax    float64
	radius         float64
	mass           float64
	elasticity     float64
	friction       float64
	color          rl.Color
	updateCallback func(*Particle)
	drawCallback   func(*Particle)
	id             uint64
	cpBody         *cp.Body
	cpShape        *cp.Shape
}

func (p *Particle) SetUpdateCallback(callback func(*Particle)) {
	p.updateCallback = callback
}

func (p *Particle) SetDrawCallback(callback func(*Particle)) {
	p.drawCallback = callback
}

func (p Particle) Radius() float64 {
	return p.radius
}

func (p *Particle) SetMass(m float64) {
	p.mass = m
	if p.cpBody != nil {
		p.cpBody.SetMass(m)
	}
}

func (p Particle) Mass() float64 {
	if p.cpBody != nil {
		return p.cpBody.Mass()
	} else {
		return p.mass
	}
}

func (p *Particle) SetElasticity(e float64) {
	p.elasticity = e
	if p.cpShape != nil {
		p.cpShape.SetElasticity(e)
	}
}

func (p Particle) Elasticity() float64 {
	if p.cpShape != nil {
		return p.cpShape.Elasticity()
	} else {
		return p.elasticity
	}
}

func (p *Particle) SetFriction(f float64) {
	p.friction = f
	if p.cpShape != nil {
		p.cpShape.SetFriction(f)
	}
}

func (p Particle) Friction() float64 {
	if p.cpShape != nil {
		return p.cpShape.Friction()
	} else {
		return p.friction
	}
}

func (p Particle) Angle() float64 {
	return p.cpBody.Angle()
}

func (p *Particle) Fix() {
	if p.cpBody != nil {
		p.SetVelocity(0, 0)
		p.cpBody.SetMass(math.Inf(1))
	}
}

func (p *Particle) Unfix() {
	if p.cpBody != nil {
		p.cpBody.SetMass(p.mass)
	}
}

func (p Particle) Position() Vector2 {
	if p.cpBody != nil {
		return Vector2{X: p.cpBody.Position().X, Y: p.cpBody.Position().Y}
	} else {
		return Vector2{X: p.Position().X, Y: p.Position().Y}
	}
}

func (p Particle) Velocity() Vector2 {
	if p.cpBody != nil {
		return Vector2{X: p.cpBody.Velocity().X, Y: p.cpBody.Velocity().Y}
	} else {
		return Vector2{X: p.Velocity().X, Y: p.Velocity().Y}
	}
}

func (p *Particle) SetVelocityMax(v float64) {
	p.velocityMax = v
}

func (p *Particle) SetVelocity(x float64, y float64) {
	p.velocity.X = x
	p.velocity.Y = y

	if p.cpBody != nil {
		p.cpBody.SetVelocity(x, y)
	}
}

func (p *Particle) SetPosition(x float64, y float64) {
	p.position.X = x
	p.position.Y = y

	if p.cpBody != nil {
		p.cpBody.SetPosition(p.Position().ToChipmunk())
	}
}

func DefaultParticleDrawFunc(p *Particle) {
	pos := p.Position()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(p.radius), p.color)
}

func NewParticle(position Vector2, radius float64, mass float64, color rl.Color) Particle {
	pOut := Particle{
		position:    position,
		radius:      radius,
		mass:        mass,
		color:       color,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
	pOut.SetDrawCallback(DefaultParticleDrawFunc)
	return pOut
}

type Box struct {
	position       Vector2
	velocity       Vector2
	velocityMax    float64
	Width          float64
	Height         float64
	Mass           float64
	elasticity     float64
	friction       float64
	Color          rl.Color
	updateCallback func(*Particle)
	drawCallback   func(*Particle)
	id             uint64
	cpBody         *cp.Body
	cpShape        *cp.Shape
}

func (b *Box) SetVelocityMax(v float64) {
	b.velocityMax = v
}

func (b *Box) SetElasticity(e float64) {
	b.elasticity = e
	if b.cpShape != nil {
		b.cpShape.SetElasticity(e)
	}
}

func (b Box) Elasticity() float64 {
	if b.cpShape != nil {
		return b.cpShape.Elasticity()
	} else {
		return b.elasticity
	}
}

func (b *Box) SetFriction(f float64) {
	b.friction = f
	if b.cpShape != nil {
		b.cpShape.SetFriction(f)
	}
}

func (b Box) Friction() float64 {
	if b.cpShape != nil {
		return b.cpShape.Friction()
	} else {
		return b.friction
	}
}

func (b Box) Velocity() Vector2 {
	if b.cpBody != nil {
		return NewVector2(b.cpBody.Velocity().X, b.cpBody.Velocity().Y)
	} else {
		return b.velocity
	}
}

func (b *Box) SetVelocity(x float64, y float64) {
	b.velocity.X = x
	b.velocity.Y = y

	if b.cpBody != nil {
		b.cpBody.SetVelocity(x, y)
	}
}

func (b Box) Position() Vector2 {
	if b.cpBody != nil {
		return NewVector2(b.cpBody.Position().X, b.cpBody.Position().Y)
	} else {
		return b.position
	}
}

func (b *Box) SetPosition(x float64, y float64) {
	b.position.X = x
	b.position.Y = y

	if b.cpBody != nil {
		b.cpBody.SetPosition(cp.Vector{X: x, Y: y})
	}
}

func (b *Box) SetKinematic() {
	if b.cpBody != nil {
		b.cpBody.SetType(cp.BODY_KINEMATIC)
	}
}

func NewBox(position Vector2, width float64, height float64, mass float64, color rl.Color) Box {
	return Box{
		position:    position,
		Width:       width,
		Height:      height,
		Mass:        mass,
		Color:       color,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
}

type Wall struct {
	Vertex1 Vector2
	Vertex2 Vector2
	Width   float64
	Color   rl.Color
	Visible bool
	id      uint64
	cpBody  *cp.Body
}

func NewWall(vertex1 Vector2, vertex2 Vector2, width float64, color rl.Color) Wall {
	return Wall{
		Vertex1: vertex1,
		Vertex2: vertex2,
		Width:   width,
		Color:   color,
		Visible: true,
	}
}
