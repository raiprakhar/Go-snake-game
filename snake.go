package main

type SnakePart struct {
	X int
	Y int
}

type SnakeBody struct {
	Parts  []SnakePart
	Xspeed int
	Yspeed int
}

func (snk *SnakeBody) UpdateDirection(verticalSpeed int, horizontalSpeed int) {
	snk.Yspeed = verticalSpeed
	snk.Xspeed = horizontalSpeed
}

func (snk *SnakeBody) UpdateLocation(width int, height int, longerSnake bool) {
	snk.Parts = append(snk.Parts, snk.Parts[len(snk.Parts)-1].GetUpdatedPart(snk, width, height))
	if !longerSnake {
		snk.Parts = snk.Parts[1:]
	}
}

func (sb *SnakeBody) ResetPos(width int, height int) {
	snakeParts := []SnakePart{
		{
			X: int(width / 2),
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 1,
			Y: int(height / 2),
		},
		{
			X: int(width/2) + 2,
			Y: int(height / 2),
		},
	}
	sb.Parts = snakeParts
	sb.Xspeed = 1
	sb.Yspeed = 0
}

func (sp *SnakePart) GetUpdatedPart(sb *SnakeBody, width int, height int) SnakePart {
	newPart := *sp
	newPart.X = (newPart.X + sb.Xspeed) % width
	if newPart.X < 0 {
		newPart.X += width
	}

	newPart.Y = (newPart.Y + sb.Yspeed) % height
	if newPart.Y < 0 {
		newPart.Y += height
	}

	return newPart
}
