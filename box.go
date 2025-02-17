package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
)

type Box struct {
	EntityBase
	rectangle      rl.Rectangle
	updateCallback func(*Box)
	drawCallback   func(*Box)
}

func NewBox(x float64, y float64, width float64, height float64, color rl.Color) Box {
	bOut := Box{
		EntityBase: EntityBase{
			position: NewVector2(x, y),
			color:    color,
			physical: false,
		},
        // rectangle origin at top left vertex
		rectangle: rl.NewRectangle(float32(x), float32(y), float32(width), float32(height)),
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
		rectangle: rl.NewRectangle(float32(x), float32(y), float32(width), float32(height)),
	}
	bOut.SetDrawCallback(defaultBoxDrawFunc)
	return bOut
}

func (b *Box) SetWidth(width float64) {
	b.rectangle.Width = float32(width)
}

func (b Box) Width() float64 {
	return float64(b.rectangle.Width)
}

func (b *Box) SetHeight(height float64) {
	b.rectangle.Height = float32(height)
}

func (b Box) Height() float64 {
	return float64(b.rectangle.Height)
}

func (b Box) Rectangle() rl.Rectangle {
    return b.rectangle
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
				body = game.space.AddBody(cp.NewBody(e.mass, cp.MomentForBox(e.mass, float64(e.rectangle.Width), float64(e.rectangle.Height))))
				body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
				body.SetVelocity(e.velocity.X, e.velocity.Y)
				shape = game.space.AddShape(cp.NewBox(body, float64(e.rectangle.Width), float64(e.rectangle.Height), 0))
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
	boxRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), float32(b.rectangle.Width), float32(b.rectangle.Height))
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

func (b *Box) OnClick(game *Game, button rl.MouseButton, state MouseState, callback func()) int {
	id := game.EventBus.CreateSubscription("input.mouse", MouseInputEvent{}, func(input MouseInputEvent) {
		var clicked = false
		switch state {
		case MousePressed:
			if input.IsButtonPressed(button) {
				clicked = true
			}
		case MouseReleased:
			if input.IsButtonReleased(button) {
				clicked = true
			}
		case MouseUp:
			if input.IsButtonUp(button) {
				clicked = true
			}
		case MouseDown:
			if input.IsButtonDown(button) {
				clicked = true
			}
		}

		if clicked {
			boxRect := rl.NewRectangle(
                float32(b.position.X-b.Width()/2.0),
				float32(b.position.Y-b.Height()/2.0),
				float32(b.Width()),
				float32(b.Height()),
			)
			if rl.CheckCollisionPointRec(game.mousePosition.ToRaylib(), boxRect) {
				callback()
			}
		}
	})

	return id
}

func (b Box) CheckMouseCollision(mousePos Vector2) bool {
    adjRect := b.rectangle
    adjRect.X -= b.rectangle.Width/2
    adjRect.Y -= b.rectangle.Height/2
    if rl.CheckCollisionPointRec(mousePos.ToRaylib(), adjRect) {
        return true
    } else {
        return false
    }
}
