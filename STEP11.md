# Interfaces

## InputHandler

Create a file called `input.go` with the following code:

```go
package main

import (
	"log"
	"math/rand"
)

type InputHandler interface {
	Read() string
}

type KBHandler struct {
	input chan string
}

func NewKBHandler() *KBHandler {
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

	return &KBHandler{input: input}
}

func (k *KBHandler) Read() string {
	var key string
	select {
	case key = <-k.input:
	default:
	}
	return key
}

type RandomHandler struct {
}

func NewRandomHandler() *RandomHandler {
	return &RandomHandler{}
}

func (r *RandomHandler) Read() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}
```

Now update `sprite.go`:

```go
package main

import "math/rand"

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
```

Now let's remove all input handling code from `main` and just call the `Update` method for player and ghosts:

```go
func main() {
	err := loadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	m, err := loadMaze(*mazeFile)
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
```

Also make sure you are using the "constructors" when creating those objects:

```go
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
```

You will also need to update the global variable `player` to the type `*Sprite`.

## Next step

Proceed to [step 11](STEP12.md).
