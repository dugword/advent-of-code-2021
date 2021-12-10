package main_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	main "github.com/dugword/advent-of-code-2021/cmd/day_03"
)

func TestLoadDiagnostics(t *testing.T) {
	t.Run("load diagnostics from file", func(t *testing.T) {
		want := []int{
			0b000111111001,
			0b111011110110,
			0b101111111000,
		}
		got, err := main.LoadDiagnostics("./testdata/input")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("fail on missing file", func(t *testing.T) {
		_, got := main.LoadDiagnostics("./testdata/missing")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
	t.Run("fail on invalid diagnostic", func(t *testing.T) {
		_, got := main.LoadDiagnostics("./testdata/invalid")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
}

func TestDecodeReport(t *testing.T) {
	testCases := []struct {
		input       []int
		wantGamma   int
		wantEpsilon int
	}{
		{
			input: []int{
				0b000000000001,
				0b000000000010,
				0b000000000011,
			},
			wantGamma:   0b000000000011,
			wantEpsilon: 0b111111111100,
		},
		{
			input: []int{
				0b000000000000,
				0b000000000010,
				0b000000000001,
			},
			wantGamma:   0b000000000000,
			wantEpsilon: 0b111111111111,
		},
		{
			input: []int{
				0b000000000001,
				0b000000000010,
				0b000000000010,
			},
			wantGamma:   0b000000000010,
			wantEpsilon: 0b111111111101,
		},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			gotGamma, gotEpsilon := main.DecodeReport(testCase.input)
			if gotGamma != testCase.wantGamma {
				t.Errorf("\ngot gamma:  %12.12b\nwant gamma: %12.12b", gotGamma, testCase.wantGamma)
			}
			if gotEpsilon != testCase.wantEpsilon {
				t.Errorf("\ngot epsilon:  %12.12b\nwant epsilon: %12.12b", gotEpsilon, testCase.wantEpsilon)
			}
		})
	}
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{{
		input: []int{
			0b000000000001,
			0b000000000010,
			0b000000000011,
		},
		want: 0b000000000011,
	}, {
		input: []int{
			0b000000000101,
			0b000000000110,
			0b000000000011,
		},
		want: 0b000000000110,
	}, {
		input: []int{
			0b100000000111,
			0b111111111000,
			0b100000000000,
		},
		want: 0b100000000111,
	}, {
		input: []int{
			0b000000000100,
			0b000000011110,
			0b000000010110,
			0b000000010111,
			0b000000010101,
			0b000000001111,
			0b000000000111,
			0b000000011100,
			0b000000010000,
			0b000000011001,
			0b000000000010,
			0b000000001010,
		},
		want: 0b000000010111,
	}}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := main.GetOxygenGeneratorRating(12, testCase.input)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ngot:  %12.12b\nwant: %12.12b", got, testCase.want)
			}
		})
	}
}

func TestGetCO2ScrubberRating(t *testing.T) {
	testCases := []struct {
		input []int
		want  int
	}{{
		input: []int{
			0b000000000001,
			0b000000000010,
			0b000000000011,
		},
		want: 0b000000000001,
	}, {
		input: []int{
			0b000000000101,
			0b000000000110,
			0b000000000011,
		},
		want: 0b000000000011,
	}, {
		input: []int{
			0b100000000111,
			0b111111111000,
			0b100000000000,
		},
		want: 0b111111111000,
	},{
		input: []int{
			0b000000000100,
			0b000000011110,
			0b000000010110,
			0b000000010111,
			0b000000010101,
			0b000000001111,
			0b000000000111,
			0b000000011100,
			0b000000010000,
			0b000000011001,
			0b000000000010,
			0b000000001010,
		},
		want: 0b000000001010,
	}}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := main.GetCO2ScrubberRating(12, testCase.input)
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ngot:  %12.12b\nwant: %12.12b", got, testCase.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("run without failure", func(t *testing.T) {
		args := []string{
			"day_03",
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
			"day_03",
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
			"day_03",
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
			"day_03",
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
			"day_03",
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
