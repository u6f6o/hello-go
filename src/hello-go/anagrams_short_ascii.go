package hello_go

import "errors"

func AreAsciiAnagrams(one, another string) (bool, error) {
	if len(one) > 9 {
		return false, errors.New("checked string must be one of len < 10")
	}

	if len(one) < 1 || len(one) != len(another) {
		return false, nil
	}

	primes := [...]int64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31,
		37, 41, 43, 47, 53, 59, 61, 67, 71, 73,
		79, 83, 89, 97, 101}

	product := int64(1)

	for _, runeVal := range one {
		if runeIdx, err := RuneToIndex(runeVal); err == nil {
			product = product * primes[runeIdx]
		} else {
			return false, err
		}
	}

	for _, runeVal := range another {
		if runeIdx, err := RuneToIndex(runeVal); err == nil {
			primeVal := primes[runeIdx]

			if product% primeVal != 0 {
				return false, nil;
			}
			product = product / primeVal
		} else {
			return false, err
		}
	}
	return product == 1, nil
}

func RuneToIndex(runeVal rune) (int64, error) {
	switch {
		case 'a' <= runeVal && runeVal <= 'z':
			return int64(runeVal - 'a' + 1), nil
		case 'A' <= runeVal && runeVal <= 'Z':
			return int64(runeVal - 'A' + 1), nil
		default:
			return -1, errors.New("rune must be one of [a-zA-Z]")
	}
}