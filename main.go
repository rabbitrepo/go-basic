// Hello World

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// }

// RUN: go run main.go
// RUN with Nodemon: nodemon --exec go run main.go --signal SIGTERM

// BUILD: go build main.go
// RUN The Binary: ./main

// -----

// Print UUID

// package main

// import (
// 	"fmt"

// 	"github.com/google/uuid"
// )

// func main() {
// 	id := uuid.New()
// 	fmt.Printf("Generated UUID: %s\n", id)
// }

// List all dependencies: go list -m all

// Function

// package main

// import (
// 	"fmt"
// )

// func sayHello() {
// 	fmt.Println("Hello World")
// }

// func main() {
// 	sayHello()
// 	sayHello()
// 	sayHello()
// }

// Importing Go Package

package main

import (
	"github.com/rabbitrepo/go-basic/rabbit"
)

func main() {
	rabbit.HelloRabbit()
	rabbit.HelloRabbit2()
	// rabbit.HelloSecret()
}

// rabbit.Hello => Public Function
// rabbit.hello => Private Function (can be used inside the module)
// 1 Folder = 1 Package | Suggested that folder name = package name
