# Game Over!

## Tracking the score

First, a few variables to track our score and lives.

```go
var score int
var numDots int
var lives = 1
```

Now, let's update `loadMaze` to count the dots:

```go
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
```

And `draw` to print them correctly... we are also adding printing the score so we can have some visual feedback.

```go
for _, line := range maze {
    for _, chr := range line {
        switch chr {
        case '#':
            fallthrough
        case '.':
            fmt.Printf("%c", chr)
        default:
            fmt.Print(" ")
        }
    }
    fmt.Println()
}

// rest of the function ommitted for brevity

simpleansi.MoveCursor(len(maze)+1, 0)
fmt.Println("Score:", score, "\tLives:", lives)
```

## Game win condition

In order to win we need to eat all the dots. So we will decrement `numDots` every time we step in one, and remove them from the board.

Create the function `processCollisions` as defined below:

```go
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
```

Change the exit condition of the main loop to:

```go
if input == "QUIT" || lives <= 0 || numDots == 0 {
    break
}
```

And add a call to `processCollisions` after player and ghosts move:

```go
for {
    draw(true)
    input, err := readInput()
    if err != nil {
        fmt.Printf("game loop: %v", err)
        break
    }

    if input == "QUIT" || lives <= 0 || numDots == 0 {
        break
    }

    player.Move(input)

    for _, g := range ghosts {
        g.RandomMove()
    }

    processCollisions()
}
```

## Game over condition

Modify the game loop to decrement the lives counter if a collision is detected.

```go
for {
    draw(true)
    input, err := readInput()
    if err != nil {
        fmt.Printf("game loop: %v", err)
        break
    }

    if input == "QUIT" || lives <= 0 || numDots == 0 {
        break
    }

    player.Move(input)

    for _, g := range ghosts {
        g.RandomMove()
    }

    died := processCollisions()
    if died {
        lives--
    }
}
```

Bonus points: implement multiple lives.

## Next step

Proceed to [step 8](STEP8.md).
