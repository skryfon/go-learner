package main

import (
	"strings"
	"unicode/utf8"
)

// Stats holds counts computed from a piece of text.
type Stats struct {
	Runes       int
	Words       int
	Lines       int
	LongestWord string
}

// Analyze computes Stats for text. Runes must be counted as Unicode code
// points, not bytes. Words are split on whitespace (spaces, tabs, newlines).
// Lines is the number of newline-separated lines (a non-empty text with no
// trailing newline still has at least 1 line).
func Analyze(text string) Stats {

	stats := Stats{}

	stats.Runes = utf8.RuneCountInString(text)
	stats.Words = len(strings.Fields(text))
	if len(text) == 0 {
		stats.Lines = 0
	} else {
		stats.Lines = strings.Count(text, "\n") + 1
	}

	longestWord := ""
	for _, word := range strings.Fields(text) {
		if utf8.RuneCountInString(word) > utf8.RuneCountInString(longestWord) {
			longestWord = word
		}
	}
	stats.LongestWord = longestWord

	return stats
}

// Merge folds other's counts into s (mutating s in place) and updates
// s.LongestWord if other.LongestWord is strictly longer (by rune count).
func (s *Stats) Merge(other Stats) {

	s.Runes += other.Runes
	s.Words += other.Words
	s.Lines += other.Lines

	if utf8.RuneCountInString(other.LongestWord) > utf8.RuneCountInString(s.LongestWord) {
		s.LongestWord = other.LongestWord
	}
}

// CombineTexts joins any number of text fragments with sep between them.
// Calling it with zero fragments returns "".
func CombineTexts(sep string, texts ...string) string {

	return strings.Join(texts, sep)
}

// WordFrequency counts occurrences of each whitespace-separated word,
// case-insensitively, and also returns the number of distinct words seen.
func WordFrequency(text string) (freq map[string]int, distinct int) {

	freq = make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		word = strings.ToLower(word)
		freq[word]++
	}

	distinct = len(freq)

	return freq, distinct
}

// NewWordCounter returns a function that, each time it's called with a word,
// returns how many times (including this call) that exact word has been
// passed to that specific returned function. Two counters created by two
// separate calls to NewWordCounter must not share state.
func NewWordCounter() func(word string) int {

	counts := make(map[string]int)

	return func(word string) int {
		counts[word]++
		return counts[word]
	}
}

// Reverse returns s with its runes in reverse order. Implement it
// recursively, and make sure multi-byte runes (accents, emoji, CJK
// characters) come back intact rather than corrupted.
func Reverse(s string) string {

	runes := []rune(s)

	var reverse func([]rune) string
	reverse = func(r []rune) string {
		if len(r) == 0 {
			return ""
		}
		return string(r[len(r)-1]) + reverse(r[:len(r)-1])
	}

	return reverse(runes)
}

// IsPalindrome reports whether s reads the same forwards and backwards,
// comparing runes (not bytes). Implement it recursively.
func IsPalindrome(s string) bool {

	runes := []rune(s)
	return checkPalindrome(runes)
}

func checkPalindrome(r []rune) bool {

	if len(r) <= 1 {
		return true
	}

	if r[0] != r[len(r)-1] {
		return false
	}

	return checkPalindrome(r[1 : len(r)-1])
}
