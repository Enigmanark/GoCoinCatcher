package core

import (
	ECS "github.com/Enigmanark/GoCoinCatcher/src/core/ecs"
	Scenes "github.com/Enigmanark/GoCoinCatcher/src/core/scenes"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameWorld struct {
	currentScene Scene
	entityManager ECS.EntityManager
	//Maybe make a world managers list
}

func NewGameWorld() GameWorld {
	w := GameWorld{}
	w.entityManager = ECS.NewEntityManager()
	w.currentScene = Scenes.NewPlayScene(&w.entityManager)
	return w
}

func (world *GameWorld) Update(delta float32) {

}

func (world *GameWorld) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}