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

// LoadDiagnostics from a file.
func LoadDiagnostics(filepath string) ([]int, error) {
	var diagnostics []int
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}
		diagnostic, err := strconv.ParseInt(line, 2, 16)
		if err != nil {
			return nil, err
		}
		diagnostics = append(diagnostics, int(diagnostic))
	}
	return diagnostics, nil
}

// DecodeReport from a slice of diagnostics.
func DecodeReport(diagnostics []int) (gamma, epsilon int) {
	bitCount := make([]int, 12)
	for _, diagnostic := range diagnostics {
		for i := range bitCount {
			mask := (1 << i)
			if diagnostic&mask != 0 {
				bitCount[len(bitCount)-1-i]++
			}
		}
	}
	d := 0
	for _, count := range bitCount {
		d <<= 1
		if count > len(diagnostics)/2 {
			d++
		}
	}
	return d, 0xfff ^ d
}

// GetOxygenGeneratorRating from a slice of diagnostics.
func GetOxygenGeneratorRating(bitLength int, diagnostics []int) int {
	bitLength--
	if bitLength < 0 || len(diagnostics) == 1 {
		return diagnostics[0]
	}
	ones, zeros := filterDiagnostics(bitLength, diagnostics)
	switch {
	case len(zeros) == 0:
		return GetOxygenGeneratorRating(bitLength, ones)
	case len(ones) == 0:
		return GetOxygenGeneratorRating(bitLength, zeros)
	case len(ones) >= len(zeros):
		return GetOxygenGeneratorRating(bitLength, ones)
	default:
		return GetOxygenGeneratorRating(bitLength, zeros)
	}
}

// GetCO2ScrubberRating from a slice of diagnostics.
func GetCO2ScrubberRating(bitLength int, diagnostics []int) int {
	bitLength--
	if bitLength < 0 || len(diagnostics) == 1 {
		return diagnostics[0]
	}
	ones, zeros := filterDiagnostics(bitLength, diagnostics)
	switch {
	case len(zeros) == 0:
		return GetCO2ScrubberRating(bitLength, ones)
	case len(ones) == 0:
		return GetCO2ScrubberRating(bitLength, zeros)
	case len(zeros) > len(ones):
		return GetCO2ScrubberRating(bitLength, ones)
	default:
		return GetCO2ScrubberRating(bitLength, zeros)
	}
}

// Run is an abstraction for main() that enables testing.
func Run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("day_01", flag.ContinueOnError)
	flags.SetOutput(stderr)
	inputFilepath := flags.String("input", "", "path to input file")
	part2 := flags.Bool("part-2", false, "use part 2 logic")
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	if *inputFilepath == "" {
		return errors.New("must provide an input file")
	}
	diagnostics, err := LoadDiagnostics(*inputFilepath)
	if err != nil {
		return fmt.Errorf("invalid input file: %w", err)
	}
	if *part2 {
		oxygenGeneratorRating := GetOxygenGeneratorRating(12, diagnostics)
		co2ScrubberRating := GetCO2ScrubberRating(12, diagnostics)
		fmt.Fprintf(stdout, "%d\n", oxygenGeneratorRating*co2ScrubberRating)
	} else {
		gamma, epsilon := DecodeReport(diagnostics)
		fmt.Fprintf(stdout, "%d\n", gamma*epsilon)
	}
	return nil
}

func filterDiagnostics(bitLength int, diagnostics []int) (ones, zeros []int) {
	mask := 1 << bitLength
	for _, diagnostic := range diagnostics {
		if diagnostic&mask == 0 {
			zeros = append(zeros, diagnostic)
		} else {
			ones = append(ones, diagnostic)
		}
	}
	return ones, zeros
}
