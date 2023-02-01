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
	RESET_KEY = "r"
)

func Run() {
	splash := `
---- START ----
@: You
&: Cargo
_: Goal
#: Wall
.: Floor

`
	fmt.Print(splash)

	s := game.GenStage()
	fmt.Println(s)
	player := s.Entities.Player()

	var poses []game.Pos
	for _, e := range s.Entities {
		poses = append(poses, *e.Pos)
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		in := scanner.Text()
		player = s.Entities.Player()

		switch in {
		case UP_KEY:
			player.Up()
		case LEFT_KEY:
			player.Left()
		case DOWN_KEY:
			player.Down()
		case RIGHT_KEY:
			player.Right()
		case RESET_KEY:
			s.ResetPos(poses)
		default:
			continue
		}

		fmt.Println(s)

		if s.IsFinish() {
			fmt.Println("solve!!")

			break
		}
	}
}
