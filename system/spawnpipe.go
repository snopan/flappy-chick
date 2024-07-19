package system

import (
	"math/rand/v2"

	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/entity"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type SpawnPipe struct {
	query *donburi.Query
}

func NewSpawnPipe() *SpawnPipe {
	return &SpawnPipe{
		query: donburi.NewQuery(filter.Contains(component.TagLastPipe, transform.Transform)),
	}
}

func (s *SpawnPipe) Update(w donburi.World) {
	if s.query.Count(w) == 0 {
		e := entity.CreatePipe(w, options.WindowWidth, randomCenterY())
		e.AddComponent(component.TagLastPipe)
	}

	s.query.Each(w, func(e *donburi.Entry) {
		position := transform.WorldPosition(e)

		// Stop if the pipe hasn't moved in from the right side of window
		collider := component.RectangleCollider.Get(e)
		topLeftX, _ := collider.GetTopLeft(position.X, position.Y)
		if topLeftX > options.WindowWidth {
			return
		}

		// Create another pipe to the right of it
		bottomRightX, _ := collider.GetBottomRight(position.X, position.Y)
		newLast := entity.CreatePipe(w, bottomRightX+options.PipesDistance, randomCenterY())
		newLast.AddComponent(component.TagLastPipe)

		// This pipe is no longer the last one
		e.RemoveComponent(component.TagLastPipe)
	})
}

func randomCenterY() float64 {
	min := (1.0 - options.PipeCenterWindowHeightFactor) * options.WindowHeight / 2.0
	max := (1.0 + options.PipeCenterWindowHeightFactor) * options.WindowHeight / 2.0
	return rand.Float64()*(max-min) + min
}
