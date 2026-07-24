package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestLetterGrade(t *testing.T) {
	cases := []struct {
		score   int
		want    string
		wantErr bool
	}{
		{100, "A", false},
		{90, "A", false},
		{89, "B", false},
		{80, "B", false},
		{79, "C", false},
		{70, "C", false},
		{69, "D", false},
		{60, "D", false},
		{59, "F", false},
		{0, "F", false},
		{-1, "", true},
		{101, "", true},
	}

	for _, c := range cases {
		got, err := LetterGrade(c.score)
		if c.wantErr {
			if !errors.Is(err, ErrInvalidScore) {
				t.Errorf("LetterGrade(%d) error = %v, want ErrInvalidScore", c.score, err)
			}
			continue
		}
		if err != nil {
			t.Errorf("LetterGrade(%d) unexpected error: %v", c.score, err)
		}
		if got != c.want {
			t.Errorf("LetterGrade(%d) = %q, want %q", c.score, got, c.want)
		}
	}
}

func TestAverage(t *testing.T) {
	if got := Average(nil); got != 0 {
		t.Errorf("Average(nil) = %v, want 0", got)
	}
	if got := Average([]int{}); got != 0 {
		t.Errorf("Average([]) = %v, want 0", got)
	}
	if got := Average([]int{80}); got != 80 {
		t.Errorf("Average([80]) = %v, want 80", got)
	}
	if got := Average([]int{70, 80, 90}); got != 80 {
		t.Errorf("Average([70,80,90]) = %v, want 80", got)
	}
}

func TestDropLowest(t *testing.T) {
	original := []int{90, 60, 100, 70, 80}
	snapshot := append([]int(nil), original...)

	got := DropLowest(original, 2)

	if !reflect.DeepEqual(original, snapshot) {
		t.Fatalf("DropLowest mutated its input: got %v, want unchanged %v", original, snapshot)
	}

	if len(got) != len(original)-2 {
		t.Fatalf("DropLowest result has len %d, want %d", len(got), len(original)-2)
	}

	sum := 0
	for _, s := range got {
		sum += s
	}
	// dropping the two lowest (60, 70) leaves 90+100+80 = 270
	if sum != 270 {
		t.Errorf("DropLowest result sums to %d, want 270 (got %v)", sum, got)
	}

	if got := DropLowest(original, len(original)+5); len(got) != 0 {
		t.Errorf("DropLowest with n >= len(scores) = %v, want empty", got)
	}
}

func TestCurveScores(t *testing.T) {
	original := []int{95, 80, 60}
	snapshot := append([]int(nil), original...)

	got := CurveScores(original, 10)

	if !reflect.DeepEqual(original, snapshot) {
		t.Fatalf("CurveScores mutated its input: got %v, want unchanged %v", original, snapshot)
	}

	want := []int{100, 90, 70} // 95 capped at 100, others +10
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CurveScores(%v, 10) = %v, want %v", snapshot, got, want)
	}
}

func TestGradeDistribution(t *testing.T) {
	dist := GradeDistribution(nil)
	if dist == nil {
		t.Fatal("GradeDistribution(nil) returned a nil map, must be usable")
	}
	if len(dist) != 0 {
		t.Errorf("GradeDistribution(nil) = %v, want empty", dist)
	}
	// reading a missing key from a non-nil empty map must not panic
	if v := dist["A"]; v != 0 {
		t.Errorf("dist[\"A\"] = %d, want 0", v)
	}

	got := GradeDistribution([]int{95, 91, 82, 71, 55})
	want := map[string]int{"A": 2, "B": 1, "C": 1, "F": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GradeDistribution(...) = %v, want %v", got, want)
	}
}

func TestTopScorers(t *testing.T) {
	original := []int{95, 60, 91, 75, 100}
	snapshot := append([]int(nil), original...)

	got := TopScorers(original, AGrade)

	if !reflect.DeepEqual(original, snapshot) {
		t.Fatalf("TopScorers mutated its input: got %v, want unchanged %v", original, snapshot)
	}

	want := []int{95, 91, 100}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopScorers(%v, %d) = %v, want %v", snapshot, AGrade, got, want)
	}
}

func TestClearRoster(t *testing.T) {
	original := [5]string{"Ada", "Grace", "Alan", "Barbara", "Linus"}
	snapshot := original

	got := ClearRoster(original)

	if original != snapshot {
		t.Fatalf("ClearRoster mutated the caller's array: got %v, want unchanged %v", original, snapshot)
	}

	want := [5]string{}
	if got != want {
		t.Errorf("ClearRoster(...) = %v, want all-empty %v", got, want)
	}
}
