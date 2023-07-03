package comp

type PositionComponent struct {
	id int
	x float64
	y float64
}

func NewPositionComponent(x float64, y float64) PositionComponent {
	pc := PositionComponent{}
	pc.x = x
	pc.y = y
	return pc
}

func (pc *PositionComponent) SetPosition(x float64, y float64) {
	pc.x = x
	pc.y = y
}

func (pc *PositionComponent) SetX(x float64) {
	pc.x = x
}

func (pc *PositionComponent) SetY(y float64) {
	pc.y = y
}

func (pc *PositionComponent) GetX() float64 {
	return pc.x
}

func (pc *PositionComponent) GetY() float64 {
	return pc.y
}

func (pc *PositionComponent) GetID() int {
	return C_POSITION
}