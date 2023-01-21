package game

import (
	"fmt"
	"math"
	"strconv"
)

type Tile struct {
	Kind int
}

type Stage struct {
	Tiles  map[Pos]Tile
	Player Entity
}

type Entity struct {
	Pos Pos
}

type Pos struct {
	X int
	Y int
}

func (s Stage) Show() {
	for k, _ := range s.Tiles {
		fmt.Printf("(%#v, %#v)", k.X, k.Y)
	}
}

// 画面表示するのに使う
func (s Stage) ToSlice() {
	l := len(s.Tiles)
	w := int(math.Sqrt(float64(l)))

	arr := make([][]Tile, w)
	for i := 0; i < w; i++ {
		arr[i] = make([]Tile, w)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			arr[i][j] = Tile{Kind: 1}
		}
	}

	fmt.Printf("\n")
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			tile := arr[i][j]
			char := strconv.Itoa(tile.Kind)
			fmt.Printf("%s ", char)
		}
		fmt.Printf("\n")
	}
}

// (x, y)の形でアクセスできる
// 一覧表示できる
// 二次元配列に変換すればよさげ。[0][1]に代入、みたいにして変換できる。

func InitStage() Stage {
	tiles := map[Pos]Tile{
		Pos{X: 0, Y: 0}: Tile{Kind: 1},
		Pos{X: 0, Y: 1}: Tile{Kind: 0},
		Pos{X: 1, Y: 0}: Tile{Kind: 1},
		Pos{X: 1, Y: 1}: Tile{Kind: 0},
	}
	player := Entity{
		Pos{
			X: 0,
			Y: 0,
		},
	}

	stage := Stage{
		Tiles:  tiles,
		Player: player,
	}
	return stage
}
