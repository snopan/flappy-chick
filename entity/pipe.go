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

// CreatePipe would create a pipe that would be positoned on the very
// right passed the edge of window and the pipes would have a hole
// at the provided centerY
func CreatePipe(w donburi.World, centerY float64) *donburi.Entry {
	pipeWidth := float64(sprite.GetSprite(sprite.PipeMiddle).Bounds().Size().X)

	// Create the pipe parent where it's position is center of pipe
	pipe := w.Entry(w.Create(
		component.Velocity,
		component.RectangleCollider,
		transform.Transform,
	))
	transform.SetWorldPosition(pipe, math.NewVec2(
		options.WindowWidth+pipeWidth*options.PipeScale/2,
		centerY,
	))
	component.SetRectangleCollider(pipe, pipeWidth*options.PipeScale, options.WindowHeight, component.AnchorCenter)
	component.SetVelocity(pipe, options.PipeSpeed, 0.0)

	CreatePipeSections(w, pipe, -options.PipeGap/2.0, pipeWidth, false)
	CreatePipeSections(w, pipe, options.PipeGap/2.0, pipeWidth, true)

	return pipe
}

func CreatePipeSections(w donburi.World, pipe *donburi.Entry, startY, pipeWidth float64, buildBottom bool) {

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
	colliderAnchor := component.AnchorBottom
	if buildBottom {
		buildDirect = 1.0
		edgeSprite = sprite.GetSprite(sprite.PipeTop)
		colliderAnchor = component.AnchorTop
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
		transform.SetWorldPosition(pipeSection, math.NewVec2(
			0.0,
			currentY+buildDirect*pipeSectionHeight*options.PipeScale/2.0,
		))
		transform.SetWorldScale(pipeSection, math.NewVec2(options.PipeScale, options.PipeScale))
		transform.AppendChild(sectionParent, pipeSection, false)
		currentY += buildDirect * pipeSectionHeight * options.PipeScale
	}

	component.SetRectangleCollider(sectionParent, pipeWidth*options.PipeScale, stdmath.Abs(currentY), colliderAnchor)
}
