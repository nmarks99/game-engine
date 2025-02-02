package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
	// "fmt"
)

type Circle struct {
	EntityBase
	radius float64
    updateCallback func(*Circle)
    drawCallback   func(*Circle)
}

func NewCircle(x float64, y float64, radius float64, color rl.Color) Circle {
	pOut := Circle{
		EntityBase: EntityBase{
			position: NewVector2(x, y),
			color:    color,
			physical: false,
		},
		radius: radius,
	}
	pOut.SetDrawCallback(defaultCircleDrawFunc)
	return pOut
}

func NewPhysicalCircle(x float64, y float64, radius float64, mass float64, color rl.Color) Circle {
	pOut := Circle{
		EntityBase: EntityBase{
			position:    NewVector2(x, y),
			mass:        mass,
			color:       color,
			physical:    true,
			elasticity:  1.0,
			friction:    1.0,
			velocityMax: 800.0,
		},
        radius: radius,
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

func (c *Circle) OnClick(game *Game, button rl.MouseButton, state MouseState, callback func()) {
	var oldUpdateCallback func(*Circle)
	if c.updateCallback != nil {
		oldUpdateCallback = c.updateCallback
	}

	c.updateCallback = func(c *Circle) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(c)
		}

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
			if rl.CheckCollisionPointCircle(game.mousePosition.ToRaylib(), c.position.ToRaylib(), float32(c.radius)) {
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

func (p *Circle) Radius() float64 {
	return p.radius
}

