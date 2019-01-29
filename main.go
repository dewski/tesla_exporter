package main

import (
	"fmt"
	"os"

	"github.com/jsgoecke/tesla"
)

func main() {
	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     os.Getenv("TESLA_CLIENT_ID"),
			ClientSecret: os.Getenv("TESLA_CLIENT_SECRET"),
			Email:        os.Getenv("TESLA_EMAIL"),
			Password:     os.Getenv("TESLA_PASSWORD"),
		})
	if err != nil {
		panic(err)
	}

	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}

	vehicle := vehicles[0]
	fmt.Println(vehicle.State)

	resp, err := vehicle.Wakeup()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// status, err := vehicle.MobileEnabled()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(status)

	vehicleState, err := vehicle.VehicleState()
	if err != nil {
		panic(err)
	}
	fmt.Println(vehicleState.Odometer)

	climateState, err := vehicle.ClimateState()
	if err != nil {
		panic(err)
	}
	fmt.Println(climateState.InsideTemp)
	fmt.Println(climateState.OutsideTemp)

}
