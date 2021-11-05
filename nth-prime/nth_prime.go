package prime

import (
	"strings"
)

func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	primesTotal := 0
	factor := 2
	searchPool := n * factor // we need at least 2*n numbers to find the nth prime
	primes := []byte{}
	for primesTotal < n {
		primes = getPrimes(searchPool)
		primesTotal = countPrimes(primes[2:])
		factor++
		searchPool = n * factor
	}
	nthPrime := calculatePrime(primes, n)
	return nthPrime, true
}

/*getPrimes uses Sieve of Eratosthenes to generte primes.
For every prime p, only factors p >= p^2 are verified.
Primes are reprsented by a byte array, where index is the number and '1' or '0' is an indicator if that number is prime
*/
func getPrimes(s int) []byte {
	primes := []byte(strings.Repeat("1", s))
	for i := 2; i < s; i++ {
		if primes[i] == '1' {
			for j := i; j <= s; j++ {
				if i*j < s {
					primes[i*j] = '0'
				} else {
					break
				}
			}
		}
	}
	return primes
}

//countPrimes is a helper function that counts how many primes we have in total
func countPrimes(primes []byte) int {
	sum := 0
	for _, p := range primes {
		if p == '1' {
			sum++
		}
	}
	return sum
}

//calculatePrime calculates the nth prime number
func calculatePrime(primes []byte, n int) int {
	count := 0
	for i := 2; i < len(primes); i++ {
		count += int(primes[i] - '0')
		if count == n {
			return i
		}
	}
	return 0
}
