package animation

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/spritesheet"
)

type Animation struct {
	sheet         *spritesheet.SpriteSheet
	startFrame    int
	endFrame      int
	frameDuration int
}

func FromSpriteSheet(sheet *spritesheet.SpriteSheet, frameDuration int) *Animation {
	return &Animation{
		sheet:         sheet,
		startFrame:    0,
		endFrame:      sheet.Rows()*sheet.Cols() - 1,
		frameDuration: frameDuration,
	}
}

func (a *Animation) GetFrame(index int) (*ebiten.Image, error) {
	if index < 0 || index > a.endFrame {
		return nil, fmt.Errorf("index %d is out of range", index)
	}

	row := int(math.Floor(float64(index) / float64(a.sheet.Cols())))
	col := index - row*a.sheet.Cols()

	return a.sheet.GetFrame(row, col)
}

func (a *Animation) StartFrame() int {
	return a.startFrame
}

func (a *Animation) EndFrame() int {
	return a.endFrame
}

func (a *Animation) FrameDuration() int {
	return a.frameDuration
}
