package game

import (
	"math"
	"math/rand"
	"time"
)

// WIP
func GenStage() *Stage {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	s := NewStagePlane(10)
	s.putCG(r.Intn(2) + 4)
	s.putWall(r.Intn(20) + 20)

	return s
}

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

	// 既にgoalになってるタイルを除く
	for _, p := range poses {
		if s.Tiles[p].Kind != Goal {
			s.Tiles[p] = Tile{Kind: Wall}
		}
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
		results = append(results, s.randomPos())
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
