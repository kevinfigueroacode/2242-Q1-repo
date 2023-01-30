package main

import (
	"testing"
)//test for the Area method
func TestTArea(t *testing.T) {
	
	t.Run("Triangles", func(t *testing.T) {
		triangle := Triangle{12.0, 6.0}
		got:=triangle.Area()
		want := 36.0
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	

}
//test for the perimeter method
//like a breaker
func TestPermimeter(t *testing.T) {
	
	t.Run("Triangles", func(t *testing.T) {
		triangle := Triangle{12.0, 6.0}
		got:=triangle.Perimeter()
		want := 31.41640786499874 
		
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}