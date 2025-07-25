package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Window constants are now in constants.go

var inputHandler *InputHandler

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		log.Fatalln("failed to create window:", err)
	}

	window.MakeContextCurrent()
	
	inputHandler = NewInputHandler()
	window.SetKeyCallback(inputHandler.HandleKeyCallback)

	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	game := NewGame()
	renderer := NewRenderer(windowWidth, windowHeight)
	renderer.SetupProjection()

	lastFrame := time.Now()

	for !window.ShouldClose() {
		currentFrame := time.Now()
		deltaTime := currentFrame.Sub(lastFrame)
		lastFrame = currentFrame

		inputHandler.ProcessGameInput(game, window)
		game.Update()

		renderer.Clear()
		renderer.DrawBoard(game.Board)
		renderer.DrawGhostPiece(game)
		renderer.DrawPiece(game.CurrentPiece)
		renderer.DrawHeldPiece(game.HeldPiece)
		renderer.DrawUI(game)

		window.SwapBuffers()
		glfw.PollEvents()

		if deltaTime < frameTargetTime {
			time.Sleep(frameTargetTime - deltaTime)
		}
	}
}

