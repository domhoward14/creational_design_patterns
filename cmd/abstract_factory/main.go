package main

import (
	"errors"
	"fmt"
	"os"
)

type SmallVehicle interface {
	DriveSmallVehicle()
}

type LargeVehicle interface {
	DriveLargeVehicle()
}

type JetSki struct{}

func (j JetSki) DriveSmallVehicle() {
	fmt.Println("Driving JetSki!")
}

type Boat struct{}

func (b Boat) DriveLargeVehicle() {
	fmt.Println("Driving Boat!")
}

type Motorcycle struct{}

func (j Motorcycle) DriveSmallVehicle() {
	fmt.Println("Driving Motorcycle!")
}

type Car struct{}

func (b Car) DriveLargeVehicle() {
	fmt.Println("Driving Car!")
}

type VehicleFactory interface {
	CreateSmallVehicle() SmallVehicle
	CreateLargeVehicle() LargeVehicle
}

type LandFactory struct{}

func (l LandFactory) CreateSmallVehicle() SmallVehicle {
	return Motorcycle{}
}

func (l LandFactory) CreateLargeVehicle() LargeVehicle {
	return Car{}
}

type SeaFactory struct{}

func (l SeaFactory) CreateSmallVehicle() SmallVehicle {
	return JetSki{}
}

func (l SeaFactory) CreateLargeVehicle() LargeVehicle {
	return Boat{}
}

func GetVehicleFactory(terrainType string) (VehicleFactory, error) {
	switch terrainType {
	case "land":
		return LandFactory{}, nil
	case "sea":
		return SeaFactory{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Invalid terrain type: %s", terrainType))
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two command line arguments: passenger type ('small' or 'large') and terrain type ('land' or 'sea').")
		return
	}

	passengerType := os.Args[1]
	terrainType := os.Args[2]

	factory, err := GetVehicleFactory(terrainType)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	switch passengerType {
	case "small":
		factory.CreateSmallVehicle().DriveSmallVehicle()
	case "large":
		factory.CreateLargeVehicle().DriveLargeVehicle()
	default:
		fmt.Printf("Invalid passenger type provided: %s\n", passengerType)
		os.Exit(1)
	}
}
