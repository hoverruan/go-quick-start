package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

type Student struct {
	People
	Class string
}

type Worker struct {
	People
	Company    string
	Department string
}

func structDemo() {
	s1 := Student{People: People{Name: "n1", Age: 11}, Class: "32"}
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("name: %s, age: %d, class: %s\n", s1.Name, s1.Age, s1.Class)

	w1 := Worker{People: People{"n1", 30}, Company: "TW", Department: "Dev"}
	fmt.Printf("w1: %v\n", w1)
}

func (p People) sayHi() {
	fmt.Printf("%s: hi\n", p.Name)
}

func (s Student) sayHi() {
	fmt.Printf("[S] %s: hi\n", s.Name)
}

func methodDemo() {
	w1 := Worker{People: People{"Tom", 25}, Company: "TW"}
	w1.sayHi()

	s1 := Student{People: People{"Petter", 12}, Class: "12"}
	s1.sayHi()
}

func NewWorker(name string, age int, company string) (w *Worker) {
	w = new(Worker)
	w.Name = name
	w.Age = age
	w.Company = company

	return
}

func constructorDemo() {
	w1 := NewWorker("John", 19, "TW")
	w1.sayHi()
}

type Speaker interface {
	Speak(something string)
}

func (p People) Speak(something string) {
	fmt.Printf("%s: %s\n", p.Name, something)
}

func saySomething(sk Speaker, something string) {
	sk.Speak(something)
}

func interfaceDemo() {
	w1 := NewWorker("Tim", 29, "SC")
	w1.Speak("Hi guys")

	saySomething(w1, "Welcome back")
}

func main() {
	structDemo()
	methodDemo()
	constructorDemo()
	interfaceDemo()
}
