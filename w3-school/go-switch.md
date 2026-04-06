The `switch` statement in Go is used to select and execute one block of code from multiple options. The expression inside the switch is evaluated only once, and its value is compared against each case.

A single-case switch checks one value per case. Each case is tested, and when a match is found, only that block runs. Go does not require a `break` statement.

```go 
switch day {
case 1:
  fmt.Println("Monday")
case 2:
  fmt.Println("Tuesday")
default:
  fmt.Println("Unknown")
}
```

A multi-case switch allows multiple values in a single case using commas. This is useful when different values should produce the same result.

```go
switch day {
case 1, 3, 5:
  fmt.Println("Odd weekday")
case 2, 4:
  fmt.Println("Even weekday")
case 6, 7:
  fmt.Println("Weekend")
}
```

The `default` case is optional and runs when no case matches the expression.

```go 
switch day {
case 1:
  fmt.Println("Monday")
default:
  fmt.Println("Not a weekday")
}
```

It is important that all case values match the type of the switch expression, otherwise an error will occur.

```go 
switch a {
case 1:
  fmt.Println("Valid")
case "b": // error: different type
  fmt.Println("Invalid")
}
```
