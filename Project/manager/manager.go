package manager

type Manager struct {
	salary   float32
	address  string
	position string
}

func (m *Manager) GetSalary() float32 {
	return m.salary
}

func (m *Manager) SetSalary(salary float32) {
	m.salary = salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(pos string) {
	m.position = pos
}
