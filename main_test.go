package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "This is a simple test case.",
			expected: []string{"this", "is", "a", "simple", "test", "case."},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "This Is Another Simple Test Case.",
			expected: []string{"this", "is", "another", "simple", "test", "case."},
		},
		{
			input:    "ALL CAPS TEST CASE",
			expected: []string{"all", "caps", "test", "case"},
		},
		{
			input:    "    Test   CRAZY    tEsT case,      that   has  ALL tyPEs   of WEirD   stuff GOIng    on.     ",
			expected: []string{"test", "crazy", "test", "case,", "that", "has", "all", "types", "of", "weird", "stuff", "going", "on."},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		failed := 0
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(len(actual), len(c.expected)) {
			failCount++
			t.Errorf(`---------------------------------
Text:     %v
Expecting: %+v
Actual:   %+v
Fail`, c.input, c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if !reflect.DeepEqual(word, expectedWord) {
				failed++
				failCount++
				t.Errorf(`---------------------------------
Text:     %v
Expecting: %+v
Actual:   %+v
Fail`, c.input, c.expected, actual)
				break
			}
		}
		if failed == 0 {
			passCount++
			t.Logf(`---------------------------------
Text:		%v
Expecting:  %+v
Actual:     %+v
Pass
`, c.input, c.expected, actual)
		}
	}

	t.Logf("---------------------------------")
	t.Logf("%d passed, %d failed\n", passCount, failCount)
}
