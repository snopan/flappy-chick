package entity

import (
	"github.com/snopan/flappy-chick/animation"
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type PlayerAnimation int

func CreatePlayer(w donburi.World) {
	entity := w.Create(
		component.Animation,
		component.Sprite,
		component.PlayerTag,
		component.Velocity,
		transform.Transform,
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
		e, transform.Transform, transform.TransformData{
			LocalPosition: math.NewVec2(
				options.WindowWidth/2.0,
				options.WindowHeight/2.0,
			),
			LocalRotation: 0,
			LocalScale: math.NewVec2(
				options.PlayerScale,
				options.PlayerScale,
			),
		})
	donburi.SetValue(
		e, component.Velocity, component.VelocityData{
			X: 0,
			Y: 0,
		})
}
