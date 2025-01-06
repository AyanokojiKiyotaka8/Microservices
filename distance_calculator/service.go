package main

import (
	"math"

	"github.com/AyanokojiKiyotaka8/Microservices/types"
)

type CalculatorServicer interface {
	CalcDistance(types.OBUData) (float64, error)
}

type CalculatorService struct {
	prevCordLat  float64
	prevCordLong float64
}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{
		prevCordLat:  -1,
		prevCordLong: -1,
	}
}

func (s *CalculatorService) CalcDistance(data types.OBUData) (float64, error) {
	var dist float64 = 0
	if s.prevCordLat > 0 {
		dist = math.Sqrt(math.Pow(s.prevCordLat-data.Lat, 2) + math.Pow(s.prevCordLong-data.Long, 2))
	}
	s.prevCordLat = data.Lat
	s.prevCordLong = data.Long
	return dist, nil
}
