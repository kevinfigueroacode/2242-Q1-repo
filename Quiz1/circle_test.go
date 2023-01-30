package main

import (
	"testing"
)
//main testing for varios methods
//like a breaker
func TestArea(t *testing.T) {

//test for the area method of a circle
t.Run("circles", func(t *testing.T) {
	circle := Circle{5.0}
	got := circle.Area()
	want  := 78.53981633974483
	if got != want {
		t.Errorf("got %g want %g", got, want)	
	}
})
}
func TestCCircumference(t *testing.T) {
	circle := Circle{5.0}
	got := circle.Circumference()
	want := 31.41592653589793

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}