# CONDITIONS
Conditions are how I make my code make decisions. In Go, a condition is just a statement that evaluates to either `true` or `false`. I use the comparison and logical operators I learned earlier to build these decision-making blocks.

* **The `if` Statement:** This is the most basic form. If the condition is true, the code inside the curly braces `{}` runs. A big thing to remember: `if` must be lowercase.
* **Handling "Otherwise" with `else`:** If the `if` condition is false, I can use `else` to run a different block of code. 
* **Multiple Options with `else if`:** When I have more than two possibilities, `else if` lets me test new conditions one after the other.
    * **Note:** If two conditions are both true, Go only runs the *first* one it hits.
* **The Curly Brace Rule:** Go is very strict about where I put my braces. The `else` or `else if` **must** be on the same line as the closing brace `}` of the previous block.
* **Nested Ifs:** I can put an `if` statement inside another `if` statement. This is useful for "double-checking" a second condition only if the first one passed.



```go
package main
import ("fmt")

func main() {
  time := 22
  temperature := 14

  // 1. Basic if-else if-else chain
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }

  // 2. Nested if example
  if temperature > 10 {
    fmt.Println("It's above freezing.")
    if temperature > 25 {
      fmt.Println("Actually, it's quite hot!")
    }
  }

  // 3. Crucial Syntax: The "} else {" must be on one line
  x := 10
  if x > 5 {
    fmt.Println("x is big")
  } else { // This line would error if 'else' was moved down
    fmt.Println("x is small")
  }
}
```

---

### My Summary Checklist
 **Brace Placement:** Always write `} else {` on the same line to avoid syntax errors.
 **Boolean Logic:** Remember that the condition inside the `if` must result in a `true` or `false`.
 **First-Match Wins:** In an `else if` chain, only the first true block executes, even if later conditions are also true.
 **Case Sensitivity:** Keep `if`, `else`, and `else if` in all lowercase.