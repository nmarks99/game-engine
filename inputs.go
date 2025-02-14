package raychip

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameInputs struct {
	MouseInputEnable    bool
	KeyboardInputEnable bool
	GamepadInputEnable  bool
	InputPublishers     []*Publisher
}

func (g *Game) EnableMouseInput() {
	g.inputs.MouseInputEnable = true
}

func (g *Game) DisableMouseInput() {
	g.inputs.MouseInputEnable = false
}

func (g *Game) EnableKeyboardInput() {
	fmt.Println("Keyboard inputs not implemented!")
	if !g.inputs.KeyboardInputEnable {
		g.inputs.KeyboardInputEnable = true
	}
}

func (g *Game) EnableGamepadInput() {
	fmt.Println("Keyboard inputs not implemented!")
	if !g.inputs.GamepadInputEnable {
		g.inputs.GamepadInputEnable = true
	}
}

type MouseButtonState struct {
	Pressed  bool
	Released bool
	Up       bool
	Down     bool
}

type MouseInputEvent struct {
	ButtonStateMap map[rl.MouseButton]MouseButtonState
	Position       Vector2
}

func (m MouseInputEvent) IsButtonReleased(button rl.MouseButton) bool {
	return m.ButtonStateMap[button].Released
}

func (m MouseInputEvent) IsButtonPressed(button rl.MouseButton) bool {
	return m.ButtonStateMap[button].Pressed
}

func (m MouseInputEvent) IsButtonUp(button rl.MouseButton) bool {
	return m.ButtonStateMap[button].Up
}

func (m MouseInputEvent) IsButtonDown(button rl.MouseButton) bool {
	return m.ButtonStateMap[button].Down
}

type KeyboardInputEvent struct {
	Key   int
	State bool
}

type GamepadInputEvent struct {
	Button int
	State  bool
}

func getMouseInputEvent() MouseInputEvent {
	mousePos := rl.GetMousePosition()
	var buttons = [4]rl.MouseButton{rl.MouseButtonLeft, rl.MouseButtonRight, rl.MouseButtonMiddle}

	mouseInputEvent := MouseInputEvent{
		ButtonStateMap: make(map[rl.MouseButton]MouseButtonState),
		Position:       Vector2FromRaylib(mousePos),
	}

	for _, button := range buttons {
		state := MouseButtonState{
			Pressed:  false,
			Released: false,
			Up:       false,
			Down:     false,
		}
		if rl.IsMouseButtonPressed(button) {
			state.Pressed = true
		}
		if rl.IsMouseButtonReleased(button) {
			state.Released = true
		}
		if rl.IsMouseButtonUp(button) {
			state.Up = true
		}
		if rl.IsMouseButtonDown(button) {
			state.Down = true
		}
		mouseInputEvent.ButtonStateMap[button] = state
	}

	return mouseInputEvent
}

// TODO:
// func getKeyboardInputEvent() KeyboardInputEvent {
// }

// TODO:
// func getGamepadInputEvent() GamepadInputEvent {
// }
