package game

import (
	"fmt"
	"math"
)

// タイル
type Tile struct {
	Kind int
}

const (
	WALL  = `#`
	FLOOR = `.`
)

func (t Tile) String() string {
	if t.Kind == 0 {
		return WALL
	} else if t.Kind == 1 {
		return FLOOR
	} else {
		return "?"
	}
}

// (x, y)の形でアクセスできる
// 一覧表示できる
// 二次元配列に変換すればよさげ。[0][1]に代入、みたいにして変換できる。
type Stage struct {
	Tiles    map[Pos]Tile
	Player   Entity
	Entities map[Pos]Entity
}

func (s Stage) String() string {
	result := ""

	l := len(s.Tiles)
	w := int(math.Sqrt(float64(l)))

	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			tile := s.Tiles[Pos{X: i, Y: j}]
			result = result + fmt.Sprintf("%s", tile)
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

// タイルの上にあるもの。プレイヤーや荷物など、移動する
type Entity struct {
	Pos *Pos
}

func (e *Entity) moveRelative(xOffset int, yOffset int) {
	e.Pos.X = e.Pos.X + xOffset
	e.Pos.Y = e.Pos.Y + yOffset
}
func (e *Entity) Left() {
	e.moveRelative(-1, 0)
}
func (e *Entity) Right() {
	e.moveRelative(1, 0)
}
func (e *Entity) Down() {
	e.moveRelative(0, 1)
}
func (e *Entity) Up() {
	e.moveRelative(0, -1)
}

// マップ上の座標
type Pos struct {
	X int
	Y int
}
