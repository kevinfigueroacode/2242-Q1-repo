package main

import (
	"testing"
)
//test for the function to find area and perimeter
func SquareArea(t *testing.T) {
got1,got2:= Square(05.) 
want1, want2:= 25.0, 20.0

if got1!= want1 {
	t.Errorf("got %.2f want %.2f", got1, want1)
}
if got2 != want2 {
	t.Errorf("got %.2f want %.2f", got2, want2)
}	

}
