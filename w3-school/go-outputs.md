In Go, the `fmt` (format) package provides three main ways to output text. Here is a simple breakdown of how they work.


## 1. The `Print()` Function
This is the most basic function. It prints text exactly as you provide it.

* **No Automatic Spaces:** It does not add spaces between strings unless you include them.
* **No Automatic New Lines:** It stays on the same line unless you use the `\n` escape character.
* **Exception:** It *does* add a space between arguments if neither one is a string (e.g., printing two numbers).

```go
fmt.Print("Hello", "World") // Output: HelloWorld
fmt.Print(10, 20)           // Output: 10 20
```


## 2. The `Println()` Function
`Println` stands for "Print Line." It is usually the most convenient for quick logging.

* **Automatic Spaces:** It always adds a space between your arguments.
* **Automatic New Line:** It automatically adds a new line at the end of the output.

```go
fmt.Println("Hello", "World") 
// Output: Hello World
// (Cursor moves to next line)
```



## 3. The `Printf()` Function
`Printf` stands for "Print Formatted." It allows you to use **Formatting Verbs** (placeholders) to control how your data looks.

### General Verbs
Used for almost any type of data:
* `%v`: The default value.
* `%T`: The type of the variable (e.g., `int` or `string`).
* `%%`: Prints a literal percent sign.

### Integer Verbs
Used to change how numbers are displayed:
* `%d`: Standard base-10 decimal.
* `%b`: Binary (base-2).
* `%x`: Hexadecimal (base-16).
* `%04d`: Pads the number with zeros until it is 4 digits long.

```go
i := 15
fmt.Printf("Value: %v, Type: %T", i, i) 
// Output: Value: 15, Type: int
```


