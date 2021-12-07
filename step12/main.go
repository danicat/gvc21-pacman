package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/danicat/simpleansi"
)

var maze []string
var player *Player
var ghosts []*Sprite

var score int
var numDots int
var lives = 1

func loadMaze(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var maze []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		maze = append(maze, text)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = NewPlayer(row, col)
			case 'G':
				ghosts = append(ghosts, NewGhost(row, col))
			case '.':
				numDots++
			}
		}
	}

	return maze, nil
}

func draw(cls bool) {
	if cls {
		simpleansi.ClearScreen()
	}

	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Print(cfg.Dot)
			case 'X':
				fmt.Print(cfg.Pill)
			default:
				fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}

	moveCursor(player.row, player.col)
	fmt.Print(cfg.Player)

	for _, g := range ghosts {
		moveCursor(g.row, g.col)
		fmt.Print(cfg.Ghost)
	}

	moveCursor(len(maze)+1, 0)
	fmt.Println("Score:", score, "\tLives:", lives)
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "QUIT", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

func prepareTerminal() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("unable to activate cbreak mode: %v", err)
	}
}

func restoreTerminal() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("unable to restore cooked mode: %v", err)
	}
}

func main() {
	m, err := loadMaze("maze.txt")
	if err != nil {
		log.Fatal(err)
	}
	maze = m

	prepareTerminal()
	defer restoreTerminal()

	for {
		draw(true)

		if quit := player.Update(); quit {
			break
		}

		if lives <= 0 || numDots == 0 {
			// game over
			if lives <= 0 {
				moveCursor(player.row, player.col)
				fmt.Print(cfg.Death)
				moveCursor(len(maze)+1, 0)
			}
			break
		}

		for _, g := range ghosts {
			g.Update()
		}

		died := processCollisions()
		if died {
			lives--
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func processCollisions() bool {
	remove := func(row, col int) {
		maze[row] = maze[row][:col] + " " + maze[row][col+1:]
	}

	switch maze[player.row][player.col] {
	case '.':
		numDots--
		score++
		remove(player.row, player.col)

	case 'X':
		score += 10
		remove(player.row, player.col)
	}

	for _, g := range ghosts {
		if player.row == g.row && player.col == g.col {
			return true
		}
	}

	return false
}

func moveCursor(row, col int) {
	if cfg.UseEmoji {
		simpleansi.MoveCursor(row, col*2)
	} else {
		simpleansi.MoveCursor(row, col)
	}
}
