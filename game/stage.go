package game

import (
	"math"
)

// (x, y)の形でアクセスできる
// 一覧表示できる
// 二次元配列に変換すればよさげ。[0][1]に代入、みたいにして変換できる。
type Stage struct {
	Tiles    map[Pos]Tile
	Entities Entities
}

func InitStage() *Stage {
	// @..#
	// .&.#
	// #_.#
	// ....

	tiles := map[Pos]Tile{
		Pos{X: 0, Y: 0}: Tile{Kind: 1},
		Pos{X: 1, Y: 0}: Tile{Kind: 1},
		Pos{X: 2, Y: 0}: Tile{Kind: 1},
		Pos{X: 3, Y: 0}: Tile{Kind: 0},

		Pos{X: 0, Y: 1}: Tile{Kind: 1},
		Pos{X: 1, Y: 1}: Tile{Kind: 1},
		Pos{X: 2, Y: 1}: Tile{Kind: 1},
		Pos{X: 3, Y: 1}: Tile{Kind: 0},

		Pos{X: 0, Y: 2}: Tile{Kind: 0},
		Pos{X: 1, Y: 2}: Tile{Kind: 1},
		Pos{X: 2, Y: 2}: Tile{Kind: 1},
		Pos{X: 3, Y: 2}: Tile{Kind: 0},

		Pos{X: 0, Y: 3}: Tile{Kind: 1},
		Pos{X: 1, Y: 3}: Tile{Kind: 1},
		Pos{X: 2, Y: 3}: Tile{Kind: 1},
		Pos{X: 3, Y: 3}: Tile{Kind: 1},
	}

	stage := Stage{
		Tiles:    tiles,
		Entities: Entities{},
	}

	// FIXME: entity初期化の順番に依存する
	player := Entity{
		&Pos{
			X: 0,
			Y: 0,
		},
		&stage,
		Player,
	}
	stage.Entities = append(stage.Entities, player)

	cargo := Entity{
		&Pos{
			X: 1,
			Y: 1,
		},
		&stage,
		Cargo,
	}
	stage.Entities = append(stage.Entities, cargo)

	goal := Entity{
		&Pos{
			X: 1,
			Y: 2,
		},
		&stage,
		Goal,
	}
	stage.Entities = append(stage.Entities, goal)

	return &stage
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
					char = es[0].String() // todo
				} else if len(es) == 2 {
					if (es[0].Kind == Cargo && es[1].Kind == Goal) || (es[1].Kind == Cargo && es[0].Kind == Goal) {
						char = `✓`
					} else if es[0].Kind == Player || es[1].Kind == Player {
						char = PlayerChar
					}
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
