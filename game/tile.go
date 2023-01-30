package game

// タイル
type Tile struct {
	Kind TileKind
}

type TileKind int

const (
	Wall TileKind = iota
	Floor
	Goal
)

const (
	WallChar    = `#`
	FloorChar   = `.`
	GoalChar    = `_`
	PassChar    = `✓`
	UnknownChar = `?`
)

func (t Tile) String() string {
	if t.Kind == Wall {
		return WallChar
	} else if t.Kind == Floor {
		return FloorChar
	} else if t.Kind == Goal {
		return GoalChar
	} else {
		return UnknownChar
	}
}
