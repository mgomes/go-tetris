package main

// Board dimensions are now in constants.go

type Board struct {
	Grid  [boardHeight][boardWidth]bool
	Colors [boardHeight][boardWidth][3]float32
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) IsValidPosition(piece *Piece) bool {
	blocks := piece.GetBlocks()
	for _, block := range blocks {
		x, y := block[0], block[1]
		
		if x < 0 || x >= boardWidth || y >= boardHeight {
			return false
		}
		
		if y >= 0 && b.Grid[y][x] {
			return false
		}
	}
	return true
}

func (b *Board) PlacePiece(piece *Piece) {
	blocks := piece.GetBlocks()
	for _, block := range blocks {
		x, y := block[0], block[1]
		if y >= 0 && y < boardHeight && x >= 0 && x < boardWidth {
			b.Grid[y][x] = true
			b.Colors[y][x] = piece.Color
		}
	}
}

func (b *Board) ClearLines() int {
	linesCleared := 0
	
	for y := boardHeight - 1; y >= 0; y-- {
		if b.isLineFull(y) {
			b.removeLine(y)
			linesCleared++
			y++
		}
	}
	
	return linesCleared
}

func (b *Board) IsPerfectClear() bool {
	for y := range boardHeight {
		for x := range boardWidth {
			if b.Grid[y][x] {
				return false
			}
		}
	}
	return true
}

func (b *Board) isLineFull(y int) bool {
	for x := range boardWidth {
		if !b.Grid[y][x] {
			return false
		}
	}
	return true
}

func (b *Board) removeLine(line int) {
	for y := line; y > 0; y-- {
		b.Grid[y] = b.Grid[y-1]
		b.Colors[y] = b.Colors[y-1]
	}
	
	b.Grid[0] = [boardWidth]bool{}
	b.Colors[0] = [boardWidth][3]float32{}
}