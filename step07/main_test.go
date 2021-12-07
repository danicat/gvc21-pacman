package main

import "testing"

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

func TestLoadMaze(t *testing.T) {
	_, err := loadMaze("maze.txt")
	if err != nil {
		t.Fatalf("failed to load maze: %v", err)
	}
}
