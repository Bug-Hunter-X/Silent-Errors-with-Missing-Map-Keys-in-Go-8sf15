```go
package main

import (

    "fmt"
)

func main() {
    m := make(map[string]int)
    m["a"] = 1
    val, ok := m["b"]
    if ok {
        fmt.Println("Value for b:", val)
    } else {
        fmt.Println("Key b not found in the map")
    }
}
```