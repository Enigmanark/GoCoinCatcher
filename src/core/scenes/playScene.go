package scenes

import (
	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type playScene struct {
	entities []ECS.Entity
	entityManager *ECS.EntityManager
}

func NewPlayScene(em *ECS.EntityManager) *playScene {
	p := playScene{}
	p.entityManager = em

	return &p
}

func (p *playScene) Load() {

}

func (p *playScene) Unload() {

}

func (p *playScene) Update(delta float32) {

}

func (p *playScene) Draw(screen *ebiten.Image) {

}