package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := Run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

// LoadDepthMeasurements from a file.
func LoadDepthMeasurements(filepath string) ([]int, error) {
	var measurements []int
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}
		measurement, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, measurement)
	}
	return measurements, nil
}

// CountMeasurementIncreases in the slice of measurements.
func CountMeasurementIncreases(measurements []int) int {
	count := 0
	if len(measurements) < 2 {
		return count
	}
	lastMeasurement := measurements[0]
	for _, measurement := range measurements[1:] {
		if measurement > lastMeasurement {
			count++
		}
		lastMeasurement = measurement
	}
	return count
}

// CountMeasurementWindowIncreases in the slice of measurements.
func CountMeasurementWindowIncreases(measurements []int) int {
	count := 0
	if len(measurements) < 4 {
		return count
	}
	measurementWindow := measurements[:3]
	lastMeasurementWindow := measurementWindow
	for _, measurement := range measurements[3:] {
		measurementWindow = append(measurementWindow[1:], measurement)
		if sum(measurementWindow) > sum(lastMeasurementWindow) {
			count++
		}
		lastMeasurementWindow = measurementWindow
	}
	return count
}

// Run is an abstraction for main() that enables testing.
func Run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("day_01", flag.ContinueOnError)
	flags.SetOutput(stderr)
	inputFilepath := flags.String("input", "", "path to measurements file")
	part2 := flags.Bool("part-2", false, "use part 2 logic")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if *inputFilepath == "" {
		return errors.New("must provide a measurements file")
	}
	measurements, err := LoadDepthMeasurements(*inputFilepath)
	if err != nil {
		return fmt.Errorf("invalid measurements file: %w", err)
	}
	count := 0
	if *part2 {
		count = CountMeasurementWindowIncreases(measurements)
	} else {
		count = CountMeasurementIncreases(measurements)
	}
	fmt.Fprintf(stdout, "%d\n", count)
	return nil
}

func sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
