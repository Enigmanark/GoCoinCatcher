package core
import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Load()
	Unload()
	Update(float64)
	Draw(*ebiten.Image)
}