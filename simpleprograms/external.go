// Main package declared
package main

// Import packages
import "fmt"

// Import external package
import "rsc.io/quote"

// Main function
func main() {
    // From external package, print certain functions
    fmt.Println(quote.Go())
    fmt.Println(quote.Glass())
    fmt.Println(quote.Opt())
}
