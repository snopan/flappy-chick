package spritesheet

import (
	"errors"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	image       *ebiten.Image
	frameWidth  int
	frameHeight int
	rows        int
	cols        int
}

func NewSpriteSheet(image *ebiten.Image, frameWidth, frameHeight, rows, cols int) (*SpriteSheet, error) {
	if image == nil {
		return nil, errors.New("empty image")
	}

	if frameWidth <= 0 || frameHeight <= 0 {
		return nil, errors.New("invalid frame dimension")
	}

	if rows <= 0 || cols <= 0 {
		return nil, errors.New("invalid rows and cols")
	}

	return &SpriteSheet{
		image:       image,
		frameWidth:  frameWidth,
		frameHeight: frameHeight,
		rows:        rows,
		cols:        cols,
	}, nil
}

func (s *SpriteSheet) GetFrame(row, col int) (*ebiten.Image, error) {
	if row < 0 || row >= s.rows {
		return nil, fmt.Errorf("row %d is out of range", row)
	}

	if col < 0 || col >= s.cols {
		return nil, fmt.Errorf("col %d is out of range", col)
	}

	sx, sy := col*s.frameWidth, row*s.frameHeight
	return s.image.SubImage(image.Rect(sx, sy, sx+s.frameWidth, sy+s.frameHeight)).(*ebiten.Image), nil
}

func (s *SpriteSheet) Rows() int {
	return s.rows
}

func (s *SpriteSheet) Cols() int {
	return s.cols
}
