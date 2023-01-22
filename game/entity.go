package game

type EntityKind int

const (
	Player EntityKind = iota
	Cargo
	Goal
)

const (
	PlayerChar = `@`
	CargoChar  = `&`
	GoalChar   = `_`
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
	case Goal:
		str = GoalChar
	}
	return str
}

func (e *Entity) Left() {
	if e.canMove(LeftD) {
		e.left()
	}
}

func (e *Entity) Right() {
	if e.canMove(RightD) {
		e.right()
	}
}

func (e *Entity) Up() {
	if e.canMove(UpD) {
		e.up()
	}
}

func (e *Entity) Down() {
	if e.canMove(DownD) {
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
	case LeftD:
		e.left()
		can = e.currentTile().Kind == Floor
		e.right()
	case RightD:
		e.right()
		can = e.currentTile().Kind == Floor
		e.left()
	case DownD:
		e.down()
		can = e.currentTile().Kind == Floor
		e.up()
	case UpD:
		e.up()
		can = e.currentTile().Kind == Floor
		e.down()
	}
	return can
}

func (e *Entity) currentTile() Tile {
	return e.Stage.Tiles[*e.Pos]
}

// func (e *Entity) OnCollisionEntity() Entity {
// 	return e.Stage.Entities[*e.Pos]
// }

// Entity同士(片方Player)が重なった状態である
func (e *Entity) isCollision() bool {
	targetE := e.Stage.Entities[*e.Pos]
	return targetE.Kind != Player
}

type Entities map[Pos]Entity

func (es Entities) Player() Entity {
	var result Entity
	for _, val := range es {
		if val.Kind == Player {
			result = val
		}
	}

	return result
}
