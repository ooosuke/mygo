package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}

func calc(x int, y int) (add int) {
	add = x + y

	return
}

func calc1(x int, y int) (add int, kakeru int) {
	add = x + y
	kakeru = x * y

	return
}

var hello1 = func(greeting string) {
	fmt.Println(greeting)
}

type Student struct {
	name string
	age int
}

func (s Student) print() {
	fmt.Println(s.name, s.age)
}

func main() {
	sayHello()
	fmt.Println(calc(1, 3))
	fmt.Println(calc1(1, 3))
	hello1("hgoehoge")

	var s Student
	s.name = "sato"
	s.age = 20

	fmt.Println(s.name, s.age)

	ss := Student{name: "kato", age: 22}
	fmt.Println(ss.name, ss.age)

	sss := Student{name: "taro", age: 23}
	sss.print()
}