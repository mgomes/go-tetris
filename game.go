package main

import (
	"math/rand"
	"time"
)

type Game struct {
	Board        *Board
	CurrentPiece *Piece
	NextPiece    *Piece
	HeldPiece    *Piece
	CanHold      bool
	Score        int
	Lines        int
	Level        int
	GameOver     bool
	Paused       bool
	LastDrop     time.Time
	DropInterval time.Duration
	LastClear    int  // Track last clear for back-to-back
	WasTetris    bool // Track if last clear was a Tetris
	rng          *rand.Rand // Random number generator
}

func NewGame() *Game {
	// Create a new random number generator with current time as seed
	source := rand.NewSource(time.Now().UnixNano())
	
	g := &Game{
		Board:        NewBoard(),
		Score:        0,
		Lines:        0,
		Level:        1,
		GameOver:     false,
		Paused:       false,
		CanHold:      true,
		LastDrop:     time.Now(),
		DropInterval: time.Second,
		rng:          rand.New(source),
	}
	
	g.CurrentPiece = g.randomPiece()
	g.NextPiece = g.randomPiece()
	g.updateDropSpeed() // Set initial speed based on level 1
	
	return g
}

func (g *Game) randomPiece() *Piece {
	pieceType := g.rng.Intn(len(pieceShapes))
	return NewPiece(pieceType)
}

func (g *Game) Update() {
	if g.GameOver || g.Paused {
		return
	}
	
	if time.Since(g.LastDrop) >= g.DropInterval {
		g.MovePiece(0, 1)
		g.LastDrop = time.Now()
	}
}

func (g *Game) MovePiece(dx, dy int) bool {
	g.CurrentPiece.X += dx
	g.CurrentPiece.Y += dy
	
	if !g.Board.IsValidPosition(g.CurrentPiece) {
		g.CurrentPiece.X -= dx
		g.CurrentPiece.Y -= dy
		
		if dy > 0 {
			g.lockPiece()
		}
		return false
	}
	return true
}

func (g *Game) RotatePiece(clockwise bool) bool {
	originalShape := make([][]bool, len(g.CurrentPiece.Shape))
	for i := range originalShape {
		originalShape[i] = make([]bool, len(g.CurrentPiece.Shape[i]))
		copy(originalShape[i], g.CurrentPiece.Shape[i])
	}
	originalX := g.CurrentPiece.X
	
	g.CurrentPiece.Rotate(clockwise)
	
	// Try rotation with wall kicks
	kicks := g.getWallKicks(len(g.CurrentPiece.Shape))
	for _, kick := range kicks {
		g.CurrentPiece.X = originalX + kick[0]
		g.CurrentPiece.Y += kick[1]
		
		if g.Board.IsValidPosition(g.CurrentPiece) {
			return true
		}
		
		g.CurrentPiece.X = originalX
		g.CurrentPiece.Y -= kick[1]
	}
	
	// If no valid position found, revert
	g.CurrentPiece.Shape = originalShape
	g.CurrentPiece.X = originalX
	return false
}

func (g *Game) getWallKicks(pieceSize int) [][2]int {
	// Basic wall kick offsets (SRS-inspired)
	if pieceSize == 4 {
		// I-piece kicks
		return [][2]int{
			{0, 0},
			{-2, 0},
			{1, 0},
			{-2, -1},
			{1, 2},
		}
	} else if pieceSize == 2 {
		// O-piece doesn't need kicks
		return [][2]int{{0, 0}}
	} else {
		// Standard kicks for 3x3 pieces
		return [][2]int{
			{0, 0},
			{-1, 0},
			{1, 0},
			{0, -1},
			{-1, -1},
			{1, -1},
		}
	}
}

func (g *Game) HardDrop() {
	for g.MovePiece(0, 1) {
	}
}

func (g *Game) lockPiece() {
	g.Board.PlacePiece(g.CurrentPiece)
	
	linesCleared := g.Board.ClearLines()
	if linesCleared > 0 {
		g.Lines += linesCleared
		
		// Calculate score based on lines cleared
		baseScore := 0
		isPerfectClear := g.Board.IsPerfectClear()
		
		if isPerfectClear {
			// Perfect clear bonuses
			switch linesCleared {
			case 1:
				baseScore = scorePerfectSingle
			case 2:
				baseScore = scorePerfectDouble
			case 3:
				baseScore = scorePerfectTriple
			case 4:
				// Check for back-to-back Tetris
				if g.WasTetris {
					baseScore = scorePerfectTetrisB2B
				} else {
					baseScore = scorePerfectTetris
				}
			}
		} else {
			// Normal scoring
			switch linesCleared {
			case 1:
				baseScore = scoreSingle
			case 2:
				baseScore = scoreDouble
			case 3:
				baseScore = scoreTriple
			case 4:
				baseScore = scoreTetris
			}
		}
		
		g.Score += baseScore * g.Level
		
		// Track Tetris for back-to-back
		g.WasTetris = (linesCleared == 4)
		g.LastClear = linesCleared
		
		// Update level
		newLevel := 1 + g.Lines/linesPerLevel
		if newLevel > g.Level {
			g.Level = newLevel
			g.updateDropSpeed()
		}
	} else {
		// No lines cleared
		g.WasTetris = false
		g.LastClear = 0
	}
	
	g.CurrentPiece = g.NextPiece
	g.NextPiece = g.randomPiece()
	g.CanHold = true
	
	if !g.Board.IsValidPosition(g.CurrentPiece) {
		g.GameOver = true
	}
}

func (g *Game) HoldPiece() {
	if !g.CanHold {
		return
	}
	
	if g.HeldPiece == nil {
		g.HeldPiece = g.CurrentPiece
		g.CurrentPiece = g.NextPiece
		g.NextPiece = g.randomPiece()
	} else {
		g.CurrentPiece, g.HeldPiece = g.HeldPiece, g.CurrentPiece
	}
	
	// Reset position for the new current piece
	g.CurrentPiece.X = 3
	g.CurrentPiece.Y = 0
	g.CurrentPiece.Rotation = 0
	
	g.CanHold = false
}

func (g *Game) updateDropSpeed() {
	g.DropInterval = calculateDropInterval(g.Level)
}