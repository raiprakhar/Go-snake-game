package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)

	snake := SnakeBody{
		Parts: []SnakePart{
			{
				X: 5,
				Y: 10,
			},
			{
				X: 6,
				Y: 10,
			},
			{
				X: 7,
				Y: 10,
			},
		},
		Xspeed: 0,
		Yspeed: -1,
	}

	game := Game{
		Screen:    screen,
		SnakeBody: snake,
	}

	// GOROUTINE FOR RUNNING THE INFDEF LOOP OF RUN()
	// AS WE NEED PARALLEL USE OF BOTH USER INPUT AND IMAGE RENDERING FUNCTIONALITIES

	// FOR UPDATING LOCATION
	go game.Run()

	// FOR UPDATING DIRECTION
	for {
		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()

		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.SnakeBody.UpdateDirection(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				game.SnakeBody.UpdateDirection(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				game.SnakeBody.UpdateDirection(0, -1)
			} else if event.Key() == tcell.KeyRight {
				game.SnakeBody.UpdateDirection(0, 1)
			} else if event.Rune() == 'y' && game.GameOver {
				go game.Run()
			} else if event.Rune() == 'n' && game.GameOver {
				game.Screen.Fini()
				os.Exit(0)
			}
		}
	}
}
