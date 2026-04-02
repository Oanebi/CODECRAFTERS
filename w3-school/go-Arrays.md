# ARRAYS
Arrays store many values of the same type in one variable. In Go, arrays always have a fixed size. This means you must decide how many items the array will hold when you create it.

* **Creating Arrays:** You can use the `var` keyword or the `:=` sign. You can set the length with a number or use `[...]` to let Go count the items for you.
* **Indexes:** Arrays start counting at 0. The first item is at index `[0]`, the second is at `[1]`, and so on.
* **Changing Items:** You can change a specific value by referring to its index number.
* **Default Values:** If you do not give an item a value, Go gives it a default one (0 for integers and "" for strings).
* **Specific Items:** You can choose to fill only certain spots in an array using the index number and a colon (like `1:10`).
* **Finding Length:** Use the `len()` function to see how many items are in an array.

```go
package main
import ("fmt")

func main() {
  
  var arr1 = [3]int{10, 20, 30}
  arr2 := [...]string{"Apple", "Banana"}
  arr1[2] = 50 

  
  arr3 := [5]int{1: 10, 2: 40}

  fmt.Println(arr1)      // [10 20 50]
  fmt.Println(len(arr2)) // 2
  fmt.Println(arr3)      // [0 10 40 0 0]
}
```