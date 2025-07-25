package main

// Color represents RGB color values
type Color [3]float32

// Point represents a 2D coordinate
type Point struct {
	X, Y int
}

// Block represents a single block position
type Block [2]int

// PieceType represents the different tetromino types
type PieceType int

const (
	PieceI PieceType = iota
	PieceO
	PieceT
	PieceS
	PieceZ
	PieceJ
	PieceL
)

// GameState represents the current state of the game
type GameState int

const (
	StateActive GameState = iota
	StatePaused
	StateGameOver
)