package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kijimaD/sokoban/game"
)

const (
	UP_KEY    = "w"
	LEFT_KEY  = "a"
	DOWN_KEY  = "s"
	RIGHT_KEY = "d"
)

func Run() {
	splash := `
---- START ----
@: You
&: Cargo
_: Goal

`
	fmt.Print(splash)

	s := game.InitStage()
	fmt.Println(s)
	player := s.Entities.Player()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		in := scanner.Text()

		switch in {
		case UP_KEY:
			player.Up()
		case LEFT_KEY:
			player.Left()
		case DOWN_KEY:
			player.Down()
		case RIGHT_KEY:
			player.Right()
		default:
			continue
		}

		fmt.Println(s)

		if s.Entities.IsFinish() {
			fmt.Println("solve!!")

			break
		}
	}
}
