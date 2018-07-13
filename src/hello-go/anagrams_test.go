package hello_go

import "testing"

func TestAreAnagrams(t *testing.T) {
	if AreAnagrams("", "") {
		t.Error("Found anagram for two empty strings")
	}
	if AreAnagrams("foo", "fooo") {
		t.Error("Found anagram for foo and fooo")
	}
	if !AreAnagrams("foo", "foo") {
		t.Error("Did not find anagram for foo and foo")
	}
	if !AreAnagrams("foo", "ofo") {
		t.Error("Did not find anagram for foo and ofo")
	}
}
