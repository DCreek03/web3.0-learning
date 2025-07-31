package main

import (
	"fmt"
	"math"
)

// 1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，
// 并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Redius float64
}

func (c *Circle) Area() float64 {
	return c.Redius * c.Redius * math.Pi
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Redius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// 2. 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	// per        Person
	// 使用匿名更好
	Person
	EmployeeId string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("员工id：%s, 员工年龄：%d, 员工姓名：%s", e.EmployeeId, e.Age, e.Name)
}

func main() {
	r := &Rectangle{Width: 5, Height: 3}
	c := &Circle{Redius: 5}
	fmt.Printf("圆形的面积是：%f, 周长是：%f\n", c.Area(), c.Perimeter())
	fmt.Printf("长方形的面积是：%f, 周长是：%f\n", r.Area(), r.Perimeter())

	// 这里的Person 作为字段名使用 也只能用Person
	e := &Employee{EmployeeId: "123", Person: Person{Name: "Aloha", Age: 11}}
	e.PrintInfo()
}
