package main

import (
	"fmt"
	"time"
)

// "unicode/utf8"

// "fmt"
// "math"
// "time"

// //  const
// const s string = "constant"

// // // Functions
// func plus(a int, b int) int {
// 	return a + b
// }

// func plusPlus(a, b, c int) int {
// 	return a + b + c
// }

// // // Multiple Return Values
// func vals() (int, int) {
// 	return 3, 7
// }

// // // Variadic Functions
// func sum(nums ...int) {
// 	fmt.Print(nums, " ")
// 	total := 0

// 	for _, num := range nums {
// 		total += num
// 	}
// 	fmt.Println(total)
// }

// // // Closures
// func intSeq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// // // Rescursion
// func fact(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * fact(n-1)
// }

// // // Pointers
// func zeroval(ival int) {
// 	ival = 0
// }

// func zeroptr(iptr *int) {
// 	*iptr = 0
// }

// // // Strings and Runes
// func examineRune(r rune) {
// 	if r == 't' {
// 		fmt.Println("found tee")
// 	} else if r == 'ส' {
// 		fmt.Println("found so sua")
// 	}
// }

// // // Structs
// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string) *person { // newPerson constructs a new person struct with the given name.
// 	p := person{name: name}
// 	p.age = 52
// 	// You can safely return a pointer to local variable as a local variable will survive the scope of the function.
// 	return &p
// }

// // // Methods
// type react struct {
// 	width, height int
// }

// func (r *react) area() int {
// 	return r.width * r.height
// }

// func (r react) perim() int {
// 	return 2*r.width + 2*r.height
// }

//  // // Interfaces

// type geometry interface {
// 	area() float64
// 	perim() float64
// }

// type rect struct {
// 	width, height float64
// }

// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }

// func (r rect) perim() float64 {
// 	return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
// 	fmt.Println(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// // // // Struct Embedding
// type base struct {
// 	num int
// }

// func (b base) describer() string {
// 	return fmt.Sprintf("base with num=%v", b.num)
// }

// type container struct {
// 	base
// 	str string
// }

// func (c container) describer() string {
// 	return fmt.Sprintf("container with num=%v, str=%v", c.num, c.str)
// }

// // //  Generics //  //
// func MapKeys[K comparable, V any](m map[K]V) []K {
// 	r := make([]K, 0, len(m))
// 	for k := range m {
// 		r = append(r, k)
// 	}
// 	return r
// }

// type List[T any] struct {
// 	head, tail *element[T]
// }
// type element[T any] struct {
// 	next *element[T]
// 	val  T
// }

// func (lst *List[T]) Push(v T) {
// 	if lst.tail == nil {
// 		lst.head = &element[T]{val: v}
// 		lst.tail = lst.head
// 	} else {
// 		lst.tail.next = &element[T]{val: v}
// 		lst.tail = lst.tail.next
// 	}
// }
// func (lst *List[T]) GetAll() []T {
// 	var elems []T
// 	for e := lst.head; e != nil; e = e.next {
// 		elems = append(elems, e.val)
// 	}
// 	return elems
// }

// // //  Errors //  //
// func f1(arg int) (int, error) {
// 	if arg == 42 {
// 		return -1, errors.New("cant' work with 42")
// 	}
// 	return arg + 3, nil
// }

// type argError struct {
// 	arg  int
// 	prob string
// }

// func (e *argError) Error() string {
// 	return fmt.Sprintf("%d - %s", e.arg, e.prob)
// }

// func f2(arg int) (int, error) {
// 	if arg == 42 {
// 		return -1, &argError{arg, "cant work with it"}
// 	}
// 	return arg + 3, nil
// }

// // // Goroutines // //
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//// // // Hello world //////////////
	// fmt.Println("hello world")

	//// // // Values ///////////
	// fmt.Println("go" + "lang")

	// fmt.Println("1+1", 1+1)

	// fmt.Println(true && false)

	// fmt.Println(true || false)

	// fmt.Println(!true)

	//// // // Variable ///////////

	// var a = "initial"
	// fmt.Println(a)

	// var b, c int = 1, 2
	// fmt.Println(b, c)

	// var g = true
	// fmt.Println(g)

	// var e int //initializatoin are zero valued
	// fmt.Println(e)

	// f := "hello" //syntax is shorthand for declaring and initializing a variable
	// fmt.Println(f)

	//// // // Constants ////////////
	// fmt.Println(s)

	// const n = 500000

	// const d = 3e20 / n
	// fmt.Println(d)

	// fmt.Println(int64(d))

	// fmt.Println(math.Sin(n))

	//// // // For loop ////////////
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	//// // //  If/Else //////////

	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	//// // // Switch /////////////
	// i = 2
	// fmt.Print("write ", i, " as ")

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

	// t := time.Now()

	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Its before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }

	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	//// //  Arrays ///////////////
	// Create array of 5 values
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])
	// fmt.Println("len:", len(a))

	// // Create array with values
	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(b)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	//// // Slices /////////
	// var s []string
	// fmt.Println("uninit: ", s, s == nil, len(s) == 0) //will return true

	// s = make([]string, 3)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", (cap(s))) // print len and capacity

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("len:", len(s))
	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("apd", s)

	// c := make([]string, len(s))
	// copy(c, s)

	// fmt.Println("cpy:", c)

	// l := s[2:5]

	// fmt.Println("sl1", l)

	// l = s[:5]
	// fmt.Println("sl2", l)

	// l = s[2:]
	// fmt.Println("sl3", l)

	// t := []string{"g", "h", "i"}
	// fmt.Println("dcl:", t)

	// twoD := make([][]int, 3)
	// for i := 0; i < 3; i++ {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2D:", twoD)

	//// // Maps //////
	// m := make(map[string]int)
	// m["k1"] = 7
	// m["k2"] = 8

	// fmt.Println("map: ", m)

	// v1 := m["k1"]
	// fmt.Println("v1", v1)

	// v3 := m["k3"]
	// fmt.Println("v3", v3)

	// fmt.Println("len:", len(m))

	// delete(m, "k2")
	// fmt.Println("map:", m)

	// _, prs := m["k2"] // key was present in the map
	// fmt.Println("prs", prs)

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println(n)

	//// // Range //////

	// nums := []int{2, 3, 4}
	// sum := 0

	// for _, num := range nums {
	// 	sum += num
	// }

	// fmt.Println("Sum:", sum)

	// for i, num := range nums {
	// 	if num == 3 {
	// 		fmt.Println("index:", i)
	// 	}
	// }

	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n ", k, v)
	// }

	// for k := range kvs {
	// 	fmt.Println("key:", k)
	// }

	// for i, c := range "go" {
	// 	fmt.Println(i, c)
	// }

	//// // Functions //////

	// res := plus(1, 2)
	// fmt.Println("1+2=", res)

	// res2 := plusPlus(1, 2, 3)
	// fmt.Println("1+2+3", res2)

	// // // Multiple Return Values // //

	// a, b := vals() // make function call with mulitple assignment's
	// fmt.Println(a)
	// fmt.Println(b)

	// _, c := vals()
	// fmt.Println(c)

	// // // Variadic function // //

	// sum(1, 2) // passing to function
	// sum(1, 2, 3)

	// nums := []int{1, 2, 3, 4} // passing array to function
	// sum(nums...)

	// // // Closuers  // //
	// nextInt := intSeq()

	// fmt.Println(nextInt()) // 1
	// fmt.Println(nextInt()) // 2
	// fmt.Println(nextInt()) // 3

	// newInts := intSeq()
	// fmt.Println(newInts()) //1

	// // // Recursion  // //

	// fmt.Println(fact(7))

	// var fib func(n int) int // Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.

	// fib = func(n int) int {
	// 	if n < 2 {
	// 		return n
	// 	}
	// 	return fib(n-1) + fib(n-2)
	// }
	// fmt.Println(fib(7))

	// // // Pointers // //

	// i := 1
	// fmt.Println("initial:", i)

	// zeroval(i)
	// fmt.Println("zeroval:", i)

	// zeroptr(&i)
	// fmt.Println("zeroptr:", i)

	// fmt.Println("pointer:", &i)

	// // // Strings and Runes // //

	// const s = "สวัสดี" //rune in Go is a data type that stores codes that represent Unicode characters

	// fmt.Println("Len:", len(s)) //len

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%x ", s[i])
	// }

	// fmt.Println()
	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// for idx, runeValue := range s {
	// 	fmt.Printf("%#U starts at %d\n", runeValue, idx)
	// }

	// fmt.Println("\nUsing DecodeRuneInString")
	// for i, w := 0, 0; i < len(s); i += w {
	//     runeValue, width := utf8.DecodeRuneInString(s[i:])
	//     fmt.Printf("%#U starts at %d\n", runeValue, i)
	//     w = width
	// 	examineRune(runeValue)
	// }

	// // // Structs  // //

	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{name: "Alice", age: 30})

	// fmt.Println(person{name: "Fred"})

	// fmt.Println(&person{name: "Ann", age: 40})

	// fmt.Println(newPerson("Jon"))

	// s := person{name: "Sean", age: 50}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp.age)

	// sp.age = 51
	// fmt.Println(sp.age)

	// dog := struct {
	// 	name   string
	// 	isGood bool
	// }{
	// 	"Rex",
	// 	true,
	// }

	// fmt.Println(dog)

	// // // Methods   // //

	// r := react{width: 10, height: 5}
	// fmt.Println("Area:", r.area())
	// fmt.Println("Perim:", r.perim())

	// rp := &r
	// fmt.Println("area:", rp.area())
	// fmt.Println("perim:", rp.perim())

	// // // Interfaces // //

	// r := rect{width: 3, height: 4}
	// c := circle{radius: 5}
	// measure(r)
	// measure(c)

	// // //  Struct Embedding //  //
	// co := container{
	// 	base: base{
	// 		num: 1,
	// 	},
	// 	str: "some name",
	// }
	// fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// fmt.Println("also num:", co.base.num)

	// fmt.Println("describe:", co.describer())

	// type describer interface {
	// 	describer() string
	// }

	// var d describer = co
	// fmt.Println("describer:", d.describer())

	// // //  Generics //  //
	// var m = map[int]string{1: "2", 2: "4", 4: "8"}
	// fmt.Println("keys", MapKeys(m))

	// _ = MapKeys[int, string](m)

	// lst := List[int]{}
	// lst.Push(10)
	// lst.Push(13)
	// lst.Push(23)
	// fmt.Println("list:", lst.GetAll())

	// // //  Errors //  //

	// for _, i := range []int{7, 42} {
	// 	if r, e := f1(i); e != nil {
	// 		fmt.Println("f1 failed:", e)
	// 	} else {
	// 		fmt.Println("f1 worked:", r)
	// 	}
	// }

	// for _, i := range []int{7, 42} {
	// 	if r, e := f2(i); e != nil {
	// 		fmt.Println("f2 failed:", e)
	// 	} else {
	// 		fmt.Println("f2 worked:", r)
	// 	}
	// }

	// _, e := f2(42)
	// if ae, ok := e.(*argError); ok {
	// 	fmt.Println(ae.arg)
	// 	fmt.Println(ae.prob)
	// }

	// // //  Goroutines //  //

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
