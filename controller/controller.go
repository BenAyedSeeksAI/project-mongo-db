package controller

import (
	"fmt"

	"github.com/BenAyedSeeksAI/project-mongo-db/db"
	"github.com/spf13/cobra"
)

func AddCarCmd(cmd *cobra.Command, args []string) {
	var model, brand, bodyWork string
	var horsePower int64
	inputLabel := []string{"Model: ", "Brand: ", "Body Work: ", "Horse Power: "}
	inputPlaceholder := []string{"%s", "%s", "%s", "%d"}
	inputValue := []interface{}{model, brand, bodyWork, horsePower}
	for idx, elem := range inputLabel {
		fmt.Printf("%s \n", elem)
		fmt.Scanf(inputPlaceholder[idx], inputValue[idx])
	}
	carObj := db.Car{Model: model, Brand: brand, BodyWork: bodyWork, HorsePower: horsePower}
	db.DBInsertCar(carObj)
}
func GetCarsCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Here are the available cars.")
	db.DBGetCars()
}
