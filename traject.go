package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Object struct {
	x float64
	y float64
}

var (
	deltaT = 0.1
	gVal   = -10.0
	m      = 0.6
	k      = 1.05
)

func csvExport(data [][]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err
		}
	}
	return nil
}

func convertToStringArray(posX float64, posY float64) []string {
	sx := fmt.Sprintf("%f", posX)
	sy := fmt.Sprintf("%f", posY)
	conv := []string{sx, sy}
	return conv
}

func EulerMethod() {
	data := [][]string{}
	position := Object{0, 0}
	velocity := Object{10, 10}
	acceleration := Object{}

	for position.y >= 0 {

		position.x += velocity.x * deltaT
		position.y += velocity.y * deltaT

		acceleration.x = (m*0 - k*velocity.x) / m
		acceleration.y = (m*gVal - k*velocity.y) / m

		velocity.x += acceleration.x * deltaT
		velocity.y += acceleration.y * deltaT

		data = append(data, convertToStringArray(position.x, position.y))
		fmt.Println("x: ", position.x, " ", "y: ", position.y)
	}
	csvExport(data, "resultEulerMethod.csv")

}

func MidPointMethod() {
	data := [][]string{}
	position := Object{0, 0}
	delta_position := Object{0, 0}
	velocity := Object{10, 10}
	velocity2 := Object{10, 10}
	delta_velocity := Object{}
	acceleration := Object{}
	acceleration2 := Object{}
	acceleration.x = (m*0 - k*velocity.x) / m
	acceleration.y = (m*gVal - k*velocity.y) / m

	for position.y >= 0 {
		position.x += delta_position.x
		position.y += delta_position.y

		velocity.x += delta_velocity.x
		velocity.y += delta_velocity.y

		acceleration.x = (m*0 - k*velocity.x) / m
		acceleration.y = (m*gVal - k*velocity.y) / m

		velocity2.x = velocity.x + acceleration.x*deltaT/2.0
		velocity2.y = velocity.y + acceleration.y*deltaT/2.0

		acceleration2.x = (m*0 - k*velocity2.x) / m
		acceleration2.y = (m*gVal - k*velocity2.y) / m

		delta_position.x = velocity2.x * deltaT
		delta_position.y = velocity2.y * deltaT

		delta_velocity.x = acceleration2.x * deltaT
		delta_velocity.y = acceleration2.y * deltaT

		data = append(data, convertToStringArray(position.x, position.y))
		fmt.Println("x: ", position.x, " ", "y: ", position.y)
	}
	csvExport(data, "resultMidPoint.csv")
}

func main() {
	fmt.Println("Euler method")
	EulerMethod()
	fmt.Println("Mid Point method")
	MidPointMethod()
}
