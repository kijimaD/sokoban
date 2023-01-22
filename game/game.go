package game

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
	RightD Direction = iota
	LeftD
	UpD
	DownD
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

// マップ上の座標
type Pos struct {
	X int
	Y int
}
