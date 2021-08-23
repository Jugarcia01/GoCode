// PALABRAS CLAVES Y RESERVADAS DE GO / GOLANG
// Go has 25 keywords like if and switch that may be used only where the syntax permits; they can’t be used as names.

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var

// In addition, there are about three dozen predeclared names like int and true for built-in constants, types, and functions:
Constants:
true  false  iota  nil

Types:
int  int8  int16  int32  int64
uint  uint8  uint16  uint32  uint64
uintptr
float32  float64  complex128  complex64
bool  byte  rune  string  error

Functions:
make  len  cap  new  append  copy  close
delete  complex  real  imag  panic  recover

// VARIABLES
La forma general de declarar variables en Golang es:
var name type = expression

Se puede omitir alguno de los dos siguientes datos ya sea el tipo o =expresion, pero no ambos a la vez.
Si se omite la expresion, entonces el valor inicial de la variable será el valor por defecto de acuerdo al tipo de variable:
  0 para numeros,
  false para booleanos,
  "" para cadenas/Strings, y
  nil para interfaces y tipos de reference (slice, pointer, map, channel, function).
  The zero value of an aggregate type like an array or a struct has the zero value of all of its elements or fields.
Por ejemplo:
  var s string
  fmt.Println(s) // ""
Imprime una cadena vacia (empty string), Go asegura la inicializacion de la variable, en lugar de causar un tipo de error o comportamiento impredecible.

En Go tambien se pueden declarar e inicializar un conjunto de variables en una simple sentencia.
Por ejem:
var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string

Declaracion corta de Variables
name := expression
El tipo/type de la variable name es determinado por el tipo de expresion.
Por ejem:
freq := rand.Float64() * 3.0
t := 0.0
Al igual quela declaraciones con var, multiples variables cortas se pueden realizar a la vez:
i, j := 0, 1-

// PUNTEROS
Un puntero es la ubicación/dirección en la cual es almacenado el valor de una variable.
Con los punteros podemos leer o actualizar indirectamente el valor de las variables, inclusive sin conocer el nombre de la variable.

If a variable is declared var x int, the expression &x (“address of x”) yields a pointer to an integer variable, that is, a value of type *int, which is pronounced “pointer to int.” If this value is called p, we say “p points to x,” or equivalently “p contains the address of x.” The variable to which p points is written *p. The expression *p yields the value of that variable, an int, but since *p denotes a variable, it may also appear on the left-hand side of an assignment, in which case the assignment updates the variable.
Por ejem:
  x := 1
  p := &x         // p, of type *int, points to x
  fmt.Println(*p) // "1"
  *p = 2          // equivalent to x = 2
  fmt.Println(x)  // "2"
Nota: El valor por defecto / The zero value para un puntero  de cualquier tipo es nil


