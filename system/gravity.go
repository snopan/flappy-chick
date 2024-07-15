package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Gravity struct {
	query *donburi.Query
}

func NewGravity() *Gravity {
	return &Gravity{
		query: donburi.NewQuery(filter.Contains(component.Velocity, component.PlayerTag)),
	}
}

func (g *Gravity) Update(w donburi.World) {
	g.query.Each(w, func(e *donburi.Entry) {
		velocity := component.Velocity.Get(e)
		velocity.Y += 9.8 * 0.01
	})
}
