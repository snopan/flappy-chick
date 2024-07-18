package entity

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/helper/sprite"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func CreateGround(w donburi.World, x float64) {
	parent := w.Entry(w.Create(
		component.TagGround,
		component.TagLastGround,
		component.Velocity,
		component.RectangleCollider,
		transform.Transform,
	))

	transform.SetWorldPosition(parent, math.NewVec2(
		x,
		options.WindowHeight,
	))

	donburi.SetValue(parent, component.Velocity, component.VelocityData{
		X: options.GroundSpeed,
		Y: 0,
	})

	currentX := 0.0
	currentY := 0.0
	groundLength := sprite.GetSprite(sprite.Dirt).Bounds().Size().X
	for currentX < options.WindowWidth {
		currentY = 0.0
		CreateGroundTile(w, parent, sprite.GetSprite(sprite.DirtLightLine), currentX, currentY, float64(groundLength))
		currentY -= float64(groundLength) * options.GroundScale
		CreateGroundTile(w, parent, sprite.GetSprite(sprite.Dirt), currentX, currentY, float64(groundLength))
		currentY -= float64(groundLength) * options.GroundScale
		CreateGroundTile(w, parent, GetRandomGroundTop(), currentX, currentY, float64(groundLength))
		currentY -= float64(groundLength) * options.GroundScale
		currentX += float64(groundLength) * options.GroundScale
	}

	donburi.SetValue(parent, component.RectangleCollider, component.RectangleColliderData{
		Width:  currentX,
		Height: -currentY,
		Anchor: component.AnchorBottomLeft,
	})

}

func CreateGroundTile(w donburi.World, parent *donburi.Entry, sprite *ebiten.Image, x, y, tileLength float64) {
	tile := w.Entry(w.Create(
		component.Sprite,
		transform.Transform,
	))
	donburi.SetValue(tile, component.Sprite, component.SpriteData{
		Image: sprite,
	})
	donburi.SetValue(tile, transform.Transform, transform.TransformData{
		LocalPosition: math.NewVec2(
			x+tileLength*options.GroundScale/2.0,
			y-tileLength*options.GroundScale/2.0,
		),
		LocalScale: math.NewVec2(
			options.GroundScale,
			options.GroundScale,
		),
	})
	transform.AppendChild(parent, tile, false)
}

func GetRandomGroundTop() *ebiten.Image {
	groundTops := []sprite.SpriteKey{
		sprite.GroundTop1,
		sprite.GroundTop2,
		sprite.GroundTop3,
	}
	randomKey := groundTops[rand.Intn(len(groundTops))]
	return sprite.GetSprite(randomKey)
}
