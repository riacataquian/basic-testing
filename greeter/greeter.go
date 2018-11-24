package greeter

import "fmt"

// Greet greets the supplied name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
