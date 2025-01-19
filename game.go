package raychip

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
	"math"
)

type Game struct {
	screenWidth     int32
	screenHeight    int32
	targetFPS       int32
	windowName      string
	space           *cp.Space
	entities        []Entity
	backgroundColor rl.Color
	updateCallback  func(*Game)
	drawCallback    func(*Game)
}

func NewGame(screenWidth int32, screenHeight int32, targetFPS int32) Game {
	space := cp.NewSpace()
	game := Game{
		screenWidth:     screenWidth,
		screenHeight:    screenHeight,
		targetFPS:       targetFPS,
		windowName:      "Game",
		backgroundColor: rl.RayWhite,
		space:           space,
	}
	rl.InitWindow(game.screenWidth, game.screenHeight, game.windowName)
	rl.SetTargetFPS(game.targetFPS)
	return game
}

func (game *Game) SetWindowName(name string) {
	game.windowName = name
}

func (game Game) EntitiesCount() int {
	return len(game.entities)
}

func (game *Game) AddPerimiterWall(width float64, color rl.Color) {
	v_tl := NewVector2(0, 0)
	v_tr := NewVector2(float64(game.screenWidth), 0)
	v_br := NewVector2(float64(game.screenWidth), float64(game.screenHeight))
	v_bl := NewVector2(0, float64(game.screenHeight))
	wallTop := NewWall(v_tl, v_tr, width, color)
	wallRight := NewWall(v_tr, v_br, width, color)
	wallBot := NewWall(v_bl, v_br, width, color)
	wallLeft := NewWall(v_tl, v_bl, width, color)
	game.AddEntity(&wallTop)
	game.AddEntity(&wallRight)
	game.AddEntity(&wallBot)
	game.AddEntity(&wallLeft)
}

func (game Game) Dt() float64 {
	return 1.0 / float64(game.targetFPS)
}

func (game *Game) SetBackgroundColor(color rl.Color) {
	game.backgroundColor = color
}

func (game *Game) SetUpdateCallback(callback func(*Game)) {
	game.updateCallback = callback
}

func (game *Game) SetDrawCallback(callback func(*Game)) {
	game.drawCallback = callback
}

func (game *Game) SetGravity(v Vector2) {
	if game.space != nil {
		game.space.SetGravity(v.ToChipmunk())
	}
}

func (game *Game) SetDamping(d float64) {
	if game.space != nil {
		game.space.SetDamping(d)
	}
}

func (game *Game) Update() {
	game.space.Step(game.Dt())
	for _, entity := range game.entities {
        entity.Update()
	}
}

func (game Game) Draw() {
	for _, entity := range game.entities {
        entity.Draw()
	}
}

func (game *Game) Run() {

	for !rl.WindowShouldClose() {

        // call default game.Update()
		game.Update() 

        // Call custom game.Update() if defined
		if game.updateCallback != nil {
			game.updateCallback(game)
		}

		// ---------- Drawing ----------
		rl.BeginDrawing()
		rl.ClearBackground(game.backgroundColor)

        // call default game.Draw()
		game.Draw()

        // call custom game.Draw() if defined
		if game.drawCallback != nil {
			game.drawCallback(game)
		}
		rl.EndDrawing()
		// -----------------------------
	}

	rl.CloseWindow()
}

// TODO: can these two functions be combined?
func (p Circle) limitVelocity(body *cp.Body, gravity cp.Vector, damping float64, dt float64) {
	maxSpeed := p.velocityMax // Maximum speed (pixels/second)
	cp.BodyUpdateVelocity(body, gravity, damping, dt)
	velocity := body.Velocity()
	speed := math.Sqrt(velocity.X*velocity.X + velocity.Y*velocity.Y)
	if speed > maxSpeed {
		scale := maxSpeed / speed
		body.SetVelocity(velocity.X*scale, velocity.Y*scale)
	}
}

func (b Box) limitVelocity(body *cp.Body, gravity cp.Vector, damping float64, dt float64) {
	maxSpeed := b.velocityMax // Maximum speed (pixels/second)
	cp.BodyUpdateVelocity(body, gravity, damping, dt)
	velocity := body.Velocity()
	speed := math.Sqrt(velocity.X*velocity.X + velocity.Y*velocity.Y)
	if speed > maxSpeed {
		scale := maxSpeed / speed
		body.SetVelocity(velocity.X*scale, velocity.Y*scale)
	}
}

// TODO: create AddEntity function for each entity similar to draw and update callbacks?
func (game *Game) AddEntity(entity Entity) {
	var body *cp.Body
	var shape *cp.Shape

	switch e := entity.(type) {
	case *Circle:
        if e.physical {
            body = game.space.AddBody(cp.NewBody(e.mass, cp.MomentForCircle(e.mass, 0.0, e.radius, cp.Vector{})))
            body.SetType(cp.BODY_DYNAMIC)
            body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
            body.SetVelocity(e.velocity.X, e.velocity.Y)
            shape = game.space.AddShape(cp.NewCircle(body, e.radius, cp.Vector{}))
            shape.SetElasticity(e.elasticity)
            shape.SetFriction(e.friction)
            body.SetVelocityUpdateFunc(e.limitVelocity)
            e.cpBody = body
            e.cpShape = shape
        }
		e.id = uint64(len(game.entities))
		game.entities = append(game.entities, e)
	case *Box:
        if e.physical {
            body = game.space.AddBody(cp.NewBody(e.mass, cp.MomentForBox(e.mass, e.width, e.height)))
            body.SetPosition(cp.Vector{X: e.position.X, Y: e.position.Y})
            body.SetVelocity(e.velocity.X, e.velocity.Y)
            shape = game.space.AddShape(cp.NewBox(body, e.width, e.height, 0))
            shape.SetElasticity(e.elasticity)
            shape.SetFriction(e.friction)
            body.SetVelocityUpdateFunc(e.limitVelocity)
            e.cpBody = body
            e.cpShape = shape
        }
		e.id = uint64(len(game.entities))
		game.entities = append(game.entities, e)
	case *Wall:
		body = cp.NewStaticBody()
		shape = game.space.AddShape(cp.NewSegment(body, cp.Vector{X: e.Vertex1.X, Y: e.Vertex1.Y}, cp.Vector{X: e.Vertex2.X, Y: e.Vertex2.Y}, e.Width/2))
		shape.SetElasticity(1)
		shape.SetFriction(1)

		e.id = uint64(len(game.entities))
		e.cpBody = body
		// e.cpShape = shape
		game.entities = append(game.entities, e)
	default:
		fmt.Println("Unknown entity type")
		return
	}
}
