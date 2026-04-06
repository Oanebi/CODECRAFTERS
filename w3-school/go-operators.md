# OPERATORS
Operators are what are used to perform calculations and comparisons between variables and values. Go groups these into five main categories, and understanding them is essential for controlling the logic of any code.

* **Arithmetic Operators:** These are for standard math. Besides the basics (`+`, `-`, `*`, `/`), I learned about:
    * **Modulus (`%`)**: Returns the remainder of a division (e.g., `5 % 2` is 1).
    * **Increment/Decrement (`++`, `--`)**: Quickly increases or decreases a value by 1.
* **Assignment Operators:** These assign values. The cool part is the "shorthand" operators like `+=` or `*=`. Instead of writing `x = x + 5`, I can just write `x += 5`.
* **Comparison Operators:** Used to check relationships. These always return a **boolean** (`true` or `false`).
    * `==` checks if values are equal.
    * `!=` checks if they are *not* equal.
* **Logical Operators:** These help me combine multiple conditions. 
    * `&&` (AND): True if *both* sides are true.
    * `||` (OR): True if *at least one* side is true.
    * `!` (NOT): Reverses the result.
* **Bitwise Operators:** These work on the binary (bit) level. While I might not use them daily, they are powerful for low-level data manipulation using `AND`, `OR`, `XOR`, and bit shifts (`<<`, `>>`).



```go
package main
import ("fmt")

func main() {
  // 1. Arithmetic
  x, y := 10, 3
  fmt.Println(x % y) // 1 (Remainder)

  // 2. Assignment Shorthand
  count := 10
  count += 5 // Same as count = count + 5
  
  // 3. Comparison
  isGreater := x > y // true
  
  // 4. Logical
  // Check if x is greater than 5 AND y is less than 10
  result := (x > 5 && y < 10) 

  fmt.Println(count)     // 15
  fmt.Println(isGreater) // true
  fmt.Println(result)    // true
}
```

---

### My Summary Checklist
**The Assignment Rule:** Remember that `=` is for assigning, but `==` is for comparing.
 **Boolean Outputs:** Comparison and Logical operators always result in `true` or `false`.
 **Order of Operations:** Just like in math, I can use parentheses `()` to make sure my logic is calculated in the right order.
  **Increment/Decrement:** These can only be used on variables, not on literal values (you can't do `5++`).