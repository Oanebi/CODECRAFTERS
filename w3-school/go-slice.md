# SLICES
Slices are much more common in Go than arrays because they are powerful and flexible. While they also store multiple values of the same type, a slice’s length can grow and shrink as needed, making them dynamic.

* **Creating Slices:** You can create a slice directly using `[]datatype{values}`, by "slicing" an existing array, or by using the `make()` function for better control.
* **Length vs. Capacity:** * `len()` returns the number of elements currently in the slice.
    * `cap()` returns the capacity (how much the slice can grow or shrink to).
* **Appending Items:** Use the `append()` function to add elements to the end. To merge two slices together, you must use the `...` operator after the second slice name.
* **Accessing and Changing:** Just like arrays, you use index numbers (starting at 0) to get or update values.
* **Memory Efficiency:** When you take a small slice from a huge array, Go keeps the whole array in memory. Use the `copy()` function to move only the needed data into a new, smaller slice to save memory.
* **The make() Function:** This is the preferred way to create a slice if you already know how much data you need to store, using the format `make([]type, length, capacity)`.



```go
package main
import ("fmt")

func main() {
  // Creating slices in different ways
  myslice1 := []int{1, 2, 3} 
  myslice2 := make([]int, 5, 10) // length 5, capacity 10

  // Creating a slice from an array
  myarr := [6]int{10, 11, 12, 13, 14, 15}
  myslice3 := myarr[2:4] // [12 13]

  // Modifying and Appending
  myslice1[0] = 100
  myslice1 = append(myslice1, 4, 5)

  // Appending one slice to another
  combined := append(myslice1, myslice3...)

  // Memory efficient copying
  src := []int{1, 2, 3, 4, 5}
  dest := make([]int, len(src))
  numCopied := copy(dest, src)

  fmt.Println(myslice1)      // [100 2 3 4 5]
  fmt.Println(cap(myslice2)) // 10
  fmt.Println(combined)     // [100 2 3 4 5 12 13]
  fmt.Println(dest)         // [1 2 3 4 5]
}
```