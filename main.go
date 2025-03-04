package main

type Item byte

const (
	quitKey = 'q'
)

func main() {

	keyPressed := make(chan Direction)
	quitCh := make(chan bool)

	keyHandler := NewGameHandler(keyPressed, quitCh)

	go keyHandler.HandleInputs()

	game := NewGame(20, 80, keyPressed, quitCh)

	go game.Run()

	<-quitCh

	keyHandler.RestoreTerm()
}
