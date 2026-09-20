// https://leetcode.com/problems/strong-password-checker/description/
package main

import "fmt"

// It has at least 6 characters and at most 20 characters.
// It contains at least one lowercase letter, at least one uppercase letter, and at least one digit.
// It does not contain three repeating characters in a row (i.e., "Baaabb0" is weak, but "Baaba0" is strong).

// The first condition: less than, 6 required to add. More than 20, required to delete characters
// The 2nd condition: Required to replace or add letters at most 3 as long as they miss one of them
// The 3rd condition: Add/Replace/Delete letters around 3 consecutive letters
//
//	Add: aaa => aBaa0c
//	Replace: aa0bBb
//	Delete: More than 20 letters

// What if a password length is 50?
// The worst case: aaaaaa..aaaa
//   - 30 letters should be removed.
//   - aabaaBaa0aacaaC: 6 replacements
//   - 36 changes

// Scan a string to
//  1. Find the length of substrings that are consecutives.
//  2. If a lowercase/uppercase/digit is included
//
// Break letters of consecutive letters more than 3
//  1. if len(password) > 20, then delete letters
//  2. If len(password) < 6, then add letters
//  3. otherwise, replace letters
func strongPasswordChecker(password string) int {
	isLowercaseIncluded := false
	isUppercaseIncluded := false
	isDigitIncluded := false
	repeatingSubstrings := []string{}

	repeatingSubstr := ""
	for i := range password {
		ch := password[i]
		// todo: simpler
		if 'a' <= ch && ch <= 'z' {
			isLowercaseIncluded = true
		}
		if 'A' <= ch && ch <= 'Z' {
			isUppercaseIncluded = true
		}
		if '0' <= ch && ch <= '9' {
			isDigitIncluded = true
		}

		if len(repeatingSubstr) > 0 && repeatingSubstr[0] == ch {
			repeatingSubstr = repeatingSubstr + string(ch)
			continue
		}

		// Store the repeating substrings
		if len(repeatingSubstr) >= 3 {
			repeatingSubstrings = append(repeatingSubstrings, repeatingSubstr)
		}
		repeatingSubstr = string(ch)
	}
	if len(repeatingSubstr) >= 3 {
		repeatingSubstrings = append(repeatingSubstrings, repeatingSubstr)
	}

	passwordLength := len(password)
	steps := 0
	changedCount := 0
	for _, substr := range repeatingSubstrings {
		substrLen := len(substr)
		if passwordLength > 20 {
			// aaaaa => Remove 3 letters (5-2)
			// a..aa (21) => Remove a letter, and others should change
			removedLength := substrLen - 2
			if passwordLength-removedLength < 20 {
				// we shouldn't remove more than necessary
				removedLength = 20 - passwordLength
			}
			steps += removedLength
			passwordLength -= removedLength
			substrLen -= removedLength
			if substrLen < 3 {
				continue
			}
		}

		// lengths to change:
		// 3-5 => 1, 6-8 => 2, 9-11 => 3
		step := substrLen / 3
		steps += step
		if passwordLength < 6 {
			// substring cannot be more than 6 length, so adding 1 characters is sufficient
			passwordLength += step
		}
		changedCount += step
	}

	fmt.Printf("%d, %d\n", steps, passwordLength)
	if passwordLength < 6 {
		changedCount += 6 - passwordLength
		steps += changedCount
	}
	if passwordLength > 20 {
		steps += passwordLength - 20
	}

	for _, included := range []bool{
		isLowercaseIncluded,
		isUppercaseIncluded,
		isDigitIncluded,
	} {
		if included {
			continue
		}
		if changedCount > 0 {
			// it's added already. Just make the flag on
			changedCount--
			continue
		}

		// add/replace a letter to contain a lowercase or uppercase
		steps++
	}
	return steps
}
