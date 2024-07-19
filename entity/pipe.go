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

	CreatePipeSections(w, pipe, -options.PipeGap/2.0, false)
	CreatePipeSections(w, pipe, options.PipeGap/2.0, true)
}

func CreatePipeSections(w donburi.World, pipe *donburi.Entry, startY float64, buildBottom bool) {

	// Create the parent for pipe section
	sectionParent := w.Entry(w.Create(
		component.RectangleCollider,
		transform.Transform,
	))
	transform.SetWorldPosition(sectionParent, math.NewVec2(0.0, startY))
	transform.AppendChild(pipe, sectionParent, false)

	// Define which direction the pipe should be built and what edge sprite to use
	buildDirect := -1.0
	edgeSprite := sprite.GetSprite(sprite.PipeBottom)
	if buildBottom {
		buildDirect = 1.0
		edgeSprite = sprite.GetSprite(sprite.PipeTop)
	}

	// Then start building the pipe sections
	currentY := 0.0
	createdEdge := false
	pipeMiddle := sprite.GetSprite(sprite.PipeMiddle)
	pipeSectionHeight := float64(pipeMiddle.Bounds().Size().Y)
	for stdmath.Abs(currentY) < options.WindowHeight {
		pipeSprite := pipeMiddle
		if !createdEdge {
			pipeSprite = edgeSprite
			createdEdge = true
		}

		pipeSection := w.Entry(w.Create(
			component.Sprite,
			transform.Transform,
		))
		component.SetSprite(pipeSection, pipeSprite)
		transform.SetWorldPosition(pipeSection, math.NewVec2(0.0, currentY))
		transform.SetWorldScale(pipeSection, math.NewVec2(options.PipeScale, options.PipeScale))
		transform.AppendChild(sectionParent, pipeSection, false)
		currentY += buildDirect * pipeSectionHeight * options.PipeScale
	}
}
