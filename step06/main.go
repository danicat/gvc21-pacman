package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

var maze []string
var player Sprite
var ghosts []*Sprite

func draw(cls bool) {
	if cls {
		simpleansi.ClearScreen()
	}

	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Print("#")
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	simpleansi.MoveCursor(player.row, player.col)
	fmt.Print("P")

	for _, g := range ghosts {
		simpleansi.MoveCursor(g.row, g.col)
		fmt.Print("G")
	}

	simpleansi.MoveCursor(len(maze)+1, 0)
}

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
				player = Sprite{row, col}
			case 'G':
				ghosts = append(ghosts, &Sprite{row, col})
			}
		}
	}

	return maze, nil
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
		input, err := readInput()
		if err != nil {
			fmt.Printf("game loop: %v", err)
			break
		}
		if input == "QUIT" {
			break
		}
		player.Move(input)

		for _, g := range ghosts {
			g.RandomMove()
		}
	}
}
