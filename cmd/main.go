package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/animation"
	"github.com/snopan/flappy-chick/drawable"
	"github.com/snopan/flappy-chick/options"
	"github.com/snopan/flappy-chick/spritesheet"
	"github.com/snopan/flappy-chick/system"
	"github.com/yohamta/donburi"
)

type System interface {
	Update(w donburi.World)
}

type Drawable interface {
	Draw(w donburi.World, screen *ebiten.Image)
}

type Game struct {
	bounds    image.Rectangle
	world     donburi.World
	systems   []System
	drawables []Drawable
}

func NewGame() *Game {
	g := &Game{
		bounds: image.Rectangle{},
		world:  donburi.NewWorld(),
	}
	g.systems = []System{
		system.NewSpawn(),
		system.NewAnimation(),
		system.NewGravity(),
		system.NewVelocity(),
		system.NewFly(),
		system.NewPlayerAnimator(),
	}
	g.drawables = []Drawable{
		drawable.NewRender(),
	}
	return g
}

func (g *Game) Update() error {
	for _, s := range g.systems {
		s.Update(g.world)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, s := range g.drawables {
		s.Draw(g.world, screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	ebiten.SetWindowSize(options.WindowWidth, options.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)

	if err := spritesheet.InitSpriteSheetLibrary(); err != nil {
		log.Fatal(err)
	}

	if err := animation.InitAnimationLibrary(); err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
