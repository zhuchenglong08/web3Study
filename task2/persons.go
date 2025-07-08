package main

import (
	"fmt"
)

//使用组合的方式创建一个 Person 结构体，
//包含 Name 和 Age 字段，再创建一个 Employee 结构体，
//组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	employeeID int
	depatment
}

type depatment struct {
	name          string
	depatmentName string
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d ,部门：%s,部门名称：%s\n", e.Person.name, e.age, e.employeeID, e.depatment.name, e.depatmentName)
}
func main() {
	em := Employee{Person{"张三", 18}, 1001, depatment{
		"IT", "开发部",
	}}
	em.PrintInfo()
}
