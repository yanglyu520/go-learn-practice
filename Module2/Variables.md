# Variables and types
## 1. Literals
We can perform arithmetic in Go with literals (or named values, covered in the next exercise) using the following operators:

- to add
* to subtract
- to multiply
/ to divide
% to take the remainder (the modulus operator) between two numbers.
## 2. Constants
use camelCase or PascalCase for constants

```go
const funFact = "Hummingbirds' wings can beat up to 200 times a second."

fmt.Println(funFact)
```

## 3. Data Type
You can use default data types that go provides or create your own data types.
- Integers, `int`
- Floating-point numbers, or `float`, can include fractional data. You would use a float to store distances, percentages, and other quantities that required division or precision.
- Complex numbers, `complex`, are pairs of floating-point numbers where the second part of the pair is marked with the “imaginary” unit i. Complex numbers are particularly useful when reasoning in 2-dimensional space and have other utilizations that make them relevant for involved calculations

## 4. Numeric Types
Go has 15 different numeric types that fall into the three categories: int, float, and complex.  This includes 11 different integer types, 2 different floating-point types, and 2 different complex number types. These types all recognize different sets of numbers as valid. 

Integers are further broken down into two categories: signed and unsigned. **Signed integers can be negative, but unsigned integers can only be positive.** This means that the minimum value for an unsigned integer is always 0. Since it can ignore the possibility of a negative value, an unsigned integer’s maximum value is much higher than a signed integer with the same number of bits.

Go also has a boolean type. Booleans are either false or true. **Go only needs one bit to store a boolean value: 0 represents false and 1 represents true**

## 5. Default int Type

Computers actually have a default length for the data in their Read-Only Memory (ROM). Some newer comps may have more processing power and can store/handle bigger chunks of data.

By providing the type int or uint, Go will check to see if the computer architecture is running on is 32-bit or 64-bit. Then it will either provide a 32-bit int (or uint) or a 64-bit one depending on the computer itself.


## 6. Strings

“string” is a term for text of any length, and the name is chosen to evoke a sequence of data.

Surround the text you want to store with **double-quotation** marks (i.e., ", the single-quote ' has a special other meaning and isn’t used to define strings)

You can use the + operator on strings to join them, this is known as string concatenation.

```go
var description string
description = nameOfSong + " is by: " + nameOfArtist + "."
fmt.Println(description)
```
## 7. Zero Values

Even before we assign anything to our variables they hold a value. Go’s designers attempted to make these “sensible defaults” that we can anticipate based on the variable’s types. All numeric variables have a value of 0 before assignment. String variables have a default value of "", which is also known as the empty string because it contains no characters. Boolean variables have a default value of false.

## 8. Variable
A variable is a named value (like a constant) with the added feature that it can change during the running of a program. If we’re working with a value in various places in our program, we can store that value in a variable to easily access it later.

Variables are defined with the var keyword and two pieces of information: **the name that we will use to refer to them and the type of data stored in the variable.** Since variables can be updated we don’t even need to assign a value initially. Here’s a couple of variable definitions:

``` go
var lengthOfSong uint16
var isMusicOver bool
var songRating float32
```
```go
var kilometersToMars int32

kilometersToMars = 62100000
```
```
var kilometersToMars int32 = 62100000
```

In addition to += (yes, pun intended), Go has other arithmetic operations that perform calculations and update the variable’s value:

-= to subtract from the variable.
\*= to multiply the variable by a factor.
/= to divide the variable by a dividend.

## 9. Multiple Variable Declaration

Go actually allows us to declare multiple variables on a single line, in fact, there’s a few different syntaxes!

```
var part1, part2 string
part1 = "To be..."
part2 = "Not to be..."
```

```
quote, fact := "Bears, Beets, Battlestar Galactica", true
fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica
fmt.Println(fact) // Prints: true
```


## 10. **Inferring Type**
There is a way to declare a variable without explicitly stating its type using the short declaration := operator.
We might use the := operator if we know what value we want our variable to store when creating it.

Floats created in this way are of type float64. Integers created in this way are either int32 or int64

```
nuclearMeltdownOccurring := true
radiumInGroundWater := 4.521
daysSinceLastWorkplaceCatastrophe := 0
externalMessage := "Everything is normal. Keep calm and carry on."
```

```
var nuclearMeltdownOccurring = true
var radiumInGroundWater = 4.521
var daysSinceLastWorkplaceCatastrophe = 0
var externalMessage = "Everything is normal. Keep calm and carry on."
```

## 11. How to read go Errors

When the Go compiler raises an error the program’s binary cannot get created and without the binary, our computer cannot execute the program’s code.

Go tries to tell you what the issue is by offering you the following pieces of information: **the name of the file where something went wrong, the line number in that file where Go noticed an issue, and the column number** (the number of characters from the left) on that line where the error occurred. One common error occurs when Go language’s compiler recognizes that _there is a defined variable that isn’t used in the program._
