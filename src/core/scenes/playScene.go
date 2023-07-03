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

	playerImage := comp.NewImageComponent("res/img/hero.png", 1)
	player.AddComponent(&playerImage)

	p.entities = append(p.entities, player)

}

func (p *playScene) Unload() {

}

func (p *playScene) Update(delta float32) {

}

func (p *playScene) Draw(surface *ebiten.Image) {
	sys.SystemRenderSprite(&p.entities, surface)
}