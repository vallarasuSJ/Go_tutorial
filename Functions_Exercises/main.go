package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {

	x := sum([]int{1, 2, 34, 4, 3, 4})
	fmt.Println("The total value is ", x)

	x1 := foo()
	fmt.Println("The value of foo is ", x1)

	x2, x3 := bar()
	fmt.Println(x2, " ", x3)

	x4 := []int{3, 4, 3, 4, 3, 4, 34}
	fmt.Println(foot(x4...))
	fmt.Println(bars(x4))

	//exercise 4- defer function
	defer fmt.Println("The defer func runs in LIFO", 3)
	defer fmt.Println("The defer func runs in LIFO", 2)
	defer fmt.Println("The defer func runs in LIFO", 1)

	p1 := Person{
		name: "kv",
		age:  21,
	}
	p1.Speak()

	c1 := Circle{
		radius: 4,
	}

	c2 := Square{
		length: 4,
		width:  4,
	}
	Info(c1)
	Info(c2)

	a1 := Add(3, 5)
	fmt.Println(a1)

	m1 := doMath(534, 32, Addition)
	m2 := doMath(43, 23, Subtraction)
	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println(Paradise("Hawai"))

	db := MockDataStore{
		Users: make(map[int]User),
	}

	serv := Service{
		ds: db,
	}

	u1 := User{
		Id:   1,
		name: "kv",
	}
	err := serv.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := serv.GetUser(u1.Id)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println(u1)
	fmt.Println(u1Returned)

	//exercise-11 anonymous function
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println("The output is ", i)
		}
	}()

	//func expression
	n1()

	//return function
	f1 := outer()
	fmt.Println("Return function returns", f1())

	fmt.Println(PrintSquare(Square1, 3))

	x5 := powInator(2)
	fmt.Println(x5())

	timeFunc(doWork)

}



//function sum with total return as a argument
// func sum(ii []int)(total int){
// 	total=0
// 	for _,v:=range ii{
// 		total+=v
// 	}
// 	return
// }

// function which returns total
func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

// exercise 2
func foo() int {
	return 42
}
func bar() (int, string) {
	return 43, "Vallarasu"
}

// exercise 3 - variadic func
func foot(i ...int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
}

func bars(i []int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
}

// exercise 5 - method
type Person struct {
	name string
	age  int
}

func (p Person) Speak() {
	fmt.Println("My name is ", p.name, " and my age is  ", p.age)
}

// exercise -6 interfaces
type Square struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.length * s.width
}
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Shape interface {
	area() float64
}

func Info(s Shape) {
	fmt.Printf("The shape of area is %.2f\n", s.area())
}

// exercise 7 - test in go
func Add(a int, b int) int {
	return a + b
}

// exercise 8- unit test in go
func Addition(a int, b int) int {
	return a + b
}

func Subtraction(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// exercise 9- unit test part 2 in go
func Paradise(location string) string {
	return fmt.Sprint("My idea of paradise is ", location)
}

//exercise -10  interface and mock testing a db
type User struct {
	Id   int
	name string
}

type MockDataStore struct {
	Users map[int]User
}

func (m MockDataStore) GetUser(id int) (User, error) {
	user, ok := m.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}
func (m MockDataStore) SaveUser(u User) error {
	_, ok := m.Users[u.Id] //if its present then ok is true
	if ok {
		return fmt.Errorf("User %v already exists ", u.Id)

	}
	m.Users[u.Id] = u
	return nil
}

type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

// exercise -12 function expression
var n1 = func() {
	fmt.Println("Function expression")
}

//exercise-13 return a function

func outer() func() int {
	return func() int {
		return 42
	}
}

// exercise-14 callback
func Square1(n int) int {
	return n * n
}
func PrintSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("The number %v squared is %v", a, x)
}

// exercise-15 closure
func powInator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

// exercise-16  wrapper function
func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}
