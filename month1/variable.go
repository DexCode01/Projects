package main

import "fmt"

func main() {

  // variable declarations and initialization 

  name := "Emmanuel"
  var score int = 3569
  var country string = " Nigeria"

  // go also infers from value type if not stated in the declaration

  var age = 60
  var sport = "golf"

  // multiple variable can also be stated on the same line of code 

  var x, y, z = "friday", "may", 8

  fmt.Println(name, score, country)
  fmt.Println(age, sport)
  fmt.Println(x, y, z)
  

}