package main

import (
	"fmt"
	_ "strings"

	"github.com/aj387/Go-Percpio/math"
)

func main() {
	vals := []float64{1, 10, 110, 1000}
	fmt.Println(math.Add(10, 20))
	fmt.Println(math.Avg(vals))
}
