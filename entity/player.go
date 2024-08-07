package entity

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/helper/animation"
	"github.com/snopan/flappy-chick/helper/sprite"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type PlayerAnimation int

func CreatePlayer(w donburi.World) {
	playerWidth := float64(sprite.GetSprite(sprite.PipeMiddle).Bounds().Size().X)

	player := w.Entry(w.Create(
		component.TagPlayer,
		component.Animation,
		component.Sprite,
		component.Velocity,
		component.RectangleCollider,
		component.Scoring,
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
	donburi.SetValue(
		player, component.Scoring, component.ScoringData{
			Score:       0,
			PassingPipe: false,
		},
	)

	component.SetRectangleCollider(player, pipeWidth*options.PipeScale, options.WindowHeight, component.AnchorCenter)
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
