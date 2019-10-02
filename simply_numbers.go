package features

//This function returns true if the number is prime, and false - otherwise.
func IsPrimeNum(number int)(result bool) {
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

//This function prints to the console primes from 1 to the specified value of the variable 'limit'.
func PrimeNumsFinder(limit int) {
  for i := 1; i <= limit; i++ {
    if IsSimpleNum(i) {
      fmt.Println(i)
    }
  }
}
