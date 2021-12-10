package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Command to give the submarine, includes a direction and a value.
type Command struct {
	Direction string
	Value     int
}

func main() {
	if err := Run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

// LoadCommands from a file.
func LoadCommands(filepath string) ([]Command, error) {
	var commands []Command
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(contents), "\n") {
		if line == "" {
			continue
		}
		re := regexp.MustCompile(`^(?P<direction>down|forward|up)\s+(?P<value>\d+)`)
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			return nil, errors.New("invalid command")
		}
		value, err := strconv.Atoi(match[re.SubexpIndex("value")])
		if err != nil {
			return nil, err
		}
		command := Command{
			Direction: match[re.SubexpIndex("direction")],
			Value:     value,
		}

		commands = append(commands, command)
	}
	return commands, nil
}

// CalculatePosition from a slice of Commands
func CalculatePosition(commands []Command) (horizontal, depth int) {
	for _, command := range commands {
		switch command.Direction {
		case "down":
			depth += command.Value
		case "up":
			depth -= command.Value
		case "forward":
			horizontal += command.Value
		}
	}
	return
}

// CalculatePositionWithAim from a slice of Commands
func CalculatePositionWithAim(commands []Command) (horizontal, depth int) {
	aim := 0
	for _, command := range commands {
		switch command.Direction {
		case "down":
			aim += command.Value
		case "up":
			aim -= command.Value
		case "forward":
			depth += aim * command.Value
			horizontal += command.Value
		}
	}
	return
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
	commands, err := LoadCommands(*inputFilepath)
	if err != nil {
		return fmt.Errorf("invalid input file: %w", err)
	}
	var horizontal int
	var depth int
	if *part2 {
		horizontal, depth = CalculatePositionWithAim(commands)
	} else {
		horizontal, depth = CalculatePosition(commands)
	}
	fmt.Fprintf(stdout, "%d\n", horizontal*depth)
	return nil
}
