package raychip

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jakecoffman/cp/v2"
)

type Game struct {
	screenWidth     int32
	screenHeight    int32
	targetFPS       int32
	windowName      string
	physical        bool
	space           *cp.Space
	entities        []Entity
	backgroundColor rl.Color
	updateCallback  func(*Game)
	drawCallback    func(*Game)
	mousePosition   Vector2
	EventBus        EventBus
	inputs          GameInputs
}

func NewGame(screenWidth int32, screenHeight int32, targetFPS int32) Game {
	space := cp.NewSpace()
	game := Game{
		screenWidth:     screenWidth,
		screenHeight:    screenHeight,
		targetFPS:       targetFPS,
		physical:        false,
		windowName:      "Game",
		backgroundColor: rl.RayWhite,
		space:           space,
		EventBus:        NewEventBus(),
	}
	rl.InitWindow(game.screenWidth, game.screenHeight, game.windowName)
	rl.SetTargetFPS(game.targetFPS)
	return game
}

type MouseState int

const (
	MouseUp MouseState = iota
	MouseDown
	MousePressed
	MouseReleased
)

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

func (game Game) MousePosition() Vector2 {
	return game.mousePosition
}

func (game *Game) SetDrawCallback(callback func(*Game)) {
	game.drawCallback = callback
}

func (game *Game) SetUpdateCallback(callback func(*Game)) {
	var oldUpdateCallback func(*Game)
	if game.updateCallback != nil {
		oldUpdateCallback = game.updateCallback
	}
	game.updateCallback = func(g *Game) {
		if oldUpdateCallback != nil {
			oldUpdateCallback(g)
		}
		callback(g)
	}
}

func (game *Game) OnClick(button rl.MouseButton, state MouseState, callback func()) {
	var oldUpdateCallback func(*Game)
	if game.updateCallback != nil {
		oldUpdateCallback = game.updateCallback
	}

	game.updateCallback = func(g *Game) {

		if oldUpdateCallback != nil {
			oldUpdateCallback(g)
		}

		switch state {
		case MouseUp:
			if rl.IsMouseButtonUp(button) {
				callback()
			}
		case MouseDown:
			if rl.IsMouseButtonDown(button) {
				callback()
			}
		case MousePressed:
			if rl.IsMouseButtonPressed(button) {
				callback()
			}
		case MouseReleased:
			if rl.IsMouseButtonReleased(button) {
				callback()
			}
		}
	}
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

    // step the physics forward if enabled
	if game.physical {
		game.space.Step(game.Dt())
	}

    // Update all game entites
	for _, entity := range game.entities {
		entity.Update()
	}

    // publish inputs if enabled
    if game.inputs.MouseInputEnable {
        mouseInputEvent := getMouseInputEvent()
        game.EventBus.Publish("input.mouse", mouseInputEvent)
    }
    if game.inputs.KeyboardInputEnable {
        // game.EventBus.Publish("input.keyboard", keyboardInputEvent)
    }
    if game.inputs.GamepadInputEnable {
        // game.EventBus.Publish("input.gamepad", gamepadInputEvent)
    }

    // TODO: remove this
    game.mousePosition = Vector2FromRaylib(rl.GetMousePosition())
}

func (game Game) Draw() {
	for _, entity := range game.entities {
		entity.Draw()
	}
}

func (game *Game) Run() {

	for !rl.WindowShouldClose() {

		game.Update()
		if game.updateCallback != nil {
			game.updateCallback(game)
		}

		// ---------- Drawing ----------
		rl.BeginDrawing()
		rl.ClearBackground(game.backgroundColor)
		game.Draw()
		if game.drawCallback != nil {
			game.drawCallback(game)
		}
		rl.EndDrawing()
		// -----------------------------
	}

	rl.CloseWindow()
}

func (game *Game) AddEntity(entity Entity) {
	var body *cp.Body
	var shape *cp.Shape
	entity.addToGame(game, body, shape)
}

func (game *Game) RemoveEntity(entity Entity) {
	var found bool = false
	var ind int
	for i, v := range game.entities {
		if v.Id() == entity.Id() {
			found = true
			ind = i
		}
	}
	if found {
		game.entities = append(game.entities[:ind], game.entities[ind+1:]...)
	}
}

func (game *Game) ClearEntities() {
	game.entities = game.entities[:0]
}

type Scene struct {
	entities []Entity
}

func NewScene() Scene {
	return Scene{}
}

func (s *Scene) AddEntity(entity Entity) {
	s.entities = append(s.entities, entity)
}

func (scene *Scene) RemoveEntity(entity Entity) {
	var found bool = false
	var ind int
	for i, v := range scene.entities {
		if v.Id() == entity.Id() {
			found = true
			ind = i
		}
	}
	if found {
		scene.entities = append(scene.entities[:ind], scene.entities[ind+1:]...)
	}
}

func (game *Game) SetScene(scene Scene) {
	game.ClearEntities()
	for i := range scene.entities {
		game.AddEntity(scene.entities[i])
	}
}
