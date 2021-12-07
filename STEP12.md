# Type embedding

## Creating the Player type

In `sprites.go` create a Player type, and change the NewPlayer function accordingly:

```go
type Player struct {
	lives, score int
	Sprite
}

func NewPlayer(row, col int) *Player {
	return &Player{Sprite: Sprite{row: row, col: col, ih: NewKBHandler()}}
}
```

Also, update the global variable `player` to be of type `*Player`.