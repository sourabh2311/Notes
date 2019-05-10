# Learning GoLang

* ` import ("math/rand") ` By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

* Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

* In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

* When two or more consecutive named function parameters share a type, you can omit the type from all but the last. So we can short (x int, y int) to (x, y int)

* A function can return any number of results. ` func swap(x, y string) (string, string) {
	return y, x
} `

* Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values. A return statement without arguments returns the named return values. This is known as a "naked" return. Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions. 
  ```go
  func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
  } 
  ```
* The var statement declares a list of variables; as in function argument lists, the type is last.

* A var declaration can include initializers, one per variable.
  If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
  ```go
  var i, j int = 1, 2
  func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)  // output: 1 2 true false no!
  }
  ```

* Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

* Go's basic types are
  ```
  bool

  string

  int  int8  int16  int32  int64
  uint uint8 uint16 uint32 uint64 uintptr

  byte // alias for uint8

  rune // alias for int32
      // represents a Unicode code point

  float32 float64

  complex64 complex128
  ```
  ```go
  import (
	"fmt"
	"math/cmplx"
  )

  var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
  )

  func main() {
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
  }
  ```
  The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

  The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

* Variables declared without an explicit initial value are given their zero value.

  The zero value is:

  0 for numeric types,

  false for the boolean type, and

  "" (the empty string) for strings.

* The expression T(v) converts the value v to the type T.

  Some numeric conversions:
  ```go
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
  // Or, put more simply:
  i := 42
  f := float64(i)
  u := uint(f)
  ```

  Unlike in C, in Go assignment between items of different type requires an explicit conversion.

* Constants are declared like variables, but with the const keyword.

  Constants can be character, string, boolean, or numeric values.

  Constants cannot be declared using the := syntax.

* **Numeric Constants** are high-precision values.

  An untyped constant takes the type needed by its context.

  ```go
  const (
    Big = 1 << 100
    Small = Big >> 99
  )

  func needInt(x int) int { return x*10 + 1 }
  func needFloat(x float64) float64 {
    return x * 0.1
  }

  func main() {
    fmt.Println(needInt(Big))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
  }
  ```
  Doing needInt(Big) will produce an error.

* Go has only one looping construct, the for loop. Syntax: `for i := 0(optional); i < 10; i++(optional) { // do something }`

* You can drop the semicolons: C's while is spelled for in Go. 
  ```go
  for sum < 1000 {
		sum += sum
	}
  ```
* ```go
  for { // If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
  }
  // remember that go has 'break' just like other languages.
  ```
* Like for, the if statement can start with a short statement to execute before the condition.
Variables declared by the statement are only in scope until the end of the if.
  ```go
  func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
      return v
    } // else {} // Note: Variables declared inside an if short statement are also available inside any of the else blocks.
    return lim
  }
  ```

* A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

  Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

* Switch without a condition is the same as switch true.

  This construct can be a clean way to write long if-then-else chains.

  ```go
  switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
  ```

* A **defer** statement defers the execution of a function until the surrounding function returns.

  The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

  ```go
  func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
  }

  // output: hello
  //         world

  ```

* Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

  ```go
  fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

  fmt.Println("done")
  /* 
    counting
    done
    9
    8
    7
    6
    5
    4
    3
    2
    1
    0 
  */
  ```

* Go has pointers. A pointer holds the memory address of a value.

  The type *T is a pointer to a T value. Its zero value is nil.

  ```go
  var p *int
  ```
  The & operator generates a pointer to its operand.

  ```go
  i := 42
  p = &i
  ```
  The * operator denotes the pointer's underlying value.

  ```go
  fmt.Println(*p) // read i through the pointer p
  *p = 21         // set i through the pointer p
  ```
  This is known as "dereferencing" or "indirecting".

  Unlike C, Go has no pointer arithmetic.

* ```go
  type Vertex struct {
    X int
    Y int
  }

  func main() {
    v := Vertex{1, 2}
    p := &v
    // To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
    p.X = 1e9
    fmt.Println(v)
  }

  // One more example (example of Struct literals)
  var (
    v1 = Vertex{1, 2}  // has type Vertex
    // You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
  )

  func main() {
    fmt.Println(v1, p, v2, v3)
  }

  ```

* ```go
  // arrays
  func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1]) // Hello World
    fmt.Println(a) // [Hello World]

    primes := [6]int{2, 3, 5, 7, 11, 13} 
    fmt.Println(primes) // [2 3 5 7 11 13]
    var s []int = primes[1:4] // slice! (Note that 4th (high) element is excluded)
    fmt.Println(s)
  }
  ```

* Slices are like references to arrays.
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.

* A slice literal is like an array literal without the length.

  This is an array literal:

  [3]bool{true, true, false}

  And this creates the same array as above, then builds a slice that references it:

  []bool{true, true, false}

* For the array
  ```go
  var a [10]int
  ```
  these slice expressions are equivalent:

  ```go
  a[0:10]
  a[:10]
  a[0:]
  a[:]
  ```

* A slice has both a length and a capacity.

  The length of a slice is the number of elements it contains.

  The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

  The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

  You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

* 
  ```go
  var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!") // nil!
  }
  ```

* The make function allocates a zeroed array and returns a slice that refers to that array
  ```go
  func main() {
    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)

    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)
  }

  func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
      s, len(x), cap(x), x)
  }

  /* -output-
  a len=5 cap=5 [0 0 0 0 0]
  b len=0 cap=5 []
  c len=2 cap=5 [0 0]
  d len=3 cap=3 [0 0 0]
  */
  ```

* Slices can contain any type, including other slices.
  ```go
  // Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
  /* -output-
  X _ X
  O _ X
  _ _ O
  */
  ```

* **Appending to a slice** 
  
  `func append(s []T, vs ...T) []T`

  The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

  The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

  If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array

* **Range**
  ```go
  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

  func main() {
    for i, v := range pow {
      fmt.Printf("2**%d = %d\n", i, v)
    }
  }
  ```

* You can skip the index or value by assigning to _.

  `for i, _ := range pow`

  `for _, value := range pow`

  If you only want the index, you can omit the second variable.

  `for i := range pow`

* **Maps** A map maps keys to values.

  The zero value of a map is nil. A nil map has no keys, nor can keys be added.

  The make function returns a map of the given type, initialized and ready for use.

  ```go 
  type Vertex struct {
    Lat, Long float64
  }

  var m map[string]Vertex

  func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
      40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
  }

  // Map Literals
    var m = map[string]Vertex{
      "Bell Labs": Vertex{
        40.68433, -74.39967,
      },
      "Google": Vertex{
        37.42202, -122.08408,
      },
    }
  ```
* Insert or update an element in map m:

  m[key] = elem

  Retrieve an element:

  elem = m[key]

  Delete an element:

  delete(m, key)

  Test that a key is present with a two-value assignment:

  elem, ok = m[key]

  If key is in m, ok is true. If not, ok is false.

  If key is not in the map, then elem is the zero value for the map's element type.

  Note: If elem or ok have not yet been declared you could use a short declaration form:

  elem, ok := m[key]

* Just like ML, functions have first class treatment in Go.
  Syntax: 
  ```go
  func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
  }

  func main() {
    hypot := func(x, y float64) float64 {
      return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(3, 4)) // output 5
    fmt.Println(compute(hypot)) // output 5
  ```

* Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

  For example, the adder function returns a closure. Each closure is bound to its own sum variable.
  ```go
  func adder() func(int) int {
    sum := 0
    return func(x int) int {
      sum += x
      return sum
    }
  }

  func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 5; i++ {
      fmt.Println(
        pos(i),
        neg(-2*i),
      )
    }
  }
  /* -output-
  0 0
  1 -2
  3 -6
  6 -12
  10 -20
  */
  ```

* Go does not have classes. However, you can define methods on types.

  A method is a function with a special receiver argument.

  The receiver appears in its own argument list between the func keyword and the method name.

  ```go
  func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }

  func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
  }
  ```
  In this example, the Abs method has a receiver of type Vertex named v.

* You can declare a method on non-struct types, too.
  ```go

  type MyFloat float64

  func (f MyFloat) Abs() float64 {
    if f < 0 {
      return float64(-f)
    }
    return float64(f)
  }

  func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
  }
  ```
  In this example we see a numeric type MyFloat with an Abs method.

  You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

* **Pointer receivers**

  ```go
  func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
  }

  func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Abs()) // Output: 50
  }
  ```
  Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). 

  With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

* As a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver (but the same doesn't hold in case of functions!).

* While methods with value receivers take either a value or a pointer as the receiver when they are called. (which again doesn't hold true for true for functions!)

* "Pointer receiver" is preferred as we can do modifications and also we save memory as copying doesn't have to occur. 

* In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

* An **interface** type is defined as a set of method signatures.

  A value of interface type can hold any value that implements those methods.

  ```go
  // example
  type Abser interface {
    Abs() float64
  }

  func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser

    // In the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser, so un-commenting it will give an error.
    // a = v

    fmt.Println(a.Abs())
  }

  type MyFloat float64

  func (f MyFloat) Abs() float64 {
    if f < 0 {
      return float64(-f)
    }
    return float64(f)
  }

  type Vertex struct {
    X, Y float64
  }

  func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  ```

* If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

  In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

  ```go
  func (t *T) M() {
    if t == nil {
      fmt.Println("<nil>")
      return
    }
    fmt.Println(t.S)
  }
  ```

* A nil interface value holds neither value nor concrete type.

  Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

* The interface type that specifies zero methods is known as the empty interface:

  interface{}
  An empty interface may hold values of any type. (Every type implements at least zero methods.)

  Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

  ```go
    func main() {
    var i interface{}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
  }

  func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
  }
  /* -output-
  (<nil>, <nil>)
  (42, int)
  (hello, string)
  */
  ```

* ```go
  func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    f, ok := i.(float64) // if false, f will be zero value of type T in i.(T)
    fmt.Println(f, ok)

    f = i.(float64) // panic (its a runtime error), should therefore check with ok 
    fmt.Println(f)
  }
  ```
* ```go
  func do(i interface{}) {
    switch v := i.(type) {
    case int:
      fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
      fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
      fmt.Printf("I don't know about type %T!\n", v)
    }
  }

  func main() {
    do(21)
    do("hello")
    do(true)
  }
  /* -output-
  Twice 21 is 42
  "hello" is 5 bytes long
  I don't know about type bool!
  */
  ```

* ```go
  // Stringers!
  type Stringer interface {
      String() string
  }
  ```
  A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
  ```go
  // example
  type Person struct {
    Name string
    Age  int
  }

  func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
  }

  func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a, z)
  }
  ```

* The error type is a built-in interface similar to fmt.Stringer:
  ```go
  type error interface {
      Error() string
  }
  ```
  A nil error denotes success; a non-nil error denotes failure.
  ```go
  // Example
  type MyError struct {
    When time.Time
    What string
  }

  func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
      e.When, e.What)
  }

  func run() error {
    return &MyError{
      time.Now(),
      "it didn't work",
    }
  }

  func main() {
    if err := run(); err != nil {
      fmt.Println(err)
    }
  }
  /* -output-
  at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work
  */
  ```

* Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.
  ```go
  import (
    "fmt"
    "io"
    "strings"
  )

  func main() {
    r := strings.NewReader("Hello, Reader!")

    b := make([]byte, 8)
    for {
      n, err := r.Read(b)
      fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
      fmt.Printf("b[:n] = %q\n", b[:n])
      if err == io.EOF {
        break
      }
    }
  }
  ```
* A *goroutine* is a lightweight thread managed by the Go runtime.

  `go f(x, y, z)`

  starts a new goroutine running

  `f(x, y, z)`

  The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

  Goroutines run in the same address space, so access to shared memory must be synchronized. 

* ```go
  func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
      count()
      wg.Done()
    }()

    wg.Wait()
  }

  func count() {
    for i := 1; i <= 5; i++ {
      fmt.Println(i)
    }
  }
  ```

* Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
  ```go
  ch <- v    // Send v to channel ch.
  v := <-ch  // Receive from ch, and
            // assign value to v.
  ```
  (The data flows in the direction of the arrow.)

  Like maps and slices, channels must be created before use:

  `ch := make(chan int)`

  By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

  The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

  ```go
  func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
      sum += v
    }
    c <- sum // send sum to c
  }

  func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
  } // output -5 17 12
  ```

* ```go
  func main() {
    c := make(chan string)
    go count("sahiwal", c)
    for {
      msg, open := <-c
      if !open {
        break
      }
      fmt.Println(msg)
    }
  }

  func count(what string, c chan string) {
    for i := 1; i <= 5; i++ {
      c <- what
    }
    close(c) // always sender should close the channel
  }  
  ```
* ```go
  // syntactic sugar
  func main() {
    c := make(chan string)
    go count("sahiwal", c)
    for msg := range c {
      fmt.Println(msg)
    }  
  }
  // rest is same as before
  ```

* ```go
  // following will give deadlock error
  func main() {
    c := make(chan string)
    c <- "hello"
    msg := <-c
    fmt.Println(msg)
  } 

* Solution is to use *buffered channels*. Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty. So if the above code is modified as `c := make(chan string, 1)` it will work. 
  ```go
    func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
  }
  ```

* Not waiting for the slower one
  ```go 
  func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
      for {
        time.Sleep(time.Millisecond * 500)
        c1 <- "Every 500ms"
      }
    }()

    go func() {
      for {
        time.Sleep(time.Second * 2)
        c2 <- "Every two seconds"
      }
    }()

    for {
      select {
      case msg1 := <-c1:
        fmt.Println(msg1)
      case msg2 := <-c2:
        fmt.Println(msg2)
      }
    }
  }
  ```

* The default case in a select is run **only if** no other case is ready.

  ```go
  func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
      select {
      case <-tick:
        fmt.Println("tick.")
      case <-boom:
        fmt.Println("BOOM!")
        return
      default:
        fmt.Println("    .")
        time.Sleep(50 * time.Millisecond)
      }
    }
  }
  /* output: 
      .
      .
  tick.
      .
      .
  tick.
      .
      .
  tick.
      .
      .
  tick.
      .
      .
  BOOM!
  */
  ```
* Mutual Exclusion!
  ```go
  // SafeCounter is safe to use concurrently.
  type SafeCounter struct {
    v   map[string]int
    mux sync.Mutex
  }

  // Inc increments the counter for the given key.
  func (c *SafeCounter) Inc(key string) {
    c.mux.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    c.v[key]++
    c.mux.Unlock()
  }

  // Value returns the current value of the counter for the given key.
  func (c *SafeCounter) Value(key string) int {
    c.mux.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    defer c.mux.Unlock()
    return c.v[key]
  }

  func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
      go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey"))
  }
  ```