package main

import "fmt"

func main() {
	text := "the quick brown fox jumps over the lazy dog"

	stats := Analyze(text)
	fmt.Printf("stats: %+v\n", stats)

	freq, distinct := WordFrequency(text)
	fmt.Println("word frequency:", freq, "distinct:", distinct)

	counter := NewWordCounter()
	fmt.Println("count 'the':", counter("the"))
	fmt.Println("count 'the':", counter("the"))

	combined := CombineTexts(" / ", "first", "second", "third")
	fmt.Println("combined:", combined)

	fmt.Println("reversed:", Reverse(text))
	fmt.Println("is 'racecar' a palindrome?", IsPalindrome("racecar"))
}
