// Go code is organized into packages, which are similar to libraries or modules in other languages. 
// A package consists of one or more .go source files in a single directory that define what the package does. 
// Each source file begins with a package declaration,
// followed by a list of other packages that it imports, 
// and then the declarations of the program that are stored in that file.

//Go comes with an extensive stand ard library of useful packages, 
// and the Go community has create d and share d many more. 
// Programming is often more about using existing packages, 
// than about writing original code of one’s own
// Before you emb ark on any new program, it’sagood ide a to see 
// if packages already exist that might help you get your job done more easily

//Go Library > Packages > Functions//

//Package declaration//
package main

//Import declaration//
import "fmt"

//Functions, Variables, Constants, Types declaration (func, const, var, type)//
func main() {

fmt.Println("Hello, World")		// calling PrintIn function from fmt package //

}