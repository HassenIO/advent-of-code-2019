package main

import (
	"testing"
)

func TestFuelCalculator(t *testing.T) {
	if fuel := FuelCalculator(12); fuel != 2 {
		t.Errorf("FuelCalculator(12) = %d; want 2", fuel)
	}
}

func TestSumFuels(t *testing.T) {
	masses := []int64{1, 2, 3}
	fuelCalculator := func(x int64) int64 { return x }

	if totalFuels := SumFuels(masses, fuelCalculator); totalFuels != 6 {
		t.Errorf("SumFuels = %d; want 6", totalFuels)
	}

	masses = []int64{}
	if totalFuels := SumFuels(masses, fuelCalculator); totalFuels != 0 {
		t.Errorf("SumFuels = %d; want 0", totalFuels)
	}
}

func TestLoadMasses(t *testing.T) {
	expectedMasses := []int64{10, 20, 40}
	fileLocation := "./tests/test.txt"
	masses := LoadMasses(fileLocation)

	if len(masses) != len(expectedMasses) {
		t.Errorf("Lengths are not the same %d != %d", len(masses), len(expectedMasses))
	}

	for i, mass := range masses {
		if mass != expectedMasses[i] {
			t.Errorf("Expected mass = %d; got %d", expectedMasses[i], mass)
		}
	}
}
