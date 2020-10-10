// Let’s start by making a basic Go program. Before anything, we should declare our package information. In Go, this is done with the line package main. 
package main

//Next we want to print out our data to the terminal, so we’ll be using the fmt package. Import fmt into our Go program.
import "fmt"

//add a main function
func main(){
  //define 4 string variables
	var publisher,writer, artist, title string
	
  publisher = "DizzyBooks Publishing Inc."
  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  //define integer variables
  var year, pageNumber int
  
  year = 1997
  pageNumber = 14
  //define condition grade
  var grade float32 = 6.5
  
  //write out
  fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year) 
  fmt.Println("let us look at page ", pageNumber)
  fmt.Println("the rating of the book is ", grade)

}
