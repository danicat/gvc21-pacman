# Run, Pac Man, Run

## Creating sprites

Create a file called `sprite.go` and add the following code:

```go
package main

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
```

## Tracking the player position

Now in `main.go` create a global variable player of type Sprite:

```go
var player Sprite
```

And add the following code to `loadMaze` so we can find its initial position:

```go
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
			}
		}
	}

	return maze, nil
}
```

## Reading the arrow keys

Now let's adapt `readInput` to also process arrow keys:

```go
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
```

## Updating the player's position

Finally: add `player.Move` to the game loop:

```go
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
	}
```

Try building the code and walking around the maze with the arrow keys! Press ESC to quit.

## Extra points

How would you test Sprite.Move()?

## Next step

Proceed to [step 6](STEP6.md).
