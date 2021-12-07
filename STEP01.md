# Step 01

## Hello world

Save the following file as `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Gophercon!")
}
```

Try running it with `go run main.go`.

Discussion: file structure, package names, imports, gofmt, visibility.

## Setting up go modules

Try running `go build`.

In older versions of Go this would work fine, but in the most recent releses it will complain about the lack of a `go.mod` file. Go modules is how we track the dependencies of our code. Initializing it is simple, just run `go mod init` with the path to your source code in version control:

```sh
$ go mod init your-vcs/your-user/your-repo
```

For example, if you are using GitHub:

```sh
$ go mod init github.com/danicat/gvc2021-pacman
```

Now try `go build` again.

## Create the game loop

Discussion: how variables work in go, how to use `for` loops, initializers and scope.

Task: rewrite the `main.go` file so that it prints `I'm a game loop` five times.

```go
for i := 0; i < 5; i++ {
    fmt.Println("I'm a game loop")
}
```

## Next step

Proceed to [step 2](STEP2.md).
