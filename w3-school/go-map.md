Maps in Go are used to store data in **key:value pairs**, where each key is unique and used to access its corresponding value. They are useful for organizing and retrieving data efficiently.

A map is **unordered**, meaning the elements are not stored in a specific order, and it does not allow duplicate keys. Maps are also changeable, so you can add, update, or remove elements. The number of elements in a map can be found using the `len()` function, and the default value of a map is `nil`.

Maps are created using either the `var` keyword or the shorthand `:=`, and you must specify the key type and value type.

```go id="m1a2p3"
var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
```

Overall, maps are a flexible way to store and manage data using unique keys, backed by an underlying hash table for efficient access.
