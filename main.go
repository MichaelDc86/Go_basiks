package main

import (
	"fmt"
	"math"
	"errors"
)

// const pi = 3.1415
// const pi = math.Pi

// func main() {
//     var phrase string = "Go рулит!"
//     fmt.Println(phrase)
// }
func main() {
	printCircleArea(2)
	printCircleArea(4)
	printCircleArea(-2)
}

func printCircleArea(radius int) {
	circlArea, err := calcCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус: %d см.\n", radius)
	fmt.Printf("Площадь круга: %f32 см. кв.\n", circlArea)
}

func calcCircleArea(radius int) (float64, error) {
	if radius <= 0 {
		return float64(0), errors.New("Радиус не может быть отрицательным!")
	}
	return math.Pow(float64(radius), 2) * float64(math.Pi), nil
}
