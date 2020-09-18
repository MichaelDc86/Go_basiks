package main

import (
	"fmt"
	"math"
)

const pi = 3.1415

// func main() {
//     var phrase string = "Go рулит!"
//     fmt.Println(phrase)
// }
func main() {
	printCircleArea()
}

func printCircleArea() {
	circleRadius := 2
	circleArea := math.Pow(float64(circleRadius), 2) * float64(pi)

	fmt.Printf("Радиус: %d см.\n", circleRadius)
	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea)
}
