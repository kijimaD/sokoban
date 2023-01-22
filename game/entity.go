package game

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

// func (e *Entity) OnCollisionEntity() Entity {
// 	return e.Stage.Entities[*e.Pos]
// }

// func (e *Entity) isCollision() {
// 	e.Stage.Entities[*e.Pos]
// }

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
