package main

import "fmt"

const (
	move SpotRequest = 1
	grow SpotRequest = 2
)

type SpotRequest byte

type Spot struct {
	x, y int
}

type SnakeState struct {
	spot      Spot
	direction Direction
	item      Item
	next      *SnakeState
	previous  *SnakeState
}

type Snake struct {
	state SnakeState
}

func createSnake(height, width int) Snake {
	headX := width / 2
	headY := height / 2
	headState := SnakeState{
		direction: down,
		spot: Spot{
			x: headX,
			y: headY,
		},
	}
	return Snake{
		state: headState,
	}
}

func (s *Snake) Move(direction Direction) {

	current := &s.state
	newDir := direction //down

	if !current.direction.IsMoveAllowed(direction) {
		newDir = current.direction
	}

	for current != nil {

		//right
		tempDir := current.direction

		current.spot = current.spot.Next(move, newDir)
		current.direction = newDir
		current.item = current.Item()

		current = current.next

		newDir = tempDir
	}
}

func (s *Snake) Grow() {

	currentHead := s.state
	newState := SnakeState{
		direction: s.state.direction,
		spot:      s.state.spot.Next(grow, s.state.direction),
	}

	s.state = newState

	currentHead.previous = &s.state
	s.state.next = &currentHead
}

func (s Snake) ReachedBoundaries(height, width int) bool {

	current := &s.state

	for current != nil {

		if current.spot.x == 0 || current.spot.x == width {
			return true
		}

		if current.spot.y == 0 || current.spot.y == height {
			return true
		}

		current = current.next
	}

	return false
}

func (s Snake) SelfColided(prevState SnakeState) bool {

	newHeadSpot := s.state.spot
	current := &prevState

	for current != nil {

		if newHeadSpot == current.spot {
			return true
		}

		current = current.next
	}

	return false
}

func (s SnakeState) Item() Item {

	if s.direction == left || s.direction == right {

		if s.previous != nil && s.previous.direction == down {
			return snake_horizontal_lower
		} else if s.previous != nil && s.previous.direction == up {
			return snake_horizontal_upper
		} else {
			return snake_horizontal_upper
		}
	}

	return snake_vertical
}

func (s SnakeState) String() string {
	fullState := "["
	current := &s
	for current != nil {
		fullState += fmt.Sprintf("dir: %d, spot: %s,", current.direction, current.spot)
		current = current.next
	}

	fullState += "]"
	return fullState
}

func (s Spot) Next(request SpotRequest, direction Direction) Spot {

	next := &Spot{
		x: s.x,
		y: s.y,
	}

	if direction == down {
		next.y++
	}

	if direction == up {
		next.y--
	}

	if direction == left {
		next.x++
	}

	if direction == right {
		next.x--
	}

	return *next
}

func (d Direction) IsMoveAllowed(other Direction) bool {

	if d == other {
		return true
	}

	if d == up || d == down {
		return other == left || other == right
	}

	return other == up || other == down
}

func (s Spot) String() string {
	return fmt.Sprintf("x: %d, y: %d", s.x, s.y)
}
