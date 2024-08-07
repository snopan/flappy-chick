package system

import (
	"fmt"

	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type Scoring struct {
	player *donburi.Query
	pipes  *donburi.Query
}

func NewScoring() *Scoring {
	return &Scoring{
		player: donburi.NewQuery(filter.Contains(component.TagPlayer, component.Scoring, component.RectangleCollider, transform.Transform)),
		pipes:  donburi.NewQuery(filter.Contains(component.TagPipe, component.RectangleCollider, transform.Transform)),
	}
}

func (s *Scoring) Update(w donburi.World) {
	player, ok := s.player.FirstEntity(w)
	if !ok {
		return
	}

	playerPosition := transform.WorldPosition(player)
	playerCollider := component.RectangleCollider.Get(player)
	scoring := component.Scoring.Get(player)

	overlap := false
	s.pipes.Each(w, func(pipe *donburi.Entry) {
		pipePosition := transform.WorldPosition(pipe)
		pipeCollider := component.RectangleCollider.Get(pipe)

		overlap = component.IsOverlap(
			playerCollider,
			pipeCollider,
			playerPosition.X,
			playerPosition.Y,
			pipePosition.X,
			pipePosition.Y,
		)
		if overlap {
			return
		}
	})

	fmt.Println(overlap)

	if scoring.PassingPipe && !overlap {
		scoring.PassingPipe = false
		scoring.Score++
		fmt.Println(scoring.Score)
	} else {
		scoring.PassingPipe = overlap
	}
}
