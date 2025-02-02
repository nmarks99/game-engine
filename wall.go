package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

type Wall struct {
    EntityBase
	vertex1 Vector2
	vertex2 Vector2
	width   float64
	// Color   rl.Color
	// id      uint64
	// cpBody  *cp.Body
}

func NewWall(vertex1 Vector2, vertex2 Vector2, width float64, color rl.Color) Wall {
	return Wall{
        EntityBase : EntityBase{
            color:   color,
        },
        vertex1: vertex1,
        vertex2: vertex2,
        width:   width,
    }
}

func (e *Wall) addToGame(game *Game, args ...any) {
	if body, ok := args[0].(*cp.Body); ok {
		if shape, ok := args[1].(*cp.Shape); ok {
			body = cp.NewStaticBody()
			shape = game.space.AddShape(cp.NewSegment(body, cp.Vector{X: e.vertex1.X, Y: e.vertex1.Y}, cp.Vector{X: e.vertex2.X, Y: e.vertex2.Y}, e.width/2))
			shape.SetElasticity(1)
			shape.SetFriction(1)
			e.id = uint64(len(game.entities))
		}
	}
	e.id = uint64(len(game.entities))
	game.entities = append(game.entities, e)
}

func (w *Wall) Update() {}

func (w *Wall) Draw() {
	rl.DrawLineEx(w.vertex1.ToRaylib(), w.vertex2.ToRaylib(), float32(w.width), w.color)
}

func (w Wall) Vertex1() Vector2 {
    return w.vertex1
}

func (w Wall) Vertex2() Vector2 {
    return w.vertex1
}

func (w Wall) Width() float64 {
    return w.width
}
