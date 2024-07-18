package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type RemoveGround struct {
	query *donburi.Query
}

func NewRemoveGround() *RemoveGround {
	return &RemoveGround{
		query: donburi.NewQuery(filter.Contains(component.TagGround, component.RectangleCollider, transform.Transform)),
	}
}

func (r *RemoveGround) Update(w donburi.World) {
	r.query.Each(w, func(e *donburi.Entry) {
		position := transform.WorldPosition(e)

		// Stop if the ground hasn't moved past the left side of window
		collider := component.RectangleCollider.Get(e)
		topLeftX, _ := collider.GetTopLeft(position.X, position.Y)
		if topLeftX > 0.0 {
			return
		}

		// Stop if the right side of ground hasn't moved past the left side of window
		bottomRightX, _ := collider.GetBottomRight(position.X, position.Y)
		if bottomRightX > 0.0 {
			return
		}

		// But if the ground is completely off screen then remove it
		e.Remove()
	})
}
