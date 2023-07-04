package comp

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SubImageComponent struct {
	currentFrame int
	image ebiten.Image
	frameWidth int
	frameHeight int
	layer int
}

func NewSubImageComponent(path string, frameWidth int, frameHeight int, layer int) SubImageComponent {
	ic := SubImageComponent{}

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
	ic.frameWidth, ic.frameHeight = ic.image.Size()

	return ic
}

func (pc *SubImageComponent) GetFrameWidth() int {
	return pc.frameWidth
}

func (pc *SubImageComponent) GetCurrentFrame() int {
	return pc.currentFrame
}

func (pc *SubImageComponent) GetFrameHeight() int {
	return pc.frameHeight
}

func (pc *SubImageComponent) GetImage() *ebiten.Image {
	return &pc.image
}

func (pc *SubImageComponent) GetID() int {
	return C_SUBIMAGE
}