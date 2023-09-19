package test

import "fmt"

type Car struct {
	Speed float64
	Name  string
}

func (c *Car) GetSpeed() float64 {
	return c.Speed
}

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("hi")
	}

	fmt.Println("Hello World!")

	var i int8 = 123
	fmt.Println(i)

	v := Car{
		Speed: 220,
		Name:  "Volvo",
	}

	fmt.Println(v.GetSpeed())
}
