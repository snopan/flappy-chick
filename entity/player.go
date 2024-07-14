package entity

import (
	"github.com/snopan/flappy-chick/animation"
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
)

type PlayerAnimation int

func CreatePlayer(w donburi.World) {
	entity := w.Create(
		component.Animation,
		component.Sprite,
	)

	e := w.Entry(entity)
	donburi.SetValue(
		e, component.Animation, component.AnimationData{
			CurrentSheet:         nil,
			CurrentFrame:         1,
			CurrentFrameDuration: 0,
			NextAnimation:        animation.PlayerFall,
			ShouldChange:         true,
		})
}
