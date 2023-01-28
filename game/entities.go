package game

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
