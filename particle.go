package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
    "fmt"
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

func limitVelocity(v rl.Vector2, vmax float32) rl.Vector2 {
    vout := rl.NewVector2(v.X, v.Y)
    if vout.X >= vmax {
        vout.X = vmax
    } else if vout.X <= -vmax {
        vout.X = -vmax
    }

    if vout.Y >= vmax {
        vout.Y = vmax
    } else if vout.Y <= -vmax {
        vout.Y = -vmax
    }
    return vout
}

func (p *Particle) Update(eng Engine) {
    if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
        if rl.CheckCollisionPointCircle(eng.mousePos, p.Position, p.Radius) {
            p.Color = p.clickedColor
            p.Position = eng.mousePos
            p.dragging = true
            p.Velocity.X = 0.0
            p.Velocity.Y = 0.0
        }
    }
    if p.dragging {
        if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
            p.Color = p.defaultColor
            mouseVel := rl.NewVector2(rl.GetMouseDelta().X / dt, rl.GetMouseDelta().Y / dt)
            fmt.Printf("before velocity = %f, %f\n", mouseVel.X, mouseVel.Y)
            mouseVel = limitVelocity(mouseVel, MAX_VELOCITY)
            fmt.Printf("limited velocity = %f, %f\n", mouseVel.X, mouseVel.Y)
            p.Velocity = mouseVel
            p.dragging = false
        } else {
            p.Position = eng.mousePos
        }
    } else {
        newVel := rl.Vector2Add(p.Velocity, rl.Vector2Scale(eng.gravity, dt))
        newVel = limitVelocity(newVel, MAX_VELOCITY)
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
