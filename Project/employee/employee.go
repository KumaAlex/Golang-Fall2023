package employee

type Employee interface {
	GetPosition() string
	GetSalary() float32
	GetAddress() string
	SetPosition(position string)
	SetSalary(salary float32)
	SetAddress(address string)
}
