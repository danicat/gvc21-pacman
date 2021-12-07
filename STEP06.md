# Who ya gonna call?

## Adding ghosts 👻

We are also going to track ghosts with a global variable, but this time we need a slice of pointers to sprite.

```go
var ghosts []*Sprite
```

Update `loadMaze` to initialize this variable:

```go
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
```

We also need to update the `draw` routine:

```go
for _, g := range ghosts {
    simpleansi.MoveCursor(g.row, g.col)
    fmt.Print("G")
}
```

## Moving the ghosts

In `sprite.go`, add the following function:

```go
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
```

## Updating the game loop

Finally, modify the game loop to update the ghosts position every time:

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

		for _, g := range ghosts {
			g.RandomMove()
		}
	}
```

Try building and running. Ghosts should move every time you move now! Scary >.<

## Next step

Proceed to [step 7](STEP7.md).
