package entity

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/helper/animation"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type PlayerAnimation int

func CreatePlayer(w donburi.World) {
	player := w.Entry(w.Create(
		component.TagPlayer,
		component.Animation,
		component.Sprite,
		component.Velocity,
		component.RectangleCollider,
		transform.Transform,
	))

	donburi.SetValue(
		player, component.Animation, component.AnimationData{
			CurrentSheet:         nil,
			CurrentFrame:         1,
			CurrentFrameDuration: 0,
			NextAnimation:        animation.PlayerFall,
			ShouldChange:         true,
		})
	component.SetSprite(player, nil, options.LayerPlayer)
	transform.SetWorldPosition(player, math.NewVec2(
		options.WindowWidth/2.0,
		options.WindowHeight/2.0,
	))
	transform.SetWorldRotation(player, 0)
	transform.SetWorldScale(player, math.NewVec2(
		options.PlayerScale,
		options.PlayerScale,
	))
}
