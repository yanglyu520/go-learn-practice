# Function
In programming, a function is a block of code designed to be reused.

## Function
Scope is a concept that refers to where the values and functions are defined and where they can be accessed. For instance, when a variable is defined within a function, that variable is only accessible within that function. When we try to access that same variable from a different function, we get an error because we can’t do it.

When we return a value, we pass the value to another place in our code. A function can be given a return type, the type of a value that will be returned by the function. At the call site, the return value can be stored within a variable of the same type as the function’s return.

```go
package main

import "fmt"
# create the func
func summonNicole() {
  fmt.Println("Hey Nicole, get over here!")
}

func main() {
// call the func in main for execution
  summonNicole()
}
```

```go
func getLengthOfCentralPark() int32 {
  var lengthInBlocks int32
  lengthInBlocks = 51
  return lengthInBlocks
}
```

```go
  func multiplier(x int32, y int32) int32 {
  return x \* y
}
```
## Functions can return multiple values

Functions also have to ability to return multiple values.
```go
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
if averageGrade > 90 {
  gradeLetter = "A"
} else if averageGrade > 80 {
  gradeLetter = "B"
} else if averageGrade > 70 {
  gradeLetter = "C"
} else if averageGrade > 60 {
  gradeLetter = "D"
} else {
  gradeLetter = "F"
}

return gradeLetter, averageGrade
}
```

## Deferring resolution

We can delay a function call to the end of the current scope by using the defer keyword. defer tells Go to run a function, but at the end of the current function. This is useful for logging, file writing, and other utilities.

```go
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!")
  taxRate := .06143
  fmt.Println("Calculating Taxes")

  if deductions == 0 || credits == 0 {
    return revenue \* taxRate
  }

  taxValue := (revenue - (deductions _ credits)) _ taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}
```
In the above example, when we call calculateTaxes() we immediately defer a message, "Taxes Calculated!". This does not print until the end of the function (after the taxes have been calculated and are about to be returned). Normally, we would consider adding fmt.Println("Taxes Calculated!") at the end of calculateTaxes(). But, we have multiple return statements in our code, instead of adding a print statement right before each return, we use defer and it prints regardless of when our function ends.
