package main

// "fmt"
// "math"
// "time"

const s string = "constant"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
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

}
