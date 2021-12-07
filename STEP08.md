# Making things realtime!

## Refactoring the input code

Add the following snippet before the game loop.

```go
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
```

Replace the input handling code in the game loop with this:

```go
select {
case key := <-input:
    if key == "QUIT" {
        lives = 0
    }
    player.Move(key)
default:
}
```

And add some delay in the main loop to make the game playable:

```go
if numDots == 0 || lives <= 0 {
    break
}

time.Sleep(200 * time.Millisecond)
```

## Next step

Proceed to [step 9](STEP9.md).
