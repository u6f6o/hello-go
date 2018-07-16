package anagram

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var (
	tt = []struct {
		first     string
		second    string
		isAnagram bool
		testName  string
	}{
		{"", "", false, "empty"},
		{" ", "", false, "empty"},
		{" ", "   ", false, "empty"},
		{"", "A", false, "differentLengthWithEmpty"},
		{"Foo", "Fooo", false, "differentLength"},
		{"Foo", "Bar", false, "differentLength"},
		{"Foo", "Foo", true, "IsAngramSameWord"},
		{"Foo", "oFo", true, "IsAngramWord"},
		{"Foo bar", "Foo bar", true, "IsAngramSameSentence"},
		{"Foo bar", "rab ooF", true, "IsAngramSameSentence"},
		{"oFo rab", "oFo rab", true, "IsAngramSameSentence"},
	}
)

func TestIsAnagram(t *testing.T) {
	customTestTable := []struct {
		first     string
		second    string
		isAnagram bool
		testName  string
	}{
		{"NAHTRIHECCUNDE GAHINNEVERAHTUNIN ZEHGESSURKLACH ZUNNUS", "CARL GUSTAV IUNG IN KUESNACH IAHR NEUNZEHNHUNDERTSECHZEHN", true, "IsAngramWikipediaSentence"},
		{"SMAISMRMILMEPOETALEVMIBVNENVGTTAVIRAS", "Altissimvm planetam tergeminvm observavi", true, "IsAngramWikipediaSentence"},
		{"Annulo cingitur tenui plano nusquam cohaerente ad eclipticam inclinato", "AAAAAAA CCCCC D EEEEE G H IIIIIII LLLL MM NNNNNNNNN OOOO PP Q RR S TTTTT UUUUU", true, "IsAngramWikipediaSentence"},
	}

	for _, tc := range append(tt, customTestTable...) {
		t.Run(tc.testName, func(t *testing.T) {
			assert.Equal(t, IsAnagram(tc.first, tc.second), tc.isAnagram)
		})
	}
}

func TestAreAsciiAnagrams(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			isAnagram, _ := IsAsciiAnagrams(tc.first, tc.second)
			assert.Equal(t, isAnagram, tc.isAnagram)
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	for _, bc := range tt {
		var isAnagram bool
		b.Run(bc.testName, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				isAnagram = IsAnagram(bc.first, bc.second)
			}
			_ = isAnagram
		})
	}
}

func BenchmarkIsAsciiAnagrams(b *testing.B) {
	for _, bc := range tt {
		var isAnagram bool
		b.Run(bc.testName, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				isAnagram, _ = IsAsciiAnagrams(bc.first, bc.second)
			}
			_ = isAnagram
		})
	}
}
