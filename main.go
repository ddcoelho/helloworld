// What does package main mean
//// packages = project, all files inside the project with the same package name
//// main is a reserved word for a an EXECUTABLE package
//// 2 type of packages:
////// -Executable - Generates a executable file
////// -Reusable - Code used as helpers - Goog for reusable code
package main

// What does import "fmt" mean
//// import gives access to all the code inside package fmt
//// fmt = format
import "fmt"

// What does func main() mean
//// func main is necessary in a package main
//// in a function
func main() {
	fmt.Print("Hello world!")
}

// How does the program run
//// go run main.go
//// go build main.go & run wth ./main

// how is the code organized
//// 1-first package declaration
//// 2-imports
//// 3-functions, var declaration and do stuff
