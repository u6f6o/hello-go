package anagram

func ExampleIsAnagram() {
	first, second := "Foo Bar", "Rab Oof"

	isAnagram := IsAnagram(first, second)
	_ = isAnagram
}

func ExampleIsAsciiAnagrams() {
	first, second := "Foo Bar", "Rab Oof"

	isAnagram, err := IsAsciiAnagrams(first, second)
	if err != nil {
		// error handling
	}
	_ = isAnagram
}
