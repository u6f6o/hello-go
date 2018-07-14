package hello_go

import "errors"

func AreAsciiAnagrams(one, another string) (bool, error) {
	runeCntOne := len([]rune(one))
	runeCntAnother := len([]rune(another))

	if runeCntOne > 8 {
		return false, errors.New("checked string must be one of len < 9")
	}

	if runeCntOne < 1 || runeCntOne != runeCntAnother {
		return false, nil
	}

	primes := [...]uint64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
		101, 103, 107, 109, 113, 127, 131, 137, 139,
		149, 151, 157, 163, 167, 173, 179, 181, 191,
		193, 197, 199, 211, 223, 227, 229, 233, 239}

	product := uint64(1)

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

			if product % primeVal != 0 {
				return false, nil
			}
			product = product / primeVal
		} else {
			return false, err
		}
	}
	return product == 1, nil
}

func RuneToIndex(runeVal rune) (int, error) {
	switch {
		case 'a' <= runeVal && runeVal <= 'z':
			return int(runeVal - 'a'), nil
		case 'A' <= runeVal && runeVal <= 'Z':
			return int(runeVal - 'A' + 26), nil
		default:
			return -1, errors.New("rune must be one of [a-zA-Z]")
	}
}