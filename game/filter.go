package game

import (
	"math"
	"math/rand"
	"time"
)

// ランダムにcargoとgoalを配置
func (s *Stage) setCG(num int) {
	cargo := NewEntity(&Pos{X: 1, Y: 1}, s, Cargo)
	s.Entities = append(s.Entities, cargo)

	s.Tiles[Pos{X: 1, Y: 1}] = Tile{Kind: Goal}
}

// ランダムに、かぶらない座標をえらぶ
func (s *Stage) randomPos() Pos {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	w := math.Sqrt(float64(len(s.Tiles)))
	ix := r.Intn(int(w))
	iy := r.Intn(int(w))

	pos := Pos{X: ix, Y: iy}
	return pos
}

// func (s *Stage) randomPoses(n int) []Pos {}
