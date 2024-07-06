package main

type Item byte

const (
	quitKey = 'q'
)

func main() {

	keyPressed := make(chan Direction)
	quitCh := make(chan bool)

	keyHandler := KeyHandler{
		keyPressed:  keyPressed,
		quitChannel: quitCh,
	}

	go keyHandler.Handle()

	game := NewGame(20, 80, keyPressed, quitCh)

	go game.Run()

	<-quitCh
}
