package security_guard

type SecurityGuard struct {
	salary   float32
	address  string
	position string
}

func (m *SecurityGuard) GetSalary() float32 {
	return m.salary
}

func (m *SecurityGuard) SetSalary(salary float32) {
	m.salary = salary
}

func (m *SecurityGuard) GetAddress() string {
	return m.address
}

func (m *SecurityGuard) SetAddress(address string) {
	m.address = address
}

func (m *SecurityGuard) GetPosition() string {
	return m.position
}

func (m *SecurityGuard) SetPosition(pos string) {
	m.position = pos
}
