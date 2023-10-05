package seller

type Seller struct {
	salary   float32
	address  string
	position string
	bonus    float32
}

func (m *Seller) GetSalary() float32 {
	return m.salary
}

func (m *Seller) SetSalary(salary float32) {
	m.salary = salary
}

func (m *Seller) GetAddress() string {
	return m.address
}

func (m *Seller) SetAddress(address string) {
	m.address = address
}

func (m *Seller) GetPosition() string {
	return m.position
}

func (m *Seller) SetPosition(pos string) {
	m.position = pos
}

func (m *Seller) GetBonus(bonus float32) {
	m.bonus = bonus
}

func (m *Seller) SellProduct(productPrice float32) {
	m.bonus += productPrice / 20
}
