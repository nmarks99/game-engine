package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const MAX_VELOCITY float32 = 3000.0 // pixels/sec
const DAMPING float32 = 1.0

type Engine struct {
	entities   []interface{}
	gravity    rl.Vector2
	mousePos   rl.Vector2
	borderRect rl.Rectangle
}

func NewEngine() Engine {
	var border_width float32 = float32(rl.GetScreenWidth()) * 1.0
	var border_height float32 = float32(rl.GetScreenHeight()) * 1.0
	var border_x float32 = (float32(rl.GetScreenWidth()) - border_width) / 2.0
	var border_y float32 = (float32(rl.GetScreenHeight()) - border_height) / 2.0
	return Engine{
		gravity:    rl.NewVector2(0.0, 0.0),
		borderRect: rl.NewRectangle(border_x, border_y, border_width, border_height),
	}
}

func (eng *Engine) SetGravity(acc rl.Vector2) {
	eng.gravity = acc
}

func (eng *Engine) AddParticle(p *Particle) {
	p.id = uint32(len(eng.entities))
	fmt.Printf("Adding particle entity with ID = %d\n", p.id)
	eng.entities = append(eng.entities, p)
}

func (eng *Engine) AddBlock(b *Block) {
	b.id = uint32(len(eng.entities))
	fmt.Printf("Adding block entity with ID = %d\n", b.id)
	eng.entities = append(eng.entities, b)
}

func (eng Engine) resolveBorderCollision(p *Particle) {
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

// update velocities for particle-particle collision
func updateParticleVelocities(p1 *Particle, p2 *Particle) {
	massFactor := (2 * p2.Mass) / (p1.Mass + p2.Mass)
	velocityDiff := rl.Vector2Subtract(p1.Velocity, p2.Velocity)
	positionDiff := rl.Vector2Subtract(p1.Position, p2.Position)
	dotProd := rl.Vector2DotProduct(velocityDiff, positionDiff)
	magSqr := rl.Vector2Length(positionDiff) * rl.Vector2Length(positionDiff)
	scalar := massFactor * dotProd / magSqr
	v1_new := rl.Vector2Subtract(p1.Velocity, rl.Vector2Scale(positionDiff, scalar))

	massFactor = (2 * p1.Mass) / (p1.Mass + p2.Mass)
	velocityDiff = rl.Vector2Subtract(p2.Velocity, p1.Velocity)
	positionDiff = rl.Vector2Subtract(p2.Position, p1.Position)
	dotProd = rl.Vector2DotProduct(velocityDiff, positionDiff)
	magSqr = rl.Vector2Length(positionDiff) * rl.Vector2Length(positionDiff)
	scalar = massFactor * dotProd / magSqr
	v2_new := rl.Vector2Subtract(p2.Velocity, rl.Vector2Scale(positionDiff, scalar))

	p1.Velocity = v1_new
	p2.Velocity = v2_new
}

// Make sure particles don't end up overlapping\
// FIX: understand this better
func resolveParticleOverlap(p1, p2 *Particle) {
	distance := rl.Vector2Length(rl.Vector2Subtract(p1.Position, p2.Position))
	overlap := p1.Radius + p2.Radius - distance
	if overlap > 0 {
		correction := rl.Vector2Scale(rl.Vector2Normalize(rl.Vector2Subtract(p1.Position, p2.Position)), overlap/2)
		p1.Position = rl.Vector2Add(p1.Position, correction)
		p2.Position = rl.Vector2Subtract(p2.Position, correction)
	}
}

// Check for particle-particle collision and resolve velocities and prevent overlapping
func (eng Engine) resolveParticleCollision(p1 *Particle) {
	for _, p2 := range eng.entities {
		switch p2 := p2.(type) {
		case *Particle:
			if p1.id != p2.id {
				if rl.CheckCollisionCircles(p1.Position, p1.Radius, p2.Position, p2.Radius) {
					resolveParticleOverlap(p1, p2)
					updateParticleVelocities(p1, p2)
				}
			}
		}
	}
}

func (eng *Engine) Update(mousePos rl.Vector2) {
	eng.mousePos = mousePos
	for _, v := range eng.entities {
		switch entity := v.(type) {
		case *Particle:
            eng.resolveBorderCollision(entity)
            eng.resolveParticleCollision(entity)
            entity.Update(*eng)
		case *Block:
            entity.Update(*eng)
        default:
            fmt.Println("Unknown entity type!")
	    }
    }
}

func (eng Engine) Draw() {
	for _, v := range eng.entities {
		switch entity := v.(type) {
		case *Particle:
            entity.Draw()
		case *Block:
            entity.Draw()
        default:
            fmt.Println("Unknown entity type!")
        }
	}
}
