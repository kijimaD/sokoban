package game

import (
	"math"
)

// タイル
type Tile struct {
	Kind TileKind
}

type TileKind int

const (
	Wall TileKind = iota
	Floor
)

const (
	PlayerChar = `@`
	WallChar   = `#`
	FloorChar  = `.`
	CargoChar  = `&`
)

type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
)

func (t Tile) String() string {
	if t.Kind == Wall {
		return WallChar
	} else if t.Kind == Floor {
		return FloorChar
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
			char := ""
			tile := s.Tiles[Pos{X: j, Y: i}]
			if s.Player.Pos.X == j && s.Player.Pos.Y == i {
				char = PlayerChar
			} else if v, ok := s.Entities[Pos{X: j, Y: i}]; ok {
				char = v.String()
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

// 衝突を判定するには、移動主体のオブジェクトが他のオブジェクトやタイルの情報を知っている必要がある?

type EntityKind int

const (
	Player EntityKind = iota
	Cargo
)

// タイルの上にあるもの。プレイヤーや荷物など、移動する
type Entity struct {
	Pos   *Pos
	Stage *Stage
	Kind  EntityKind
}

func (e *Entity) String() string {
	var str string
	switch e.Kind {
	case Player:
		str = PlayerChar
	case Cargo:
		str = CargoChar
	}
	return str
}

func (e *Entity) Left() {
	if e.canMove(Left) {
		e.left()
	}
}

func (e *Entity) Right() {
	if e.canMove(Right) {
		e.right()
	}
}

func (e *Entity) Up() {
	if e.canMove(Up) {
		e.up()
	}
}

func (e *Entity) Down() {
	if e.canMove(Down) {
		e.down()
	}
}

func (e *Entity) left() {
	e.moveRelative(-1, 0)
}

func (e *Entity) right() {
	e.moveRelative(1, 0)
}

func (e *Entity) down() {
	e.moveRelative(0, 1)
}

func (e *Entity) up() {
	e.moveRelative(0, -1)
}

func (e *Entity) moveRelative(xOffset int, yOffset int) {
	e.Pos.X = e.Pos.X + xOffset
	e.Pos.Y = e.Pos.Y + yOffset
}

func (e Entity) canMove(d Direction) bool {
	var can bool
	switch d {
	case Left:
		e.left()
		can = e.currentTile().Kind == Floor
		e.right()
	case Right:
		e.right()
		can = e.currentTile().Kind == Floor
		e.left()
	case Down:
		e.down()
		can = e.currentTile().Kind == Floor
		e.up()
	case Up:
		e.up()
		can = e.currentTile().Kind == Floor
		e.down()
	}
	return can
}

func (e *Entity) currentTile() Tile {
	return e.Stage.Tiles[*e.Pos]
}

func (e *Entity) OnCollision() {}

// マップ上の座標
type Pos struct {
	X int
	Y int
}
