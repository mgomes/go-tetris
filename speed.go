package main

import "time"

// speedCurve contains the Tetris Worlds speed curve G values (rows per frame)
var speedCurve = []float64{
	0.01667,  // Level 1
	0.021017, // Level 2
	0.026977, // Level 3
	0.035256, // Level 4
	0.04693,  // Level 5
	0.06361,  // Level 6
	0.0879,   // Level 7
	0.1236,   // Level 8
	0.1775,   // Level 9
	0.2598,   // Level 10
	0.388,    // Level 11
	0.59,     // Level 12
	0.92,     // Level 13
	1.46,     // Level 14
	2.36,     // Level 15
	3.91,     // Level 16
	6.61,     // Level 17
	11.43,    // Level 18
	20.23,    // Level 19
	36.6,     // Level 20+
}

const (
	framesPerSecond = 60
	frameTime       = 16.67 // milliseconds per frame at 60 FPS
)

// calculateDropInterval converts G value (rows per frame) to drop interval
func calculateDropInterval(level int) time.Duration {
	// Get G value for current level
	gValue := speedCurve[len(speedCurve)-1] // Default to max speed
	if level > 0 && level <= len(speedCurve) {
		gValue = speedCurve[level-1]
	}
	
	// Convert G (rows per frame) to milliseconds per row
	if gValue > 0 {
		framesPerRow := 1.0 / gValue
		millisecondsPerRow := framesPerRow * frameTime
		return time.Duration(millisecondsPerRow) * time.Millisecond
	}
	
	return time.Second // Fallback
}