package sys

import (
	"image"

	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Enigmanark/GoCoinCatcher/src/core/ecs/comp"
)

func SystemRenderSprite(entities []ECS.Entity, surface *ebiten.Image) {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if e.HasComponent(comp.C_IMAGE) && e.HasComponent(comp.C_POSITION) {
			comp_p := *e.GetComponent(comp.C_POSITION)
			var p *comp.PositionComponent = comp_p.(*comp.PositionComponent)
			
			comp_i := *e.GetComponent(comp.C_IMAGE)
			var img *comp.ImageComponent = comp_i.(*comp.ImageComponent)

			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(float64(p.GetX()), float64(p.GetY()))
			surface.DrawImage(img.GetImage(), drawOp)
		} else if e.HasComponent(comp.C_SUBIMAGE) && e.HasComponent(comp.C_POSITION) {
			comp_p := *e.GetComponent(comp.C_POSITION)
			var p *comp.PositionComponent = comp_p.(*comp.PositionComponent)
			
			comp_si := *e.GetComponent(comp.C_SUBIMAGE)
			var img *comp.SubImageComponent = comp_si.(*comp.SubImageComponent)

			x := img.GetCurrentFrame() * 32
			y := 0
			width := img.GetFrameWidth()
			height := img.GetFrameHeight() 

			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(float64(p.GetX()), float64(p.GetY()))
			surface.DrawImage(img.GetImage().SubImage(image.Rect(x, y, width, height)).(*ebiten.Image), drawOp)
		}
	}
}