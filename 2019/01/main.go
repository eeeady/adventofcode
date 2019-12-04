package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"./fuel"
)

func main() {
	totalFuel := 0
	totalFuelPlusFuel := 0

	args := os.Args[1:]
	inputFile := args[0]

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		totalFuel += fuel.MassToFuel(mass)
		totalFuelPlusFuel += fuel.MassToFuelPlusFuel(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalFuel)
	fmt.Println(totalFuelPlusFuel)
}
