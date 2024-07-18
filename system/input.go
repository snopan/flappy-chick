package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Input struct {
	query *donburi.Query
}

func NewInput() *Input {
	return &Input{
		query: donburi.NewQuery(filter.Contains(component.TagPlayer, component.Velocity)),
	}
}

func (f *Input) Update(w donburi.World) {
	f.query.Each(w, func(e *donburi.Entry) {
		if !inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			return
		}

		velocity := component.Velocity.Get(e)
		velocity.Y = options.FlyUpSpeed
	})
}
