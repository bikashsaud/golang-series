### Day 2

#### Documentation

> go help

- Usage: go <command> [arguments]
- go help gopath --> to get path of go envs



#### Lexer

> https://go.dev/ref/spec#Lexical_elements
- one line comment -->  //
- multi line comment --> /* */

...

#### Types

- Case insensitive; almost
- Veriable type should be known in advance
- every thing is type
- String, 
- Bool, 
- integer[uint8, uint64, int8, uint64, uintptr] 
- Float, [float32, float 64]
- COmplex


- Array
- SLices
- Maps
- Structs
- Pointer

- FUnction
- Channel


### Numeric type values
```

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

```

> Importent Note:
- for public variable use First latter is Capital 
- EG: const Token = "ashdkj-s0-aisdas"
- 