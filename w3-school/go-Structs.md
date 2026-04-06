A struct in Go is used to group different types of data into a single variable. Unlike arrays, which store values of the same type, structs allow you to combine multiple data types together, making them useful for representing real-world entities or records.

To create a struct, you use the `type` and `struct` keywords, and define its fields with their respective data types.

```go 
type Person struct {
  name string
  age int
  job string
  salary int
}
```

Once a struct is declared, you can create variables from it and assign values to its fields using the dot (`.`) operator.

```go 
var p Person
p.name = "Hege"
p.age = 45
```

You can also access and print the values of struct members using the same dot operator.

```go 
fmt.Println(p.name)
fmt.Println(p.age)
```

Structs can be passed as arguments to functions, allowing you to work with structured data inside functions.

```go 
func printPerson(p Person) {
  fmt.Println(p.name, p.age)
}
```

```go 
printPerson(p)
```

Overall, structs help organize related data into a single unit, making programs easier to manage and understand.
