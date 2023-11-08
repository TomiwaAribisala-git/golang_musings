### Install Go
- [Install GO](https://go.dev/doc/install)

### Visual Studio Code

### VS Code Go Extension

### Create Working Directory and Go file
```
mkdir go_tutorial
```
```
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
```
go mod init <module path>
```
Module path can correspond to a repository you plan to publish your module.

### Go Commands
- Go Build
```
go build helloworld.go
```
```
./helloworld
```
- Go Run
```
go run helloworld.go
```
- Go Format
```
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
- Variables and fucntions defined outside any function, can be accessed in all other files within the same package
- Execute all files in a package
```
go run .
```
- When working with multiple packages, create a directory/file as regards each package, you can also import packages between each other eg. `import <module_path>/package` then `package.(any function in the helper package)`
- For importing functions between packages, make sure to capitalize the first letter of the function, this action exports the function and makes it importable in another package
- You can also export variables, functions, constants, types; makes sure to capitalize the first letter of the data type to make it importable

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
- Floating point numbers
- Complex Numbers 
- Booleans: `bool`---`true`, `false`
- Strings: Packages for manipulating strings--`strings`, `strconv`, `bytes`, `unicode`


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
// ...use f...
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
- Defined inside a function or block
- Local variables can be accessed only inside that function or block of code
- Best practice is to define variable as local as possible, create the variable where you need it 

### Package Level Variables
- Defined at the top outside all functions
- They can be accessed inside any of the functions, and in all files/functions which are in the same `package ...`

### Global Variables
- Declaration outside all functions and uppercase first letter
- Global variables can be used everywhere across all packages 

### Assignments 
- The value held by a variable can be updated by an assignment statement, which in its simplest form has a variable on the left of the = sign and an expression on the right
### Data Types
Go is a statically typed language, you need to tell the Go compiler the data type when declaring a variable; however, in most cases, Go can infer the data type of a variable when you assign a value.
- Strings   eg. `var UserName string`
- Integers  eg. `var UserTickets int`, `var UserTickets uint`
- Slices
- Booleans  
- Maps       
- Arrays     

### User Input 
- Using `fmt` package to collect user input; 
    `var userName string`
    `fmt.Println("Please enter your first name: ")`
    `fmt.Scan(&userName)`

### Pointers
- Pointers are special variables that points to the memory address of another variable(value)
    `var matchPrice = 50`
    `fmt.Println(&matchPrice)`
    `&matchPrice--address of matchPrice`
    `p := &matchPrice`
    `*p = 70` # this expression sets the value of the `matchPrice` value to 70
    `fmt.Println(matchPrice)`
    `z := new(int)`
    `fmt.Println(*z)`
    `*z = 2`
    `fmt.Println(*z)`

### Arrays and Slices 
- Data structures to store collection of elements in a single variable, index-based
- Arrays: 
    `var bookings [50] string`
    `var firstName string`
    `var lastName string`
    `fmt.Println("Please enter your first name: ")`
    `fmt.Scan(&firstName)`
    `fmt.Println("Please enter your last name: ")`
    `fmt.Scan(&lastName)`
    `bookings[0] = firstName + " " + lastName`
    `fmt.Println("Welcome", firstName, lastName)`
- Slices: an abstraction of an array, more flexible and powerful--variable-length or get a sub-array of its own, slices have a size like an array, but can be resized when needed. 
	`var bookings [] string`
	`var firstName string`
	`var lastName string`
	`fmt.Println("Please enter your first name: ")`
	`fmt.Scan(&firstName)`
	`fmt.Println("Please enter your last name: ")`
	`fmt.Scan(&lastName)`
	`bookings = append(bookings, firstName + " " + lastName)`
	`fmt.Println("Thank you", firstName, lastName, "for inputting your name!")`
	`fmt.Println("Print all our bookings:", bookings)`

### Loops
- Loops provide various control structures to control the applications flow
- A loop statement allows us to execute code multiple times, in a loop
- `for loop`: execute code multiple times
```go
for {
		var surName string
		var entryPrice uint
		fmt.Println("Please enter your surName: ")
		fmt.Scan(&surName)
		fmt.Println("Please enter your entryPrice: ")
		fmt.Scan(&entryPrice)
		fmt.Println("Thank you for coming to the cinema", surName, "your entry price is", entryPrice, "dollars")
	}
```
- `for-each loop`: iterating over a list 
```go
for condition {

}
```

### Conditionals
- If/Else Statements
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
if condition; condition {

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

### Functions
```go
func greetUsers() {
    
}
```
```go
func greetUsers(confName string) {
    fmt.Println("Welcome everybody to the", confName, "conference")
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

### Maps
- Maps maps unique keys to values, you can retrieve the value later by using its key
- All keys and values have the same data type, map supports only one data type at once
```go
var userData = []map[string]string
userData["firstname"] = firstname
userData["lastname"] = lastname
userData["email"] = email
```

### Struct
- Struct is a data structure that allows us to mix different data types, defines a structure(which fields) of the User Type
```go
type UserData struct {
    firstName string
    lastName string
    email string
    numberofTickets uint
}
```
```go
var userData = UserData {
    firstName string: firstName,
    lastName: lastName, 
    email: email,
    numberofTickets: userTickets,
}
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