package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func main() {
	m, err := loadMaze("maze.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range m {
		fmt.Println(line)
	}
}
