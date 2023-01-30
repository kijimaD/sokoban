package game

import (
	"errors"
	"fmt"
	"math"
)

var (
	StageInvalidError = errors.New("invalid stage")
)

// (x, y)の形でアクセスできる
// 一覧表示できる
// 二次元配列に変換すればよさげ。[0][1]に代入、みたいにして変換できる。
type Stage struct {
	Tiles    map[Pos]Tile
	Entities Entities
}

func NewStageByString(tiles string, entities string) *Stage {
	stage := Stage{Tiles: map[Pos]Tile{}}

	err, tileArr := stageStrToArray(tiles)
	if err != nil {
		panic(err)
	}
	for i, col := range tileArr {
		for j, rune := range col {
			switch string(rune) {
			case WallChar:
				stage.Tiles[Pos{X: j, Y: i}] = Tile{Kind: Wall}
			case FloorChar:
				stage.Tiles[Pos{X: j, Y: i}] = Tile{Kind: Floor}
			case GoalChar:
				stage.Tiles[Pos{X: j, Y: i}] = Tile{Kind: Goal}
			default:
				fmt.Printf("`%s`は不正な文字です\n", string(rune))
			}
		}
	}

	err, entityArr := stageStrToArray(entities)
	if err != nil {
		panic(err)
	}
	for i, col := range entityArr {
		for j, rune := range col {
			switch string(rune) {
			case PlayerChar:
				player := NewEntity(Pos{X: j, Y: i}, &stage, Player)
				stage.Entities = append(stage.Entities, player)
			case CargoChar:
				cargo := NewEntity(Pos{X: j, Y: i}, &stage, Cargo)
				stage.Entities = append(stage.Entities, cargo)
			case "~":
			default:
				fmt.Printf("`%s`は不正な文字です\n", string(rune))
			}
		}
	}

	return &stage
}

func NewStagePlane(w int) *Stage {
	stage := Stage{Tiles: map[Pos]Tile{}}
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			stage.Tiles[Pos{X: j, Y: i}] = Tile{Kind: Floor}
		}
	}
	return &stage
}

func stageStrToArray(s string) (error, []string) {
	// "12\n34"
	// >>>>
	// ["12",
	//  "34"]

	arr := []string{}
	var row string
	for _, rune := range s {
		s := string(rune)
		if s == "\n" {
			arr = append(arr, row)
			row = ""
		} else {
			row += s
		}
	}

	// validation
	l := len(arr)
	for _, r := range arr {
		if l != len(r) {
			return StageInvalidError, nil
		}
	}

	return nil, arr
}

// テスト用
func InitStage() *Stage {
	// @..#
	// .&.#
	// #_.#
	// ....

	tiles := `...#
...#
#_.#
....
`
	entities := `@~~~
~&~~
~~~~
~~~~
`
	stage := NewStageByString(tiles, entities)
	return stage
}

func InitBigStage() *Stage {
	tiles := `...#.....
...#.....
#_.#.._..
...#_...#
...#.....
..#......
...#.###_
.........
.........
`
	entities := `~~~~~~~~~
~~~~~&~~~
~&~~~~&~~
~~~~&~~~~
~~~~~~~~~
~~~~~~@~~
~~~~~~~~&
~~~~~~~~~
~~~~~~~~~
`
	stage := NewStageByString(tiles, entities)
	return stage
}

func (s Stage) String() string {
	result := ""

	l := len(s.Tiles)
	w := int(math.Sqrt(float64(l)))

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			char := ""
			tile := s.Tiles[Pos{X: j, Y: i}]
			if ok, es := s.Entities.GetEntitiesByPos(Pos{X: j, Y: i}); ok {
				if len(es) == 1 {
					if es[0].Kind == Cargo && tile.Kind == Goal {
						char = PassChar
					} else {
						char = es[0].String()
					}
				} else if len(es) == 2 && (es[0].Kind == Player || es[1].Kind == Player) {
					char = PlayerChar
				} else if len(es) > 1 {
					char = UnknownChar
				}
			} else {
				char = tile.String()
			}
			result = result + char
		}
		result = result + "\n"
	}

	return result
}

// 二次元配列に変換する
func (s Stage) ToSlice() [][]Tile {
	l := len(s.Tiles)
	w := int(math.Sqrt(float64(l)))

	arr := make([][]Tile, w)
	for i := 0; i < w; i++ {
		arr[i] = make([]Tile, w)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			tile := s.Tiles[Pos{X: i, Y: j}]
			arr[i][j] = Tile{Kind: tile.Kind}
		}
	}

	return arr
}

// すべてのゴール上に荷物が置かれていればクリア
func (s Stage) IsFinish() bool {
	finish := true
	for k, v := range s.Tiles {
		tileFinish := false
		if v.Kind == Goal {
			_, es := s.Entities.GetEntitiesByPos(k)
			for _, e := range es {
				if e.Kind == Cargo {
					tileFinish = true
					break
				}
			}

			if !tileFinish {
				finish = false
				break
			}
		}
	}
	return finish
}
