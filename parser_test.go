package stringutils

import (
	"reflect"
	"testing"
)

type testData struct {
	in   string
	want []string
}

func TestNumericValues(t *testing.T) {
	if !reflect.DeepEqual(SplitIntoNumericFields("1/12."), []string{"1", "12"}) {
		t.Error("Numeric Text Parser Error")
	}
}

func TestRegularValues(t *testing.T) {
	for _, test := range []testData{
		{"A1;A2,B1/", []string{"A1", "A2", "B1"}},
		{"Epic Records / Sony Music", []string{"Epic Records", "Sony Music"}},
		{"(Epic Records) Sony Music", []string{"Epic Records", "Sony Music"}},
		// {"2xHD - Storyville Records", []string{"2xHD", "Storyville Records"}},
	} {
		got := SplitIntoRegularFields(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Numeric text parser error: want %v, but result %v", test.want, got)
		}
	}
}

func TestNaiveStringToInt(t *testing.T) {
	if NaiveStringToInt("1977") != 1977 {
		t.Fail()
	}
	if NaiveStringToInt("A1977") != 0 {
		t.FailNow()
	}
}
