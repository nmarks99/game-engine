package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
    Position rl.Vector2
    Size rl.Vector2
    Text string
    Color rl.Color
    action func()
}

func NewButton(text string, action func(), position rl.Vector2, size rl.Vector2, color rl.Color) Button {
    return Button {
        Text: text,
        Position: position,
        Size: size,
        Color: color,
        action: action,
    }
}

func (b Button) Draw(mousePos rl.Vector2) {
    const HOVER_ALPHA uint8 = 128
    const DEFAULT_ALPHA uint8 = 255
    if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(b.Position.X, b.Position.Y, b.Size.X, b.Size.Y)) {
        b.Color.A = HOVER_ALPHA
        if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
            b.action()
        }
    } else {
        b.Color.A = DEFAULT_ALPHA
    }
    rl.DrawRectangleV(b.Position, b.Size, b.Color)
    rl.DrawText(b.Text, int32(b.Position.X)+2, int32(b.Position.Y), int32(b.Size.Y), rl.White)
}

