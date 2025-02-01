package raychip

import (
	// "math"

	rl "github.com/gen2brain/raylib-go/raylib"
	// "fmt"
)

type Button struct {
	position       Vector2
	width          float64
	height         float64
	color          rl.Color
	label          string
	labelColor     rl.Color
	updateCallback func(*Button)
	drawCallback   func(*Button)
	id             uint64
}

func NewButton(x float64, y float64, width float64, height float64, color rl.Color) Button {
	bOut := Button{
		position:   NewVector2(x, y),
		width:      width,
		height:     height,
		color:      color,
		label:      " ",
		labelColor: rl.White,
	}
	bOut.SetDrawCallback(defaultButtonDrawFunc)
	return bOut
}

func (b *Button) Id() uint64 {
	return b.id
}

func (b *Button) Update() {
	if b.updateCallback != nil {
		b.updateCallback(b)
	}
}

func (b *Button) Draw() {
	if b.drawCallback != nil {
		b.drawCallback(b)
	}
}

func defaultButtonDrawFunc(b *Button) {
    angle := 0.0
    pos := b.Position()
	buttonRect := rl.NewRectangle(float32(pos.X), float32(pos.Y), float32(b.width), float32(b.height))
	rl.DrawRectanglePro(buttonRect, rl.NewVector2(buttonRect.Width/2, buttonRect.Height/2), float32(angle), b.color)
	rl.DrawText(b.label, int32(b.position.X)+1-int32(b.width)/2, int32(b.position.Y)-int32(b.height)/2, int32(b.height*0.7), b.labelColor)
}

func (b Button) DefaultDraw() {
	defaultButtonDrawFunc(&b)
}

func (b *Button) SetDrawCallback(callback func(*Button)) {
	b.drawCallback = callback
}

func (b *Button) SetUpdateCallback(callback func(*Button)) {
	b.updateCallback = callback
}

func (b *Button) SetPosition(x float64, y float64) {
	b.position.X = x
	b.position.Y = y
}

func (b *Button) Position() Vector2 {
	return b.position
}

func (b *Button) SetLabel(text string, color rl.Color) {
	b.label = text
}

func (b *Button) SetColor(color rl.Color) {
	b.color = color
}

func (b *Button) Color() rl.Color {
	return b.color
}

func (e *Button) addToGame(game *Game, args ...any) {
	e.id = uint64(len(game.entities))
	game.entities = append(game.entities, e)
}

func (b *Button) OnClick(game *Game, button rl.MouseButton, state MouseState, callback func()) {
	var oldUpdateCallback func(*Button)
	if b.updateCallback != nil {
		oldUpdateCallback = b.updateCallback
	}

	b.updateCallback = func(b *Button) {
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

func (b *Button) OnHover(game *Game, callbackOn func(), callbackOff func()) {
	var oldUpdateCallback func(*Button)
	if b.updateCallback != nil {
		oldUpdateCallback = b.updateCallback
	}

	b.updateCallback = func(b *Button) {
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

// func (b *Button) OnHoverAlpha(game *Game, alphaBase uint8, alphaHover uint8) {
    // var oldUpdateCallback func(*Button)
    // if b.updateCallback != nil {
        // oldUpdateCallback = b.updateCallback
    // }
//
    // b.updateCallback = func(b *Button) {
        // if oldUpdateCallback != nil {
            // oldUpdateCallback(b)
        // }
        // mousePos := game.mousePosition.ToRaylib()
        // buttonRect := rl.NewRectangle(float32(b.position.X-b.width/2.0), float32(b.position.Y-b.height/2.0), float32(b.width), float32(b.height))
        // if rl.CheckCollisionPointRec(mousePos, buttonRect) {
            // b.SetColor(rl.ColorAlpha(b.Color(), float32(alphaHover)/math.MaxUint8))
        // } else {
            // b.SetColor(rl.ColorAlpha(b.Color(), float32(alphaBase)/math.MaxUint8))
        // }
//
    // }
// }
