package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if err := Run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

// LoadXXX from a file.
func LoadXXX(filepath string) error {
	contents, err := os.ReadFile(filepath)
	if err != nil {
	}
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}
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
	_, err := LoadXXX(*inputFilepath)
	if err != nil {
		return fmt.Errorf("invalid input file: %w", err)
	}
	if *part2 {
	} else {
	}
	return nil
}
