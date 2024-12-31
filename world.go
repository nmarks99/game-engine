package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
    "math"
)


type World struct {
    particles [] *Particle
    accel rl.Vector2
    mousePos rl.Vector2 
}

func NewWorld() World {
    return World {
        accel: rl.NewVector2(0.0, 0.0),
    }
}

func (w *World) SetAcceleration(acc rl.Vector2) {
    w.accel = acc
}

func (w *World) AddParticle(p *Particle) {
    w.particles = append(w.particles, p)
}

func (w *World) Step (mousePos rl.Vector2) {
    w.mousePos = mousePos
    for _, p := range w.particles {
        p.Step(*w)

        // handle collisions
        // make separate functions for this and check all collisions
        // implement particle-particle collisions
        const FLOORY float32 = BORDER_HEIGHT + BORDER_Y
        const DAMPING float32 = 0.7
        if math.Abs(float64(FLOORY) - float64(p.Position.Y)) <= float64(p.Radius) {
            p.Position.Y = FLOORY - p.Radius
            p.Velocity.Y = -p.Velocity.Y * DAMPING
        }
    }
}

func (w World) Draw () {
    for _, p := range w.particles {
        p.Draw()
    }
}
