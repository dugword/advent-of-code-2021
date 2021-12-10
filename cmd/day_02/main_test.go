package main_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	main "github.com/dugword/advent-of-code-2021/cmd/day_02"
)

func TestLoadCommands(t *testing.T) {
	t.Run("load commands from file", func(t *testing.T) {
		want := []main.Command{{
			Direction: "forward",
			Value:     1,
		}, {
			Direction: "down",
			Value:     2,
		}, {
			Direction: "up",
			Value:     3,
		}}
		got, err := main.LoadCommands("./testdata/input")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("fail on missing file", func(t *testing.T) {
		_, got := main.LoadCommands("./testdata/missing")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
	t.Run("fail on invalid command", func(t *testing.T) {
		_, got := main.LoadCommands("./testdata/invalid")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
}

func TestCalculatePosition(t *testing.T) {
	testCases := []struct {
		input []main.Command
		want  int
	}{
		{
			input: []main.Command{{
				Direction: "up",
				Value:     1,
			}, {
				Direction: "down",
				Value:     1,
			}},
			want: 0,
		},
		{
			input: []main.Command{{
				Direction: "up",
				Value:     2,
			}, {
				Direction: "forward",
				Value:     2,
			}},
			want: -4,
		},
		{
			input: []main.Command{{
				Direction: "up",
				Value:     2,
			}, {
				Direction: "down",
				Value:     4,
			}, {
				Direction: "forward",
				Value:     5,
			}},
			want: 10,
		},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			horizontal, depth := main.CalculatePosition(testCase.input)
			got := horizontal * depth
			if got != testCase.want {
				t.Errorf("got: %d, want: %d", got, testCase.want)
			}
		})
	}
}

func TestCalculatePositionWithAim(t *testing.T) {
	testCases := []struct {
		input []main.Command
		want  int
	}{
		{
			input: []main.Command{{
				Direction: "up",
				Value:     1,
			}, {
				Direction: "down",
				Value:     1,
			}},
			want: 0,
		},
		{
			input: []main.Command{{
				Direction: "up",
				Value:     2,
			}, {
				Direction: "forward",
				Value:     2,
			}},
			want: -8,
		},
		{
			input: []main.Command{{
				Direction: "up",
				Value:     2,
			}, {
				Direction: "down",
				Value:     4,
			}, {
				Direction: "forward",
				Value:     5,
			}},
			want: 50,
		},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			horizontal, depth := main.CalculatePositionWithAim(testCase.input)
			got := horizontal * depth
			if got != testCase.want {
				t.Errorf("got: %d, want: %d", got, testCase.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("run without failure", func(t *testing.T) {
		args := []string{
			"day_02",
			"-input", "./testdata/input",
		}
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		if err := main.Run(args, &stdout, &stderr); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("run part 2 without failure", func(t *testing.T) {
		args := []string{
			"day_02",
			"-input", "./testdata/input",
			"-part-2",
		}
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		if err := main.Run(args, &stdout, &stderr); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("fail on missing argument", func(t *testing.T) {
		want := "must provide an input file"
		args := []string{
			"day_02",
		}
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		got := main.Run(args, &stdout, &stderr)
		if got == nil {
			t.Fatal("did not fail as expected")
		}
		if got.Error() != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
	t.Run("fail on invalid argument", func(t *testing.T) {
		want := "flag provided but not defined: -invalid"
		args := []string{
			"day_02",
			"-invalid",
		}
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		got := main.Run(args, &stdout, &stderr)
		if got == nil {
			t.Fatal("did not fail as expected")
		}
		if got.Error() != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
	t.Run("fail on invalid input file", func(t *testing.T) {
		want := "invalid input file:"
		args := []string{
			"day_02",
			"-input", "./testdata/invalid",
		}
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		got := main.Run(args, &stdout, &stderr)
		if got == nil {
			t.Fatal("did not fail as expected")
		}
		if !strings.HasPrefix(got.Error(), want) {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
