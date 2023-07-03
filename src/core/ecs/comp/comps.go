package comp

const (
	C_POSITION = 0
	C_VELOCITY = 1
	C_SPATIAL = 2
	C_IMAGE = 3
	C_TEXT = 4
)

type Component interface {
	GetID() int
}