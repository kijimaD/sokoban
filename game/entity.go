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
	var canTile bool
	canEntity := true
	switch d {
	case LeftD:
		e.left()
		canTile = e.currentTile().Kind == Floor
		e.right()
	case RightD:
		e.right()
		canTile = e.currentTile().Kind == Floor
		e.left()
	case DownD:
		e.down()
		canTile = e.currentTile().Kind == Floor
		if e.Kind == Player && e.isCollision() {
			_, target := e.collisionEntity()
			if target.Kind == Cargo && target.canMove(DownD) {
				target.Down()
			} else {
				canEntity = false
			}
		} else {
			canEntity = true
		}
		e.up()
	case UpD:
		e.up()
		canTile = e.currentTile().Kind == Floor
		e.down()
	}
	return canTile && canEntity
}

func (e *Entity) currentTile() Tile {
	return e.Stage.Tiles[*e.Pos]
}

// 同じ座標にいるもう1つのentityを取得する
// 重なってないと失敗を返す
func (e *Entity) collisionEntity() (bool, *Entity) {
	var result *Entity
	var ok bool

	_, es := e.Stage.Entities.GetEntitiesByPos(*e.Pos)

	if len(es) == 2 {
		for _, ent := range es {
			if *e != ent {
				result = &ent
				ok = true
				break
			}
		}
	}
	return ok, result
}

// Entity同士が重なった状態である
func (e *Entity) isCollision() bool {
	var result bool

	ok, es := e.Stage.Entities.GetEntitiesByPos(*e.Pos)
	if ok && len(es) > 1 {
		result = true
	}
	return result
}

type Entities []Entity

func (es Entities) Player() *Entity {
	var result *Entity
	for _, val := range es {
		if val.Kind == Player {
			result = &val
			break
		}
	}

	return result
}

func (es Entities) NotPlayer() *Entity {
	var result *Entity
	for _, val := range es {
		if val.Kind != Player {
			result = &val
			break
		}
	}

	return result
}

// 同じ座標にあるEntitiesを取得する
func (es Entities) GetEntitiesByPos(p Pos) (bool, Entities) {
	var result Entities
	var success bool
	for _, e := range es {
		if e.Pos.X == p.X && e.Pos.Y == p.Y {
			result = append(result, e)
			success = true
		}
	}

	return success, result
}
