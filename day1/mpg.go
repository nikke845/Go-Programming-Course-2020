package main

import "fmt"
import "flag"

func main(){
	var nationality string
	var distance float64
	var fuelUsed float64
	var fuelCost float64
	flag.StringVar(&nationality, "nationality", "European", "American or European style units i.e. gallons vs. litres")
	flag.Float64Var(&distance, "distance", 600, "Distance driven")
	flag.Float64Var(&fuelUsed, "fuel-used", 33, "Amount of fuel refueled")
	flag.Float64Var(&fuelCost, "fuel-cost", 1.3, "Cost of fuel per gallon or litre")
	flag.Parse()
	if nationality == "European" {
		mpg := distance / fuelUsed
		litresPer100km := fuelUsed / distance * 100
		costPerMile := fuelCost / mpg
		fmt.Println("Average consumption is", litresPer100km, "L/100km and cost per km is", costPerMile, "€")
	} else if nationality == "American" {
		mpg := distance / fuelUsed
		costPerMile := fuelCost / mpg
		fmt.Println("Fuel economy is", mpg, "miles per gallon and cost per mile is", costPerMile, "$")
	} else {
		fmt.Println("Incorrect syntax for nationality, choose either American or European")
	}
}