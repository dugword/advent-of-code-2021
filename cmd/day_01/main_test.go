package main_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	main "github.com/dugword/advent-of-code-2021/cmd/day_01"
)

func TestLoadDepthMeasurements(t *testing.T) {
	t.Run("load depth measurements from file", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, err := main.LoadDepthMeasurements("./testdata/input")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("fail on missing file", func(t *testing.T) {
		_, got := main.LoadDepthMeasurements("./testdata/missing")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
	t.Run("fail on invalid measurement", func(t *testing.T) {
		_, got := main.LoadDepthMeasurements("./testdata/invalid")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
}

func TestCountMeasurementIncreases(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1},
			want:  0,
		},
		{
			input: []int{1, 2},
			want:  1,
		},
		{
			input: []int{3, 2, 1},
			want:  0,
		},
		{
			input: []int{1, 3, 1},
			want:  1,
		},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := main.CountMeasurementIncreases(testCase.input)
			if got != testCase.want {
				t.Errorf("got: %d, want: %d", got, testCase.want)
			}
		})
	}
}

func TestCountMeasurementWindowIncreases(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 2, 3},
			want:  0,
		},
		{
			input: []int{1, 2, 3, 4},
			want:  1,
		},
		{
			input: []int{4, 3, 2, 1},
			want:  0,
		},
		{
			input: []int{1, 3, 1, 4},
			want:  1,
		},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := main.CountMeasurementWindowIncreases(testCase.input)
			if got != testCase.want {
				t.Errorf("got: %d, want: %d", got, testCase.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("run without failure", func(t *testing.T) {
		args := []string{
			"day_01",
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
			"day_01",
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
		want := "must provide a measurements file"
		args := []string{
			"day_01",
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
			"day_01",
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
	t.Run("fail on invalid measurements file", func(t *testing.T) {
		want := "invalid measurements file:"
		args := []string{
			"day_01",
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
