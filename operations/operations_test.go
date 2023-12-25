package operations

import (
	"regexp"
	"testing"
)

func TestGetAgeUsingNameRishit(t *testing.T) {
	name := "rishit"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := GetAgeUsingName(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatal("Function GetAgeUsingNameRishit() failed.")
	}
}

func TestGetGenderUsingNameRishit(t *testing.T) {
	name := "rishit"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := GetGenderUsingName(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatal("Function GetGenderUsingNameRishit() failed.")
	}
}

func TestGetAgeUsingNameEmpty(t *testing.T) {
	name := ""

	msg, err := GetAgeUsingName(name)

	if msg != "" || err.Error() != "The name entered is empty\n" {
		t.Fatal("Function GetAgeUsingNameEmpty() failed.")
	}
}

func TestGetGenderUsingNameEmpty(t *testing.T) {
	name := ""

	msg, err := GetGenderUsingName(name)

	if msg != "" || err.Error() != "The name entered is empty\n" {
		t.Fatal("Function GetGenderUsingNameEmpty() failed.")
	}
}
