package entity

import (
	"github.com/snopan/flappy-chick/animation"
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
)

type PlayerAnimation int

func CreatePlayer(w donburi.World) {
	entity := w.Create(
		component.Animation,
		component.Sprite,
		component.Transform,
		component.PlayerTag,
		component.Velocity,
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
	donburi.SetValue(
		e, component.Transform, component.TransformData{
			Position: component.PositionData{
				X: options.WindowWidth / 2,
				Y: options.WindowHeight / 2,
			},
			Scale: component.ScaleData{
				X: options.PlayerScale,
				Y: options.PlayerScale,
			},
			Rotation: component.RotationData{
				Theta: 0,
			},
		})
	donburi.SetValue(
		e, component.Velocity, component.VelocityData{
			X: 0,
			Y: 0,
		})
}
