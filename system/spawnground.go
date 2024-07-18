package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/entity"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type SpawnGround struct {
	query *donburi.Query
}

func NewSpawnGround() *SpawnGround {
	return &SpawnGround{
		query: donburi.NewQuery(filter.Contains(component.TagLastGround, component.RectangleCollider, transform.Transform)),
	}
}

func (s *SpawnGround) Update(w donburi.World) {
	if s.query.Count(w) == 0 {
		entity.CreateGround(w, 0.0)
	}

	s.query.Each(w, func(e *donburi.Entry) {
		position := transform.WorldPosition(e)

		// Stop if the ground hasn't moved past the left side of window
		collider := component.RectangleCollider.Get(e)
		topLeftX, _ := collider.GetTopLeft(position.X, position.Y)
		if topLeftX > 0.0 {
			return
		}

		// Create another ground to the right of it
		bottomRightX, _ := collider.GetBottomRight(position.X, position.Y)
		entity.CreateGround(w, bottomRightX)

		// This ground is no longer the last one
		e.RemoveComponent(component.TagLastGround)
	})
}
