package riego

import "time"

type RegistroClimatico struct {
	fecha         time.Time
	temperatura   float64
	humedad       float64
	viento        float64
	radiacion     float64
	precipitacion float64
	eto           float64
}