# Go lang learning

## Advantages of Go lang

- Code runs fast
- Garbage collection
- Simpler objects
- Concurrency is efficient

###= Software Translation

- **Machine Language** : CPU instructions represented in binary
- **Assembly Language** : CPU instructions with mnemonics 
    - Easier to read
    - Equivalent to machine language
- **High Level Language** : Commonly used languages (C, C++, Java, Python, Go)
    - Much easier to use


> All software must be translated into the machine language of processors

## Compiled vs Interpreted

- **Compilation** : Translate instructions once before running the code 
    - C, C++, Java (partially)
    - Translation occurs only once, saves time

- **Interpretation** :  Translate instructions while code is executed
    - Python, Java (partially)
    - Translation occurs every execution
    - Requires an interpreter


## Efficiency vs Ease-of-Use

- Compiled code is fast
- Interpreters make coding easier 
    - Manage memory automatically
    - Infer variable types 

- Go is a good compromise

## Garbage Collection

- Automatic memory management
    - where should memory be allocated?
    - when can memory be deallocated?

- Manual memory management is hard
    - Deallocate too early, false memory access
    - Deallocate too late, wasted memory

**Go includes garbage collection**
- Typically only done by interpreters

## Object-Oriented Programming

- Organize your code through encapsulation
- Group together data and functions which are related
- User-defined type which is specific to an application

## Objects in Go

- Go does not use the term `class`
- Go uses `structs` with associated methods
- Simplified implementation of classes
    - No inheritance
    - No constructor
    - No generics

## Performance limits

- _Moore's law_ used to help performance
    - Number of transistors doubles every 18 months

- More transistors used to lead to higher clock frequencies
- **Power/temperature constraints** limit clock frequencies now

### Parallelism
- Number of cores still increases over time
- Multiple tasks may be performed at same time on different cores

- Difficulties with parallelism
    - when do tasks start/stop?
    - What if one task need data from another task?
    - do tasks conflict in memory?

### Concurrent programming
- **Concurrency** is the mngmt of multiple tasks at the same time
- Key requirements of large systems
- Concurrent programming enables parallelism
    - Management of task execution
    - Communication between tasks
    - Synchronization between tasks

## Concurrency in GO

- Go includes concurrency primitives
- **Goroutines** represent concurrent tasks
- **Channels** are used to communicate between tasks
- **Select** enable task synchronization
- Concurrency primitives are efficient and easy to use

## Variables

- Data stored in memory
- Must have a **name** and a **type**
- All variables must have **declarations**
- Most basic variable declaration 
```
                var       x       int
                /         |         \
               /          |          \
           keyword      name        type
```

- Can declare many on the same line
    var x, y int

### Variable types
- Types define the values a variable may take and operations that can be performed on it.
- Integer
    - Integral values
    - Integer arithmetic (+, -, *, /, ...)

- Floating point
    - Fractional (decimal) values
    - Floating point arithmetic (may use different hardware)

- Strings
    - Byte (character) sequences
    - String comparison, search ...

## Pointers
- A pointer is an address to data in memory
- **&** operator returns the address of a variable/function
- _*_ operator returns data at an address(dereferencing)

```go
var x int = 1
var y int
var ip *int // ip is a pointer to int

ip = &x // ip now points to x

y = *ip // y is now 1
```

## New 

- Alternate way to create a new variable
- **new()** function creates a variable and returns a pointer to the variable
- variable is initialized to zero
    ```go
    ptr = new(int)
    *ptr = 3
    ```