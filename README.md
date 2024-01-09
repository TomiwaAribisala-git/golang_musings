### Install Go
- [Install GO](https://go.dev/doc/install)

### Visual Studio Code

### VS Code Go Extension

### Create Working Directory and Go file
```sh
mkdir go_tutorial
```
```sh
touch main.go
```

### Go Documentation
[Golang Documentation](https://go.dev/doc/)

### Go Packages
[Golang Standard Library](https://pkg.go.dev/std)
[Golang Packages](https://pkg.go.dev/)

### Go Modules
- A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The go.mod file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.
- Go modules are group of related projects versioned and distributed together, they specify the requirements of a project, list all the required dependencies, and help us keep track of the specific versions of installed dependencies
- Make the directory of a Go project a module using `go mod init (module path)`
- Go Commands 
```sh 
go list -m all # print the current module dependencies
```
```go
import "module"
```
```sh
go get (module)
```

```sh
go mod tidy # removes unused dependencies
```
- Adding a package/subdirectory(mypackage) to a module/root directory(mymodule)
```sh
mkdir mypackage
```
```sh
cd mypackage
nano mypackage.go
```go
package mypackage

import "fmt"

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}
```
```go
package main

import (
	"fmt"

	"mymodule/mypackage"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
}
```
- Adding a specific module version 
Since Go modules are distributed from a version control repository, they can use version control features such as tags, branches, and even commits. You can reference these in your dependencies using the @ symbol at the end of the module path along with the version you’d like to use.
```sh
go get github.com/spf13/cobra@v1.1.1
```

### Initialize Working Directory
```sh
go mod init <module path>
```
Module path can correspond to a repository you plan to publish your module.

### Go Commands
- Go Build
```sh
go build helloworld.go
```
```sh
./helloworld
```
- Go Run
```sh
go run helloworld.go
```
```sh
go run .
```
- Go Format
```sh
gofmt
```

### Notes
- Your Go file must belong to a package. the first statement in a Go file must be `package ...`, each source file begins with a `package` declaration--which states the package the file belongs to, then followed by a list of other packages the file imports
- The `import` declaration must follow the `package` declaration
- The `func main()` is the typical entrypoint of a Go Program
- A program can only have one `main function`, because you can only have one entrypoint
- Declarations: `var`, `const`, `type`, `func` 

### Go Packages
- Modularize your application
- Go programs are organized into packages, a package is a collection of Go files
- Variables and functions defined outside any function, can be accessed in all other files within the same package
- Execute all files in a package
```sh
go run .
```
- When working with multiple packages, create a directory/file as regards each package, you can also import packages between each other eg. `import <module_path>/package` then `package.(any function in the helper package)`
- For importing functions between packages, make sure to capitalize the first letter of the function, this action exports the function and makes it importable in another package
- You can also export variables, functions, constants, types; make sure to capitalize the first letter of the data type to make it importable

### Calling code in an external package
- Import package and use function in the package; packages are published in modules
```go
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```
- Add new module and requirements
```sh
go mod tidy
```
- Run code 
```sh
go run .
```

### Data Types
- Type Declarations
```go
import "fmt"
type Celsius float64
type Fahrenheit float64
const (
AbsoluteZeroC Celsius = -273.15
FreezingC Celsius = 0
BoilingC Celsius = 100
)
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
- Go binary operators
```go
* / % << >> & &^
+ - | ^
== != < <= > >= +=
&&
||
```
- Types
- Integers: `int`, `unit`
- Floating point numbers: `float32`, `float64`
- Complex Numbers: `complex64`, `complex128` 
- Booleans: `bool`, `true`, `false`
- Strings: packages for manipulating strings--`strings`, `strconv`, `bytes`, `unicode`

### Variables and Constants
- `isValidEmail := tomiwaaribisala@gmail.com`   ## short variable declarations for initializing local variables
- `isValidName := len(firstName) >= 2 || len(surName) >= 10` ## short variable declarations for initializing local variables
- `isValidTC := boughtTickets > 0 && leftTickets <= 50` ## short variable declarations for initializing local variables
- A short variable description can be used to call a function such as the example below where os.Open returns two values
```go
f, err := os.Open(name)
if err != nil {
return err
}
f.Close()
```
- A short variable declaration does not neccessarily declare all the variables of the left hand side, in the code below, the first statement declares both `in` and `err`. The second statement declares `out` but only assigns a value to the existing `err` variable.
```go
in, err := os.Open(infile)
// ...
out, err := os.Create(outfile)
```
- `var name string = "tomiwa"`, `var name type = expression`, either the type or the = expression part may be omitted, but not both.
- `i, j := 0, 1`
- `var i, j, k init`
- `var name = "tomiwa"`
- `var name string`, a variable `type` determines the characteristics of values it can take on, and the type of intrinsic operations that can be done on those values
- `const ticketPrice = $70`
- `s`, `s++`, `s--`, `s +=`
- Printing formatted data in Go using `Println`; `fmt.Println("My name is", name, " and my ticket price is", ticketPrice, "dollars")`
- Printing formatted data in Go using `printf`; `fmt.printf("The ticket price is %v.\n, ticketPrice")`

### Local Variables
- Defined inside a function or block; `a := 7`
- Local variables can be accessed only inside that function or block of code
- Best practice is to define variable as local as possible, create the variable where you need it 

### Package Level Variables
- Defined at the top outside all functions
- They can be accessed inside any of the functions, and in all files/functions which are in the same `package ...`

### Global Variables
- Variables denoted by uppercase first letter, declaration outside all functions
- Global variables can be used everywhere across all packages 

### Assignments 
- The value held by a variable can be updated by an assignment statement, which in its simplest form has a variable on the left of the = sign and an expression on the right

### Data Types
Go is a statically typed language, you need to tell the Go compiler the data type when declaring a variable; however, in most cases, Go can infer the data type of a variable when you assign a value.
- Strings   eg. `var UserName string`
- Integers  eg. `var UserTickets int`, `var UserTickets uint`
- Floats
- Booleans  
- Complex Numbers
- Arrays  
- Slices
- Maps          
- Pointers
- Structs
- Methods
- Interfaces
- Closures
- Channels
- Goroutines

### User Input 
- Using `fmt` package to collect user input; 
```go
var userName string
fmt.Println("Please enter your first name: ")
fmt.Scan(&userName)
```

### Methods
- Object Oriented Programming
### Pointers
- Pointers are special variables that points to the memory address of another variable(value)
```go
var matchPrice = 50
fmt.Println(&matchPrice) // &matchPrice is the memory address of matchPrice

p := &matchPrice   
*p = 70 // this expression sets the value of the matchPrice value to 70
fmt.Println(matchPrice)

z := new(int)
fmt.Println(*z)
*z = 2
fmt.Println(*z)
```

### Arrays and Slices 
- Data structures to store collection of elements in a single variable, index-based
- Arrays
```go
var bookings [50]string // an array of 50 values of type string
a := [10]int{}
var firstName string
var lastName string
fmt.Println("Please enter your first name: ")
fmt.Scan(&firstName)
fmt.Println("Please enter your last name: ")
fmt.Scan(&lastName)
bookings[0] = firstName + " " + lastName
fmt.Println("Welcome", firstName, lastName)
```
- Slices: an abstraction of an array, more flexible and powerful--variable-length or get a sub-array of its own, slices have a size like an array, but can be resized when needed. 
```go
a := []int{}
printSlice(a)
var bookings []string
var firstName string
var lastName string
fmt.Println("Please enter your first name: ")
fmt.Scan(&firstName)
fmt.Println("Please enter your last name: ")
fmt.Scan(&lastName)
bookings = append(bookings, firstName + " " + lastName) // the append function
fmt.Println("Thank you", firstName, lastName, "for inputting your name!")
fmt.Println("Print all our bookings:", bookings)
```
- A slice can be formed by specifying two indices, a low and high bound, separated by a colon, `a[1:4]`, describing a section of an underlying array
- Changing the elements of a slice modifies the corresponding elements of its underlying array
- Integrating a Slice into a function
- Creating a slice with the `make()` function
```go
slice_name := make([]type, length, capacity)
```
- Length of a Slice 
- Capcacity of a Slice

### Maps
- Maps maps unique keys to values, you can retrieve the value later by using its key
- All keys and values have the same data type, map supports only one data type at once
```go
map[K]V
var userData = []map[string]string 
userData["firstname"] = firstname
userData["lastname"] = lastname
userData["email"] = email

delete(userData, "firstname") // deleting an element from a map

var timeZone = map[string]int{
    "dennis": 42
    "carly": 21
}
dennisAge := timeZone["dennis"]
carlyAge := timeZone["carly"]

j := map[string]int{
    "shaggy": 12
    "froy": 15
}

m := make([]map[string]int)
m["Question"] = 21
m["Answer"] = 42
// The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.
```
- Adding/Updating/Retrieving elements in a map
- Deleting elements in a map: `delete(map, key_name)`
- Checking for specific elements in a map 
```go
val, ok := map_name[key_name]
fmt.Println(val, ok)

// To check the existence of a certain key, use the blank identifier
_, ok1 := map_name[key_name]
fmt.Println(ok1)
```
- Iterating over a map
```go
m := make([]map[string]int)

for i, v := range m {
    fmt.Println(i) // Print index
}
```
- Integrating a Map into a function

### Make()
- Go in-built function `make` which helps create/intialize slices, maps and channels
```go
// Initializing a slice
var intSlice = make([]int, 10)        
var strSlice = make([]string, 10, 20)

//Initializing a map
 var employee = make(map[string]int)
employee["Mark"] = 10
employee["Sandy"] = 20

// Initializing a channel
channelName := make(chan int)
```

### Structs
- Struct is a data structure that allows us to mix different data types, defines a structure(which fields) of the User Type
```go
type UserData struct {
    firstName string
    lastName string
    email string
    numberofTickets int
}

tomiwainfo := UserData{"tomiwa", "sala", "tm@Gmail.com", 10}
fmt.Println(tomiwainfo)

var dennisinfo = UserData {
    firstName : "dennis",
    lastName: "roy", 
    email: "royd@gmail.com",
    numberofTickets: 1,
}
fmt.Println(dennisinfo)

// Access struct fields with a dot
fmt.Println(dennisinfo.firstName)
fmt.Println(dennisinfo.lastName)

// Pointers to structs
p := &dennisinfo
p.firstName = "dane" 
p.lastName = "ova"
fmt.Println(dennisinfo)
```

### Interfaces 
- Interfaces are named collections of method signatures
- [Interfaces](https://gobyexample.com/interfaces)
```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```
```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}
// In Go, when you see `*T` as a receiver for a method, it means that the method is associated with a pointer to an instance of the T type, `(t *T)` is the receiver of the method M(). It specifies the type on which the method operates. In this case, the receiver is a pointer to an instance of the T type.
// When a method has a pointer receiver, it means that the method can modify the fields of the receiver, and changes made to the receiver inside the method will affect the original instance of the type.
// For example, if you have an instance of T:
instance := T{"Hello, World!"}
// You can call the method M() on this instance:
instance.M()
// However, since M() has a pointer receiver, Go will automatically take the address of instance for you. So, the call is effectively the same as if you did:
(&instance).M()
// This allows the method to modify the fields of the original instance if needed.

In Go, *T represents a pointer to a type T. Let me break it down:

T: This is a type, and *T indicates a pointer to that type.
*: The asterisk * is the pointer-indirection operator in Go.
For example, if you have a type T:

type T struct {
    // fields
}
Then, *T would represent a pointer to this type:

var tInstance T
var tPointer *T

tPointer = &tInstance // Assign the address of tInstance to tPointer
In this case, tPointer is a pointer variable that can hold the memory address of a variable of type T. The & operator is used to obtain the address of a variable.

When you see a method receiver like (t *T) in a method declaration:

func (t *T) methodName() {
    // Method implementation
}
It means that the method is associated with a pointer to an instance of type T. This allows the method to modify the fields of the original instance and not just receive a copy of the instance. The * in (t *T) denotes that the receiver is a pointer.


type F float64

func (f F) M() {
	fmt.Println(f)
}
```

```go
package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
```
### Context
- Concept of Context

### Type Assertions
- Type assertions in Golang provide access to the exact type of variable of an interface. If already the data type is present in the interface, then it will retrieve the actual data type value held by the interface.
```go
// Golang program to illustrate  
// the concept of type assertions 
package main 
  
import ( 
    "fmt"
) 
  
// main function 
func main() { 
      
    // an interface that has  
    // a string value 
    var value interface{} = "GeeksforGeeks"
      
    // retrieving a value 
    // of type string and assigning 
    // it to value1 variable 
    var value1 string = value.(string) 
      
    // printing the concrete value 
    fmt.Println(value1) 
      
    // this will panic as interface 
    // does not have int type 
    var value2 int = value.(int) 
      
    fmt.Println(value2) 
} 
```

### Goroutines(Concurrency)
- A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program. Or in other words, every concurrently executing activity in Go language is known as a Goroutines
- A `goroutine` is a lightweight thread of execution
- Executing different parts of code in each separate thread(goroutine) managed concurrently by the Go runtime e.g. function calls running asychronously in separate goroutines, both `main` and `new` goroutines work concurrently
```go
package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}
```

### Channels
- A channel is a medium through which a goroutine communicates with another goroutine of which the communication is lock-free, a channel is a technique which let one goroutine to send data to another goroutine, the process is bidrectional
- A channel only allows transfer of data of the same type, different data types are not allowed
```go
// Creating a channel
var Channel_name chan Type
channel_name:= make(chan Type)

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
```go
package main

import "fmt"

func main() {

    messages := make(chan string)
    
    // Sending data to a channel
    go func() { messages <- "ping" }()

// Receiving data from a channel
    msg := <-messages
    fmt.Println(msg)
}
```

### Buffer
- The `buffer` belongs to the byte package of the Go language, and we can use these package to manipulate the byte of the string
- [Golang-Biffer](https://www.educba.com/golang-buffer/)
- [Bytes Buffer package](https://pkg.go.dev/bytes@go1.21.5)

### Select
- The `select` statement lets a goroutine wait on multiple communication operations
```go
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second) // 1 second
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second) // 2 second
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

### Working with JSON
- [Marshalling and Unmarshalling data in Go](https://betterstack.com/community/guides/scaling-go/json-in-go/)
- [Encoding and Decoding JSON data in Go](https://go.dev/blog/json)


### Loops
- Loops provide various control structures to control application flow
- A loop statement allows us to execute a block of code repeatedly
- `for loop`: execute code multiple times, for loop has three components--init statement, condition expression, post statement
```go
for condition {

}
```
```go
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
```
- Looping over a map/slice/array/string or reading from a channel, a `range` clause can manage the loop`, for-each range loop--three iterations
```go
for key, value := range oldMap {
    newMap[key] = value
}
```

### Range
- `Range` form for a `for loop` iterates over a slice/map/string/array, `range` iterates over elements in a variety of data structures
```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow { // i is the index, v is he value of each entry in the slice
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

### Conditionals
- Execute code if a certain condition is true
- If and If/Else Statements
```go
if condition {
    // end program
    break
}
```
```go
if condition {

} else {

}
```
```go
if declaration; condition {

} else {

}
```
```go
if condition {

} else if condition {

} else {

}
```
```go
if condition {

} else if !condition {

} else {

}
```
- If/Else statements for User Input Validation: ==, !=, >=, <=, AND, OR, a := b && c, a := b || c, a := b == c, a := b != c
```go
if condition && condition {

} else {

}
```
```go
if condition || condition {

} else {

}
```
```go
if !condition && !condition {

} else {

}
```
```go
if !condition || !condition {

} else {

}
```

### Switch Statement
- Another alternative to `if-else statement` is the `switch statement`, `switch` improves the application flow better
- An example using `switch statement`, given a user selects a specific city, a certain code is executed for that city.
```go
switch city {
   case "New York":
    // execute code for New York
    case "Singapore":
    // execute code for Singapore
    case "London":
    // execute code for London
    case "Berlin":
    // execute code for Berlin
    case "Lagos", "Ontario":
    // execute code for Lagos and Ontario
    default:
        fmt.Println("No valid city selected")
}
```

### Error Handling, Defer, Panic, Recover
- Built-in `error type` in Go
```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
f.Close()
```
- Error handling is important is golang as sometimes code success is not assured, which depends on factors beyond the programmer control. For example, any function that does I/O must confront the possibility of an error
- An `error` may be `nil` or `non-nil`, `nil` implies success and `non-nil` implies failure, a `non-nil` error has an error message string which we can obtain by calling its Error method or print by calling `fmt.Println(err)` or `fmt.Printf("%v", err)`.
- Two things to note in the Go Programming Language book about errors
- [Error Handling in Go](https://go.dev/blog/error-handling-and-go, https://earthly.dev/blog/golang-errors/)
- [Go Error Package](https://pkg.go.dev/errors@go1.17.5)
- When designing error messages, be deliberate, so that each one is a meaningful description of the problem with sufficient and relevant detail, and be consistent, so that errors returned by the same function or by a group of functions in the same package are similar in form and can be dealt wit h in the same way.
- For example, the os package guarantees that every error returned by a file operat ion, such as os.Open or the Read, Write, or Close methods of an open file, des rib es not just the nature of the failure (permission denied, no such directory, and so on) but also the name of the file, so the caller needn’t include this information in the error message it constructs.
```go
doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```
- Deferred function calls: `defer`, a defer statement is often used with paired operations like open and close, connect and disconnect, or lock and unlock to ensure that resources are released in all cases, no matter how complex the control flow.
- Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
- Recover() is a built-in function in Go that is used to regain control of a panicking goroutine. When a panic() is called, the normal flow of the program is interrupted, and the deferred functions in the same goroutine are executed. You can use recover() within a deferred function to catch the panic value, handle the error, and prevent the program from crashing. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
- [Defer, Panic and Recover](https://go.dev/blog/defer-panic-and-recover)

### Functions
- A function is a group of statements that exist within a program for the purpose of performing a specific task
```go
func greetUsers() {
    fmt.Println("Hello World")
}
```
```go
func main() {
    greetUsers()
}
```
- A function with parameters
```go
func greetUsers(confName string) {
    fmt.Println("Welcome everybody to the", confName, "conference")
}
```
- A function with return value
```go 
func values(x, y int) int {
    return x, y
}
// The execution flow is a return statement that returns a value of the declared type
func main() {
    values(3, 4)
}
```
- Encapsulate code into own container(`=function`), which logically belong together
- Give a `function` a descriptive name
- Call a `function` by its name whenever you want to execute its block of code
```go
greetUsers()
```
```go
greetUsers(confName)
```
- Every function has at least one `function()`, which is the `main()` function
- Concept of bare `return`
- Variadic functions: A variadic function is one that has a varying number of parameters which is written as `func sum(vals ...int) int` and can be called with a varying number of arguments
```go
func sum(vals ...int) int {
total := 0
for _, val := range vals {
total += val
}
return total
}

fmt.Println(sum()) // "0"
fmt.Println(sum(3)) // "3"
fmt.Println(sum(1, 2, 3, 4)) // "10"
```
### Go Type Conversions
- Assignment between items of different type in Go requires explicit conversion
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

### Return Values
- A `function` can `return` data as a result, a function can take an input and return an output
- Function is only executed, when "called"
- Function is also used to reduce code duplication
- In Go, you have to define the input and output parameters including its type explicitly
- Go also allows return multiple values unlike other programming languages
```go
func greetUsers(confName string, firstName string, surName string) (bool, bool, bool) {
    return confName, firstName, surName
}
```

### Building CLIs
- [Cobra Website](https://cobra.dev/), [Cobra Library Documentation](https://pkg.go.dev/github.com/spf13/cobra)
- [urfave/cli Website](https://cli.urfave.org/), [urfave/cli Website](https://github.com/urfave/cli)

### GORM 
- The GORM is fantastic ORM library for Golang, aims to be developer friendly. It is an ORM library for dealing with relational databases. 
- [GORM Website](https://gorm.io/docs/index.html), [GORM Package](https://pkg.go.dev/gorm.io/gorm)