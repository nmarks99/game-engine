package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
    "fmt"
)

type Block struct {
	Position rl.Vector2
	Size     rl.Vector2
	Mass     float32
	Color    rl.Color
	id       uint32
}

func NewBlock(position rl.Vector2, size rl.Vector2, mass float32, color rl.Color) Block {
	return Block{
		Position: position,
		Size:     size,
		Mass:     mass,
		Color:    color,
	}
}

func (b *Block) Update(eng Engine) {
    fmt.Println("TODO: update block here")
}

func (b Block) Draw() {
	rl.DrawRectangleV(b.Position, b.Size, b.Color)
}
