package main

import "time"

// Window configuration
const (
	windowWidth  = 600
	windowHeight = 800
	windowTitle  = "Go Tetris"
)

// Board dimensions
const (
	boardWidth  = 10
	boardHeight = 20
)

// Rendering constants
const (
	cellSize     = 30
	boardOffsetX = 50
	boardOffsetY = 50
	depthOffset  = 4
)

// UI layout constants
const (
	holdBoxX      = boardOffsetX + boardWidth*cellSize + 50
	holdBoxY      = boardOffsetY + 50
	nextBoxX      = boardOffsetX + boardWidth*cellSize + 50
	nextBoxY      = boardOffsetY + 200
	scoreBoxX     = boardOffsetX + boardWidth*cellSize + 50
	scoreBoxY     = boardOffsetY + 350
	levelBoxX     = boardOffsetX + boardWidth*cellSize + 50
	levelBoxY     = boardOffsetY + 430
	infoBoxWidth  = 100
	infoBoxHeight = 60
	miniBlockSize = 20
)

// Game timing constants
const (
	frameTargetTime        = 16 * time.Millisecond
	keyJustPressedWindow   = 50 * time.Millisecond
	keyRepeatDelay         = 400 * time.Millisecond
	keyRepeatInterval      = 200 * time.Millisecond
	keyRepeatFastInterval  = 50 * time.Millisecond
	keyFastThreshold       = 700 * time.Millisecond
)

// Scoring constants
const (
	linesPerLevel = 10
	
	// Normal line clear scores
	scoreSingle = 100
	scoreDouble = 300
	scoreTriple = 500
	scoreTetris = 800
	
	// Perfect clear scores
	scorePerfectSingle       = 800
	scorePerfectDouble       = 1200
	scorePerfectTriple       = 1800
	scorePerfectTetris       = 2000
	scorePerfectTetrisB2B    = 3200
)

// Rendering style constants
const (
	lineWidthThin   = 1.0
	lineWidthMedium = 2.0
	lineWidthThick  = 3.0
	
	// Grid spacing
	gridHorizontalSpacing = 40
	gridVerticalSpacing   = 30
	
	// Text rendering
	digitWidth    = 15
	digitHeight   = 20
	letterSpacing = 35
	letterSize    = 25
)

// Color intensity multipliers
const (
	colorDimFactor    = 0.3
	colorMediumFactor = 0.5
	colorBrightFactor = 1.2
	colorGlowFactor   = 0.8
)