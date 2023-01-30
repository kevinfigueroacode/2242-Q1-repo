package main

import (
	//"fmt"
	 "math"
)
//assign variables to Triangle
type Triangle struct {
	Base   float64
	Height float64
	
}
//method to calculate for the area of a Triangle
//returns Area
func (t Triangle) Area() float64 {
	var Area = .5 * (t.Base * t.Height)
	return Area
}
//method to calulare the perimeter of a Triangle with its missing side
//returns Perimeter
func (t Triangle) Perimeter() float64 {
var Perimeter = (math.Sqrt(math.Pow(t.Base, 2) + math.Pow(t.Height, 2))) + (t.Base + t.Height)
	return Perimeter
}
/*func Perimeter(t Triangle) float64 {
	var Perimeter = t.sideA + t.sideB + t.sideC
	return Perimeter
}*/
/*func main() {
	tria1 := Triangle{
		Height: 4.4,
		Base:   2.3,
	}
	tria2 := Triangle{
		sideA: 2.1,
		sideB: 3.7,
		sideC: 6.2,
	}

	fmt.Println("Area of Triangle: ", tria1.Area())
	fmt.Println("Perimeter of Triangle: ", tria2.Perimeter())
}
*/