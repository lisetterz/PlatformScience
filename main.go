package main

import (
	"bufio"
	"example/lramirez/sde/structs"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	addressesFile = "data/10-list-addresses.txt"
	driversFile   = "data/10-list-drivers.txt"
)

func main() {

	// read the input files
	addresses, err := readFile(addressesFile)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	drivers, err := readFile(driversFile)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	// get list of relations and scores
	scoreMap, err := getScore(addresses, drivers)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}

	sortedMap := sortByScore(scoreMap)

	output := assignShipments(drivers, addresses, sortedMap)

	// print sorted map
	var totalSS float32
	for _, j := range output {
		totalSS = totalSS + j.Ss
		fmt.Printf("%s %s\n", j.Shipment.Address, j.Shipment.Driver)
	}

	fmt.Printf("\n Total SS: %v", totalSS)
}

func sortByScore(score map[*structs.Shipment]float32) []*structs.File {
	a := []*structs.File{}
	for i, j := range score {
		a = append(a, &structs.File{
			Shipment: i,
			Ss:       j,
		})
	}
	sort.Slice(a, func(d, e int) bool {
		return a[d].Ss > a[e].Ss
	})

	return a

}
func searchList(list []string, name string) (bool, int) {

	for i, j := range list {
		if j == name {
			return true, i
		}
	}

	return false, 0

}

func removeItem(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func assignShipments(drivers, addresses []string, combinations []*structs.File) []*structs.File {
	var output []*structs.File

	for _, combination := range combinations {

		driverExists, driverIndex := searchList(drivers, combination.Shipment.Driver)

		addressExists, addressIndex := searchList(addresses, combination.Shipment.Address)

		if driverExists && addressExists {
			output = append(output, combination)
			drivers = removeItem(drivers, driverIndex)
			addresses = removeItem(addresses, addressIndex)
		}

	}

	return output
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func readFile(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %s", fileName, err)
	}

	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	return result, nil

}

func getScore(addresses, drivers []string) (map[*structs.Shipment]float32, error) {
	shippingMap := make(map[*structs.Shipment]float32)

	var shipment *structs.Shipment

	for _, address := range addresses {
		for _, driver := range drivers {

			// calculate SS
			ss, _ := calculateSS(address, driver)

			// create struct
			shipment = &structs.Shipment{
				Address: address,
				Driver:  driver,
			}
			// add combination to map with SS
			shippingMap[shipment] = ss

		}
	}

	return shippingMap, nil
}

func calculateSS(address, driver string) (float32, error) {
	var ss float32
	addressLength := len(strings.Trim(address, " "))
	driverLength := len(strings.Trim(driver, " "))

	// calculate drivers vowels
	driversVowels, driversConsonants := countVowelsAndConsonants(driver)

	// calculus
	if isEven(addressLength) {
		ss = float32(driversVowels) * 1.5
	} else {
		ss = float32(driversConsonants)
	}

	// common factor 50% bonus
	if addressLength%driverLength == 0 {
		ss = ss * 1.5
	}

	return ss, nil
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func isVowel(char rune) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return true

	} else {

		return false

	}

}

func countVowelsAndConsonants(str string) (int, int) {
	var vowelsCount, consonantsCount int

	str = strings.Trim(str, " ")

	for _, char := range str {

		if isVowel(char) {
			vowelsCount = vowelsCount + 1
		} else {
			consonantsCount = consonantsCount + 1
		}

	}

	return vowelsCount, consonantsCount
}
