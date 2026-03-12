# Learning Go Programming language through [BootDev](https://boot.dev) platform

## Datatypes

- `bool`
- `int`
- `byte` = Just like `buffer` in nodejs. It can contains 8 bit of data
- `float64`
- `string`

## Zero Values

In GO programming language, zero value of any variable is the value of that variable's type when only declared without initialization.

| Type                                           | Zero Value                   |
| ---------------------------------------------- | ---------------------------- |
| `int`, `float64`                               | `0`                          |
| `bool`                                         | `false`                      |
| `string`                                       | `""` (empty string)          |
| `pointer`, `slice`, `map`, `chan`, `interface` | `nil`                        |
| `struct`                                       | each field is its zero value |

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

> [!TIP]
> Empty Structs are the smallest possible type in Go. It takes 0 bytes of memory

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

> [!NOTE]
> In Go, all values are passed by value unless explicitly passed by reference. When you pass an array to a function, a copy of the entire array is made, which can be inefficient for large arrays. However, when you pass a slice, you are passing a copy of the slice header (which contains a pointer to the underlying array, its length, and capacity). Since the slice header is small and points to the same underlying array, it behaves like passing by reference. Modifying the slice within the function will affect the original data in the underlying array.

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

> [!CAUTION]
> Always use `append` function on the same slice the result is assigned to. Otherwise, it can lead to unexpected behavior if underlying array has enough capacity to fill new elements, and more than 1 slice is pointing to same underlying array. As, `append` function changes underlying array if it has enough capacity for newer elements.

![](/assets/2026-03-08-14-29-33.png)

## Maps

- In GO programming language, maps are built-in data structures that provide a way to store and retrieve key-value pairs.
- Similar to slices, they also behave like passed by reference to a function call due to its internal structure.
- In GO, except `slices`, `maps` and `functions`, all other types are comparable using `==` operator such as `bool`, `int`, `string`, `pointer`, `channel`, and `interface`. Even `struct` and `array` are also comparable if all their fields or elements are comparable. According to language specification, all comparable types are eligible to be used as keys in a map.
- We can distinguish between zero value and non existence of key in map using `comma ok` idiom. When we try to access a key in map, it returns two values: the value associated with the key (or zero value if key not present) and a `bool` indicating whether the key exists in the map or not.
- We can also implement a set like data structure using a map with value type `bool` or empty structs. As, indexing of maps returns zero value of type along with `false` indicating non existence of key using `comma ok` idiom. Using empty struct as value type is more memory efficient than using `bool` as value type, as empty struct takes 0 bytes of memory.

```go
type Key struct {
    Path, Country string
}

// for storing number of hits by each path and country combination
hits := make(map[Key]int)
hits[Key{"/", "vn"}]++
n := hits[Key{"/ref/spec", "ch"}]
```

## Strings, bytes, runes and characters in Go

- [ ] [GO Strings, Bytes, Characters and Runes](https://go.dev/blog/strings)
- [ ] [Absolute Minimum to Know about Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

## Pointers

Recalling, I learnt following in C++.

![](/assets/2026-03-08-17-56-54.png)

- In GO, there is no concept of reference variable. So, `&` operator is only used to get the memory address of a variable.
- In GO, there are two use cases of `*` operator:
  - To declare a pointer variable, we use `*` before the type. For example, `var p *int` declares a pointer to an integer.
  - To dereference a pointer, we use `*` before the pointer variable. For example, if `p` is a pointer to an integer, `*p` gives us the value stored at that memory address.

> [!NOTE]
> Selector expressions = Shorthand syntax to access fields of structs. Eg. `analytics.MessagesTotal = (*analytics).MessagesTotal`. We were using this unconsciously using receiver functions on structs. They are also called **Pointer Receivers**.

**_NOTE_**: Deferencing a `nil` pointer can lead to crashing program at runtime due to `panic`. It is always good practice to check if pointer is `nil` before dereferencing it.

## Packages and Modules

> [!WARNING]
> Only capitalized names are exported, meaning they can be accessed by other packages. Uncapitalized names are private.

- Every go file belongs to a package. Till now, we were writting `package main` in our go files. This main package generally have entry point `main()` function. Any other package name is just a library to export a set of functionalitues. There is no concept of entry point in other packages. Only, the `main` package is converted into executable binary when compiled. A package is just a directory of go code that's all compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package

> Package names convention = The package name should be same as the last element of the import path. For example, if we have a package located at `github.com/username/mypackage`, the package name should be `mypackage`. Although, any package name is allowed by compiler, but it is discourage for the sake of readability and consistency in codebase. We can use local aliases like `import m "github.com/username/mypackage"` to import package with different name in our codebase.

- A repository can have multiple modules, and each module can have multiple packages. A module is a collection of related Go packages that are versioned together as a single unit. A file named `go.mod` is used to define a module and its dependencies. It has the following things mainly:
  - The module path, which is the import path prefix for all packages within the module.
  - The Go version that the module is compatible with.
  - A optional list of dependencies, which are other modules that the current module depends on. Each dependency is specified with its module path and version.

> **_IMPORT PATH = MODULE PATH + PACKAGE PATH._**. The packages in the standard library do not have a module path prefix, as they are built into the go distribution itself. The path of package's directory determines the import path and download path of the package in case of remote packages.

> [!TIP]
> `go run` command is used to quickly compile and run go package named `main`. The compiled binary is stored in temporary directory and is deleted after execution. Use it for quick debugging and testing on local

#### Remote Packages

For local development, we can use `replace` directive in `go.mod` file to replace the module path with a local directory path. This is useful when we are developing a module locally and want to test it without publishing it to a remote repository. But, it is advisable to publish the modules to remote repositories like github, gitlab etc.

```
WHERE SHOULD GO ACTUALLY FIND IT?
replace example.com/username/mystrings v0.0.0 => ../mystrings =========> This directive is used to replace the module path with a local directory path.

WHAT DOES MY MODULE DEPEND ON?
require example.com/username/mystrings v0.0.0
```

#### Clean Packages

We should follow the OOPS principle of encapsulation while designing packages.

- We should only export the necessary functions, types, variables and constants that are required for the users of the package.
- We should also avoid exporting any internal implementation details that are not required for the users of the package. This will help in maintaining the integrity of the package and prevent any unintended usage of the package's internal details.
- The `main` package should not be treated as library, and there's no need to export functions from it. A clean package should never have specific knowledge about a particular application that uses it.

> [!NOTE]
> The unexported functions within a package can and should change often for testing, refactoring, and bug fixing. A well-designed library must have a stable API so that users don't get breaking changes each time they update the package version. In Go, this means not changing exported function’s signatures.
