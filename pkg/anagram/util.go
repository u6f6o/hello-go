package anagram

import (
	"errors"
	"strings"
)

func IsAnagram(first, second string) (isAnagram bool) {
	first = strings.Replace(first, " ", "", -1)
	second = strings.Replace(second, " ", "", -1)

	if len(first) < 1 || len(first) != len(second) {
		return
	}

	runeCount := make(map[rune]int)
	for _, runeVal := range strings.ToLower(first) {
		runeCount[runeVal] += 1
	}

	for _, runeVal := range strings.ToLower(second) {
		if _, present := runeCount[runeVal]; present {
			if runeCount[runeVal] -= 1; runeCount[runeVal] < 0 {
				return
			}
		} else {
			return
		}
	}

	return true
}

var primes = []int{
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31,
	37, 41, 43, 47, 53, 59, 61, 67, 71, 73,
	79, 83, 89, 97, 101}

func IsAsciiAnagrams(first, second string) (bool, error) {
	first = strings.Replace(first, " ", "", -1)
	second = strings.Replace(second, " ", "", -1)

	switch {
	case len(first) > 9:
		return false, errors.New("")
	case len(first) < 1 || len(first) != len(second):
		return false, nil
	default:
		product := 1

		for _, runeVal := range first {
			if runeIdx, err := runeToIndex(runeVal); err == nil {
				product = product * primes[runeIdx]
			} else {
				return false, err
			}
		}

		for _, runeVal := range second {
			if runeIdx, err := runeToIndex(runeVal); err == nil {
				primeVal := primes[runeIdx]

				if product%primeVal != 0 {
					return false, nil
				}
				product = product / primeVal
			} else {
				return false, err
			}
		}
		return product == 1, nil

	}
}

func runeToIndex(runeVal rune) (int, error) {
	switch {
	case 'a' <= runeVal && runeVal <= 'z':
		return int(runeVal - 'a' + 1), nil
	case 'A' <= runeVal && runeVal <= 'Z':
		return int(runeVal - 'A' + 1), nil
	default:
		return -1, errors.New("rune must be one of [a-zA-Z]")
	}
}
