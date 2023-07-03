package sys

import (
	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/Enigmanark/GoCoinCatcher/src/core/ecs/comp"
)

func SystemRenderSprite(entities *[]ECS.Entity, surface *ebiten.Image) {
	for i := 0; i < len(*entities); i++ {
		es := *entities
		e := es[i]
		if e.HasComponent(comp.C_IMAGE) && e.HasComponent(comp.C_POSITION) {
			comp_p := *e.GetComponent(comp.C_POSITION)
			var p *comp.PositionComponent = comp_p.(*comp.PositionComponent)
			
			comp_i := *e.GetComponent(comp.C_IMAGE)
			var i *comp.ImageComponent = comp_i.(*comp.ImageComponent)

			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(p.GetX(), p.GetY())
			surface.DrawImage(i.GetImage(), drawOp)
		}
	}
}