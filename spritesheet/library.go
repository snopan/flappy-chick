package spritesheet

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheetKey int

var SpriteSheetLibrary = map[SpriteSheetKey]*SpriteSheet{}

const (
	ChickenDie SpriteSheetKey = iota
	ChickenFall
	ChickenFly
)

var spriteSheetData = map[SpriteSheetKey]struct {
	src         string
	frameWidth  int
	frameHeight int
	rows        int
	cols        int
}{
	ChickenDie: {
		src:         "assets/images/chicken_die.png",
		frameWidth:  40,
		frameHeight: 40,
		rows:        1,
		cols:        10,
	},
	ChickenFall: {
		src:         "assets/images/chicken_fall.png",
		frameWidth:  40,
		frameHeight: 40,
		rows:        1,
		cols:        2,
	},
	ChickenFly: {
		src:         "assets/images/chicken_fly.png",
		frameWidth:  40,
		frameHeight: 40,
		rows:        1,
		cols:        4,
	},
}

func InitSpriteSheetLibrary() error {
	for key, data := range spriteSheetData {
		f, err := os.Open(data.src)
		if err != nil {
			return fmt.Errorf("failed to opem file for %d sprite sheet: %w", key, err)
		}

		img, _, err := image.Decode(f)
		if err != nil {
			return fmt.Errorf("failed to decode image for %d sprite sheet: %w", key, err)
		}

		ebitenImage := ebiten.NewImageFromImage(img)
		spriteSheet, err := NewSpriteSheet(
			ebitenImage,
			data.frameWidth,
			data.frameHeight,
			data.rows,
			data.cols,
		)
		if err != nil {
			return fmt.Errorf("failed to create sprite %d: %w", key, err)
		}

		SpriteSheetLibrary[key] = spriteSheet

		f.Close()
	}

	return nil
}
