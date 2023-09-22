package main

import (
	"Project/manager"
	"Project/security_guard"
	"Project/seller"
	"fmt"
)

func main() {
	s := seller.Seller{}
	m := manager.Manager{}
	sg := security_guard.SecurityGuard{}

	s.SetPosition("seller")
	m.SetPosition("manager")
	sg.SetAddress("Tole Bi 59")
	s.SetSalary(666.666)

	fmt.Println(s.GetPosition())
	fmt.Println(m.GetPosition())
	fmt.Println(sg.GetAddress())
	fmt.Println(s.GetSalary())
}
