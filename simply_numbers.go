package features

import ( 
    "fmt"
)

//This function returns true if the number is simple, and false - otherwise.
func IsSimpleNum(number int)(result bool) {
  counter := 0
  for i := 1; i <= number; i++ {
    if number % i == 0 {
      counter++
    }
  }

  if counter == 2 {
    return true
  }

  return false
}
