package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type GameState struct {
	stateMap   map[Spot]Item
	blankSpots []Spot
}

type Level struct {
	score int
	speed time.Duration
}

type Food struct {
	item Item
	spot Spot
}

type Game struct {
	snake      Snake
	food       Food
	state      GameState
	board      [][]Item
	level      Level
	height     int
	width      int
	keyPressed chan Direction

	gameOver chan bool
}

func NewGame(height, width int, keyPressed chan Direction, gameOver chan bool) Game {

	board := make([][]Item, height)
	snake := createSnake(height, width)

	for h := 0; h < height; h++ {
		board[h] = make([]Item, width)
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			board[h][w] = GetBoardItem(h, w, height-1, width-1)
		}
	}

	game := Game{
		board:      board,
		height:     height,
		width:      width,
		snake:      snake,
		keyPressed: keyPressed,
		gameOver:   gameOver,
		level: Level{
			score: 1,
			speed: time.Second / 5,
		},
		state: GameState{
			stateMap: make(map[Spot]Item),
		},
	}

	game.UpdateStateMap()
	game.UpdateBlankSpots()

	return game
}

func (g *Game) Run() {

	g.Update()

	time.Sleep(time.Second)
	ticker := time.NewTicker(g.level.speed)

	g.SpawnFood()

	direction := down

	for {

		select {
		case <-ticker.C:
			g.MoveSnake(direction)
			g.Update()
		case direction = <-g.keyPressed:
			ticker.Stop()
			g.MoveSnake(direction)
			ticker.Reset(g.level.speed)
		case <-g.gameOver:
			ticker.Stop()
			return
		}
	}
}

func (g Game) ClearBoard() {
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			g.board[h][w] = GetBoardItem(h, w, g.height-1, g.width-1)
		}
	}
}

func (g *Game) MoveSnake(direction Direction) {

	prevState := g.snake.state

	g.snake.Move(direction)
	g.HandleCollision(prevState)
	g.UpdateStateMap()
	g.UpdateBlankSpots()
}

func (g *Game) HandleCollision(prevState SnakeState) {

	if g.snake.SelfColided(prevState) {
		g.gameOver <- true
	}

	if g.FoodEaten() {
		g.SpawnFood()
		g.snake.Grow()
	}

	if g.snake.ReachedBoundaries(g.height, g.width) {
		g.gameOver <- true
	}
}

func (g Game) FoodEaten() bool {

	headSpot := g.snake.state.spot
	foodSpot := g.food.spot

	return headSpot == foodSpot
}

func (g *Game) UpdateStateMap() {
	clear(g.state.stateMap)
	snakeTrack := &g.snake.state

	for snakeTrack != nil {
		g.state.stateMap[snakeTrack.spot] = snakeTrack.item
		snakeTrack = snakeTrack.next
	}
}

func (g *Game) UpdateBlankSpots() {

	i := 0
	blankSize := ((g.height - 4) * (g.width - 4)) - (len(g.state.stateMap) - 1)
	g.state.blankSpots = make([]Spot, blankSize)

	for h := 2; h < g.height-2; h++ {
		for w := 2; w < g.width-2; w++ {

			spot := Spot{x: w, y: h}
			_, ok := g.state.stateMap[spot]

			if ok {
				continue
			}

			g.state.blankSpots[i] = spot
			i++
		}
	}
}

func (g *Game) SpawnFood() {

	i := rand.Intn(len(g.state.blankSpots) - 2)

	spot := g.state.blankSpots[i]

	g.food = Food{
		item: food,
		spot: spot,
	}

}

func (g Game) Display() {
	buf := new(bytes.Buffer)
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			buf.WriteString(GetUnicode(g.board[h][w]))
		}

		buf.WriteString("\r\n")
	}
	fmt.Print("\033[H\033[2J")
	fmt.Println(buf.String())
}

func (g Game) Update() {

	g.ClearBoard()

	current := &g.snake.state

	for current != nil {
		y := current.spot.y
		x := current.spot.x

		g.board[y][x] = current.item

		current = current.next
	}

	g.board[g.food.spot.y][g.food.spot.x] = g.food.item

	g.Display()
}
