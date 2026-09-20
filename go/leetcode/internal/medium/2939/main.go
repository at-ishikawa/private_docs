// https://leetcode.com/problems/maximum-xor-product/description/
// https://leetcode.com/problems/maximum-xor-product/solutions/4304306/python-explained/
package main

// a=12 => 11000 b=5 => 101, x=10
// a=6 (110), b=7(111), 2^5=32, 25(11001)
func maximumXorProduct(a int64, b int64, n int) int {
	for i := n - 1; i >= 0; i-- {
		mask := int64(1 << i)
		// 1 << 3 => b1000
		isABitOne := a&mask == mask
		isBBitOne := b&mask == mask

		if isABitOne && isBBitOne {
			continue
		}
		if !isABitOne && !isBBitOne {
			// x = x | mask
			a ^= mask
			b ^= mask
			continue
		}
		if isABitOne && a > b {
			// x = x | mask
			a ^= mask
			b ^= mask
			continue
		}
		if isBBitOne && b > a {
			// x = x | mask
			a ^= mask
			b ^= mask
			continue
		}
	}

	mod := int64(1e9 + 7)
	a %= mod
	b %= mod
	return int((a * b) % mod)
}
