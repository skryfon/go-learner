package main

import "testing"

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name string
		text string
		want Stats
	}{
		{
			name: "ascii",
			text: "the quick brown fox",
			want: Stats{Runes: 19, Words: 4, Lines: 1, LongestWord: "quick"},
		},
		{
			name: "multibyte runes",
			// "héllo wörld" has fewer runes than bytes; a byte-based
			// implementation will overcount Runes here.
			text: "héllo wörld",
			want: Stats{Runes: 11, Words: 2, Lines: 1, LongestWord: "héllo"},
		},
		{
			name: "emoji",
			// each 😀 is 4 bytes but 1 rune.
			text: "😀😀 test",
			want: Stats{Runes: 7, Words: 2, Lines: 1, LongestWord: "test"},
		},
		{
			name: "multiple lines",
			text: "first line\nsecond line here\nthird",
			want: Stats{Runes: 33, Words: 6, Lines: 3, LongestWord: "second"},
		},
		{
			name: "empty",
			text: "",
			want: Stats{Runes: 0, Words: 0, Lines: 0, LongestWord: ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Analyze(tt.text)
			if got != tt.want {
				t.Errorf("Analyze(%q) = %+v, want %+v", tt.text, got, tt.want)
			}
		})
	}
}

func TestStatsMerge(t *testing.T) {
	s := Stats{Runes: 5, Words: 1, Lines: 1, LongestWord: "hello"}
	other := Stats{Runes: 3, Words: 1, Lines: 1, LongestWord: "wörld"}

	s.Merge(other)

	if s.Runes != 8 || s.Words != 2 || s.Lines != 2 {
		t.Errorf("Merge counts = %+v, want Runes=8 Words=2 Lines=2", s)
	}
	// "wörld" (5 runes) is longer than "hello" (5 runes) in byte length but
	// equal in rune length; use a clearly longer word to force the update.
	other2 := Stats{LongestWord: "extraordinarily"}
	s.Merge(other2)
	if s.LongestWord != "extraordinarily" {
		t.Errorf("LongestWord after merge = %q, want %q", s.LongestWord, "extraordinarily")
	}

	// original argument must not be mutated by the pointer receiver call.
	if other.Runes != 3 {
		t.Errorf("Merge must not mutate its argument, other.Runes = %d", other.Runes)
	}
}

func TestCombineTexts(t *testing.T) {
	tests := []struct {
		name  string
		sep   string
		texts []string
		want  string
	}{
		{name: "none", sep: ", ", texts: nil, want: ""},
		{name: "one", sep: ", ", texts: []string{"solo"}, want: "solo"},
		{name: "many", sep: " | ", texts: []string{"a", "b", "c"}, want: "a | b | c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineTexts(tt.sep, tt.texts...); got != tt.want {
				t.Errorf("CombineTexts(%q, %v) = %q, want %q", tt.sep, tt.texts, got, tt.want)
			}
		})
	}
}

func TestWordFrequency(t *testing.T) {
	freq, distinct := WordFrequency("the Cat sat on the mat the cat ran")
	want := map[string]int{
		"the": 3, "cat": 2, "sat": 1, "on": 1, "mat": 1, "ran": 1,
	}
	if len(freq) != len(want) {
		t.Fatalf("WordFrequency distinct words = %d, want %d (freq=%v)", len(freq), len(want), freq)
	}
	for word, count := range want {
		if freq[word] != count {
			t.Errorf("freq[%q] = %d, want %d", word, freq[word], count)
		}
	}
	if distinct != len(want) {
		t.Errorf("distinct = %d, want %d", distinct, len(want))
	}
}

func TestNewWordCounter(t *testing.T) {
	counter := NewWordCounter()

	seq := []struct {
		word string
		want int
	}{
		{"go", 1},
		{"rust", 1},
		{"go", 2},
		{"go", 3},
		{"rust", 2},
	}
	for _, s := range seq {
		if got := counter(s.word); got != s.want {
			t.Errorf("counter(%q) = %d, want %d", s.word, got, s.want)
		}
	}

	// A fresh counter must not share state with the first one.
	other := NewWordCounter()
	if got := other("go"); got != 1 {
		t.Errorf("independent counter(%q) = %d, want 1 (state leaked between closures)", "go", got)
	}
}

func TestReverse(t *testing.T) {
	tests := []struct{ in, want string }{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"héllo", "olléh"},
		{"😀😁", "😁😀"},
		{"日本語", "語本日"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.in); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"racecar", true},
		{"hello", false},
		{"héllo", false},
		{"😀a😀", true},
		{"日本語本日", true},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.in); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
