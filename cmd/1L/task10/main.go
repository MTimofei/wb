// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.

package main

var list = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
var group1 = []float64{}
var group2 = []float64{}
var group3 = []float64{}
var group4 = []float64{}

func main() {
	for _, v := range list {
		switch {
		case v < float64(-20):
			group1 = append(group1, v)
		case v > float64(10) && v < float64(20):
			group2 = append(group2, v)
		case v > float64(20) && v < float64(30):
			group3 = append(group3, v)
		case v > float64(30):
			group4 = append(group4, v)
		}
	}
}
