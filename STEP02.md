# Loading the Maze

## Reading from a file

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
        line := scanner.Text()
        maze = append(maze, line)
    }

    return maze, nil
}
```

```go
func main() {
	m, err := loadMaze("maze.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range m {
		fmt.Println(line)
	}
}
```

Discussion: functions, idiomatic error handling, range loop.

## Next step

Proceed to [step 3](STEP3.md).
