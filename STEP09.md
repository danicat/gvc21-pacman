# Finally, Emojis!

## Handling configurations

Create a file called `config.go` and add the following code:

```go
package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config holds the emoji configuration
type Config struct {
	Player   string `json:"player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
	UseEmoji bool   `json:"use_emoji"`
}

var cfg Config

func loadConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
}
```

## Correct coordinates when using emojis

Create a new function called `moveCursor`:

```go
func moveCursor(row, col int) {
    if cfg.UseEmoji {
        simpleansi.MoveCursor(row, col*2)
    } else {
        simpleansi.MoveCursor(row, col)
    }
}
```

We need to replace any current calls to `simpleansi.MoveCursor` with this function, to take into account that emojis are two bytes wide and will misalign the board when drawn alongside characters 1 byte wide.

## Replace hard coded chars with config values

```go
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
```

## Game over and pills

Add game over emoji to the game over condition on main loop:

```go
if numDots == 0 || lives <= 0 {
    if lives <= 0 {
        moveCursor(player.row, player.col)
        fmt.Print(cfg.Death)
        moveCursor(len(maze)+2, 0)
    }
    break
}
```

And give pills a score, updating `processCollisions`:

```go
func processCollisions() bool {
    remove := func(row, col int) {
        maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
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
		if player == *g {
			return true
		}
	}

	return false
}

```

## Next step

Proceed to [step 10](STEP10.md).
