package moderator

type Moderator struct {
	salary   float32
	address  string
	position string
}

func (m *Moderator) GetSalary() float32 {
	return m.salary
}

func (m *Moderator) SetSalary(salary float32) {
	m.salary = salary
}

func (m *Moderator) GetAddress() string {
	return m.address
}

func (m *Moderator) SetAddress(address string) {
	m.address = address
}

func (m *Moderator) GetPosition() string {
	return m.position
}

func (m *Moderator) SetPosition(pos string) {
	m.position = pos
}
