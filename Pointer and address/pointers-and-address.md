# The Point of Pointers and Addresses
 
Go is a pass-by-value language. In other words, we’re passing functions the value of an argument. In a technical sense, when we’re calling a function with an argument, the Go compiler is strictly using the value of the argument rather than the argument itself. Because of this feature (pass-by-value), the changes that take place in our function, stay within that function.

But, we do have the ability to change values from different scopes. To do so, we need to make use of:

- addresses
- pointers
- dereferencing

## addresses

Every time we use a variable, what we’re doing is retrieving the value stored at the variable’s address. 

To find a variable’s address we use the & operator followed by the variable name, like so

When we see the 0x prefix, this means that the number is in formatted in hexadecimal, which is a way of representing 16 digit numbers. Thus, the 0x414020 is actually the address of x in hexadecimal format.

```go
x := "My very first address"
fmt.Println(&x) // Prints 0x414020
```

## pointers

Pointers are variables that specifically store addresses. To break it down further, the `*` operator signifies that this variable will store an address and the int portion means that the address contains an integer value.

```go
var pointerForInt *int

minutes := 525600

pointerForInt = &minutes

fmt.Println(pointerForInt) // Prints 0xc000018038
```

```go
minutes := 55

pointerForInt := &minutes
```

## Dereferencing

We know that addresses are where values are stored and pointers allow us to keep track of addresses. But what if we want the address to store a different value? Well, we can actually use our pointer to access the address and change its value! This action is called dereferencing or indirecting.

```go
lyrics := "Moments so dear"
pointerForStr := &lyrics

\*pointerForStr = "Journeys to plan"

fmt.Println(lyrics) // Prints: Journeys to plan
```
In our example, we have our variables: lyrics that has the value of "Moments so dear" and pointerForStr which is a pointer for lyrics. Then we use the \* operator on pointerForStr to dereference it and assign a new value of "Journeys to plan". When we check the value of lyrics it’s now "Journeys to plan"!

## Changing Values in Different Scopes


```go
func addHundred(num int) {
num += 100
}

func main() {
x := 1
addHundred(x)
fmt.Println(x) // Prints 1
}
```
Even though we call addHundred(x), the value of x doesn’t change! Why is that?

Remember, Go is a pass-by-value language. When we call addHundred(x) we’re providing addHundred() with a value of 1. We’re not actually providing the address of x for addHundred() to go in and change the value stored there.

If we want to change the value of x using a function, we’re going to need to first change our function:
Our new function now has a parameter of a pointer for an integer. By passing the value of a pointer (which is an address) to addHundred(), we can also dereference the address and add 100 to its value. But now that addHundred() expects a pointer for an argument, we’re also going to need to change our main()! The complete code is as follows:

```go
func addHundred (numPtr *int) {
*numPtr += 100
}

func main() {
x := 1
addHundred(&x)
fmt.Println(x) // Prints 101
}

```