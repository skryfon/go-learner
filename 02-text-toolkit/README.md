# 02 - Text Toolkit

A small text-analysis toolkit: word/rune/line stats, word frequency, a
stateful word counter, and recursive string reversal/palindrome checks.

## Topics covered

- Functions
- Multiple return values
- Variadic functions
- Closures
- Recursion
- Range over built-in types (strings, slices, maps)
- Pointers
- Strings and runes (Unicode-correct handling, not just bytes)
- Structs

## Setup

```bash
go run .          # runs the demo in main.go (will panic until you implement textkit.go)
go test ./...     # run the test suite — should start red
go vet ./...
gofmt -l .        # should print nothing
```

## Goals

Implement every function in `textkit.go` (replace the `panic("TODO: implement")`
bodies) until `go test ./...` passes. Along the way you should be able to
explain:

- [ ] Why `Analyze` must range over runes, not bytes, to get `Runes` right —
      what breaks on "héllo" or "😀😀" if you use `len(text)` instead.
- [ ] Why `Merge` takes a pointer receiver — what happens to the caller's
      `Stats` if you use a value receiver instead.
- [ ] Why `CombineTexts` works fine when called with zero extra arguments —
      what a variadic parameter actually is inside the function body.
- [ ] Why `WordFrequency` returns two values instead of just the map — what
      information the caller would have to recompute otherwise.
- [ ] Why `NewWordCounter` returns a *function*, and why two separate calls
      to it produce counters with independent state (this is the closure
      capturing its own variable, not a shared global).
- [ ] Why `Reverse` and `IsPalindrome` need to recurse over runes, not
      bytes or a plain `for i := range s` byte index — what a naive
      byte-slice reversal does to a multi-byte character.

Ground rules from the repo root apply: write it yourself before comparing,
run `gofmt`/`go vet` before calling it done, and don't reach for AI
autocomplete while solving.
