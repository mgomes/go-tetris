package main

import (
	"time"
	
	"github.com/go-gl/glfw/v3.3/glfw"
)

type InputHandler struct {
	keyStates     map[glfw.Key]bool
	keyTimers     map[glfw.Key]time.Time
	keyPressed    map[glfw.Key]bool
	keyRepeatTimers map[glfw.Key]time.Time
}

func NewInputHandler() *InputHandler {
	return &InputHandler{
		keyStates:     make(map[glfw.Key]bool),
		keyTimers:     make(map[glfw.Key]time.Time),
		keyPressed:    make(map[glfw.Key]bool),
		keyRepeatTimers: make(map[glfw.Key]time.Time),
	}
}

func (ih *InputHandler) HandleKeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		ih.keyStates[key] = true
		ih.keyTimers[key] = time.Now()
		ih.keyPressed[key] = true
		delete(ih.keyRepeatTimers, key)
	} else if action == glfw.Release {
		ih.keyStates[key] = false
		delete(ih.keyTimers, key)
		delete(ih.keyPressed, key)
		delete(ih.keyRepeatTimers, key)
	}
}

func (ih *InputHandler) IsKeyJustPressed(key glfw.Key) bool {
	if pressed, ok := ih.keyStates[key]; ok && pressed {
		if timer, ok := ih.keyTimers[key]; ok {
			if time.Since(timer) < keyJustPressedWindow {
				return true
			}
		}
	}
	return false
}

func (ih *InputHandler) IsKeyPressed(key glfw.Key) bool {
	return ih.keyPressed[key]
}

func (ih *InputHandler) ConsumeKeyPress(key glfw.Key) {
	delete(ih.keyPressed, key)
}

func (ih *InputHandler) IsKeyRepeating(key glfw.Key) bool {
	return ih.isKeyRepeatingWithSpeed(key, keyRepeatDelay, keyRepeatInterval, keyRepeatFastInterval, keyFastThreshold)
}

func (ih *InputHandler) isKeyRepeatingWithSpeed(key glfw.Key, repeatDelay, repeatInterval, fastInterval, fastThreshold time.Duration) bool {
	if pressed, ok := ih.keyStates[key]; ok && pressed {
		if timer, ok := ih.keyTimers[key]; ok {
			elapsed := time.Since(timer)
			if elapsed > repeatDelay {
				interval := repeatInterval
				if elapsed > fastThreshold {
					interval = fastInterval
				}
				
				if lastRepeat, ok := ih.keyRepeatTimers[key]; !ok || time.Since(lastRepeat) >= interval {
					ih.keyRepeatTimers[key] = time.Now()
					return true
				}
			}
		}
	}
	return false
}

func (ih *InputHandler) ProcessGameInput(game *Game, window *glfw.Window) {
	// System controls
	if ih.IsKeyPressed(glfw.KeyEscape) {
		window.SetShouldClose(true)
		ih.ConsumeKeyPress(glfw.KeyEscape)
		return
	}

	if ih.IsKeyPressed(glfw.KeyP) {
		game.Paused = !game.Paused
		ih.ConsumeKeyPress(glfw.KeyP)
	}

	// Game over controls
	if game.GameOver {
		if ih.IsKeyPressed(glfw.KeyR) {
			*game = *NewGame()
			ih.ConsumeKeyPress(glfw.KeyR)
		}
		return
	}

	if game.Paused {
		return
	}

	// Movement controls
	ih.processMovementInput(game)
	
	// Rotation controls
	ih.processRotationInput(game)
	
	// Special action controls
	ih.processActionInput(game)
}

func (ih *InputHandler) processMovementInput(game *Game) {
	if ih.IsKeyPressed(glfw.KeyLeft) {
		game.MovePiece(-1, 0)
		ih.ConsumeKeyPress(glfw.KeyLeft)
	} else if ih.IsKeyRepeating(glfw.KeyLeft) {
		game.MovePiece(-1, 0)
	}
	
	if ih.IsKeyPressed(glfw.KeyRight) {
		game.MovePiece(1, 0)
		ih.ConsumeKeyPress(glfw.KeyRight)
	} else if ih.IsKeyRepeating(glfw.KeyRight) {
		game.MovePiece(1, 0)
	}
	
	if ih.IsKeyPressed(glfw.KeyDown) {
		game.MovePiece(0, 1)
		ih.ConsumeKeyPress(glfw.KeyDown)
	} else if ih.IsKeyRepeating(glfw.KeyDown) {
		game.MovePiece(0, 1)
	}
}

func (ih *InputHandler) processRotationInput(game *Game) {
	if ih.IsKeyPressed(glfw.KeyUp) {
		game.RotatePiece(true)
		ih.ConsumeKeyPress(glfw.KeyUp)
	}
	
	if ih.IsKeyPressed(glfw.KeyLeftShift) || ih.IsKeyPressed(glfw.KeyRightShift) {
		game.RotatePiece(false)
		ih.ConsumeKeyPress(glfw.KeyLeftShift)
		ih.ConsumeKeyPress(glfw.KeyRightShift)
	}
}

func (ih *InputHandler) processActionInput(game *Game) {
	if ih.IsKeyPressed(glfw.KeySpace) {
		game.HardDrop()
		ih.ConsumeKeyPress(glfw.KeySpace)
	}
	
	if ih.IsKeyPressed(glfw.KeyLeftControl) || ih.IsKeyPressed(glfw.KeyRightControl) {
		game.HoldPiece()
		ih.ConsumeKeyPress(glfw.KeyLeftControl)
		ih.ConsumeKeyPress(glfw.KeyRightControl)
	}
}