package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

const (
	up    Direction = 1
	down  Direction = 2
	left  Direction = 3
	right Direction = 4
)

type Direction byte

type GameHandler struct {
	keyPressed  chan Direction
	quitChannel chan bool

	termState *term.State
}

func NewGameHandler(keyCh chan Direction, quitCh chan bool) GameHandler {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("Error setting terminal to raw mode: %v\n", err)
		return GameHandler{}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()

	return GameHandler{
		keyPressed:  keyCh,
		quitChannel: quitCh,
		termState:   oldState,
	}
}

func (k *GameHandler) HandleInputs() {
	buf := make([]byte, 3)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			break
		}

		if buf[0] == 'q' {
			k.quitChannel <- true
		}

		if buf[0] == '\x1b' && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				k.keyPressed <- up
			case 'B':
				k.keyPressed <- down
			case 'C':
				k.keyPressed <- left
			case 'D':
				k.keyPressed <- right
			}
		}
	}
}

func (k GameHandler) RestoreTerm() {
	term.Restore(int(os.Stdin.Fd()), k.termState)
}
