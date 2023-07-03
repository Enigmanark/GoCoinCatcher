package ecs

import (
	"math/rand"
	"strconv"
	"time"
)

//Write an entity manager that combines random letters and numbers as an ID
//then check to make sure no other has that ID that currentlye exists

type EntityHelper struct {

}

func NewEntityHelper() EntityHelper {
	em := EntityHelper{}
	return em
}

func (eh *EntityHelper) NewEntity(entities *[]Entity) Entity {
	e := Entity{}
	e.id = generateRandomID()
	for doesIDExist(&e, entities) {
		e.id = generateRandomID()
		println("Message: Generating new ID for Entity, becauese copy already existed.")
	}
	return e
}

func doesIDExist(entity *Entity, otherEntities *[]Entity) bool {
	for i := 0; i < len(*otherEntities); i++ {
		ents := *otherEntities
		e := ents[i]
		if e.GetID() == entity.GetID() {
			return true
		}
	}
	return false
}

func generateRandomID() string {
	rand.Seed(time.Now().UnixNano())

	id := ""
	counter := 0

	for counter < 75 {
		random := rand.Intn(3-1) + 1

		//If 1 generate a random number between 0 and 9
		if random == 1 {
			random = rand.Intn(9-0) + 0
			id += strconv.Itoa(random)
		} else {
			//else generate a random letter
			random = rand.Intn(27-1) + 1
			if random == 1 {
				id += "A"
			} else if random == 2 {
				id += "B"
			} else if random == 3 {
				id += "C"
			} else if random == 4 {
				id += "D"
			} else if random == 5 {
				id += "E"
			} else if random == 6 {
				id += "F"
			} else if random == 7 {
				id += "G"
			} else if random == 8 {
				id += "H"
			} else if random == 9 {
				id += "I"
			} else if random == 10 {
				id += "J"
			} else if random == 11 {
				id += "K"
			} else if random == 12 {
				id += "L"
			} else if random == 13 {
				id += "M"
			} else if random == 14 {
				id += "N"
			} else if random == 15 {
				id += "O"
			} else if random == 16 {
				id += "P"
			} else if random == 17 {
				id += "Q"
			} else if random == 18 {
				id += "R"
			} else if random == 19 {
				id += "S"
			} else if random == 20 {
				id += "T"
			} else if random == 21 {
				id += "U"
			} else if random == 22 {
				id += "V"
			} else if random == 23 {
				id += "W"
			} else if random == 24 {
				id += "X"
			} else if random == 25 {
				id += "Y"
			} else if random == 26 {
				id += "Z"
			}
		}
		counter += 1	
	}

	return id
}