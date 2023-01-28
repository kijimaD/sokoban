package game

type Direction int

const (
	RightD Direction = iota
	LeftD
	UpD
	DownD
)

// マップ上の座標
type Pos struct {
	X int
	Y int
}
