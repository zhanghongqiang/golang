package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{Width: 5, Height: 3}

	var shape Shape

	shape = rect

	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())

	circ := Circle{Radius: 4}

	shape = circ

	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())

	e := Employee{EmployeeID: 3,
		Person: Person{Name: "Lisu", Age: 25}}
	e.PrintInfo()

}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Person struct {
	Name string
	Age  int64
}

type Employee struct {
	EmployeeID int64
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 工号: %d\n", e.Name, e.Age, e.EmployeeID)
}
