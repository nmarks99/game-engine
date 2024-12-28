package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
    Position rl.Vector2
    Velocity rl.Vector2
    Radius float32
    Color rl.Color
    Mass float32
}

func NewParticle(position rl.Vector2, radius float32, mass float32, color rl.Color) Particle {
    return Particle {
        Position : position,
        Radius : radius,
        Mass : mass,
        Color : color,
    }
}

func (p Particle) Draw() {
    rl.DrawCircleV(p.Position, p.Radius, p.Color);
}

func (p *Particle) IncrementX(inc float32) {
    p.Position.X += inc
}

func (p *Particle) IncrementY(inc float32) {
    p.Position.Y += inc
}

func (p *Particle) SetY(y float32) {
    p.Position.Y = y
}

func (p *Particle) SetX(x float32) {
    p.Position.X = x
}

func (p *Particle) SetPosition(position rl.Vector2) {
    p.Position = position
}
