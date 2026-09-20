// https://leetcode.com/problems/number-of-digit-one/

package main

func countDigitOne(n int) int {
	return countDigitOne2(n)
}

// 13 => 13, 12, 1-1, 10, 1 (6)
// 220 => 219-210(11), 209-200 (1 each), 190-120 (11), 109-100 (11), 119-110 (21), 90-20 (1 each), 19-10 (11), 9-1 (1 each)
func countDigitOne1(n int) int {
	// 1 < n < 10 => 1
	// 10 <= n <= 99 => 19
	//   10 <= n <= 19: 11
	//   20 <= n <= 99: 8
	// 100 <= n <= 999: 280
	//   100 <= n <= 199:  100 + (0 < n < 99) => 120
	//   200 <= n <= 999:  20 * 8 => 160
	// 1,000 <= n < 9,999
	//   1,000 <= n < 1,999: 1000 + (280) (0 < n < 999)

	result := 0

	// Each position
	// 1 digit: 1 is appeared (n+1)/10 + n%10 >= 1 == 1)
	// 10-th digit: 1 is appeared (n+1)/100 + (0 < n%100 < 20 => module n%10, 20 < n%100 => 10)
	// 100-th digit: 1 is appeared (n+1)/1000 +

	// count 1 on each digit
	for divider := 1; divider <= n; divider *= 10 {
		// count 1 appeared on the higher digits
		// Let's say 1,234 and count the 2nd digit "3", the 1 appeared 12 times (from 100 - 1200, 130 times)
		result += n / (divider * 10) * divider

		if (n/divider)%10 == 1 {
			// For 1XX like 134, 1 appeared 34 + 1 times, so add them
			result += n%divider + 1
		} else if (n/divider)%10 >= 2 {
			// For 2XX to 9XX, the 1 appeared in 1XX are already shown, so added the count
			result += divider
		}
	}
	return result
}

func countDigitOne2(n int) int {
	// digitOnes (n): the number of digit ones < 10^n
	// oneX(n): The number of digit ones where the first digit is 1 and the value is < 10^n
	// nonOneX(n): The number of digit ones where the first digit is not 1 and the value is < 10^n
	// oneX(0): 1
	// nonOneX(0): 0
	// digitOnes(0): oneX(0) + nonOneX(0)

	// digitOnes(1): digitOnes(0) + oneX(1) + nonOneX(1)
	// oneX(1): 10^1 + digitOes(0) = 10^1 + 1n
	// nonOneX(1): 8 * digitOnes(0) = 8
	// digitOnes(2): digitOnes(1) + oneX(2) + nonX(2)
	// oneX(2): 10^2 + digitOnes(1) = 10^2 + 11
	// nonX(2): 8 * digitOnes(1) = 8 * 8

	// Examples:
	// Example 1: 1875
	// 1000 - 1875: 875 + digitOnes(2)
	// 0-999: digitOnes(2)

	// Example 2: 2345
	// 1000-2000: 1000 + digitOnes(2)
	// -999: digitOnes(2)
	// 2000-2300: oneX(2) + nonOneX(2) * 2
	// 2300-2340: oneX(1) + nonOneX(1)* 3
	// 2340-2345: oneX(0) + nonOneX(0) * 5

	// 134:
	//   (134-100+1) + f(99) + f(34)
	//   (n+1)%100 +  f(99) + f(9) * (n/10)%10 + f(4)
	//   (n+1)%100 + f(99) + f(9) * 3 + f(4)

	// 888
	//   100 + f(99) (for 100-199) + (888-200)/100*f(99) + (888-800)/10*f(9)
	//   (n-div)

	digit := 1
	digitOnes := []int{1}
	oneX := []int{1}
	nonOneX := []int{0}

	divisor := 1
	i := 0
	for ; divisor <= n; i++ {
		digit *= 10
		oneX = append(oneX, digit+digitOnes[i])
		nonOneX = append(nonOneX, 8*digitOnes[i])
		digitOnes = append(digitOnes, digitOnes[i]+oneX[i+1]+nonOneX[i+1])

		divisor *= 10
	}

	result := 0
	for divisor /= 10; i >= 0 && divisor > 0; divisor /= 10 {
		i--
		// Add 1 appeared in the current digit
		// For example, for 234, in XX, 1XX, the digitOnes(i-1) shows twice
		if i > 0 {
			result += (n / divisor) * digitOnes[i-1]
		}

		mod := (n / divisor) % 10
		if mod == 1 {
			// The current digit is 1 like 134, add the count when 100-134 => 35 times
			result += n%divisor + 1
		} else if mod > 1 {
			result += divisor
		}

		n = n % divisor
		// fmt.Printf("i: %d, n: %d, divisor: %d, result: %d\n", i, n, divisor, result)
	}

	return result
}
