# Learning Go Programming language through [BootDev](https://boot.dev) platform

## Datatypes

- `bool`
- `int`
- `byte` = Just like `buffer` in nodejs. It can contains 8 bit of data
- `float64`
- `string`

## CPU Performance and Memory Management

- Go compiles faster than other compiled languages like C, Rust, C++ but its execution speed is slower than them due to go runtime. Whenever a go file is compiled, a small runtime is also always added into every go executable. This runtime is responsible for memory management, goroutine scheduling and other features of the language. But, it is much faster then interpreted languages like Python and JavaScript.

![](/assets/2026-02-28-10-53-16.png)

- In terms of execution speed, it is faster than interpreted languages like Python and JavaScript but slower than compiled languages like C, C++ and Rust due to the presence of go runtime (automatic memory management e.g.) which adds some overhead to the execution.

![](/assets/2026-02-28-11-28-15.png)

- In terms of memory usage as well, it is performant then Java, C# as it doesn't need a virtual machine to run but have automatic garbage collection like those, but is less performant then C, C++ and Rust as they don't have a runtime and automatic garbage collection but needs manual memory management. It also does allocate more data in stack more often than heap which is faster to access and manage.

![](/assets/2026-02-28-10-55-46.png)

#### Memory Layout of programs

![](/assets/2026-02-28-11-06-19.png)

1. **Stack**: It is a region of memory that stores local variables, function parameters, return addresses. It is typically faster to access than the heap because of its contiguous memory allocation. It has limited size and can overflow if too much data is stored on it (like deep recursion), which is known as a stack overflow. It is automatically allocated when functions are called and deallocated when they return.

2. **Heap**: It is a region of memory that is used for dynamic memory allocation. It is managed by the runtime and can grow or shrink as needed. The heap is used for storing data that needs to persist beyond the scope of a function, such as objects or data structures. It is typically slower to access than the stack because of its non-contiguous memory allocation. It can also lead to memory leaks if not managed properly, and it required GC overhead to manage memory.
3. **Data Segment**: It is a region of memory that stores global variables and static variables. It is divided into two parts: the initialized data segment, which contains variables that are initialized with a value, and the uninitialized data segment (also known as the BSS segment), which contains variables that are not initialized. Its lifecycle is the entire duration of the program.

4. **Text Segment**: It is a region of memory that contains the executable code of the program. It is typically read-only and is shared among all instances of the program. Its lifecycle is the entire duration of the program.

## Constants in Go

Constants are the variables which are immutable while the program is running. The walrus operator `:=` cannot be used to declare constants. Go has kept this syntactical difference to make it clear that the value is a constant and its value must be known at the compile time, but the `var` are mutable and their values can be changed while runtime. We can use constants for computed values as long as they can be known to compiler at compile time.

![](/assets/2026-02-28-11-22-46.png)

## Printing in Go

Go follows C style of printing using `fmt` package along with formatting verbs like `%v`, `%d`, `%f` etc.

## Defer keyword in Go

`defer` keyword is a special feature in Go lang to defer a function's execution until the enclosing function returns. All deferred functions are executed in `LIFO`. Think of like all deferred functions goes into stack, and each function is popped and executed one by one before enclosing function returns itself. It is best way to run a piece of code just before each return statement by writting it once like closing db connections, or cleaning up resources.

## Datastructures in GO

#### Structs

Structs are custom data types that allow you to group together related data. They are similar to objects in JS or dictionaries in Python. Structs can also contains other structs (known as nested structs) as fields for creating complex data structures.

```go
type message struct { // Uppercase Name = Exported (Public), Lowercase Name = Unexported (Private)
    text: string
    sender: user
    recipient: user
}

type user struct {
    name string
    age int
}

type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct = Always prefer named struct. Use this only when you need it just once, and not reusable anywhere else in the codebase.
 wheel struct {
    radius int
    material string
  }
}

type car struct {
  brand string
  model string
}

type truck struct {
    // embedded struct for just data only inheritance (actually composition). the fields of embedded struct can be accessed directly from the top level unlike nested struct
    car
    bedSize int
}

// Methods on structs
func (c car) upgradeModel() {
    c.model = "New Model" // This will not change original model, as the whole struct is passed by value (completely new copy of struct is passed)
}

func (c *car) upgradeModel() {
    c.model = "New Model" // This will change original model, as the struct is now passed by reference (pointer to original struct is passed)
}
```

###### Memory layout of struct

In Go, the struct sit in memory in contiguous block with fields one after another. For efficient execution speed, Go may add padding between fields to ensure proper memory alignment.

> Alignment = Size of the primitive type (upto CPU's word size). For eg. for 64 bit architecture, the word size is 8 bytes. Modern CPUs don't read 1 byte at a time, but the fixed size chunks called words. When the CPU wants to fetch data, it fetches an entire word at a time from an address that is a multiple of the word size. This is called memory alignment. Thus, it means the each field of struct must start on address i.e. mutiple of its own size (upto word size). The total size of the struct is finally rounded up to a multiple of largest alignment of any of its fields.

![](/assets/2026-03-01-16-52-41.png)

So, the order of fields in struct can affect the memory layout and performance. To minimize padding and optimize memory usage, it is generally recommended to order fields from largest to smallest size.

> Empty Structs are the smallest possible type in Go. It takes 0 bytes of memory
>
> ```go
> type empty struct {}
> empty 1 := empty{}
> empty2 := struct{}{}
> ```

## Unique way of error handling in GO

The error handling in Go is very unique compared to other programming languages. There is no try catch thing... or catching exceptions. Instead, an error is just treated as straightword normal return value.

```go
type error interface {
    Error() string
}
```

Each function should be responsible for its error handling, and returning the error appropriately `nil` or non `nil` value. The caller of the function can decide the behavior of each function's error handling differently based on whether the error value returned is `nil` or not.

#### Panic and Recover in error handling

- There is a mechanism for creating panic in call stack using `panic` function in GO. Doing so, will result in unwinding the calls from call stacks by popping them out one by one. But, while this happens, Go ensure to call each defer function in order to ensure proper cleanup of resources created within those function calls.
- We can also recover a `panic` using `recover` function. The `recover` function can only be called in deferred function. Functions till `panic` is recovered are popped out (ensuring their corresponding deferred functions are invoked in correct order), and normal execution of parent function (of function that recovered the `panic`) resumes.
- In case of no function defering the `recover`, the goroutine crashes after all function calls are popped and their deferred functions are executed.
- Using `log.Fatal` is instead cleaner way to exit the code. In this case, no deferred functions are executed. In is basically similar to `Print` followed by a call to `os.Exit(1)`.

## Slices

Arrays are fixed size contiguous block of memory, while slices are dynamic size and more flexible. Slices are built on top of arrays and provide a more convenient way to work with sequences of data. They have three components: a pointer to the underlying array, the length of the slice (number of elements in the slice), and the capacity of the slice (the maximum number of elements that can be stored in the underlying array starting from the pointer). When you append to a slice and it exceeds its capacity, Go automatically creates a new underlying array with double the capacity, copies the existing elements to it, and updates the slice's pointer and capacity accordingly.

#### Variadic Function

Just like a spread operator in javascript. We can any aribitrary number of arguments to a function using `...` syntax. The arguments becomes a slice of the specified type. Eg. `fmt.Println`, `fmt.Sprintf` etc. are variadic functions.

```go
func concat(strs ...string) string { // SPREAD OPERATOR - to accept variable number of arguments as slice
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
    names := []string{"bob", "sue", "alice"}
    printStrings(names...) // UNSPREAD OPERATOR - to pass slice as variadic arguments
}
```

#### `make` and `append` functions

- The `make` function is used to create slices, maps, and channels in Go. This is used to pre-allocate slice with a specified length and capacity. It is more efficient than using `append` to create a slice from an array when you know the size of the slice in advance, as it avoids the overhead of multiple allocations and copying that can occur when appending to a slice.
- The `append` function is used to add elements to the end of a slice. It is variadic function as well. It takes a slice and one or more values to append, and returns a new slice with the values added. If the capacity of the original slice is exceeded, `append` will create a new underlying array with double the capacity, copy the existing elements to it, and update the slice's pointer and capacity accordingly.
