package comp

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageComponent struct {
	image ebiten.Image
	width int
	height int
	layer int
}

func NewImageComponent(path string, layer int) ImageComponent {
	ic := ImageComponent{}

	if layer > 4 && layer < 1 {
		panic("ImageComponent not passed a proper rendering layer! Layer must be passed 1, 2, 3, or 4!")
	} else {
		ic.layer = layer
	}
	img, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		panic(err)
	}

	ic.image = *ebiten.NewImageFromImage(img)
	ic.width, ic.height = ic.image.Size()

	return ic
}

func (pc *ImageComponent) GetWidth() int {
	return pc.width
}

func (pc *ImageComponent) GetHeight() int {
	return pc.height
}

func (pc *ImageComponent) GetImage() *ebiten.Image {
	return &pc.image
}

func (pc *ImageComponent) GetID() int {
	return C_IMAGE
}