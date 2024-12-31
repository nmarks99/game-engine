package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
    Position rl.Vector2
    Velocity rl.Vector2
    Radius float32
    Color rl.Color
    defaultColor rl.Color
    clickedColor rl.Color
    Mass float32
    dragging bool
}

func NewParticle(position rl.Vector2, radius float32, mass float32, color rl.Color) Particle {
    return Particle {
        Position : position,
        Radius : radius,
        Mass : mass,
        Color : color,
        defaultColor : color,
        clickedColor : rl.Green,
    }
}

func (p *Particle) Step(w World) {
    if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
        if rl.CheckCollisionPointCircle(w.mousePos, p.Position, p.Radius) {
            p.Color = p.clickedColor
            p.Position = w.mousePos
            p.dragging = true
        }
    }
    if p.dragging {
        if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
            p.dragging = false
            p.Color = p.defaultColor
        } else {
            p.Position = w.mousePos
        }
    } else {
        newVel := rl.Vector2Add(p.Velocity, rl.Vector2Scale(w.accel, dt))
        newPos := rl.Vector2Add(p.Position, rl.Vector2Scale(newVel, dt))
        p.Velocity = newVel
        p.Position = newPos 
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
