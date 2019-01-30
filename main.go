package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/jsgoecke/tesla"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Celsius represents C temperatures
type Celsius float32

// Fahrenheit represents F temperatures
type Fahrenheit float32

// ToFahrenheit converts Celsius temperate to equivelent Fahrenheit
func (t Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(t*9/5 + 32)
}

func buildClient() (*tesla.Client, error) {
	clientID := os.Getenv("TESLA_CLIENT_ID")
	if clientID == "" {
		return nil, errors.New("Need to set TESLA_CLIENT_ID")
	}

	clientSecret := os.Getenv("TESLA_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, errors.New("Need to set TESLA_CLIENT_SECRET")
	}

	teslaEmail := os.Getenv("TESLA_EMAIL")
	if teslaEmail == "" {
		return nil, errors.New("Need to set TESLA_EMAIL")
	}

	teslaPassword := os.Getenv("TESLA_PASSWORD")
	if teslaPassword == "" {
		return nil, errors.New("Need to set TESLA_PASSWORD")
	}

	client, err := tesla.NewClient(
		&tesla.Auth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Email:        teslaEmail,
			Password:     teslaPassword,
		})

	if err != nil {
		return nil, err
	}

	return client, nil
}

var (
	odometer = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_odometer",
			Help: "Vehicle odometer reading",
		},
		[]string{"vin", "name"},
	)

	temperatureInside = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_temperature_inside",
			Help: "Vehicle temperature readings",
		},
		[]string{"vin", "name", "format"},
	)

	temperatureOutside = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_temperature_outside",
			Help: "Vehicle temperature readings",
		},
		[]string{"vin", "name", "format"},
	)

	temperatureDriver = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_temperature_driver",
			Help: "Vehicle temperature readings",
		},
		[]string{"vin", "name", "format"},
	)

	temperaturePassenger = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_temperature_passenger",
			Help: "Vehicle temperature readings",
		},
		[]string{"vin", "name", "format"},
	)

	batteryLevel = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_battery_level",
			Help: "Vehicle charge readings",
		},
		[]string{"vin", "name"},
	)

	batteryLevelUsable = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_battery_level_usable",
			Help: "Vehicle charge readings",
		},
		[]string{"vin", "name"},
	)

	batteryRange = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_battery_range",
			Help: "Vehicle range readings",
		},
		[]string{"vin", "name"},
	)

	batteryRangeIdeal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_battery_range_ideal",
			Help: "Vehicle range readings",
		},
		[]string{"vin", "name"},
	)

	batteryRangeEstimated = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tesla_battery_range_estimaated",
			Help: "Vehicle range readings",
		},
		[]string{"vin", "name"},
	)
)

func init() {
	prometheus.MustRegister(odometer)
	prometheus.MustRegister(temperatureInside)
	prometheus.MustRegister(temperatureOutside)
	prometheus.MustRegister(temperatureDriver)
	prometheus.MustRegister(temperaturePassenger)
	prometheus.MustRegister(batteryLevel)
	prometheus.MustRegister(batteryLevelUsable)
	prometheus.MustRegister(batteryRange)
	prometheus.MustRegister(batteryRangeIdeal)
	prometheus.MustRegister(batteryRangeEstimated)
}

func gaugeOdometer(v *tesla.Vehicle) error {
	state, err := v.VehicleState()
	if err != nil {
		return err
	}

	odometer.With(prometheus.Labels{"vin": v.Vin, "name": v.DisplayName}).Set(state.Odometer)

	return nil
}

func gaugeClimate(v *tesla.Vehicle) error {
	state, err := v.ClimateState()
	if err != nil {
		return err
	}

	temperatureInside.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "c",
	}).Set(state.InsideTemp)

	insideTempC := Celsius(state.InsideTemp)
	temperatureInside.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "f",
	}).Set(float64(insideTempC.ToFahrenheit()))

	temperatureOutside.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "c",
	}).Set(state.OutsideTemp)

	outsideTempC := Celsius(state.OutsideTemp)
	temperatureOutside.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "f",
	}).Set(float64(outsideTempC.ToFahrenheit()))

	temperatureDriver.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "c",
	}).Set(state.DriverTempSetting)

	driverTempC := Celsius(state.DriverTempSetting)
	temperatureDriver.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "f",
	}).Set(float64(driverTempC.ToFahrenheit()))

	temperaturePassenger.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "c",
	}).Set(state.PassengerTempSetting)

	passengerTempC := Celsius(state.PassengerTempSetting)
	temperaturePassenger.With(prometheus.Labels{
		"vin":    v.Vin,
		"name":   v.DisplayName,
		"format": "f",
	}).Set(float64(passengerTempC.ToFahrenheit()))

	return nil
}

func gaugeChargeState(v *tesla.Vehicle) error {
	state, err := v.ChargeState()
	if err != nil {
		return err
	}

	batteryRange.With(prometheus.Labels{
		"vin":  v.Vin,
		"name": v.DisplayName,
	}).Set(state.BatteryRange)

	batteryRangeEstimated.With(prometheus.Labels{
		"vin":  v.Vin,
		"name": v.DisplayName,
	}).Set(state.EstBatteryRange)

	batteryRangeIdeal.With(prometheus.Labels{
		"vin":  v.Vin,
		"name": v.DisplayName,
	}).Set(state.IdealBatteryRange)

	batteryLevel.With(prometheus.Labels{
		"vin":  v.Vin,
		"name": v.DisplayName,
	}).Set(float64(state.BatteryLevel))

	batteryLevelUsable.With(prometheus.Labels{
		"vin":  v.Vin,
		"name": v.DisplayName,
	}).Set(float64(state.UsableBatteryLevel))

	return nil
}

func main() {
	client, err := buildClient()
	if err != nil {
		panic(err)
	}

	vehicles, err := client.Vehicles()
	if err != nil {
		panic(err)
	}

	for _, vehicle := range vehicles {
		if vehicle.State != "online" {
			_, err = vehicle.Wakeup()
			if err != nil {
				panic(err)
			}
		}

		_ = gaugeOdometer(vehicle.Vehicle)
		_ = gaugeClimate(vehicle.Vehicle)
		_ = gaugeChargeState(vehicle.Vehicle)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8752" // TSLA
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
