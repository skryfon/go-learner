package main

import "fmt"

func main() {
	scores := []int{95, 82, 71, 58, 90, 67, 100}

	avg := Average(scores)
	fmt.Printf("class average: %.2f\n", avg)

	dist := GradeDistribution(scores)
	fmt.Println("grade distribution:", dist)

	top := TopScorers(scores, AGrade)
	fmt.Println("top scorers (A grade):", top)

	dropped := DropLowest(scores, 2)
	fmt.Println("after dropping lowest 2:", dropped)

	curved := CurveScores(scores, 5)
	fmt.Println("curved scores:", curved)
}
