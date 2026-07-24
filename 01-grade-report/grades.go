package main

import "errors"

// Grade boundaries. A score >= AGrade is an "A", >= BGrade but < AGrade is
// a "B", and so on. Anything below DGrade is an "F".
const (
	AGrade = 90
	BGrade = 80
	CGrade = 70
	DGrade = 60
)

// ErrInvalidScore is returned when a score is outside the valid 0-100 range.
var ErrInvalidScore = errors.New("score must be between 0 and 100")

// LetterGrade converts a numeric score (0-100) into a letter grade using the
// AGrade/BGrade/CGrade/DGrade boundaries. It returns ErrInvalidScore if score
// is outside the 0-100 range.
func LetterGrade(score int) (string, error) {
	panic("TODO: implement")
}

// Average returns the arithmetic mean of scores. It returns 0 for an empty
// slice (no division by zero).
func Average(scores []int) float64 {
	panic("TODO: implement")
}

// DropLowest returns a new slice containing scores with its n lowest values
// removed. It must not modify the backing array of the scores argument. If n
// is greater than or equal to len(scores), it returns an empty slice.
func DropLowest(scores []int, n int) []int {
	panic("TODO: implement")
}

// CurveScores returns a new slice where every score has been increased by
// points, capped at 100. It must not modify the backing array of the scores
// argument.
func CurveScores(scores []int, points int) []int {
	panic("TODO: implement")
}

// GradeDistribution returns a map from letter grade to how many scores fall
// into that grade. The returned map is always usable (never nil), even for
// an empty scores slice.
func GradeDistribution(scores []int) map[string]int {
	panic("TODO: implement")
}

// TopScorers returns a new slice containing only the scores >= threshold, in
// their original relative order. It must not modify the backing array of the
// scores argument.
func TopScorers(scores []int, threshold int) []int {
	panic("TODO: implement")
}

// ClearRoster returns a copy of roster with every name replaced by "". Since
// Go arrays are values, the caller's original roster must be unaffected by
// this call regardless of how ClearRoster is implemented.
func ClearRoster(roster [5]string) [5]string {
	panic("TODO: implement")
}
