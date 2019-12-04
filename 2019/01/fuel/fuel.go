package fuel

func MassToFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func MassToFuelPlusFuel(mass int) int {
	totalFuelRequired := 0
	fuel := MassToFuel(mass)
	totalFuelRequired += fuel
	for fuel > 0 {
		fuel = MassToFuel(fuel)
		totalFuelRequired += fuel
	}
	return totalFuelRequired
}
