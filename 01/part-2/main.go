package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FuelCalculatorType func(int64) int64

func FuelCalculator(m int64) int64 {
	fuel := m/3 - 2

	if fuel > 0 {
		return fuel + FuelCalculator(fuel)
	}

	return 0
}

func LoadMasses(fileLocation string) []int64 {
	file, err := os.Open(fileLocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var masses []int64
	for scanner.Scan() {
		text := scanner.Text()
		mass, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			panic(err)
		}
		masses = append(masses, mass)
	}
	return masses
}

func SumFuels(masses []int64, fuelCalculator FuelCalculatorType) int64 {
	var totalFuels int64 = 0
	for _, mass := range masses {
		fuel := fuelCalculator(mass)
		totalFuels += fuel
	}
	return totalFuels
}

func main() {
	masses := LoadMasses("./inputs.txt")
	totalFuels := SumFuels(masses, FuelCalculator)
	fmt.Println(totalFuels)
}
