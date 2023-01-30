package main

import (
	//"fmt"
	"math"
)//function to find area and perimeter of a square
//returns the Area and Perimeter of a Square
func Square(s float64) (float64, float64) {
var Area =  math.Pow(s, 2) //side x side
var Perimeter = 4 * s //side + side + side + side
//fmt.Println("The Area of the Square is: ", Area)
//fmt.Println("The Perimeter of the square is: ", Perimeter)
return Area, Perimeter 
}
/*func main() {

	//fmt.Println(Square(2))
fmt.Println(Square(2))

}*/