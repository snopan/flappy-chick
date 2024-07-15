package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Tilt struct {
	query *donburi.Query
}

func NewTilt() *Tilt {
	return &Tilt{
		query: donburi.NewQuery(filter.Contains(component.PlayerTag, component.Velocity, component.Transform)),
	}
}

func (t *Tilt) Update(w donburi.World) {
	t.query.Each(w, func(e *donburi.Entry) {
		velocity := component.Velocity.Get(e)
		transform := component.Transform.Get(e)

	})
}
