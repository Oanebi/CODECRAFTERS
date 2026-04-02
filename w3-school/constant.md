
CONSTANTS
 **constants** are immutable values that cannot be changed once they are defined. They are created using the `const` keyword and are ideal for values that should remain fixed throughout the life of a program, such as mathematical constants or configuration settings. Unlike variables, you must assign a value to a constant at the moment you declare it.

### **Key Rules and Characteristics**
* **Read-only:** Once declared, the value is locked and cannot be reassigned.
* **Naming:** They follow standard variable naming rules, but are typically written in **UPPERCASE** to distinguish them from regular variables.
* **Scope:** You can define constants both globally (outside functions) or locally (inside functions).
* **Categorization:** * **Typed:** The data type is explicitly defined (e.g., `int`, `string`).
    * **Untyped:** The type is inferred based on the assigned value.

### **Code Example**
Here is a quick look at how to declare both typed and untyped constants:

```go
package main
import ("fmt")

// Untyped constant
const PI = 3.14 

// Typed constant
const MAX_USERS int = 100 

func main() {
  fmt.Println(PI)
  fmt.Println(MAX_USERS)
}
```