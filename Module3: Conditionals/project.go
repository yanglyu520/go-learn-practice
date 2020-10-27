package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  var isHeistOn = true
  var eludeGuards = rand.Intn(100)

  if eludeGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    fmt.Println("Plan a better disguise next time?")
  }

  var openedVault = rand.Intn(100)

  fmt.Println(isHeistOn)
  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn && openedVault < 70 {
    isHeistOn = false
    fmt.Println("Vault cannot be opened")
  } 

  if leftSafely := rand.Intn(5); isHeistOn {
    switch leftSafely {
      case 0: 
        isHeistOn = false
        fmt.Println("Heist failed")
      case 1: 
        isHeistOn = false 
        fmt.Print("Turns out vault doors don't open from the inside...")
      case 2: 
        isHeistOn = false 
        fmt.Print("Heist failed case 2")
      case 3: 
        fmt.Print("success")
      case 4: 
        fmt.Print("success")
      default:
        fmt.Print("Start the getaway car!")
    }
  }

  if isHeistOn {
    var amtStolen = 10000 + rand.Intn(1000000)
    fmt.Println("$", amtStolen, "not bad!")
  }
  
  }
