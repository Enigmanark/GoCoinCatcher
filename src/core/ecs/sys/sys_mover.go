package sys

import (
	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	"github.com/Enigmanark/GoCoinCatcher/src/core/ecs/comp"
)

func SystemUpdateMovers(entities []ECS.Entity, delta float64) {
	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if e.HasComponent(comp.C_POSITION) && e.HasComponent(comp.C_VELOCITY) {
			p_com := *e.GetComponent(comp.C_POSITION)
			var p *comp.PositionComponent = p_com.(*comp.PositionComponent)

			v_com := *e.GetComponent(comp.C_VELOCITY)
			var v *comp.VelocityComponent = v_com.(*comp.VelocityComponent)

			move_x := p.GetX() + (v.GetXVelocity() * v.GetSpeed() * delta) 
			move_y := p.GetY() + (v.GetYVelocity() * v.GetSpeed() * delta)
			
			p.SetX(move_x)
			p.SetY(move_y)
		}
	}
}