package main

const (
	blank            Item = 0
	upperLeftCorner  Item = 1
	lowerLeftCorner  Item = 2
	upperRightCorner Item = 3
	lowerRightCorner Item = 4
	upperBlock       Item = 5
	lowerBlock       Item = 6
	leftBlock        Item = 7
	rightBlock       Item = 8

	snake_vertical Item = 10

	snake_horizontal_upper Item = 11
	snake_horizontal_lower Item = 12

	food = 13
)

func GetUnicode(value Item) string {

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
	case snake_vertical:
		return "█"
	case snake_horizontal_upper:
		return "▀"
	case snake_horizontal_lower:
		return "▄"
	case food:
		return "⬤"
	default:
		return " "
	}
}

func GetBoardItem(h, w, maxHeight, maxWidth int) Item {

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
