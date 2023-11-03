### Install Go

### Visual Studio Code

### VS Code Go Extension

### Create Working Directory and Go file
```
mkdir go_tutorial
```
```
touch main.go
```

### Initialize Working Directory
```
go mod init <module path>
```
Module path can correspond to a repository you plan to publish your module.

### Go Commands
```
go run <file name>
```

### Notes
- Your Go code must belong to a package. the first statement in a Go file must be `package ...`
- The `func main()` is the typical entrypoint of a Go Program
- A program can only have one `main function`, because you can only have one entrypoint

### Variables and Constants
- `isValidEmail := tomiwaaribisala@gmail.com`
- `isValidName := len(firstName) >= 2 || len(surName) >= 10`
- `isValidTC := boughtTickets > 0 && leftTickets <= 50`
- `var name = "tomiwa"`
- `const ticketPrice = $70`
- Printing formatted data in Go using `Println`; `fmt.Println("My name is", name, "my ticket price is", ticketPrice)`
- Printing formatted data in Go using `printf`; `fmt.printf("The ticket price is %v.\n, ticketPrice")`

### Package Level Variables
- Defined at the top outside all functions
- They can be accessed inside any of the functions, and in all files/functions which are in the same `package`

### Data Types
Go is a statically typed language, you need to tell the Go compiler the data type when declaring a variable; however, in most cases, Go can infer the data type of a variable when you assign a value.
- Strings   eg. `var UserName string`
- Integers  eg. `var UserTickets int`, `var UserTickets uint`
- Booleans  eg.
- Maps      eg. 
- Arrays    eg. 

### User Input 
- Using `fmt` package to collect user input; 
    `var userName string`
    `fmt.Println("Please enter your first name: ")`
    `fmt.Scan(&userName)`

### Pointers
- Pointers are special variables that points to the memory address of another variable(value)
    `var matchPrice = 50`
    `fmt.Println(&matchPrice)`

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
if condition {

} else if condition {

} else {

}
```
- If/Else statements for User Input Validation: ==, !=, >=, <=, AND, OR, a := b && c, a := b || c
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
- In Go you have to define the input and output parameters including its type explicitly
- Go also allows return multiple values
```go
func greetUsers(confName string, firstName string, surName string) (bool, bool, bool) {
    return confName, firstName, surName
}
```