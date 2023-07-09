package main

import (
    "fmt"
    // "math"
    "time"
)
const s string = "constant"

func main() {
    // Hello world // 
    fmt.Println("hello world")

    Values // 
    fmt.Println("go"+"lang")

    fmt.Println("1+1",1+1)

    fmt.Println(true && false)

    fmt.Println(true || false)

    fmt.Println(!true)

    // Variable // 

    var a = "initial"
    fmt.Println(a)

    var b,c int  = 1,2
    fmt.Println(b,c);

    var d = true
    fmt.Println(d)

    var e int  //initializatoin are zero valued 
    fmt.Println(e)

    f:= "hello" //syntax is shorthand for declaring and initializing a variable
    fmt.Println(f)

    Constants //
    fmt.Println(s)

    const n = 500000

    const d = 3e20 / n 
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

    For loop //
    i:=1
    for i <= 3 {
         fmt.Println(i)
         i=i+1
    }

    for j:=7; j<=9; j++ {
        fmt.Println(j)
    }    

    for {
        fmt.Println("loop")
        break
    }

    for n:=0; n<=5;n++{
        if n%2 == 0{
            continue
        }
        fmt.Println(n)
    }

    //  If/Else // 

    if 7%2 == 0{
        fmt.Println("7 is even")
    }else {
        fmt.Println("7 is odd")
    }

    	
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num:=9; num < 0{
        fmt.Println(num, "is negative")
    }else if num < 10 {
        fmt.Println(num, "has 1 digit")
    }else {
        fmt.Println(num , "has multiple digits")
    }

    // Switch // 
    i:=2 
    fmt.Print("write ",i," as ")

    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday(){
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t:=time.Now()

    switch{
    case t.Hour() <12:
        fmt.Println("Its before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI:= func(i interface{}){
        switch t:= i.(type){
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")


}
