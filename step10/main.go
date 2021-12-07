package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/danicat/simpleansi"
)

var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	mazeFile   = flag.String("maze-file", "maze.txt", "path to a custom maze file")
)

var maze []string
var player Sprite
var ghosts []*Sprite

var score int
var numDots int
var lives = 1

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
			case '.':
				numDots++
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
	flag.Parse()

	err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	m, err := loadMaze(*mazeFile)
	if err != nil {
		log.Fatal(err)
	}
	maze = m

	prepareTerminal()
	defer restoreTerminal()

	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := readInput()
			if err != nil {
				log.Println("error reading input:", err)
				ch <- "ESC"
			}
			ch <- input
		}
	}(input)

	for {
		draw(true)

		select {
		case key := <-input:
			if key == "QUIT" {
				lives = 0
			}
			player.Move(key)
		default:
		}

		if lives <= 0 || numDots == 0 {
			break
		}

		for _, g := range ghosts {
			g.RandomMove()
		}

		died := processCollisions()
		if died {
			lives--
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func processCollisions() bool {
	if maze[player.row][player.col] == '.' {
		numDots--
		score++
		// Remove dot from the maze
		maze[player.row] = maze[player.row][0:player.col] + " " + maze[player.row][player.col+1:]
	}

	for _, g := range ghosts {
		if player == *g {
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
