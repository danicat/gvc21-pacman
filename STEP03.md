# Examples and testing

## Isolate the draw function

```go
var maze []string

func draw() {
	for _, line := range maze {
		fmt.Println(line)
	}
}
```

```go
func main() {
	m, err := loadMaze("maze.txt")
	if err != nil {
		log.Fatal(err)
	}
	maze = m

	for {
		draw()
        break
	}
}
```

## Testing the draw function

Create a file called `main_test.go`

```go
package main

func ExampleDraw() {
	maze = []string{
		"####",
		"#  #",
		"####",
	}

	draw()
	//output:
	//####
	//#  #
	//####
}
```

Run `go test` to run the tests. 

Try also `go test -v` and `go test -cover`.

## Testing load maze

Add a new test to `main_test.go`:

```go
func TestLoadMaze(t *testing.T) {
	_, err := loadMaze("maze.txt")
	if err != nil {
		t.Fatalf("failed to load maze: %v", err)
	}
}
```

Noticed any changes in `go test -v` and `go test -cover`?

## Using external packages

```go
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)
```

```go
func draw(cls bool) {
	if cls {
		simpleansi.ClearScreen()
	}

	for _, line := range maze {
		fmt.Println(line)
	}
}
```

You will need to download the package either by:
- `go get github.com/danicat/simpleansi`
- `go mod tidy`

Update the test code to call draw with `false`, and the code in `main` to call it with `true`:

```go
func ExampleDraw() {
	maze = []string{
		"####",
		"#  #",
		"####",
	}

	draw(false)
	//output:
	//####
	//#  #
	//####
}
```
We need to do this because clearing the screen will break the expected output.

## Next step

Proceed to [step 4](STEP4.md).
