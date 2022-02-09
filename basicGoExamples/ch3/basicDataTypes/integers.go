// Integers

package main

import (
  "fmt"
)

func main () {
  
  // Limit and overflow of uint8
  var u uint8 = 255
  fmt.Println(u, u+1, u*u) // "255 0 1"

  // Limit and overflow of int8
  var i int8 = 127
  fmt.Println(i, i+1, i*i) // "127 -128 1"
}

/*
Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of signed
integers—8, 16, 32, and 64 bits—represented by the types int8, int16, int32, and int64, and
corresponding unsigned versions uint8, uint16, uint32, and uint64.
There are also two types called just int and uint that are the natural or most efficient size for
signed and unsigned integers on a particular platform; int is by far the most widely used numeric
type. Both these types have the same size, either 32 or 64 bits, but one must not make assumptions
about which; different compilers may make different choices even on identical hardware.
The type rune is a synonym for int32 and conventionally indicates that a value is a Unicode code
point. The two names may be used interchangeably. Similarly, the type byte is a synonym for
uint8, and emphasizes that the value is a piece of raw data rather than a small numeric quantity.
Finally, there is an unsigned integer type uintptr, whose width is not specified but is sufficient to
hold all the bits of a pointer value. The uintptr type is used only for low-level programming, such
as at the boundary of a Go program with a C library or an operating system.

Regardless of their size, int, uint, and uintptr are different types from their explicitly sized
siblings. Thus int is not the same type as int32, even if the natural size of integers is 32 bits, and
an explicit conversion is required to use an int value where an int32 is needed, and vice versa.
Signed numbers are represented in 2’s-complement form, in which the high-order bit is reserved for
the sign of the number and the range of values of an n-bit number is from −2n−1 to 2n−1−1. Unsigned
integers use the full range of bits for non-negative values and thus have the range 0 to 2n−1. For
instance, the range of int8 is −128 to 127, whereas the range of uint8 is 0 to 255.

Each operator in the first two lines of the table above, for instance +, has a corresponding assignment
operator like += that may be used to abbreviate an assignment statement.
The integer arithmetic operators +, -, *, and / may be applied to integer, floating-point, and complex
numbers, but the remainder operator % applies only to integers. The behavior of % for negative
numbers varies across programming languages. In Go, the sign of the remainder is always the same
as the sign of the dividend, so -5%3 and -5%-3 are both -2. The behavior of / depends on whether
its operands are integers, so 5.0/4.0 is 1.25, but 5/4 is 1 because integer division truncates the
result toward zero.
If the result of an arithmetic operation, whether signed or unsigned, has more bits than can be
represented in the result type, it is said to overflow. The high-order bits that do not fit are silently
discarded. If the original number is a signed type, the result could be negative if the leftmost bit is a 1.
*/

