package game

import (
	"math"
	"math/rand"
	"time"
)

func GenStage() *Stage {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	s := NewStagePlane(10)
	s.putCG(r.Intn(2) + 4)
	s.putWall(r.Intn(10) + 20)
	s.putPlayer(1)

	for i := 0; i < 1000; i++ {
		s.Entities.Player().randomPull()
	}

	return s
}

// TODO: put系は似通ってるので共通化したい
// ランダムにcargoとgoalを配置
func (s *Stage) putCG(num int) {
	poses := s.randomPoses(num)
	for _, p := range poses {
		s.Entities = append(s.Entities, NewEntity(p, s, Cargo))
		s.Tiles[p] = Tile{Kind: Goal}
	}
}

// ランダムに壁を配置
func (s *Stage) putWall(num int) {
	poses := s.randomPoses(num)

	for _, p := range poses {
		s.Tiles[p] = Tile{Kind: Wall}
	}
}

func (s *Stage) putPlayer(num int) {
	poses := s.randomPoses(num)

	for _, p := range poses {
		s.Entities = append(s.Entities, NewEntity(p, s, Player))
	}
}

// ランダムに、座標をえらぶ
func (s *Stage) randomPos() Pos {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	w := math.Sqrt(float64(len(s.Tiles)))
	ix := r.Intn(int(w))
	iy := r.Intn(int(w))

	pos := Pos{X: ix, Y: iy}
	return pos
}

// ランダムにかぶらない座標をえらぶ
func (s *Stage) randomPoses(n int) []Pos {
	max := int(math.Pow(float64(len(s.Tiles)), 2))
	if n > max || n == 0 {
		panic("invalid n value")
	}
	var results []Pos
	for {
		pos := s.randomPos()
		_, es := s.Entities.GetEntitiesByPos(pos)
		if s.Tiles[pos].Kind == Floor && len(es) == 0 {
			results = append(results, pos)
		}
		results = uniq(results)
		if len(results) == n {
			break
		}
	}
	return results
}

// mapに同じキーは存在できないことを利用する
func uniq(poses []Pos) []Pos {
	m := make(map[Pos]bool)
	uniq := []Pos{}

	for _, ele := range poses {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return uniq
}
