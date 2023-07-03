package ecs
//Write an entity manager that combines random letters and numbers as an ID
//then check to make sure no other has that ID that currentlye exists

type EntityManager struct {

}

func NewEntityManager() EntityManager {
	em := EntityManager{}
	return em
}