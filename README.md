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
[Organizing a Go Module](https://go.dev/doc/modules/layout#)

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
bookings = append(bookings, firstName + " " + lastName)
fmt.Println("Thank you", firstName, lastName, "for inputting your name!")
fmt.Println("Print all our bookings:", bookings)
```
- A slice can be formed by specifying two indices, a low and high bound, separated by a colon, `a[1:4]`, describing a section of an underlying array
- Changing the elements of a slice modifies the corresponding elements of its underlying array
- Length and Capacity of a slice
- Creating a slice with the `make()` function
```go
slice_name := make([]type, length, capacity)
```

### Maps
- Maps maps unique keys to values, you can retrieve the value later by using its key
- All keys and values have the same data type, map supports only one data type at once
```go
var userData = []map[string]string 
userData["firstname"] = firstname
userData["lastname"] = lastname
userData["email"] = email

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
- Adding/Updating elements in a map
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

### Struct
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

### Goroutines(Concurrency)
- Sequential code execution, make our program more efficient
- Executing different parts of code in each separate thread(goroutine)
```go
func main() {

}
```
```go
go main()   // Go starts a new goroutine
```

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
```
- Deferred function calls: `defer`, a defer statement is often used with paired operations like open and close, connect and disconnect, or lock and unlock to ensure that resources are released in all cases, no matter how complex the control flow.
- Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
- Recover() is a built-in function in Go that is used to regain control of a panicking goroutine. When a panic() is called, the normal flow of the program is interrupted, and the deferred functions in the same goroutine are executed. You can use recover() within a deferred function to catch the panic value, handle the error, and prevent the program from crashing.

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