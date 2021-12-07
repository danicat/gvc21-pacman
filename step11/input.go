package main

import (
	"log"
	"math/rand"
)

type InputHandler interface {
	Read() string
}

type KBHandler struct {
	input chan string
}

func NewKBHandler() *KBHandler {
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

	return &KBHandler{input: input}
}

func (k *KBHandler) Read() string {
	var key string
	select {
	case key = <-k.input:
	default:
	}
	return key
}

type RandomHandler struct {
}

func NewRandomHandler() *RandomHandler {
	return &RandomHandler{}
}

func (r *RandomHandler) Read() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}
