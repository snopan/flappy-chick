package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Velocity struct {
	query *donburi.Query
}

func NewVelocity() *Velocity {
	return &Velocity{
		query: donburi.NewQuery(filter.Contains(component.Transform, component.Velocity)),
	}
}

func (v *Velocity) Update(w donburi.World) {
	v.query.Each(w, func(e *donburi.Entry) {
		transform := component.Transform.Get(e)
		velocity := component.Velocity.Get(e)

		transform.Position.X += velocity.X
		transform.Position.Y += velocity.Y
	})
}
