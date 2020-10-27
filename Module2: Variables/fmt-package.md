# More about fmt package

## 1. fmt.Println() vs fmt.Print()
fmt.Println() allows us to print to the terminal and view the data that we’re working with. It has some defaulting styling built-in that makes viewing data easier for us. fmt.Println() prints its arguments (data provided within its parentheses ( )) with **an included space in between each argument and adds a line break at the end.**

However, there are times we might not want the default formatting. In such cases, using fmt.Print() would be more appropriate:
```
fmt.Print("The answer is", ": ")
fmt.Print("12")
```

**Notice that there’s no default spacing added when fmt.Print() has multiple arguments.** Also, since fmt.Print() doesn’t add a line break after printing, the argument for the second print statement print on the same line as the first print statement’s.

## 2. fmt.Printf
Using fmt.Println() and fmt.Print() we have the ability to concatenate strings, i.e. combine different strings into a single string.

```
guess := "C"
fmt.Println("Is", guess, "your final answer?")
// Prints: Is C your final answer?
```
**With fmt.Printf(), we can interpolate strings, or leave placeholders in a string and use values to fill in the placeholders.**
```
guess := "C"
fmt.Printf("Is %v your final answer?", guess)
// Prints: Is C your final answer?
```

The %v portion is our placeholder and is known as a verb in Go. Verbs are identified by the combination of a % character followed by a letter. The specific letter informs what goes fills in the placeholder, in this case, %v gets the value of "C" from our second argument, guess.

To check out more fmt print verbs, check [here](https://golang.org/pkg/fmt/#hdr-Printing)
Ex:
-  %T prints out the type of the  argument
- %d is for integers
- %f is for floats
- %v is for string

```
specialNum := 42
fmt.Printf("This value's type is %T.", specialNum)
// Prints: This value's type is int.

quote := "To do or not to do"
fmt.Printf("This value's type is %T.", quote)
// Prints: This value's type is string.

votingAge := 18
fmt.Printf("You must be %d years old to vote.", votingAge)
// Prints: You must be 18 years old to vote.


gpa := 3.8
fmt.Printf("You're averaging: %f.", gpa)
// Prints: You're averaging 3.800000.

```
## 3. Sprint and Sprintln
fmt is the formatter package and we have other methods that don't print strings but format them instead 
fmt.Sprintln() works like fmt.Sprint() but it automatically includes spaces between the arguments for us (just like fmt.Println() vs fmt.Print()!):

```
grade := "100"
compliment := "Great job!"
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

fmt.Print(teacherSays)
// Prints: You scored a 100 on the test! Great job!

```

## 4. Sprintf

fmt.Sprintf() works very similarly to fmt.Printf(), the major difference is that fmt.Sprintf() returns its value instead of printing it out!

## 5. Scan()
Another helpful method from the fmt package is .Scan() which allows us to get user input! If we were expecting two values, we’d need to declare two variables ahead of time, fmt.Scan() expects addresses for arguments, hence the & before response1 and response2


```
fmt.Println("How are you doing?")

var response string
fmt.Scan(&response)

fmt.Printf("I'm %v.", response)

//two values
fmt.Println("How are you doing?")

var response1 string
var response2 string
fmt.Scan(&response1)
fmt.Scan(&response2)

fmt.Printf("I'm %v %v", response1, response2)

```