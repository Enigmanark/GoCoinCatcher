package scenes

import (
	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	"github.com/Enigmanark/GoCoinCatcher/src/core/ecs/comp"
	"github.com/Enigmanark/GoCoinCatcher/src/core/ecs/sys"
	"github.com/hajimehoshi/ebiten/v2"
)

type playScene struct {
	entities []ECS.Entity
	entityHelper *ECS.EntityHelper
}

func NewPlayScene(eh *ECS.EntityHelper) *playScene {
	p := playScene{}
	p.entityHelper = eh

	return &p
}

func (p *playScene) Load() {
	player := p.entityHelper.NewEntity(&p.entities)

	playerPos := comp.NewPositionComponent(10, 10)
	player.AddComponent(&playerPos)

	playerImage := comp.NewSubImageComponent("res/img/hero.png", 32, 32, 1)
	player.AddComponent(&playerImage)

	playerVelocity := comp.NewVelocityComponent(1, 1)
	playerVelocity.SetSpeed(50)
	player.AddComponent(&playerVelocity)

	p.entities = append(p.entities, player)

}

func (p *playScene) Unload() {

}

func (p *playScene) Update(delta float64) {
	sys.SystemUpdateMovers(p.entities, delta)
}

func (p *playScene) Draw(surface *ebiten.Image) {
	sys.SystemRenderSprite(p.entities, surface)
}