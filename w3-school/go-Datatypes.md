## Summary of Go Data Types

Go is a **statically typed** language, meaning a variable's type is decided when it is created and cannot change. Data types tell the computer how much memory to use and what kind of data is being stored. Every data type in Go has a **zero value**, which is the default value assigned if you declare a variable without giving it an initial value.


### Core Data Types

* **Boolean (`bool`):** Used for logic. It only accepts `true` or `false`. 
    * **Zero Value:** `false`
* **Integer (`int`):** Used for whole numbers. They are split into **signed** (positive and negative) and **unsigned** (positive only). Common types include `int`, `int8`, `int64`, `uint`, and `uint8`.
    * **Zero Value:** `0`
* **Float (`float32`, `float64`):** Used for numbers with decimals. `float64` is the default and provides higher precision.
    * **Zero Value:** `0`
* **String (`string`):** Used for sequences of text characters. String values must be wrapped in double quotes (`""`).
    * **Zero Value:** `""` (an empty string)


### Code Example

```go
package main
import ("fmt")

func main() {
    
    var isReady bool = true
    var count int = 10
    var price float64 = 19.99
    var greeting string = "Hello"

    var defaultBool bool     // false
    var defaultInt int       // 0
    var defaultFloat float64 // 0
    var defaultString string // "

}
```