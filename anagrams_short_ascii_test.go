package hello_go

import "testing"


func TestAreAsciiAnagrams(t *testing.T) {
	if result, _ := AreAsciiAnagrams("", ""); result {
		t.Error("Found anagram for two empty strings")
	}
	if result, _ := AreAsciiAnagrams("foo", "fooo"); result {
		t.Error("Found anagram for foo and fooo")
	}
	if result, _ := AreAsciiAnagrams("fOo", "foo"); result {
		t.Error("Found anagram for fOo and foo")
	}
	if result, _ := AreAsciiAnagrams("foo", "foo"); !result {
		t.Error("Did not find anagram for foo and foo")
	}
	if result, _ := AreAsciiAnagrams("foo", "ofo"); !result  {
		t.Error("Did not find anagram for foo and ofo")
	}
	if result, _ := AreAsciiAnagrams("foO", "fOo"); !result {
		t.Error("Did not find anagram for foO and fOo")
	}
	if _, err := AreAsciiAnagrams("123456789", "fOo"); err == nil {
		t.Error("Accepted string with len > 8")
	}
	if result, err := AreAsciiAnagrams("fÖo", "foo"); err == nil {
		println(result)
		t.Error("Accepted non-ascii string")
	}
}

func TestRuneToIndex(t *testing.T) {
	if result, _ := RuneToIndex('a'); result != 0 {
		t.Errorf("a should be 0 but was %d", result)
	}
	if result, _ := RuneToIndex('m'); result != 12 {
		t.Errorf("a should be 12 but was %d", result)
	}
	if result, _ := RuneToIndex('z'); result != 25 {
		t.Errorf("z should be 25 but was %d", result)
	}
	if result, _ := RuneToIndex('A'); result != 26 {
		t.Errorf("A should be 26 but was %d", result)
	}
	if result, _ := RuneToIndex('N'); result != 39 {
		t.Errorf("A should be 39 but was %d", result)
	}
	if result, _ := RuneToIndex('Z'); result != 51 {
		t.Errorf("Z should be 51 but was %d", result)
	}
	if _, err := RuneToIndex('Ö'); err == nil {
		t.Error("Accepted non-ascii string")
	}
}
