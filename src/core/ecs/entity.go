package ecs

import "github.com/Enigmanark/GoCoinCatcher/src/core/ecs/comp"

type Entity struct {
	components []comp.Component
	id string
	tag string
}

func (e *Entity) GetID() string {
	return e.id
}

func (e *Entity) GetComponents() *[]comp.Component {
	return &e.components
}

func (e *Entity) AddComponent(component comp.Component) {
	e.components = append(e.components, component)
}

func (e *Entity) HasComponent(c_id int) bool {
	for i := 0; i < len(e.components); i++ {
		if e.components[i].GetID() == c_id {
			return true
		}
	}
	return false
}

func (e *Entity) GetComponent(c_id int) *comp.Component {
	for i := 0; i < len(e.components); i++ {
		if e.components[i].GetID() == c_id {
			return &e.components[i]
		}
	}
	return nil
}