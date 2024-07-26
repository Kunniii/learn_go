package main

import (
  "fmt"
  "strconv"
)

func main() {
  var n string;
  fmt.Print("Enter a float: ")
  fmt.Scan(&n)

  
  // Convert a string into float32
  var f_number, _ = strconv.ParseFloat(n, 32)

  // Remove floating point by convert to int
  var i_number = int(f_number)

  fmt.Println(i_number)
}
