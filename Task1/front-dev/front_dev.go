package front_dev

type Front_dev struct {
	salary   float32
	address  string
	position string
}

func (m *Front_dev) GetSalary() float32 {
	return m.salary
}

func (m *Front_dev) SetSalary(salary float32) {
	m.salary = salary
}

func (m *Front_dev) GetAddress() string {
	return m.address
}

func (m *Front_dev) SetAddress(address string) {
	m.address = address
}

func (m *Front_dev) GetPosition() string {
	return m.position
}

func (m *Front_dev) SetPosition(pos string) {
	m.position = pos
}
