# 01 - Grade Report

A small grade-book calculator. You're given a class's test scores and need
to compute letter grades, averages, curves, and distributions.

## Topics covered

- Values, variables, constants
- `for` loops
- `if`/`else`
- `switch`
- Arrays (fixed-size, value semantics)
- Slices (dynamic, shared backing arrays)
- Maps
- Functions (including multiple return values and errors)

## Setup

```bash
go run .          # runs the demo in main.go (will panic until you implement grades.go)
go test ./...     # run the test suite — should start red
go vet ./...
gofmt -l .        # should print nothing
```

## Goals

Implement every function in `grades.go` (replace the `panic("TODO: implement")`
bodies) until `go test ./...` passes. Along the way you should be able to
explain:

- [ ] Why `LetterGrade` needs a `switch` (or `if`/`else` chain) instead of a
      single formula, and how you handle the invalid-score case.
- [ ] Why `Average` has to guard against an empty slice.
- [ ] Why `DropLowest`, `CurveScores`, and `TopScorers` must **not** mutate
      the caller's original slice — what happens if you sort or reslice
      the input in place instead of building a new result slice.
- [ ] Why `GradeDistribution` needs `make(map[string]int)` (or a map
      literal) instead of a `var m map[string]int` zero value, if you intend
      to write into it.
- [ ] Why `ClearRoster` can safely take a `[5]string` by value and return a
      cleared copy without ever touching the caller's array — what's
      different about arrays vs. slices here.

Ground rules from the repo root apply: write it yourself before comparing,
run `gofmt`/`go vet` before calling it done, and don't reach for AI
autocomplete while solving.
