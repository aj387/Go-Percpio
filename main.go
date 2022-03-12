package main

import (
	"fmt"
	"math"
	"os"
	_ "strings"

	m "github.com/aj387/Go-Percpio/math"
)

func main() {
	vals := []float64{1, 10, 110, 1000}
	fmt.Println(m.Add(10, 20))
	fmt.Println(m.Sub(10, 20))
	fmt.Println(m.Pi)
	fmt.Println(math.Pi)
	fmt.Fprintln(os.Stdout, "Hello !!!")
	fmt.Println(m.Avg(vals))
}
