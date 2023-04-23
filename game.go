package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	SnakeBody SnakeBody
	FoodPos   SnakePart
	Score     int
	GameOver  bool
}

func drawParts(s tcell.Screen, parts []SnakePart, foodPos SnakePart, style tcell.Style, foodStyle tcell.Style) {
	s.SetContent(foodPos.X, foodPos.Y, '\u25CF', nil, foodStyle)
	for _, part := range parts {
		s.SetContent(part.X, part.Y, ' ', nil, style)
	}
}

func checkCollision(parts []SnakePart, otherPart SnakePart) bool {
	for _, part := range parts {
		if part.X == otherPart.X && part.Y == otherPart.Y {
			return true
		}
	}
	return false
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (g *Game) UpdateFoodPos(width int, height int) {
	g.FoodPos.X = rand.Intn(width)
	g.FoodPos.Y = rand.Intn(height)
	if g.FoodPos.Y == 1 && g.FoodPos.X < 10 {
		g.UpdateFoodPos(width, height)
	}
}

func (g *Game) Run() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	g.SnakeBody.ResetPos(width, height)
	g.UpdateFoodPos(width, height)
	g.GameOver = false
	g.Score = 0
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)
	for {
		longerSnake := false
		g.Screen.Clear()
		// cehck if snake has ate food
		if checkCollision(g.SnakeBody.Parts[len(g.SnakeBody.Parts)-1:], g.FoodPos) {
			g.UpdateFoodPos(width, height)
			longerSnake = true
			g.Score++
		}
		// cehck if snake has collided with it's own tail
		if checkCollision(g.SnakeBody.Parts[:len(g.SnakeBody.Parts)-1], g.SnakeBody.Parts[len(g.SnakeBody.Parts)-1]) {
			break
		}
		g.SnakeBody.UpdateLocation(width, height, longerSnake)
		drawParts(g.Screen, g.SnakeBody.Parts, g.FoodPos, snakeStyle, defStyle)
		drawText(g.Screen, 1, 1, 8+len(strconv.Itoa(g.Score)), 1, "Score: "+strconv.Itoa(g.Score))
		time.Sleep(30 * time.Millisecond)
		g.Screen.Show()
	}
	g.GameOver = true
	drawText(g.Screen, width/2-20, height/2, width/2+20, height/2, "Game Over OOPSI, Score: "+strconv.Itoa(g.Score)+", Play Again? y/n")
	g.Screen.Show()
}
