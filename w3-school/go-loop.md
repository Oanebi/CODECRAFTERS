
The `for` loop in Go is used to repeatedly execute a block of code for a specific number of times. It is the only looping structure available in Go, and each repetition of the loop is called an iteration.

The loop can have three parts: initialization, condition, and increment. The initialization sets the starting value, the condition determines how long the loop runs, and the increment updates the value after each iteration. If the condition becomes false, the loop stops.

```go 
for i := 0; i < 5; i++ {
  fmt.Println(i)
}
```

The three parts of the loop are flexible and can be adjusted. For example, you can increase the counter differently, such as counting by tens instead of ones.

```go 
for i := 0; i <= 100; i += 10 {
  fmt.Println(i)
}
```

The `continue` statement is used to skip a particular iteration and move to the next one, usually when a condition is met.

```go
for i := 0; i < 5; i++ {
  if i == 3 {
    continue
  }
  fmt.Println(i)
}
```

The `break` statement is used to stop the loop completely when a condition is met.

```go 
for i := 0; i < 5; i++ {
  if i == 3 {
    break
  }
  fmt.Println(i)
}
```

Go also allows nested loops, where one loop runs inside another. The inner loop runs completely for each iteration of the outer loop.

```go
for i := 0; i < 2; i++ {
  for j := 0; j < 3; j++ {
    fmt.Println(i, j)
  }
}
```

The `range` keyword provides an easier way to loop through arrays, slices, or maps. It returns both the index and the value during each iteration.

```go 
fruits := [3]string{"apple", "orange", "banana"}
for idx, val := range fruits {
  fmt.Println(idx, val)
}
```

If you only need the value or the index, you can ignore the other using an underscore (`_`).

```go 
for _, val := range fruits {
  fmt.Println(val)
}
```

Overall, the `for` loop in Go is flexible and powerful, allowing different styles of looping, control with `break` and `continue`, and easy iteration using `range`.
