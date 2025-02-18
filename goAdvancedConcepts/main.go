package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurement.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rowData := scanner.Text()
		semiColon := strings.Index(rowData, ";")
		location := rowData[:semiColon]
		rawtemp := rowData[semiColon+1:]

		temp, _ := strconv.ParseFloat(rawtemp, 64)

		measurement, ok := data[location]
		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}

		data[location] = measurement
	}
	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)
	}
	sort.Strings(locations)
	fmt.Printf("{")
	for _, name := range locations {
		measurement := data[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurement.Max,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")
}
