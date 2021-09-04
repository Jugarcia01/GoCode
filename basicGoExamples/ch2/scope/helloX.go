// Scope
/* Within a function, lexical blocks may be nested to arbitrary depth, so one local declaration can shadow another. Most blocks are created by control-flow constructs like if statements and for loops. The program below has three different variables called x because each declaration appears in a different lexical block. (This example illustrates scope rules, not good style!) */

package main

import("fmt")

func main() {
    x := "hello!"
    for i := 0; i < len(x); i++ {
        x := x[i]
        if x != '!' {
            x := x + 'A' - 'a'
            fmt.Printf("%c", x) // "HELLO" (una letra por cada iterac)
        }
    }
}
