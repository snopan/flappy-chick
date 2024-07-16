package spritesheet

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheetKey int

var spriteSheetLibrary = map[SpriteSheetKey]*SpriteSheet{}

const (
	ChickenDie SpriteSheetKey = iota
	ChickenFall
	ChickenFly
	Pipes
	Ground
)

var spriteSheetData = map[SpriteSheetKey]struct {
	src          string
	frameWidth   int
	frameHeight  int
	offsetWidth  int
	offsetHeight int
	rows         int
	cols         int
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
	Pipes: {
		src:         "assets/images/terrain.png",
		frameWidth:  32,
		frameHeight: 16,
		rows:        3,
		cols:        4,
	},
	Ground: {
		src:          "assets/images/terrain.png",
		frameWidth:   16,
		frameHeight:  16,
		offsetHeight: 48,
		rows:         3,
		cols:         4,
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
			data.offsetWidth,
			data.offsetHeight,
			data.rows,
			data.cols,
		)
		if err != nil {
			return fmt.Errorf("failed to create sprite %d: %w", key, err)
		}

		spriteSheetLibrary[key] = spriteSheet

		f.Close()
	}

	return nil
}

func GetSpriteSheet(key SpriteSheetKey) *SpriteSheet {
	spriteSheet, ok := spriteSheetLibrary[key]
	if !ok {
		panic(fmt.Sprintf("failed to get sprite sheet: %d", key))
	}

	return spriteSheet
}
