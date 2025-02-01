package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
    // "fmt"
)

type Entity interface {
	Update()
	Draw()
	Id() uint64
	addToGame(game *Game, args ...any)
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

func NewPhysicalCircle(x float64, y float64, radius float64, mass float64, color rl.Color) Circle {
	pOut := Circle{
		position:    NewVector2(x, y),
		radius:      radius,
		mass:        mass,
		color:       color,
		physical:    true,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
	pOut.SetDrawCallback(defaultCircleDrawFunc)
	return pOut
}

func NewCircle(x float64, y float64, radius float64, color rl.Color) Circle {
	pOut := Circle{
		position:    NewVector2(x, y),
		radius:      radius,
		color:       color,
		physical:    false,
		velocityMax: 800.0,
	}
	pOut.SetDrawCallback(defaultCircleDrawFunc)
	return pOut
}

func (p Circle) limitVelocity(body *cp.Body, gravity cp.Vector, damping float64, dt float64) {
	maxSpeed := p.velocityMax // Maximum speed (pixels/second)
	cp.BodyUpdateVelocity(body, gravity, damping, dt)
	velocity := body.Velocity()
	speed := math.Sqrt(velocity.X*velocity.X + velocity.Y*velocity.Y)
	if speed > maxSpeed {
		scale := maxSpeed / speed
		body.SetVelocity(velocity.X*scale, velocity.Y*scale)
	}
}

func (e *Circle) addToGame(game *Game, args ...any) {
	if e.physical {
		if body, ok := args[0].(*cp.Body); ok {
			if shape, ok := args[1].(*cp.Shape); ok {
				game.physical = true
				body = game.space.AddBody(cp.NewBody(e.mass, cp.MomentForCircle(e.mass, 0.0, e.radius, cp.Vector{})))
				body.SetType(cp.BODY_DYNAMIC)
				body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
				body.SetVelocity(e.velocity.X, e.velocity.Y)
				shape = game.space.AddShape(cp.NewCircle(body, e.radius, cp.Vector{}))
				shape.SetElasticity(e.elasticity)
				shape.SetFriction(e.friction)
				body.SetVelocityUpdateFunc(e.limitVelocity)
				e.cpBody = body
				e.cpShape = shape
			}
		}
	}
	e.id = uint64(len(game.entities))
	game.entities = append(game.entities, e)
}

func (c *Circle) Id() uint64 {
	return c.id
}

func defaultCircleDrawFunc(p *Circle) {
	pos := p.Position()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), float32(p.radius), p.color)
}

func (c Circle) DefaultDraw() {
	defaultCircleDrawFunc(&c)
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

func (p *Circle) SetDrawCallback(callback func(*Circle)) {
	p.drawCallback = callback
}

func (c *Circle) SetUpdateCallback(callback func(*Circle)) {
	var oldUpdateCallback func(*Circle)
	if c.updateCallback != nil {
		oldUpdateCallback = c.updateCallback
	}

	c.updateCallback = func(c *Circle) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(c)
		}
		callback(c)
	}
}

func (c *Circle) OnClick(button rl.MouseButton, state MouseState, callback func()) {
	var oldUpdateCallback func(*Circle)
	if c.updateCallback != nil {
		oldUpdateCallback = c.updateCallback
	}

	c.updateCallback = func(c *Circle) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(c)
		}

		mousePos := rl.GetMousePosition()

		var clicked bool = false
		switch state {
		case MouseUp:
			if rl.IsMouseButtonUp(button) {
				clicked = true
			}
		case MouseDown:
			if rl.IsMouseButtonDown(button) {
				clicked = true
			}
		case MousePressed:
			if rl.IsMouseButtonPressed(button) {
				clicked = true
			}
		case MouseReleased:
			if rl.IsMouseButtonReleased(button) {
				clicked = true
			}
		}

		if clicked {
			if rl.CheckCollisionPointCircle(mousePos, c.position.ToRaylib(), float32(c.radius)) {
				callback()
			}
		}

	}

}

func (c *Circle) SetTexture(texture rl.Texture2D) {
	// not sure if we want to do this?
	// c.SetDrawCallback(func (c *Circle){
	// if c.drawCallback != nil {
	// c.drawCallback(c)
	// }
	// })
	c.SetDrawCallback(func(c *Circle) {
		pos := c.Position()
		textureWidth := float32(texture.Width)
		textureHeight := float32(texture.Height)
		srcRect := rl.NewRectangle(0, 0, textureWidth, textureHeight)
		destRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), textureWidth, textureHeight)
		origin := rl.NewVector2(textureWidth/2, textureHeight/2)
		angle := float32(c.Angle() * 180.0 / math.Pi)
		rl.DrawTexturePro(texture, srcRect, destRect, origin, float32(angle), rl.White)
	})

}

func (c *Circle) SetColor(color rl.Color) {
	c.color = color
}

func (p Circle) IsPhysical() bool {
	return p.physical
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

func (p *Circle) SetAngle(a float64) {
	p.angle = a
	if p.cpBody != nil {
		p.cpBody.SetAngle(a)
	}
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

func (p *Circle) SetPosition(x float64, y float64) {
	p.position.X = x
	p.position.Y = y
	if p.cpBody != nil {
		p.cpBody.SetPosition(p.Position().ToChipmunk())
	}
}

func (p *Circle) Position() Vector2 {
	if p.cpBody != nil {
		p.position = Vector2(p.cpBody.Position())
	}
	return p.position
}

func (p *Circle) SetVelocity(x float64, y float64) {
	p.velocity.X = x
	p.velocity.Y = y
	if p.cpBody != nil {
		p.cpBody.SetVelocity(x, y)
	}
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

func (p *Circle) VelocityMax() float64 {
	return p.velocityMax
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

func NewPhysicalBox(x float64, y float64, width float64, height float64, mass float64, color rl.Color) Box {
	bOut := Box{
		position:    NewVector2(x, y),
		width:       width,
		height:      height,
		mass:        mass,
		color:       color,
		physical:    true,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
	bOut.SetDrawCallback(defaultBoxDrawFunc)
	return bOut
}

func NewBox(x float64, y float64, width float64, height float64, color rl.Color) Box {
	bOut := Box{
		position:    NewVector2(x, y),
		width:       width,
		height:      height,
		color:       color,
		physical:    false,
		elasticity:  1.0,
		friction:    1.0,
		velocityMax: 800.0,
	}
	bOut.SetDrawCallback(defaultBoxDrawFunc)
	return bOut
}

func (b Box) limitVelocity(body *cp.Body, gravity cp.Vector, damping float64, dt float64) {
	maxSpeed := b.velocityMax // Maximum speed (pixels/second)
	cp.BodyUpdateVelocity(body, gravity, damping, dt)
	velocity := body.Velocity()
	speed := math.Sqrt(velocity.X*velocity.X + velocity.Y*velocity.Y)
	if speed > maxSpeed {
		scale := maxSpeed / speed
		body.SetVelocity(velocity.X*scale, velocity.Y*scale)
	}
}

func (e *Box) addToGame(game *Game, args ...any) {
	if e.physical {
		if body, ok := args[0].(*cp.Body); ok {
			if shape, ok := args[1].(*cp.Shape); ok {
				game.physical = true
				body = game.space.AddBody(cp.NewBody(e.mass, cp.MomentForBox(e.mass, e.width, e.height)))
				body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
				body.SetVelocity(e.velocity.X, e.velocity.Y)
				shape = game.space.AddShape(cp.NewBox(body, e.width, e.height, 0))
				shape.SetElasticity(e.elasticity)
				shape.SetFriction(e.friction)
				body.SetVelocityUpdateFunc(e.limitVelocity)
				e.cpBody = body
				e.cpShape = shape
			}
		}
	}
	e.id = uint64(len(game.entities))
	game.entities = append(game.entities, e)
}

func (b *Box) Id() uint64 {
	return b.id
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

func defaultBoxDrawFunc(b *Box) {
	angle := b.Angle() * 180.0 / math.Pi
	pos := b.Position()
	boxRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), float32(b.width), float32(b.height))
	rl.DrawRectanglePro(boxRect, rl.NewVector2(boxRect.Width/2, boxRect.Height/2), float32(angle), b.color)
}

func (b Box) DefaultDraw() {
	defaultBoxDrawFunc(&b)
}

func (b *Box) SetDrawCallback(callback func(*Box)) {
	b.drawCallback = callback
}

func (b *Box) SetUpdateCallback(callback func(*Box)) {
	b.updateCallback = callback
}

func (b *Box) OnClick(button rl.MouseButton, state MouseState, callback func()) {
	var oldUpdateCallback func(*Box)
	if b.updateCallback != nil {
		oldUpdateCallback = b.updateCallback
	}

	b.updateCallback = func(b *Box) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(b)
		}

		mousePos := rl.GetMousePosition()

		var clicked bool = false
		switch state {
		case MouseUp:
			if rl.IsMouseButtonUp(button) {
				clicked = true
			}
		case MouseDown:
			if rl.IsMouseButtonDown(button) {
				clicked = true
			}
		case MousePressed:
			if rl.IsMouseButtonPressed(button) {
				clicked = true
			}
		case MouseReleased:
			if rl.IsMouseButtonReleased(button) {
				clicked = true
			}
		}

		if clicked {
			boxRect := rl.NewRectangle(float32(b.position.X-b.width/2.0), float32(b.position.Y-b.height/2.0), float32(b.width), float32(b.height))
			if rl.CheckCollisionPointRec(mousePos, boxRect) {
				callback()
			}
		}

	}
}

// func (b *Box) OnHoverAlpha(alphaBase uint8, alphaHover uint8) {
	// var oldUpdateCallback func(*Box)
	// if b.updateCallback != nil {
		// oldUpdateCallback = b.updateCallback
	// }
//
	// b.updateCallback = func(b *Box) {
		// if oldUpdateCallback != nil {
			// oldUpdateCallback(b)
		// }
		// mousePos := rl.GetMousePosition()
        // alphaBase = b.color.A
		// buttonRect := rl.NewRectangle(float32(b.position.X-b.width/2.0), float32(b.position.Y-b.height/2.0), float32(b.width), float32(b.height))
		// if rl.CheckCollisionPointRec(mousePos, buttonRect) {
            // b.SetColor(rl.ColorAlpha(b.Color(), float32(alphaHover)/math.MaxUint8))
		// } else {
            // b.SetColor(rl.ColorAlpha(b.Color(), float32(alphaBase)/math.MaxUint8))
        // }
//
	// }
// }

func (b *Box) SetColor(color rl.Color) {
	b.color = color
}

func (b Box) Color() rl.Color {
	return b.color
}

func (b Box) IsPhysical() bool {
	return b.physical
}

func (b *Box) SetVelocityMax(v float64) {
	b.velocityMax = v
}

func (b Box) VelocityMax() float64 {
	return b.velocityMax
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

func (b *Box) SetVelocity(x float64, y float64) {
	b.velocity.X = x
	b.velocity.Y = y
	if b.cpBody != nil {
		b.cpBody.SetVelocity(x, y)
	}
}

func (b *Box) Velocity() Vector2 {
	if b.cpBody != nil {
		b.velocity = Vector2(b.cpBody.Velocity())
	}
	return b.velocity
}

func (b *Box) SetAngle(a float64) {
	b.angle = a
	if b.cpBody != nil {
		b.cpBody.SetAngle(a)
	}
}

func (b *Box) Angle() float64 {
	if b.cpBody != nil {
		b.angle = b.cpBody.Angle()
	}
	return b.angle
}

func (b *Box) SetPosition(x float64, y float64) {
	b.position.X = x
	b.position.Y = y
	if b.cpBody != nil {
		b.cpBody.SetPosition(cp.Vector{X: x, Y: y})
	}
}

func (b *Box) Position() Vector2 {
	if b.cpBody != nil {
		b.position = Vector2(b.cpBody.Position())
	}
	return b.position
}

type Wall struct {
	Vertex1 Vector2
	Vertex2 Vector2
	Width   float64
	Color   rl.Color
	id      uint64
	cpBody  *cp.Body
}

func NewWall(vertex1 Vector2, vertex2 Vector2, width float64, color rl.Color) Wall {
	return Wall{
		Vertex1: vertex1,
		Vertex2: vertex2,
		Width:   width,
		Color:   color,
	}
}

func (e *Wall) addToGame(game *Game, args ...any) {
	if body, ok := args[0].(*cp.Body); ok {
		if shape, ok := args[1].(*cp.Shape); ok {
			body = cp.NewStaticBody()
			shape = game.space.AddShape(cp.NewSegment(body, cp.Vector{X: e.Vertex1.X, Y: e.Vertex1.Y}, cp.Vector{X: e.Vertex2.X, Y: e.Vertex2.Y}, e.Width/2))
			shape.SetElasticity(1)
			shape.SetFriction(1)
			e.id = uint64(len(game.entities))
		}
	}
	e.id = uint64(len(game.entities))
	game.entities = append(game.entities, e)
}

func (w *Wall) Id() uint64 {
	return w.id
}

func (w *Wall) Update() {}
func (w *Wall) Draw() {
	rl.DrawLineEx(w.Vertex1.ToRaylib(), w.Vertex2.ToRaylib(), float32(w.Width), w.Color)
}
