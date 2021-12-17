package main_test

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	main "github.com/dugword/advent-of-code-2021/cmd/day_0x"
)

func TestLoadXXX(t *testing.T) {
	t.Run("load xxx from file", func(t *testing.T) {
		want := []int{
			0b000111111001,
			0b111011110110,
			0b101111111000,
		}
		got, err := main.LoadXXX("./testdata/input")
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("fail on missing file", func(t *testing.T) {
		_, got := main.LoadXXX("./testdata/missing")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
	t.Run("fail on invalid diagnostic", func(t *testing.T) {
		_, got := main.LoadXXX("./testdata/invalid")
		if got == nil {
			t.Error("did not fail as expected")
		}
	})
}

func TestX(t *testing.T) {
	testCases := []struct {
		input       
		want
	}{
		{
			input: 
			want:
		},
	}
	for i, testCase := range testCases {
		testName := fmt.Sprintf("test case %d", i)
		t.Run(testName, func(t *testing.T) {
			got := main.XXX()
			if got!= testCase.want{
				t.Errorf("got: %q, want: %q", got, testCase.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Run("run without failure", func(t *testing.T) {
		args := []string{
			"day_0x",
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
			"day_0x",
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
			"day_0x",
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
			"day_0x",
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
			"day_0x",
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
