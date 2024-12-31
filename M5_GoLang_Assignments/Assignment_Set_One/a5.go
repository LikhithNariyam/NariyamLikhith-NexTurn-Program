package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CityData represents climate data for a city
type CityData struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

// Function to find the city with the highest temperature
func findHighestTemperature(cities []CityData) CityData {
	highest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

// Function to find the city with the lowest temperature
func findLowestTemperature(cities []CityData) CityData {
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

// Function to calculate the average rainfall
func calculateAverageRainfall(cities []CityData) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter cities by rainfall threshold
func filterCitiesByRainfall(cities []CityData, threshold float64) []CityData {
	var filteredCities []CityData
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

// Function to search for a city by name
func searchCityByName(cities []CityData, name string) (CityData, error) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, nil
		}
	}
	return CityData{}, errors.New("city not found")
}

// Main function
func main() {
	// Climate data for Indian cities
	cities := []CityData{
		{"Mumbai", 27.5, 2000},
		{"Delhi", 31.2, 700},
		{"Bengaluru", 24.3, 850},
		{"Chennai", 29.1, 1400},
		{"Kolkata", 28.4, 1700},
		{"Hyderabad", 27.0, 800},
		{"Jaipur", 33.0, 600},
		{"Pune", 26.5, 750},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Climate Data Analysis Menu ---")
		fmt.Println("1. View City with Highest Temperature")
		fmt.Println("2. View City with Lowest Temperature")
		fmt.Println("3. Calculate Average Rainfall")
		fmt.Println("4. Filter Cities by Rainfall Threshold")
		fmt.Println("5. Search City by Name")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			highest := findHighestTemperature(cities)
			fmt.Printf("City with the highest temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)

		case "2":
			lowest := findLowestTemperature(cities)
			fmt.Printf("City with the lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)

		case "3":
			averageRainfall := calculateAverageRainfall(cities)
			fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)

		case "4":
			fmt.Print("Enter rainfall threshold (mm): ")
			thresholdInput, _ := reader.ReadString('\n')
			thresholdInput = strings.TrimSpace(thresholdInput)

			threshold, err := strconv.ParseFloat(thresholdInput, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			filteredCities := filterCitiesByRainfall(cities, threshold)
			if len(filteredCities) == 0 {
				fmt.Println("No cities found with rainfall above the given threshold.")
			} else {
				fmt.Println("Cities with rainfall above the threshold:")
				for _, city := range filteredCities {
					fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
				}
			}

		case "5":
			fmt.Print("Enter city name: ")
			cityName, _ := reader.ReadString('\n')
			cityName = strings.TrimSpace(cityName)

			city, err := searchCityByName(cities, cityName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("City: %s\nTemperature: %.2f°C\nRainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			}

		case "6":
			fmt.Println("Exiting Climate Data Analysis. Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
