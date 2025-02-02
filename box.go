package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
)

type Box struct {
    EntityBase
	width          float64
	height         float64
	updateCallback func(*Box)
	drawCallback   func(*Box)
}

func NewBox(x float64, y float64, width float64, height float64, color rl.Color) Box {
	bOut := Box{
        EntityBase: EntityBase{
            position:    NewVector2(x, y),
            color:       color,
            physical:    false,
        },
        width:       width,
        height:      height,
	}
	bOut.SetDrawCallback(defaultBoxDrawFunc)
	return bOut
}


func NewPhysicalBox(x float64, y float64, width float64, height float64, mass float64, color rl.Color) Box {
	bOut := Box{
        EntityBase: EntityBase{
            position:    NewVector2(x, y),
            mass:        mass,
            color:       color,
            physical:    true,
            elasticity:  1.0,
            friction:    1.0,
            velocityMax: 800.0,
        },
        width:       width,
        height:      height,
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

func (b *Box) OnClick(game *Game, button rl.MouseButton, state MouseState, callback func()) {
	var oldUpdateCallback func(*Box)
	if b.updateCallback != nil {
		oldUpdateCallback = b.updateCallback
	}

	b.updateCallback = func(b *Box) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(b)
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
			boxRect := rl.NewRectangle(float32(b.position.X-b.width/2.0), float32(b.position.Y-b.height/2.0), float32(b.width), float32(b.height))
			if rl.CheckCollisionPointRec(game.mousePosition.ToRaylib(), boxRect) {
				callback()
			}
		}

	}
}

func (b *Box) OnHover(game *Game, callbackOn func(), callbackOff func()) {
	var oldUpdateCallback func(*Box)
	if b.updateCallback != nil {
		oldUpdateCallback = b.updateCallback
	}

	b.updateCallback = func(b *Box) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(b)
		}
		buttonRect := rl.NewRectangle(float32(b.position.X-b.width/2.0), float32(b.position.Y-b.height/2.0), float32(b.width), float32(b.height))
		if rl.CheckCollisionPointRec(game.mousePosition.ToRaylib(), buttonRect) {
			callbackOn()
		} else {
			callbackOff()
        }

	}
}
