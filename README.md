Myers
=====

Golang implementation of Myers diff algorithm

This is a direct port to golang of the algorithm implemented by Robert Elder software [here](https://blog.robertelder.org/diff-algorithm/)

Usage:
```
// go get github.com/MFAshby/myers

package main

import (
    "fmt"
    "github.com/MFAshby/myers"
)

func main() {
    e := []string{"bar", "baz"}
    f := []string{"foo", "baz"}
    ops := myers.DiffStr(e, f)
    fmt.Printf("%v\n", ops) // []Op{{OpDelete, 0, -1, "bar"}, {OpInsert, 1, 0, "foo"}}},
}
```