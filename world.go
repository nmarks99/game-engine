package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RigidBody interface {
    Step(w World)
    Draw()
}

type World struct {
    bodies []RigidBody
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

func (w *World) AddRigidBody(body RigidBody) {
    w.bodies = append(w.bodies, body)
}

func (w *World) Step (mousePos rl.Vector2) {
    w.mousePos = mousePos
    for _, body := range w.bodies {
        body.Step(*w)
    }
}

func (w World) Draw () {
    for _, body := range w.bodies {
        body.Draw()
    }
}
