package game

import (
	"math"
	"math/rand"
	"time"
)

// ランダムにcargoとgoalを配置
func (s *Stage) setCG(num int) {
	poses := s.randomPoses(num)

	for _, p := range poses {
		s.Entities = append(s.Entities, NewEntity(&p, s, Cargo))

		s.Tiles[p] = Tile{Kind: Goal}
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
	w := len(s.Tiles)
	if n > w || n == 0 {
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
