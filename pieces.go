package main

type Piece struct {
	Shape      [][]bool
	Color      [3]float32
	X, Y       int
	Rotation   int
}

var pieceShapes = [][][]bool{
	// I-piece
	{
		{false, false, false, false},
		{true, true, true, true},
		{false, false, false, false},
		{false, false, false, false},
	},
	// O-piece
	{
		{true, true},
		{true, true},
	},
	// T-piece
	{
		{false, true, false},
		{true, true, true},
		{false, false, false},
	},
	// S-piece
	{
		{false, true, true},
		{true, true, false},
		{false, false, false},
	},
	// Z-piece
	{
		{true, true, false},
		{false, true, true},
		{false, false, false},
	},
	// J-piece
	{
		{true, false, false},
		{true, true, true},
		{false, false, false},
	},
	// L-piece
	{
		{false, false, true},
		{true, true, true},
		{false, false, false},
	},
}

var pieceColors = [][3]float32{
	{0.0, 0.9, 1.0},   // I - Neon Cyan
	{1.0, 0.0, 0.5},   // O - Hot Pink
	{0.5, 0.0, 1.0},   // T - Electric Purple
	{0.0, 1.0, 0.5},   // S - Neon Green
	{1.0, 0.0, 0.8},   // Z - Magenta
	{0.2, 0.5, 1.0},   // J - Electric Blue
	{1.0, 0.3, 0.7},   // L - Sunset Pink
}

func NewPiece(pieceType int) *Piece {
	shape := make([][]bool, len(pieceShapes[pieceType]))
	for i := range shape {
		shape[i] = make([]bool, len(pieceShapes[pieceType][i]))
		copy(shape[i], pieceShapes[pieceType][i])
	}
	
	return &Piece{
		Shape:    shape,
		Color:    pieceColors[pieceType],
		X:        3,
		Y:        0,
		Rotation: 0,
	}
}

func (p *Piece) Rotate(clockwise bool) {
	n := len(p.Shape)
	rotated := make([][]bool, n)
	for i := range rotated {
		rotated[i] = make([]bool, n)
	}
	
	if clockwise {
		for i := range n {
			for j := range n {
				rotated[j][n-1-i] = p.Shape[i][j]
			}
		}
	} else {
		for i := range n {
			for j := range n {
				rotated[n-1-j][i] = p.Shape[i][j]
			}
		}
	}
	
	p.Shape = rotated
}

func (p *Piece) GetBlocks() [][2]int {
	var blocks [][2]int
	for y, row := range p.Shape {
		for x, filled := range row {
			if filled {
				blocks = append(blocks, [2]int{p.X + x, p.Y + y})
			}
		}
	}
	return blocks
}