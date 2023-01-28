package game

type EntityKind int

const (
	Player EntityKind = iota
	Cargo
)

const (
	PlayerChar = `@`
	CargoChar  = `&`
)

// タイルの上にあるもの。プレイヤーや荷物など、移動する
type Entity struct {
	Pos   *Pos
	Stage *Stage
	Kind  EntityKind
}

func NewEntity(pos *Pos, stage *Stage, kind EntityKind) Entity {
	e := Entity{
		Pos:   pos,
		Stage: stage,
		Kind:  kind,
	}

	return e
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

// FIXME: 重複がたくさんあるのをどうにかする
// 方向別に微妙に違うので、DRYにする方法がわからない
func (e *Entity) canMove(d Direction) bool {
	var canTile bool
	canEntity := true
	initialPosX := e.Pos.X
	initialPosY := e.Pos.Y

	switch d {
	case LeftD:
		e.left()
		canTile = e.currentTile().Kind == Floor || e.currentTile().Kind == Goal
		if e.Kind == Player {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind != Cargo {
					break
				}

				if t.Kind == Cargo && t.canMove(LeftD) {
					t.Left()
					break
				} else {
					canEntity = false
				}
			}
		} else {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind == Cargo {
					canEntity = false
					break
				}
			}
		}
	case RightD:
		e.right()
		canTile = e.currentTile().Kind == Floor || e.currentTile().Kind == Goal
		if e.Kind == Player {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind != Cargo {
					break
				}

				if t.Kind == Cargo && t.canMove(RightD) {
					t.Right()
					break
				} else {
					canEntity = false
				}
			}
		} else {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind == Cargo {
					canEntity = false
					break
				}
			}
		}
	case DownD:
		e.down()
		canTile = e.currentTile().Kind == Floor || e.currentTile().Kind == Goal
		if e.Kind == Player {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind != Cargo {
					break
				}

				if t.Kind == Cargo && t.canMove(DownD) {
					t.Down()
					break
				} else {
					canEntity = false
				}
			}
		} else {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind == Cargo {
					canEntity = false
					break
				}
			}
		}
	case UpD:
		e.up()
		canTile = e.currentTile().Kind == Floor || e.currentTile().Kind == Goal
		if e.Kind == Player {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind != Cargo {
					break
				}

				if t.Kind == Cargo && t.canMove(UpD) {
					t.Up()
					break
				} else {
					canEntity = false
				}
			}
		} else {
			targets := e.collisionEntities()
			for _, t := range targets {
				if t.Kind == Cargo {
					canEntity = false
					break
				}
			}
		}
	}
	e.Pos.X = initialPosX
	e.Pos.Y = initialPosY
	return canTile && canEntity
}

func (e *Entity) currentTile() Tile {
	return e.Stage.Tiles[*e.Pos]
}

// 同じ座標にいるもう1つのentityを取得する
// 重なってないと失敗を返す
// 一時的な移動を含めて、最大で3つ重なる可能性がある。player, cargo, goal
func (e *Entity) collisionEntities() Entities {
	var result Entities

	_, es := e.Stage.Entities.GetEntitiesByPos(*e.Pos)

	for _, ent := range es {
		if *e != ent {
			result = append(result, ent)
		}
	}
	return result
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
