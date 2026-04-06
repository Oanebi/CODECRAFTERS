A function in Go is a block of code that can be reused to perform a specific task. Functions do not run automatically, they only execute when they are called. You define a function using the `func` keyword, followed by its name, parentheses, and code inside curly braces.

```go 
func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage()
}
```

Functions can be called multiple times, making code more reusable and organized.

Functions can also return values. To do this, you specify the return type and use the `return` keyword inside the function.

```go 
func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(1, 2))
}
```

Go allows named return values, where you define the return variable in the function signature and return it directly.

```go 
func add(x int, y int) (result int) {
  result = x + y
  return
}
```

A function can return multiple values, which is a powerful feature in Go. These values can be stored in variables, and unwanted values can be ignored using an underscore (`_`).

```go
func calc(x int, y string) (int, string) {
  return x + x, y + " World!"
}

func main() {
  a, b := calc(5, "Hello")
  fmt.Println(a, b)
}
```

Functions can take parameters, which act as inputs. When calling the function, you pass arguments that match the parameters in number, type, and order.

```go
func greet(name string) {
  fmt.Println("Hello", name)
}

func main() {
  greet("Liam")
}
```

You can also define multiple parameters in a function.

```go 
func info(name string, age int) {
  fmt.Println(name, age)
}
```

Go supports recursion, where a function calls itself. A stopping condition is important to prevent infinite loops.

```go 
func count(x int) {
  if x == 5 {
    return
  }
  fmt.Println(x)
  count(x + 1)
}
```

Recursion is useful for solving problems like factorials, but it must be used carefully to avoid excessive memory use or non-terminating loops.

```go 
func factorial(x int) int {
  if x == 0 {
    return 1
  }
  return x * factorial(x-1)
}
```

Overall, functions help organize code, make it reusable, support input and output through parameters and return values, and allow advanced techniques like recursion.
