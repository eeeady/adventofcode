package fuel

import "testing"

func TestMassToFuel(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		output := MassToFuel(table.mass)
		if output != table.fuel {
			t.Errorf("Nope: for mass %d, got %d, should be %d", table.mass, output, table.fuel)
		}
	}
}

func TestMassToFuelPlusFuel(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		output := MassToFuelPlusFuel(table.mass)
		if output != table.fuel {
			t.Errorf("Nope: for mass %d, got %d, should be %d", table.mass, output, table.fuel)
		}
	}
}
