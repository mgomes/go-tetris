package main

import (
	"fmt"
	
	"github.com/go-gl/gl/v2.1/gl"
)

// Rendering constants are now in constants.go

type Renderer struct {
	windowWidth  int
	windowHeight int
}

func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		windowWidth:  width,
		windowHeight: height,
	}
}

func (r *Renderer) Clear() {
	// Dark purple/blue gradient background
	gl.ClearColor(0.05, 0.0, 0.15, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	
	// Draw synthwave grid
	r.drawSynthwaveGrid()
}

func (r *Renderer) DrawBoard(board *Board) {
	r.drawBorder()
	
	for y := range boardHeight {
		for x := range boardWidth {
			if board.Grid[y][x] {
				color := board.Colors[y][x]
				r.drawBlock(x, y, color[0], color[1], color[2])
			}
		}
	}
}

func (r *Renderer) DrawPiece(piece *Piece) {
	blocks := piece.GetBlocks()
	for _, block := range blocks {
		x, y := block[0], block[1]
		if y >= 0 {
			r.drawBlock(x, y, piece.Color[0], piece.Color[1], piece.Color[2])
		}
	}
}

func (r *Renderer) DrawGhostPiece(game *Game) {
	if game.Paused {
		return
	}
	
	ghost := &Piece{
		Shape:    game.CurrentPiece.Shape,
		Color:    game.CurrentPiece.Color,
		X:        game.CurrentPiece.X,
		Y:        game.CurrentPiece.Y,
		Rotation: game.CurrentPiece.Rotation,
	}
	
	for game.Board.IsValidPosition(ghost) {
		ghost.Y++
	}
	ghost.Y--
	
	blocks := ghost.GetBlocks()
	for _, block := range blocks {
		x, y := block[0], block[1]
		if y >= 0 {
			r.drawBlock(x, y, ghost.Color[0]*0.3, ghost.Color[1]*0.3, ghost.Color[2]*0.3)
		}
	}
}

func (r *Renderer) DrawHeldPiece(piece *Piece) {
	holdX := holdBoxX
	holdY := holdBoxY
	
	// Draw neon cyan hold box border
	gl.LineWidth(2.0)
	gl.Color3f(0.0, 1.0, 1.0)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(holdX-10), float32(holdY-10))
	gl.Vertex2f(float32(holdX+4*miniBlockSize+10), float32(holdY-10))
	gl.Vertex2f(float32(holdX+4*miniBlockSize+10), float32(holdY+4*miniBlockSize+10))
	gl.Vertex2f(float32(holdX-10), float32(holdY+4*20+10))
	gl.End()
	gl.LineWidth(1.0)
	
	if piece == nil {
		return
	}
	
	// Draw the held piece (scaled down) with simple blocks
	for y, row := range piece.Shape {
		for x, filled := range row {
			if filled {
				pixelX := float32(holdX + x*miniBlockSize)
				pixelY := float32(holdY + y*miniBlockSize)
				
				// Simple flat blocks for UI
				gl.Color3f(piece.Color[0]*0.8, piece.Color[1]*0.8, piece.Color[2]*0.8)
				gl.Begin(gl.QUADS)
				gl.Vertex2f(pixelX, pixelY)
				gl.Vertex2f(pixelX+18, pixelY)
				gl.Vertex2f(pixelX+18, pixelY+18)
				gl.Vertex2f(pixelX, pixelY+18)
				gl.End()
				
				// Outline
				gl.Color3f(piece.Color[0], piece.Color[1], piece.Color[2])
				gl.Begin(gl.LINE_LOOP)
				gl.Vertex2f(pixelX, pixelY)
				gl.Vertex2f(pixelX+18, pixelY)
				gl.Vertex2f(pixelX+18, pixelY+18)
				gl.Vertex2f(pixelX, pixelY+18)
				gl.End()
			}
		}
	}
}

func (r *Renderer) DrawUI(game *Game) {
	// Draw score box
	r.drawInfoBox(scoreBoxX, scoreBoxY, infoBoxWidth, infoBoxHeight, 0.0, 1.0, 0.5) // Neon green
	r.drawNumber(scoreBoxX+10, scoreBoxY+20, game.Score, 0.0, 1.0, 0.5)
	
	// Draw level box
	r.drawInfoBox(levelBoxX, levelBoxY, infoBoxWidth, infoBoxHeight, 1.0, 0.5, 0.0) // Orange
	r.drawNumber(levelBoxX+10, levelBoxY+20, game.Level, 1.0, 0.5, 0.0)
	
	// Draw next piece preview
	nextX := nextBoxX
	nextY := nextBoxY
	
	// Draw neon magenta next box border
	gl.LineWidth(2.0)
	gl.Color3f(1.0, 0.0, 1.0)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(nextX-10), float32(nextY-10))
	gl.Vertex2f(float32(nextX+4*miniBlockSize+10), float32(nextY-10))
	gl.Vertex2f(float32(nextX+4*miniBlockSize+10), float32(nextY+4*miniBlockSize+10))
	gl.Vertex2f(float32(nextX-10), float32(nextY+4*20+10))
	gl.End()
	gl.LineWidth(1.0)
	
	// Draw next piece (scaled down)
	for y, row := range game.NextPiece.Shape {
		for x, filled := range row {
			if filled {
				pixelX := float32(nextX + x*miniBlockSize)
				pixelY := float32(nextY + y*miniBlockSize)
				
				// Simple flat blocks for UI
				gl.Color3f(game.NextPiece.Color[0]*0.8, game.NextPiece.Color[1]*0.8, game.NextPiece.Color[2]*0.8)
				gl.Begin(gl.QUADS)
				gl.Vertex2f(pixelX, pixelY)
				gl.Vertex2f(pixelX+18, pixelY)
				gl.Vertex2f(pixelX+18, pixelY+18)
				gl.Vertex2f(pixelX, pixelY+18)
				gl.End()
				
				// Outline
				gl.Color3f(game.NextPiece.Color[0], game.NextPiece.Color[1], game.NextPiece.Color[2])
				gl.Begin(gl.LINE_LOOP)
				gl.Vertex2f(pixelX, pixelY)
				gl.Vertex2f(pixelX+18, pixelY)
				gl.Vertex2f(pixelX+18, pixelY+18)
				gl.Vertex2f(pixelX, pixelY+18)
				gl.End()
			}
		}
	}
	
	// Draw pause overlay if paused
	if game.Paused {
		// Semi-transparent overlay
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.Color4f(0.0, 0.0, 0.0, 0.7)
		gl.Begin(gl.QUADS)
		gl.Vertex2f(0, 0)
		gl.Vertex2f(float32(r.windowWidth), 0)
		gl.Vertex2f(float32(r.windowWidth), float32(r.windowHeight))
		gl.Vertex2f(0, float32(r.windowHeight))
		gl.End()
		gl.Disable(gl.BLEND)
	}
	
	// Draw game over overlay
	if game.GameOver {
		r.drawGameOverBanner(game)
	}
}

func (r *Renderer) drawGameOverBanner(game *Game) {
	// Semi-transparent dark overlay
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Color4f(0.0, 0.0, 0.0, 0.8)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(float32(r.windowWidth), 0)
	gl.Vertex2f(float32(r.windowWidth), float32(r.windowHeight))
	gl.Vertex2f(0, float32(r.windowHeight))
	gl.End()
	
	// Banner background
	bannerY := float32(r.windowHeight/2 - 100)
	bannerHeight := float32(200)
	
	// Gradient banner with neon glow
	gl.Begin(gl.QUADS)
	// Top edge (darker)
	gl.Color4f(0.1, 0.0, 0.2, 0.9)
	gl.Vertex2f(0, bannerY)
	gl.Vertex2f(float32(r.windowWidth), bannerY)
	// Bottom edge (lighter)
	gl.Color4f(0.3, 0.0, 0.5, 0.9)
	gl.Vertex2f(float32(r.windowWidth), bannerY+bannerHeight)
	gl.Vertex2f(0, bannerY+bannerHeight)
	gl.End()
	
	// Neon border lines
	gl.LineWidth(3.0)
	gl.Color3f(1.0, 0.0, 1.0)
	gl.Begin(gl.LINES)
	gl.Vertex2f(0, bannerY)
	gl.Vertex2f(float32(r.windowWidth), bannerY)
	gl.Vertex2f(0, bannerY+bannerHeight)
	gl.Vertex2f(float32(r.windowWidth), bannerY+bannerHeight)
	gl.End()
	
	// Additional glow lines
	gl.LineWidth(1.0)
	gl.Color3f(0.0, 1.0, 1.0)
	gl.Begin(gl.LINES)
	gl.Vertex2f(0, bannerY+5)
	gl.Vertex2f(float32(r.windowWidth), bannerY+5)
	gl.Vertex2f(0, bannerY+bannerHeight-5)
	gl.Vertex2f(float32(r.windowWidth), bannerY+bannerHeight-5)
	gl.End()
	
	// "GAME OVER" text (stylized with lines)
	r.drawGameOverText(r.windowWidth/2, int(bannerY+50))
	
	// Score display
	scoreY := int(bannerY + 120)
	r.drawCenteredText(r.windowWidth/2, scoreY, "SCORE", 0.0, 1.0, 0.5)
	r.drawCenteredNumber(r.windowWidth/2, scoreY+25, game.Score, 0.0, 1.0, 0.5)
	
	// Instructions
	r.drawCenteredText(r.windowWidth/2, scoreY+60, "PRESS R TO RESTART", 1.0, 0.0, 0.8)
	
	gl.Disable(gl.BLEND)
}

func (r *Renderer) drawGameOverText(centerX, y int) {
	// Large stylized "GAME OVER" using lines
	gl.LineWidth(3.0)
	gl.Color3f(1.0, 0.0, 0.5)
	
	letterSpacing := 35
	letterSize := 25
	startX := centerX - 4*letterSpacing
	
	// G
	x := startX
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize/2))
	gl.Vertex2f(float32(x+letterSize/2), float32(y+letterSize/2))
	gl.End()
	
	// A
	x += letterSpacing
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.Vertex2f(float32(x+letterSize/2), float32(y))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.End()
	gl.Begin(gl.LINES)
	gl.Vertex2f(float32(x+5), float32(y+letterSize/2))
	gl.Vertex2f(float32(x+letterSize-5), float32(y+letterSize/2))
	gl.End()
	
	// M
	x += letterSpacing
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x+letterSize/2), float32(y+letterSize/2))
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.End()
	
	// E
	x += letterSpacing
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.End()
	gl.Begin(gl.LINES)
	gl.Vertex2f(float32(x), float32(y+letterSize/2))
	gl.Vertex2f(float32(x+letterSize-5), float32(y+letterSize/2))
	gl.End()
	
	// Space
	x += letterSpacing + 20
	
	// O
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.End()
	
	// V
	x += letterSpacing
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x+letterSize/2), float32(y+letterSize))
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.End()
	
	// E
	x += letterSpacing
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.End()
	gl.Begin(gl.LINES)
	gl.Vertex2f(float32(x), float32(y+letterSize/2))
	gl.Vertex2f(float32(x+letterSize-5), float32(y+letterSize/2))
	gl.End()
	
	// R
	x += letterSpacing
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex2f(float32(x), float32(y+letterSize))
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x+letterSize), float32(y))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize/2))
	gl.Vertex2f(float32(x), float32(y+letterSize/2))
	gl.End()
	gl.Begin(gl.LINES)
	gl.Vertex2f(float32(x+letterSize/2), float32(y+letterSize/2))
	gl.Vertex2f(float32(x+letterSize), float32(y+letterSize))
	gl.End()
	
	gl.LineWidth(1.0)
}

func (r *Renderer) drawCenteredText(centerX, y int, text string, red, green, blue float32) {
	// Simple centered text - just draw a line for now (would need proper text rendering)
	gl.Color3f(red, green, blue)
	// This is simplified - in a real implementation you'd render actual text
}

func (r *Renderer) drawCenteredNumber(centerX, y int, number int, red, green, blue float32) {
	digits := fmt.Sprintf("%d", number)
	digitWidth := 15
	totalWidth := len(digits) * digitWidth
	startX := centerX - totalWidth/2
	
	r.drawNumber(startX, y, number, red, green, blue)
}

func (r *Renderer) drawBlock(x, y int, red, green, blue float32) {
	pixelX := float32(boardOffsetX + x*cellSize)
	pixelY := float32(boardOffsetY + y*cellSize)
	depth := float32(4) // 3D depth offset
	
	// Draw back face (darker)
	gl.Color3f(red*0.3, green*0.3, blue*0.3)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(pixelX+depth, pixelY+depth)
	gl.Vertex2f(pixelX+cellSize, pixelY+depth)
	gl.Vertex2f(pixelX+cellSize, pixelY+cellSize)
	gl.Vertex2f(pixelX+depth, pixelY+cellSize)
	gl.End()
	
	// Draw right face (medium)
	gl.Color3f(red*0.6, green*0.6, blue*0.6)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY)
	gl.Vertex2f(pixelX+cellSize, pixelY+depth)
	gl.Vertex2f(pixelX+cellSize, pixelY+cellSize)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY+cellSize-depth)
	gl.End()
	
	// Draw bottom face (medium)
	gl.Color3f(red*0.5, green*0.5, blue*0.5)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(pixelX, pixelY+cellSize-depth)
	gl.Vertex2f(pixelX+depth, pixelY+cellSize)
	gl.Vertex2f(pixelX+cellSize, pixelY+cellSize)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY+cellSize-depth)
	gl.End()
	
	// Draw front face (brightest)
	gl.Color3f(red, green, blue)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(pixelX, pixelY)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY+cellSize-depth)
	gl.Vertex2f(pixelX, pixelY+cellSize-depth)
	gl.End()
	
	// Draw neon outline
	gl.LineWidth(2.0)
	gl.Color3f(red*1.2, green*1.2, blue*1.2)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(pixelX, pixelY)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY)
	gl.Vertex2f(pixelX+cellSize-depth, pixelY+cellSize-depth)
	gl.Vertex2f(pixelX, pixelY+cellSize-depth)
	gl.End()
	
	// Draw edge highlights
	gl.Begin(gl.LINES)
	// Top-right edge
	gl.Vertex2f(pixelX+cellSize-depth, pixelY)
	gl.Vertex2f(pixelX+cellSize, pixelY+depth)
	// Bottom-right edge
	gl.Vertex2f(pixelX+cellSize-depth, pixelY+cellSize-depth)
	gl.Vertex2f(pixelX+cellSize, pixelY+cellSize)
	// Bottom-left edge
	gl.Vertex2f(pixelX, pixelY+cellSize-depth)
	gl.Vertex2f(pixelX+depth, pixelY+cellSize)
	gl.End()
	gl.LineWidth(1.0)
}

func (r *Renderer) drawBorder() {
	// Neon pink border with glow effect
	gl.LineWidth(3.0)
	gl.Color3f(1.0, 0.0, 0.8)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(boardOffsetX-5), float32(boardOffsetY-5))
	gl.Vertex2f(float32(boardOffsetX+boardWidth*cellSize+5), float32(boardOffsetY-5))
	gl.Vertex2f(float32(boardOffsetX+boardWidth*cellSize+5), float32(boardOffsetY+boardHeight*cellSize+5))
	gl.Vertex2f(float32(boardOffsetX-5), float32(boardOffsetY+boardHeight*cellSize+5))
	gl.End()
	
	// Inner glow
	gl.LineWidth(1.0)
	gl.Color3f(1.0, 0.3, 0.9)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(boardOffsetX-3), float32(boardOffsetY-3))
	gl.Vertex2f(float32(boardOffsetX+boardWidth*cellSize+3), float32(boardOffsetY-3))
	gl.Vertex2f(float32(boardOffsetX+boardWidth*cellSize+3), float32(boardOffsetY+boardHeight*cellSize+3))
	gl.Vertex2f(float32(boardOffsetX-3), float32(boardOffsetY+boardHeight*cellSize+3))
	gl.End()
}

func (r *Renderer) drawSynthwaveGrid() {
	// Draw horizontal grid lines with perspective
	gl.Color3f(0.5, 0.0, 0.8)
	gl.LineWidth(1.0)
	
	// Horizontal lines
	for y := 0; y < r.windowHeight; y += 40 {
		intensity := float32(y) / float32(r.windowHeight)
		gl.Color3f(0.3*intensity, 0.0, 0.5*intensity)
		gl.Begin(gl.LINES)
		gl.Vertex2f(0, float32(y))
		gl.Vertex2f(float32(r.windowWidth), float32(y))
		gl.End()
	}
	
	// Vertical lines with perspective effect
	centerX := float32(r.windowWidth) / 2
	for x := -20; x <= 20; x++ {
		xPos := centerX + float32(x*30)
		if xPos >= 0 && xPos <= float32(r.windowWidth) {
			gl.Begin(gl.LINES)
			gl.Color3f(0.2, 0.0, 0.4)
			gl.Vertex2f(xPos, 0)
			gl.Color3f(0.1, 0.0, 0.2)
			gl.Vertex2f(centerX + float32(x*15), float32(r.windowHeight))
			gl.End()
		}
	}
}

func (r *Renderer) drawInfoBox(x, y, width, height int, red, green, blue float32) {
	// Draw neon border
	gl.LineWidth(2.0)
	gl.Color3f(red, green, blue)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x+width), float32(y))
	gl.Vertex2f(float32(x+width), float32(y+height))
	gl.Vertex2f(float32(x), float32(y+height))
	gl.End()
	gl.LineWidth(1.0)
	
	// Draw background
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Color4f(red*0.2, green*0.2, blue*0.2, 0.3)
	gl.Begin(gl.QUADS)
	gl.Vertex2f(float32(x), float32(y))
	gl.Vertex2f(float32(x+width), float32(y))
	gl.Vertex2f(float32(x+width), float32(y+height))
	gl.Vertex2f(float32(x), float32(y+height))
	gl.End()
	gl.Disable(gl.BLEND)
}

func (r *Renderer) drawNumber(x, y int, number int, red, green, blue float32) {
	// Simple 7-segment style number rendering
	digits := fmt.Sprintf("%d", number)
	digitWidth := 15
	digitX := x
	
	for _, digit := range digits {
		r.drawDigit(digitX, y, digit, red, green, blue)
		digitX += digitWidth
	}
}

func (r *Renderer) drawDigit(x, y int, digit rune, red, green, blue float32) {
	gl.Color3f(red, green, blue)
	gl.LineWidth(2.0)
	
	// Simple digit patterns
	switch digit {
	case '0':
		gl.Begin(gl.LINE_LOOP)
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.End()
	case '1':
		gl.Begin(gl.LINES)
		gl.Vertex2f(float32(x+6), float32(y+2))
		gl.Vertex2f(float32(x+6), float32(y+18))
		gl.End()
	case '2':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.End()
	case '3':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.End()
		gl.Begin(gl.LINES)
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.End()
	case '4':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.End()
		gl.Begin(gl.LINES)
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.End()
	case '5':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.End()
	case '6':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.End()
	case '7':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.End()
	case '8':
		gl.Begin(gl.LINE_LOOP)
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.End()
		gl.Begin(gl.LINES)
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.End()
	case '9':
		gl.Begin(gl.LINE_STRIP)
		gl.Vertex2f(float32(x+2), float32(y+18))
		gl.Vertex2f(float32(x+10), float32(y+18))
		gl.Vertex2f(float32(x+10), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+2))
		gl.Vertex2f(float32(x+2), float32(y+10))
		gl.Vertex2f(float32(x+10), float32(y+10))
		gl.End()
	}
	gl.LineWidth(1.0)
}

func (r *Renderer) SetupProjection() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(r.windowWidth), float64(r.windowHeight), 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}