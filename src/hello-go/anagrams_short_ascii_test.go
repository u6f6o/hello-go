package hello_go

import "testing"

func TestAreAsciiAnagrams(t *testing.T) {
	if result, err := AreAsciiAnagrams("", ""); err == nil && result {
		t.Error("Found anagram for two empty strings")
	}
	if result, err := AreAsciiAnagrams("foo", "fooo"); err == nil && result {
		t.Error("Found anagram for foo and fooo")
	}
	if result, err := AreAsciiAnagrams("foo", "foo"); err == nil && !result {
		t.Error("Did not find anagram for foo and foo")
	}
	if result, err := AreAsciiAnagrams("foo", "ofo"); err == nil && !result  {
		t.Error("Did not find anagram for foo and ofo")
	}
}

func TestRuneToIndex(t *testing.T) {
	if result, err := RuneToIndex('a'); err == nil && result == 0 {
		t.Errorf("a should be 0 but was %d", result)
	}
	if result, err := RuneToIndex('z'); err == nil && result == 25 {
		t.Errorf("z should be 0 but was %d", result)
	}
	if result, err := RuneToIndex('A'); err == nil && result == 0 {
		t.Errorf("A should be 0 but was %d", result)
	}
	if result, err := RuneToIndex('Z'); err == nil && result == 25 {
		t.Errorf("Z should be 0 but was %d", result)
	}
}
