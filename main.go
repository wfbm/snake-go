package main

import (
	"bytes"
	"fmt"
	"time"
)

const (
	blank            = 0
	upperLeftCorner  = 1
	lowerLeftCorner  = 2
	upperRightCorner = 3
	lowerRightCorner = 4
	upperBlock       = 5
	lowerBlock       = 6
	leftBlock        = 7
	rightBlock       = 8

	player = 10
)

type Item byte

type Spot struct {
	x, y int
}

type SnakePosition struct {
	spot     Spot
	next     *SnakePosition
	previous *SnakePosition
}

type Snake struct {
	position SnakePosition
}

type Level struct {
	score int
	speed time.Duration
}

type GameState struct {
	stateMap   map[Spot]Item
	blankSpots []Spot
}

type Game struct {
	snake  Snake
	state  GameState
	board  [][]byte
	height int
	width  int
}

func main() {
	game := NewGame(20, 80)
	game.Update()
}

func NewGame(height, width int) Game {

	board := make([][]byte, height)
	snake := Snake{
		position: SnakePosition{
			spot: Spot{
				x: width / 2,
				y: height / 2,
			},
		},
	}

	for h := 0; h < height; h++ {
		board[h] = make([]byte, width)
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			board[h][w] = GetBoardItem(h, w, height-1, width-1)
		}
	}

	return Game{
		board:  board,
		height: height,
		width:  width,
		snake:  snake,
	}
}

func (g Game) Display() {
	buf := new(bytes.Buffer)
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			buf.WriteString(GetUnicode(g.board[h][w]))
		}

		buf.WriteString("\n")
	}

	fmt.Println(buf.String())
}

func (g Game) Update() {

	clearScreen()

	playerY := g.snake.position.spot.y
	playerX := g.snake.position.spot.x

	g.board[playerY][playerX] = player

	g.Display()
}

func GetUnicode(value byte) string {

	switch value {
	case upperLeftCorner:
		return "▛"
	case lowerLeftCorner:
		return "▙"
	case upperRightCorner:
		return "▜"
	case lowerRightCorner:
		return "▟"
	case upperBlock:
		return "▀"
	case lowerBlock:
		return "▄"
	case leftBlock:
		return "▌"
	case rightBlock:
		return "▐"
	case player:
		return "■"
	default:
		return " "
	}
}

func GetBoardItem(h, w, maxHeight, maxWidth int) byte {

	if h == 0 && w == 0 {
		return upperLeftCorner
	} else if h == maxHeight && w == 0 {
		return lowerLeftCorner
	} else if h == 0 && w == maxWidth {
		return upperRightCorner
	} else if h == maxHeight && w == maxWidth {
		return lowerRightCorner
	} else if h == 0 {
		return upperBlock
	} else if h == maxHeight {
		return lowerBlock
	} else if w == 0 {
		return leftBlock
	} else if w == maxWidth {
		return rightBlock
	}

	return blank
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
