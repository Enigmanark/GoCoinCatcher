package engine

import (
	"log"

	Core "github.com/Enigmanark/GoCoinCatcher/src/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type engine struct {
	WindowWidth int
	WindowHeight int
	ScreenWidth int
	ScreenHeight int
	WndowTitle string
	World Core.GameWorld
	frameCount int
}

func Run() {
	if err := ebiten.RunGame(newEngine()); err != nil {
		log.Fatal(err)
	}
}

func newEngine() *engine {
	eng := engine{}
	eng.WindowWidth = 1024
	eng.WindowHeight = 768
	eng.WndowTitle = "Go Ebitengine Coin Catcher by Enigmanark 2023"
	eng.ScreenWidth = eng.WindowWidth / 2
	eng.ScreenHeight = eng.WindowHeight / 2
	ebiten.SetWindowSize(eng.WindowWidth, eng.WindowHeight)
	ebiten.SetWindowTitle(eng.WndowTitle)

	eng.World = Core.NewGameWorld()
	return &eng
}

func (e *engine) Update() error {
	//Calculate delta time
	var delta float64 = 1.0/60.0

	//Update world
	e.World.Update(delta)

	return nil
}

func (e *engine) Draw(screen *ebiten.Image) {
	//Create backBuffer
	backBuffer := ebiten.NewImage(e.ScreenWidth, e.ScreenHeight)

	//Draw World to backbuffer
	e.World.Draw(backBuffer)

	//Render backBuffer to screen, scaled by 2
	drawOps := &ebiten.DrawImageOptions{}
	drawOps.GeoM.Scale(2, 2)
	screen.DrawImage(backBuffer, drawOps)
}

func (e *engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return e.WindowWidth, e.WindowHeight
}


