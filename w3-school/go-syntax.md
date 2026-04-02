A Go program is organized into package declarations, imported packages, functions, and statements. Every executable Go file must belong to the **main** package to run.

* **package main:** Defines the package as the starting point for the program.
* **import ("fmt"):** Includes the "fmt" package, which contains functions for formatting and printing text.
* **func main():** The main function where the program's execution begins; code inside the curly brackets `{}` is executed.
* **fmt.Println():** A statement used to output text to the terminal.
* **Line Endings:** Statements are separated by newlines or semicolons. Pressing Enter adds a semicolon implicitly to the source code.
* **Syntax Rule:** The left curly bracket `{` cannot be placed at the start of a new line; it must follow the function declaration.

```go
package main
import ("fmt")

func main() { 
  fmt.Println("Hello World!") 
}
```

Whitespace is used throughout the code to improve readability, but it is ignored by the Go compiler during execution.