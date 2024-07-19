package entity

import (
	stdmath "math"

	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/helper/sprite"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreatePipe(w donburi.World, centerY float64) {

	// Create the pipe parent where it's position is center of pipe
	pipe := w.Entry(w.Create(
		component.Velocity,
		transform.Transform,
	))
	tempX := options.WindowWidth - 100.0
	transform.SetWorldPosition(pipe, math.NewVec2(tempX, centerY))
	component.SetVelocity(pipe, options.PipeSpeed, 0.0)

	// Have a collider in the center so we can check when player pas ses the pipe
	center := w.Entry(w.Create(
		component.RectangleCollider,
		transform.Transform,
	))
	component.SetRectangleCollider(center, 0.0, options.PipeGap, component.AnchorCenter)
	transform.AppendChild(pipe, center, false)

	CreateTopPipe(w, pipe, -options.PipeGap/2.0)
	CreateBottomPipe(w, pipe, options.PipeGap/2.0)
}

func CreateTopPipe(w donburi.World, pipe *donburi.Entry, startY float64) {

	// Create the parent for top pipe
	top := w.Entry(w.Create(
		component.RectangleCollider,
		transform.Transform,
	))
	transform.SetWorldPosition(top, math.NewVec2(0.0, startY))
	transform.AppendChild(pipe, top, false)

	// Then start building the pipe sections
	currentY := 0.0
	createdEdge := false
	pipeMiddle := sprite.GetSprite(sprite.PipeMiddle)
	pipeSectionHeight := float64(pipeMiddle.Bounds().Size().Y)
	for stdmath.Abs(currentY) < options.WindowHeight {
		pipeSprite := pipeMiddle
		if !createdEdge {
			pipeSprite = sprite.GetSprite(sprite.PipeBottom)
			createdEdge = true
		}

		pipeSection := w.Entry(w.Create(
			component.Sprite,
			transform.Transform,
		))
		component.SetSprite(pipeSection, pipeSprite)
		transform.SetWorldPosition(pipeSection, math.NewVec2(0.0, currentY))
		transform.SetWorldScale(pipeSection, math.NewVec2(options.PipeScale, options.PipeScale))
		transform.AppendChild(top, pipeSection, false)
		currentY -= pipeSectionHeight * options.PipeScale
	}
}

func CreateBottomPipe(w donburi.World, pipe *donburi.Entry, startY float64) {

	// Create the parent for bottom pipe
	bottom := w.Entry(w.Create(
		component.RectangleCollider,
		transform.Transform,
	))
	transform.SetWorldPosition(bottom, math.NewVec2(0.0, startY))
	transform.AppendChild(pipe, bottom, false)

	// Then start building the pipe sections
	currentY := 0.0
	createdEdge := false
	pipeMiddle := sprite.GetSprite(sprite.PipeMiddle)
	pipeSectionHeight := float64(pipeMiddle.Bounds().Size().Y)
	for stdmath.Abs(currentY) < options.WindowHeight {
		pipeSprite := pipeMiddle
		if !createdEdge {
			pipeSprite = sprite.GetSprite(sprite.PipeTop)
			createdEdge = true
		}

		pipeSection := w.Entry(w.Create(
			component.Sprite,
			transform.Transform,
		))
		component.SetSprite(pipeSection, pipeSprite)
		transform.SetWorldPosition(pipeSection, math.NewVec2(0.0, currentY))
		transform.SetWorldScale(pipeSection, math.NewVec2(options.PipeScale, options.PipeScale))
		transform.AppendChild(bottom, pipeSection, false)
		currentY += pipeSectionHeight * options.PipeScale
	}
}
