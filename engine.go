package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "math"
)

type Engine struct {
	particles  []*Particle
	gravity    rl.Vector2
	mousePos   rl.Vector2
	borderRect rl.Rectangle
}

const SCREEN_WIDTH int32 = 700
const SCREEN_HEIGHT int32 = 600
const BORDER_WIDTH float32 = float32(SCREEN_WIDTH) * 0.95
const BORDER_HEIGHT float32 = float32(SCREEN_HEIGHT) * 0.95
const BORDER_X float32 = (float32(SCREEN_WIDTH) - BORDER_WIDTH) / 2.0
const BORDER_Y float32 = (float32(SCREEN_HEIGHT) - BORDER_HEIGHT) / 2.0
const MAX_VELOCITY float32 = 2000.0 // pixels/sec

func NewEngine() Engine {
	return Engine{
		gravity:    rl.NewVector2(0.0, 0.0),
		borderRect: rl.NewRectangle(BORDER_X, BORDER_Y, BORDER_WIDTH, BORDER_HEIGHT),
	}
}

func (eng *Engine) SetGravity(acc rl.Vector2) {
	eng.gravity = acc
}

func (eng *Engine) AddParticle(p *Particle) {
	eng.particles = append(eng.particles, p)
}

func (eng Engine) checkBorderCollision(p *Particle) {
	const DAMPING float32 = 0.5
	topLeft := rl.NewVector2(eng.borderRect.X, eng.borderRect.Y)
	topRight := rl.NewVector2(eng.borderRect.X+eng.borderRect.Width, eng.borderRect.Y)
	bottomLeft := rl.NewVector2(eng.borderRect.X, eng.borderRect.Y+eng.borderRect.Height)
	bottomRight := rl.NewVector2(eng.borderRect.X+eng.borderRect.Width, eng.borderRect.Y+eng.borderRect.Height)

	// Top border
	if rl.CheckCollisionCircleLine(p.Position, p.Radius, topLeft, topRight) {
		p.Position.Y = eng.borderRect.Y + p.Radius
		p.Velocity.Y = -p.Velocity.Y * DAMPING
	}

	// bottom border
	if rl.CheckCollisionCircleLine(p.Position, p.Radius, bottomLeft, bottomRight) {
		p.Position.Y = (eng.borderRect.Height + eng.borderRect.Y) - p.Radius
		p.Velocity.Y = -p.Velocity.Y * DAMPING
	}

	// Left border
	if rl.CheckCollisionCircleLine(p.Position, p.Radius, topLeft, bottomLeft) {
		p.Position.X = eng.borderRect.X + p.Radius
		p.Velocity.X = -p.Velocity.X * DAMPING
	}

	// Right border
	if rl.CheckCollisionCircleLine(p.Position, p.Radius, topRight, bottomRight) {
		p.Position.X = (eng.borderRect.Width + eng.borderRect.X) - p.Radius
		p.Velocity.X = -p.Velocity.X * DAMPING
	}

}

func (eng *Engine) Update(mousePos rl.Vector2) {
	eng.mousePos = mousePos
	for _, p := range eng.particles {
		p.Update(*eng)
		eng.checkBorderCollision(p)
	}
}

func (eng Engine) Draw() {
	for _, p := range eng.particles {
		p.Draw()
		rl.DrawRectangleLinesEx(eng.borderRect, 5.0, rl.Black)
	}
}
