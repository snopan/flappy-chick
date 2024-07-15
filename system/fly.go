package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Fly struct {
	query *donburi.Query
}

func NewFly() *Fly {
	return &Fly{
		query: donburi.NewQuery(filter.Contains(component.PlayerTag, component.Velocity)),
	}
}

func (f *Fly) Update(w donburi.World) {
	f.query.Each(w, func(e *donburi.Entry) {
		if !ebiten.IsKeyPressed(ebiten.KeySpace) {
			return
		}

		velocity := component.Velocity.Get(e)
		velocity.Y = options.FlyUpSpeed
	})
}
