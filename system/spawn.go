package system

import (
	"github.com/snopan/flappy-chick/entity"
	"github.com/yohamta/donburi"
)

type Spawn struct {
	spawned bool
}

func NewSpawn() *Spawn {
	return &Spawn{
		spawned: false,
	}
}

func (s *Spawn) Update(w donburi.World) {
	if s.spawned {
		return
	}
	entity.CreatePlayer(w)
	s.spawned = true
}
