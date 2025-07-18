# GO in Development (go-in-dev) fast microservices
---
## Python vs Go
### *Similarities Between Go and Python**

| Aspect                   | Go and Python                                                                            |
| ------------------------ | ---------------------------------------------------------------------------------------- |
| **Open-source**          | Both are open-source and have large, active communities.                                 |
| **Cross-platform**       | Both can run on major OSs like Linux, macOS, and Windows.                                |
| **Garbage Collection**   | Both languages use garbage collection for memory management.                             |
| **Standard Libraries**   | Both have rich standard libraries for networking, file I/O, etc.                         |
| **Concurrency Support**  | Both can handle concurrency (Python with threading/multiprocessing; Go with goroutines). |
| **Readable Syntax**      | Both emphasize simplicity and readability in code design.                                |
| **Interpreted/Compiled** | Go compiles to binary, but both can run without a separate VM (unlike Java).             |
| **Popular for Web Dev**  | Both are widely used for backend web development (e.g., Flask for Python, Gin for Go).   |

---

### **Differences Between Go and Python**

| Feature / Aspect          | **Python**                              | **Go (Golang)**                                  |
| ------------------------- | --------------------------------------- | ------------------------------------------------ |
| **Typing**                | Dynamically typed                       | Statically typed                                 |
| **Performance**           | Slower (interpreted, dynamic typing)    | Faster (compiled to machine code)                |
| **Compilation**           | Interpreted (via CPython, etc.)         | Compiled (to native binaries)                    |
| **Concurrency Model**     | Threads, asyncio (more complex)         | Goroutines and channels (lightweight and simple) |
| **Syntax**                | Very high-level, flexible               | Simple, strict (e.g., no generics until v1.18)   |
| **Error Handling**        | Exceptions (try/except)                 | Explicit error returns (no exceptions)           |
| **Use Cases**             | Scripting, data science, ML, automation | Systems programming, microservices, DevOps       |
| **Object-Oriented**       | Yes, supports OOP fully                 | Limited OOP (no inheritance, only composition)   |
| **Runtime Dependency**    | Requires Python interpreter installed   | Binaries are standalone (no interpreter needed)  |
| **Learning Curve**        | Easier for beginners                    | Slightly harder, but very beginner-friendly      |
| **Community & Ecosystem** | Larger, especially in data/ML fields    | Growing, especially in cloud/devops              |
| **Tooling**               | Virtualenv, pip                         | Built-in tooling (go build, go test, go fmt)     |
| **Memory Management**     | Garbage collected                       | Garbage collected, but with more control         |
| **Development Speed**     | Faster for prototyping                  | Slower for prototyping, faster in production     |

---

### **When to Choose Which?**

* **Use Python if**:

    * You’re building data science or ML applications
    * Rapid prototyping and flexibility are priorities
    * You prefer dynamic typing and extensive libraries (NumPy, TensorFlow, etc.)

* **Use Go if**:

    * You need high performance with minimal memory footprint
    * You’re building web servers, networking tools, or CLI apps
    * Concurrency and scalability are key concerns (e.g., microservices)

---

Here’s a clean side‑by‑side comparison of the **“Python Basics”** section from that repo versus how you’d approach the same in **Go**:

---

### 1. **Dynamic typing & variable declaration**

| Python                                                                     | Go                                                                                                                        |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `a = 1 + 1` – type inferred at runtime; `type(a)` shows `int` | `var a = 1 + 1` or `a := 1 + 1`. Types inferred at compile‑time, static typing. To inspect type: `fmt.Printf("%T\n", a)`. |

---

### 2. **Strings & escapes**

| Python                                                                                      | Go     |
| ------------------------------------------------------------------------------------------- | ------ |
| `"hello\nworld"`, indexing `myString[0]`, slicing with steps `mystring[2:7:2]` | In Go: |

```go
s := "hello\nworld"  
fmt.Println(s[0])               // prints byte value  
r := []rune(s)[0]              // prints first rune  
substr := string([]rune(s)[2:7])  
```

No step slicing; convert to `[]rune` for proper indexing. |

---

### 3. **Immutable strings & repetition**

| Python                                                              | Go                                |
| ------------------------------------------------------------------- | --------------------------------- |
| Strings are immutable; `'a'*10` repeats `'aaaaaaaaaa'` | Strings immutable too. To repeat: |

```go
s := strings.Repeat("a", 10)  
```

Need `strings` import. |

---

### 4. **String methods & formatting**

| Python                                                                                                 | Go     |
| ------------------------------------------------------------------------------------------------------ | ------ |
| `'x'.upper()`, `'hi there'.split()`; formatted strings: `.format()`, `f"He said {name}"` | In Go: |

```go
strings.ToUpper("x")  
strings.Split("hi there", " ")  
fmt.Sprintf("He said his name is %s.", name)
```

Use `fmt` and `strings` packages. |

---

### 5. **Length, concatenation**

| Python                                         | Go                                                      |
| ---------------------------------------------- | ------------------------------------------------------- |
| `len(myString)` or `len(myList)` | Go uses the same built‑in: `len(s)`, `len(slice)`, etc. |

---

### 6. **Lists**

| Python                                                                            | Go          |
| --------------------------------------------------------------------------------- | ----------- |
| `mylist = ['one','two']`; `append()`, `pop()`, `sort()`, `reverse()`| Use slices: |

```go
list := []string{"one", "two"}  
list = append(list, "three")  
last := list[len(list)-1]  
list = list[:len(list)-1]  // pop  
sort.Strings(list)  
// reverse manually using a loop
```

Use `sort` package for sorting. |

---

### 7. **Dictionaries**

| Python                                                                               | Go       |
| ------------------------------------------------------------------------------------ | -------- |
| `dic = {'apple':2.99}`; access `dic['apple']`; `dic.keys()`, `values()` | Go maps: |

````go
m := map[string]float64{"apple":2.99}
m["apple"]
for k := range m { /* keys */ }
for _, v := range m { /* values */ }
````

---

### 8. **Tuples**  
| Python | Go |
|---|---|
| Immutable `(1,3,4)`; unpacking: `x1,x2 = coord` | No tuples, but multiple returns:  
```go
return x1, x2
// capture: a, b := f()
````

For constant pairs, define a `struct`.

---

### 9. **Sets**

| Python                                          | Go                          |
| ----------------------------------------------- | --------------------------- |
| `myset = set()`; `add()`, unique from iterable  | Use `map[T]struct{}` idiom: |

```go
set := make(map[rune]struct{})
set['a'] = struct{}{}
```

Iterate keys for unique elements. 

---

### 10. **Booleans**

| Python           | Go                     |
| ---------------- | ---------------------- |
| `True`, `False`  | `true`, `false` in Go. |

---

### 11. **File I/O**

| Python                                                                       | Go  |
| ---------------------------------------------------------------------------- | --- |
| Uses `open()`, `.read()`, `.seek()`, context manager `with open(...) as f:`  | Go: |

````go
data, err := os.ReadFile("file.txt")
f, err := os.Open("file.txt"); defer f.Close()
buf, err := io.ReadAll(f)
// For writing or streaming, use bufio.Reader/Writer
````

### ✅ **Summary**  
Despite differences in typing, concurrency, and some language idioms, **most Python basics map clearly to Go equivalents**. Go offers:

- **Static typing**, compile-time safety
- **Slicing & indexing** via `[]rune`
- **String methods** via `strings` and formatting with `fmt`
- **Lists** as slices, with manual operations
- **Maps**, `structs`, and idiomatic set patterns
- **Built‑in I/O** with `os` and `io`

### Method and Function in both

**Python**

```python
# Function
def greet(name):
    return f"Hello, {name}"

# Method inside a class
class Greeter:
    def __init__(self, name):
        self.name = name

    def greet(self):  # 'self' refers to instance
        return f"Hello, {self.name}"
```

**Go**

```go
// Function
func Greet(name string) string {
    return "Hello, " + name
}

// Method on a struct
type Greeter struct {
    Name string
}

func (g Greeter) Greet() string {  // receiver: 'g' is like 'self'
    return "Hello, " + g.Name
}
```

## Method in Go

A **method** in Go is just a **function with a receiver** — meaning it's **bound to a specific type**, similar to how methods in Python are bound to objects via `self`.

---

### 🛠️ Go Method Syntax

```go
func (receiver ReceiverType) MethodName(params...) ReturnType {
    // method body
}
```

* `receiver`: like `self` in Python — it’s how the method gets access to the data inside the struct.
* `ReceiverType`: the type you're attaching the method to (usually a `struct`).
* `MethodName`: the name of the method.
* `params...`: any additional arguments.
* `ReturnType`: what the method returns.

---

### 🔍 Example: Value Receiver

```go
type Person struct {
    Name string
}

// Method with a value receiver
func (p Person) Greet() string {
    return "Hello, " + p.Name
}
```

* Here, `p` is a **value receiver**.
* It gets a **copy** of the `Person`, not the original.

Usage:

```go
john := Person{Name: "John"}
fmt.Println(john.Greet())  // Output: Hello, John
```

---

### Example: Pointer Receiver

```go
func (p *Person) Rename(newName string) {
    p.Name = newName
}
```

* `*Person` means it's a **pointer receiver**.
* It modifies the **original** `Person` instance (not a copy).

Usage:

```go
john := Person{Name: "John"}
john.Rename("Johnny")
fmt.Println(john.Name)  // Output: Johnny
```

---

### When to Use Value vs Pointer Receivers

| Use Value Receiver         | Use Pointer Receiver         |
| -------------------------- | ---------------------------- |
| Method doesn't modify data | Method modifies data         |
| Struct is small            | Struct is large (efficiency) |
| You want a copy            | You want to avoid copying    |

---

### 🔗 Summary

* A **method** in Go is a **function tied to a type** (usually a struct).
* The method gets the type via a **receiver** (value or pointer).
* This is Go’s way of supporting behavior on types — similar to object-oriented programming.

## Interface in Go

Let’s dive into how **Go methods** relate to **interfaces**, which is one of the most powerful and idiomatic parts of the language.

---
An **interface** in Go defines a **set of method signatures**.
Any type that implements **all the methods** in an interface **implicitly satisfies it** — no need for explicit `implements`.

---

### Interface Syntax

```go
type Greeter interface {
    Greet() string
}
```

* This interface says: "Any type with a `Greet() string` method is a `Greeter`."

---

### Implementing the Interface (implicitly)

```go
type Person struct {
    Name string
}

// Person has a method Greet() string
func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}
```

Now, `Person` **satisfies the `Greeter` interface**, even though we never said so explicitly.

---

### Using the Interface

```go
func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    john := Person{Name: "John"}
    SayHello(john)  // Works because Person implements Greeter
}
```

* The function `SayHello` doesn’t care what the actual type is — as long as it satisfies `Greeter`.

---

### Why This is Powerful

* Go has **no inheritance** — but interfaces give you **polymorphism**.
* Go uses **composition over inheritance**.
* Interfaces enable **loose coupling** and clean code.

---


In Python:

```python
class Greeter:
    def greet(self):
        raise NotImplementedError

class Person(Greeter):
    def greet(self):
        return "Hi, I'm John"
```

In Go:

```go
type Greeter interface {
    Greet() string
}

type Person struct{ Name string }

func (p Person) Greet() string {
    return "Hi, I'm " + p.Name
}
```

Except: **no base class needed**. Just define the method, and you're in.

---

## ⚠️ Interface Tips

| Best Practice                                                | Why?                                   |
| ------------------------------------------------------------ | -------------------------------------- |
| **Program to interfaces**                                    | Makes your code flexible and testable  |
| **Keep interfaces small**                                    | Prefer 1–2 methods (e.g., `io.Reader`) |
| **Use interfaces in function parameters, not struct fields** | Avoid tight coupling                   |


## Go env setup
- install go binaries in mcbook from golang.org
- install vs code - native support for go development
- `go run main.go`
- `go build main.go`
- `./main`

## Go topics

- types of package in go
  - executable i.e package called 'main' is reserved for executable ones
  - reusable
- go built-in packages
  - math
  - fmt: see `golang.org/pkg`
  - io
  - encoding
  - crypto
  - debug

```
package main

import "fmt"

// This is a simple Go program that prints "Hello, World!" to the console.
// The main function serves as the entry point of the program.

func main() {
	fmt.Println("Hello, World!")
}
```
- go data types:
  - string
  - int
  - bool
  - float64
  - array - fixed size 
  - slice - an array that can grow or shrink

