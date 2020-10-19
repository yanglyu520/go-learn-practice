# Conditionals

## Here are a list of the content below
- if, else if, and else statements.
- the switch statement.
- short declarations in conditionals.
- how to use conditionals in randomized conditions.

## if else if else statement
The parentheses are optional and can make it easier to read, but typically, we’ll see the if statement written without parentheses.

```go
if (alarmRinging) {
fmt.Println("Turn off the alarm!!")
}
```
if else statement
```go
isHungry := false
if isHungry {
fmt.Println("Eat the cookie")
} else {
fmt.Println("Step away from the cookie...")
}
```
The if...else if...else statement
``` go
position := 2

if position == 1 {
fmt.Println("You won the gold!")
} else if position == 2 {
fmt.Println("You got the silver medal.")
} else if position == 3 {
fmt.Println("Great job on bronze.")
} else {
fmt.Println("Sorry, better luck next time?")
}
```
## Switch Statement
```go
clothingChoice := "sweater"

switch clothingChoice {
case "shirt":
fmt.Println("We have shirts in S and M only.")
case "polos":
fmt.Println("We have polos in M, L, and XL.")
case "sweater":
fmt.Println("We have sweaters in S, M, L, and XL.")
case "jackets":
fmt.Println("We have jackets in all sizes.")
default:
fmt.Println("Sorry, we don't carry that")
}
// Prints: We have sweaters in S, M, L, and XL.
```
1. The switch keyword initiates the statement and is followed by a value. In the example, the value after switch is compared to the value after each case, until there is a match.
2. Inside the switch block, { ... }, there are multiple cases. The case keyword checks if the expression matches the specified value that comes after it. The value following the first case is "shirt". Since the value of clothingChoice is not the same as "shirt", that case‘s code does not run.
3. The value of clothingChoice is "sweater", so the third case‘s code runs and "We have sweaters in S, M, L, and XL." is printed. No more case statements are checked.
4. At the end of our switch statement is a default statement. If none of the cases are true, then the code in the default statement will run.

## Scoped Short Declaration Statement
We can also include a short variable declaration before we provide a condition in either if or switch statements. Here’s the syntax:

```go
x := 8
y := 9
if product := x \* y; product > 60 {
fmt.Println(product, " is greater than 60")
}
```

```go
switch season := "summer" ; season {
case "summer"
fmt.Println("Go out and enjoy the sun!")
}
```

## random number
So let’s introduce some uncertainty to our code by generating a random number. Go has a math/rand library that helps us generate a random integer
```go
import (
"math/rand"
"fmt"
)

func main() {
fmt.Println(rand.Intn(100))
}
```
Our random numbers weren’t entirely random. The reason for this behavior is due to how Go seeds or chooses a number as a starting point for generating random numbers. **By default, the seed value is 1.** We can provide a new seed value using the method rand.Seed() and passing in an integer.

However, we encounter the same problem if we pass in a constant, i.e. pass in 5. Each time we run our program, we’ll always print the same set of numbers. **Therefore, each time we run our program, we also need a unique number as a seed.** One way to get this unique number each time we run our program is to use time. The reason why we can use time to be our unique number is because it’s always a different time when our program is run! Take for example:


