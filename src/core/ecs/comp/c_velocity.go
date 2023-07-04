package comp

type VelocityComponent struct {
	x_velocity float64
	y_velocity float64
	speed float64
}

func NewVelocityComponent(velx float64, vely float64) VelocityComponent {
	v := VelocityComponent{}
	v.x_velocity = velx
	v.y_velocity = vely
	return v
}

func (vel *VelocityComponent) SetSpeed(sp float64) {
	vel.speed = sp
}

func (vel *VelocityComponent) GetSpeed() float64 {
	return vel.speed
}

func (vel *VelocityComponent) GetXVelocity() float64 {
	return vel.x_velocity
}

func (vel *VelocityComponent) SetXVelocity(v float64) {
	vel.x_velocity = v
}

func (vel *VelocityComponent) GetYVelocity() float64 {
	return vel.y_velocity
}

func (vel *VelocityComponent) SetYVelocity(v float64) {
	vel.y_velocity = v
}

func (vel *VelocityComponent) GetID() int {
	return C_VELOCITY
}