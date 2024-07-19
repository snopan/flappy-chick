package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData struct {
	Image *ebiten.Image
}

var Sprite = donburi.NewComponentType[SpriteData]()

func SetSprite(entry *donburi.Entry, image *ebiten.Image) {
	donburi.SetValue(entry, Sprite, SpriteData{
		Image: image,
	})
}
