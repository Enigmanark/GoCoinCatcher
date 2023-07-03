package core

import (
	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	"github.com/Enigmanark/GoCoinCatcher/src/core/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameWorld struct {
	currentScene Scene
	entityManager ECS.EntityHelper
	//Maybe make a world managers list
}

func NewGameWorld() GameWorld {
	w := GameWorld{}
	w.entityManager = ECS.NewEntityHelper()
	w.currentScene = scenes.NewPlayScene(&w.entityManager)
	w.currentScene.Load()
	return w
}

func (world *GameWorld) Update(delta float32) {

}

func (world *GameWorld) Draw(surface *ebiten.Image) {
	world.currentScene.Draw(surface)
}