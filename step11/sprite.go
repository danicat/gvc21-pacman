package main

type Sprite struct {
	row, col int
	ih       InputHandler
}

func NewPlayer(row, col int) *Sprite {
	return &Sprite{row: row, col: col, ih: NewKBHandler()}
}

func NewGhost(row, col int) *Sprite {
	return &Sprite{row: row, col: col, ih: NewRandomHandler()}
}

func (s *Sprite) Update() bool {
	key := s.ih.Read()
	if key == "QUIT" {
		return true
	}

	s.Move(key)

	return false
}

func (s *Sprite) Move(direction string) {
	newRow, newCol := s.row, s.col

	switch direction {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] != '#' {
		s.row = newRow
		s.col = newCol
	}
}
