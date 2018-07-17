package hello_go

import (
	"github.com/u6f6o/hello-go/pkg/anagram"
)

func ExampleIsAnagram() {
	isAnagram := anagram.IsAnagram("Foo Bar", "Bar Oof")
	_ = isAnagram
}
