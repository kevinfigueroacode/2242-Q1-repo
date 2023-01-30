
package main
import (
	//call math package
	"math"
	
)
//initializing valriable
type Circle struct {
	radius float64
}
//method for Area of a circle
//returns Area
func (c Circle) Area() float64 {
	var Area = math.Pi * math.Pow(c.radius, 2)
	return Area
}
//method for the Circumfernce of Circle
//returns Circumference
func (c Circle) Circumference() float64 {
	var Circumference = 2 * math.Pi * c.radius
	return Circumference
}

/*func main() {
	circ1:= Circle {
		radius: 4,
	}
	circ2:= Circle {
		radius:6,
	}
	fmt.Println("The Area of the Circle: ",circ1.Area())
	fmt.Println("The Circumfernce of the Circle is: ", circ2.Circumference())
}*/
