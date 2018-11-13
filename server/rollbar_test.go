package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func helperLoadJsonFile(t *testing.T, name string) *Rollbar {
	path := filepath.Join("testdata", name)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	var rollbar Rollbar

	if err := json.Unmarshal(bytes, &rollbar); err != nil {
		t.Fatal(err)
	}

	return &rollbar
}

type InterpolateMessageTest struct {
	filename string
	template string
	expected string
}

func TestInterpolateMessage(t *testing.T) {
	everyOccurrenceMessage := "new_item - [production] - error"
	highOccurrenceMessage := "item_velocity - [production] - error"
	newItemMessage := "new_item - [production] - error"
	tenNthMessage := "exp_repeat_item - [production] - error"

	testcases := []InterpolateMessageTest{
		{"new_item.json", "", ""},
		{"every_occurrence.json", DefaultTemplate, everyOccurrenceMessage},
		{"high_occurrence_rate.json", DefaultTemplate, highOccurrenceMessage},
		{"new_item.json", DefaultTemplate, newItemMessage},
		{"ten_nth.json", DefaultTemplate, tenNthMessage},
	}

	for _, test := range testcases {
		rollbar := helperLoadJsonFile(t, test.filename)
		actual, err := rollbar.interpolateMessage(test.template)
		if err != nil {
			t.Errorf("Failed to interpolate message. %s", err)
		}

		if actual != test.expected {
			t.Errorf("Expected: %s\nActual: %s", test.expected, actual)
		}
	}
}