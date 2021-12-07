package main

import "math/rand"

type Sprite struct {
	row, col int
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

func (s *Sprite) RandomMove() {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}

	s.Move(move[dir])
}
