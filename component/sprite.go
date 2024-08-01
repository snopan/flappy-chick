package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
	Layer int
}

var Sprite = donburi.NewComponentType[SpriteData]()

func SetSprite(entry *donburi.Entry, image *ebiten.Image, layer int) {
	donburi.SetValue(entry, Sprite, SpriteData{
		Image: image,
		Layer: layer,
	})
}
