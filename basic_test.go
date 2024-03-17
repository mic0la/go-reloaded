package reloaded

import (
	"testing"
)

// func TestCorrectAll(t *testing.T) {
// 	correct := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
// 	want := regexp.MustCompile(correct)
// 	message := CorrectAll("it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.")
// 	if !want.MatchString(message) {
// 		t.Fatalf("Expected %s, got %s", want, message)
// 	}
// }

func testCases(t *testing.T, data []struct {
	input    string
	expected string
}) {
	for _, v := range data {
		message := CorrectAll(v.input)
		if message != v.expected {
			t.Fatalf("I expect %s, but got %s", v.expected, v.input)
		}
	}
}

var testTableData = []struct {
	input    string
	expected string
}{
	{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},
	{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
	{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
	{"Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
}

func TestCorrectAll(t *testing.T) {
	testCases(t, testTableData)
}
