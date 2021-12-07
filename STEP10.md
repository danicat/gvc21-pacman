# Fun with flags

## Create flag variables

Add as global variables:

```go
var (
    configFile = flag.String("config-file", "config.json", "path to custom configuration file")
    mazeFile   = flag.String("maze-file", "maze.txt", "path to a custom maze file")
)
```

Add to the first line of your `main` function:

```go
func main() {
    flag.Parse()

    // ...
}
```

And then replace the hardcoded strings `maze.txt` and `config.json`. What happens with the code running in the `init` function? Hint: you might want to move it to `main`.

## Next step

Proceed to [step 11](STEP11.md).
