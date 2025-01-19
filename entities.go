package raychip

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

type Entity interface{
    Update()
    Draw()
}

type Circle struct {
	position       Vector2
	angle          float64
	velocity       Vector2
	velocityMax    float64
	radius         float64
	color          rl.Color
	updateCallback func(*Circle)
	drawCallback   func(*Circle)
	id             uint64
	physical       bool
	mass           float64
	elasticity     float64
	friction       float64
	cpBody         *cp.Body
	cpShape        *cp.Shape
}

func NewPhysicalCircle(position Vector2, radius float64, mass float64, color rl.Color) Circle {
	pOut := Circle{
		position:    position,
		radius:      radius,
		mass:        mass,
		color:       color,
		physical:    true,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
	pOut.SetDrawCallback(DefaultCircleDrawFunc)
	return pOut
}

func NewCircle(position Vector2, radius float64, color rl.Color) Circle {
	pOut := Circle{
		position:    position,
		radius:      radius,
		color:       color,
		physical:    false,
		velocityMax: 800.0,
	}
	pOut.SetDrawCallback(DefaultCircleDrawFunc)
	return pOut
}

func (c *Circle) Update() {
    if c.updateCallback != nil {
        c.updateCallback(c)
    }
}

func (c *Circle) Draw() {
    if c.drawCallback != nil {
        c.drawCallback(c)
    }
}

func (p Circle) IsPhysical() bool {
	return p.physical
}

func DefaultCircleDrawFunc(p *Circle) {
	pos := p.Position()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(p.radius), p.color)
}

func (p *Circle) SetUpdateCallback(callback func(*Circle)) {
	p.updateCallback = callback
}

func (p *Circle) SetDrawCallback(callback func(*Circle)) {
	p.drawCallback = callback
}

func (p *Circle) Radius() float64 {
	return p.radius
}

func (p *Circle) SetMass(m float64) {
	p.mass = m
	if p.cpBody != nil {
		p.cpBody.SetMass(m)
	}
}

func (p *Circle) Mass() float64 {
	if p.cpBody != nil {
		p.mass = p.cpBody.Mass()
	}
	return p.mass
}

func (p *Circle) SetElasticity(e float64) {
	p.elasticity = e
	if p.cpShape != nil {
		p.cpShape.SetElasticity(e)
	}
}

func (p *Circle) Elasticity() float64 {
	if p.cpShape != nil {
		p.elasticity = p.cpShape.Elasticity()
	}
	return p.elasticity
}

func (p *Circle) SetFriction(f float64) {
	p.friction = f
	if p.cpShape != nil {
		p.cpShape.SetFriction(f)
	}
}

func (p *Circle) Friction() float64 {
	if p.cpShape != nil {
		p.friction = p.cpShape.Friction()
	}
	return p.friction
}

func (p *Circle) Angle() float64 {
	if p.cpBody != nil {
		p.angle = p.cpBody.Angle()
	}
	return p.angle
}

func (p *Circle) Fix() {
	if p.cpBody != nil {
		p.SetVelocity(0, 0)
		p.cpBody.SetMass(math.Inf(1))
	}
}

func (p *Circle) Unfix() {
	if p.cpBody != nil {
		p.cpBody.SetMass(p.mass)
	}
}

func (p *Circle) Position() Vector2 {
	if p.cpBody != nil {
		p.position = Vector2(p.cpBody.Position())
	}
	return p.position
}

func (p *Circle) Velocity() Vector2 {
	if p.cpBody != nil {
		p.velocity = Vector2(p.cpBody.Velocity())
	}
	return p.velocity
}

func (p *Circle) SetVelocityMax(v float64) {
	p.velocityMax = v
}

func (p *Circle) SetVelocity(x float64, y float64) {
	p.velocity.X = x
	p.velocity.Y = y
	if p.cpBody != nil {
		p.cpBody.SetVelocity(x, y)
	}
}

func (p *Circle) SetPosition(x float64, y float64) {
	p.position.X = x
	p.position.Y = y
	if p.cpBody != nil {
		p.cpBody.SetPosition(p.Position().ToChipmunk())
	}
}

type Box struct {
	position       Vector2
	angle          float64
	velocity       Vector2
	velocityMax    float64
	width          float64
	height         float64
	color          rl.Color
	updateCallback func(*Box)
	drawCallback   func(*Box)
	id             uint64
	physical       bool
	mass           float64
	elasticity     float64
	friction       float64
	cpBody         *cp.Body
	cpShape        *cp.Shape
}

func NewPhysicalBox(position Vector2, width float64, height float64, mass float64, color rl.Color) Box {
    bOut := Box{
		position:    position,
		width:       width,
		height:      height,
		mass:        mass,
		color:       color,
		physical:    true,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
    bOut.SetDrawCallback(DefaultBoxDrawFunc)
    return bOut
}

func NewBox(position Vector2, width float64, height float64, color rl.Color) Box {
    bOut := Box{
		position:    position,
		width:       width,
		height:      height,
		color:       color,
		physical:    false,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
    bOut.SetDrawCallback(DefaultBoxDrawFunc)
    return bOut
}

func DefaultBoxDrawFunc(b *Box) {
    angle := b.Angle() * 180.0 / math.Pi
    pos := b.Position()
    boxRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), float32(b.width), float32(b.height))
    rl.DrawRectanglePro(boxRect, rl.NewVector2(boxRect.Width/2, boxRect.Height/2), float32(angle), b.color)
}

func (b *Box) SetUpdateCallback(callback func(*Box)) {
	b.updateCallback = callback
}

func (b *Box) SetDrawCallback(callback func(*Box)) {
	b.drawCallback = callback
}

func (b *Box) Update() {
    if b.updateCallback != nil {
        b.updateCallback(b)
    }
}

func (b *Box) Draw() {
    if b.drawCallback != nil {
        b.drawCallback(b)
    }
}

// func (b *Box) SetOnClick(callback func(*Box)) {
// }

func (b *Box) SetVelocityMax(v float64) {
	b.velocityMax = v
}

func (b *Box) SetMass(m float64) {
	b.mass = m
	if b.cpBody != nil {
		b.cpBody.SetMass(m)
	}
}

func (b *Box) Mass() float64 {
	if b.cpBody != nil {
		b.mass = b.cpBody.Mass()
	}
	return b.mass
}

func (b *Box) SetElasticity(e float64) {
	b.elasticity = e
	if b.cpShape != nil {
		b.cpShape.SetElasticity(e)
	}
}

func (b *Box) Elasticity() float64 {
	if b.cpShape != nil {
		b.elasticity = b.cpShape.Elasticity()
	}
	return b.elasticity
}

func (b *Box) SetFriction(f float64) {
	b.friction = f
	if b.cpShape != nil {
		b.cpShape.SetFriction(f)
	}
}

func (b *Box) Friction() float64 {
	if b.cpShape != nil {
		b.friction = b.cpShape.Friction()
	}
	return b.friction
}

func (b *Box) Velocity() Vector2 {
	if b.cpBody != nil {
		b.velocity = Vector2(b.cpBody.Velocity())
	}
	return b.velocity
}

func (b *Box) SetVelocity(x float64, y float64) {
	b.velocity.X = x
	b.velocity.Y = y
	if b.cpBody != nil {
		b.cpBody.SetVelocity(x, y)
	}
}

func (b *Box) Angle() float64 {
	if b.cpBody != nil {
		b.angle = b.cpBody.Angle()
	}
	return b.angle
}

func (b *Box) Position() Vector2 {
	if b.cpBody != nil {
		b.position = Vector2(b.cpBody.Position())
	}
	return b.position
}

func (b *Box) SetPosition(x float64, y float64) {
	b.position.X = x
	b.position.Y = y
	if b.cpBody != nil {
		b.cpBody.SetPosition(cp.Vector{X: x, Y: y})
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

func (w *Wall) Update() {}
func (w *Wall) Draw() {
    if w.Visible {
        rl.DrawLineEx(w.Vertex1.ToRaylib(), w.Vertex2.ToRaylib(), float32(w.Width), w.Color)
    }
}
