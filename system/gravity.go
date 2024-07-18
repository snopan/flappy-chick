package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Gravity struct {
	query *donburi.Query
}

func NewGravity() *Gravity {
	return &Gravity{
		query: donburi.NewQuery(filter.Contains(component.TagPlayer, component.Velocity)),
	}
}

func (g *Gravity) Update(w donburi.World) {
	g.query.Each(w, func(e *donburi.Entry) {
		velocity := component.Velocity.Get(e)
		velocity.Y += options.Graivty
	})
}
