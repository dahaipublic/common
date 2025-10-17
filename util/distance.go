package util

import (
	"fmt"
	"math"
	"strconv"
)

const (
	earthRadiusKm = 6371.0
)

// degreesToRadians 将角度转换为弧度
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// haversineDistance 使用Haversine公式计算两个坐标点之间的球面距离（单位：千米）
func HaversineDistance(latitude1, longitude1, latitude2, longitude2 float64) (distance float64) {
	lat1 := degreesToRadians(latitude1)
	lon1 := degreesToRadians(longitude1)
	lat2 := degreesToRadians(latitude2)
	lon2 := degreesToRadians(longitude2)

	dlat := lat2 - lat1
	dlon := lon2 - lon1

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadiusKm * c
	distance, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", distance), 64)
	return
}
