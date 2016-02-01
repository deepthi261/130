/* obtain the smallest positive number that is evenly divisible by all of the numbers from 1 to 20
 solution -> find the prime factorization of every number from 1 to 20 and keep a count of how many of each prime factor is necessary.
 */

package main

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	result := 1
	var div int
	for i := range primes {
		div = primes[i]
		for div <= 20 {
			div *= primes[i]
		}
		div /= primes[i]
		result *= div
	}
	println(result)
}
