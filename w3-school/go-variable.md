VARIABLES.
In Go, variables act as containers for storing data. Every variable has a specific type that determines what kind of data it can hold, such as whole numbers (**int**), decimals (**float32**), text (**string**), or true/false values (**bool**). Go is flexible, allowing you to either state the data type explicitly or let the computer figure it out based on the value you provide.


### Key Features of Go Variables

* **Declaration Methods:** You can create variables using the `var` keyword (e.g., `var name string = "John"`) or the short assignment operator `:=` (e.g., `name := "John"`).
* **Type Inference:** If you don't specify a type but provide a value, Go automatically assigns the correct type for you.
* **Multiple Variables:** You can declare several variables at once on a single line or group them together in a `var ()` block for better organization.
* **Naming Strictness:** Names must start with a letter or underscore, cannot contain spaces, and are case-sensitive (`age` and `Age` are different).
* **Naming Styles:** To handle names with multiple words, programmers typically use **Camel Case** (`myVar`), **Pascal Case** (`MyVar`), or **Snake Case** (`my_var`).


### Code Examples

#### 1. Basic Declaration
This shows the different ways to create variables and assign values.

```go
package main
import ("fmt")

func main() {
    // Using var with a specific type
    var student1 string = "John"
    
    // Using var where the type is inferred
    var student2 = "Jane"
    
    // Short hand declaration (only works inside functions)
    x := 2

    fmt.Println(student1, student2, x)
}
```

#### 2. Multiple Variables and Blocks
You can declare many variables at once, even if they have different data types.

```go
package main
import ("fmt")

func main() {
    // Multiple variables on one line
    var a, b, c = 5, 10, "Hello"
    
    // Declaring variables in a block
    var (
        name string = "Alice"
        age  int    = 25
        isReady bool = true
    )

    fmt.Println(a, b, c)
    fmt.Println(name, age, isReady)
}
```