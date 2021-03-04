package version_test

import (
	"testing"

	"github.com/mileusna/version"
)

func TestParse(t *testing.T) {

	v1 := version.Parse("1.2.3")
	if v1.Major != 1 {
		t.Error("Major not correct")
	}
}

type VerDemo struct {
	ParseString   string
	ExpectError   bool
	VersionString string
}

func TestVersion(t *testing.T) {
	data := []VerDemo{
		{ParseString: "v2", ExpectError: false, VersionString: "2.0.0"},
		{ParseString: "", ExpectError: true},
		{ParseString: "z", ExpectError: true},
		{ParseString: "3", ExpectError: false, VersionString: "3.0.0"},
		{ParseString: "v2.1", ExpectError: false, VersionString: "2.1.0"},
		{ParseString: "chrome-3.10.2", ExpectError: false, VersionString: "3.10.2"},
		{ParseString: "asdkasj", ExpectError: true},
		{ParseString: "2.asdas.233", ExpectError: true},
		{ParseString: "ios-2.sdd.2", ExpectError: true},
		{ParseString: "5.0.5a", ExpectError: false, VersionString: "5.0.5"},
	}

	for _, d := range data {
		var v version.Version
		var err error
		v = version.Parse(d.ParseString)
		if d.ExpectError {
			if v.String(true) != "0.0.0" {
				t.Error("Error expected but not found " + d.ParseString)
			}
			continue
		}
		if err != nil {
			t.Error(err, d.ParseString)
			continue
		}

		if v.String(true) != d.VersionString {
			t.Error("Versions don't match", v.String(true), d.VersionString)
		}
	}
}

func TestCompare(t *testing.T) {
	v := version.Parse("2.0.5")

	x := version.Parse("2.0.5")
	if !v.Equal(x) {
		t.Error("Should be equal")
	}
	if !v.EqualOrHigherThan(x) {
		t.Error("Should be equal or higher")
	}

	x = version.Parse("2.0.4")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrHigherThan(x) {
		t.Error("Should be equal or higher #2")
	}

	x = version.Parse("2.1.0")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrLowerThan(x) {
		t.Error("Should be equal or lower #2")
	}

	x = version.Parse("2.0.6")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrLowerThan(x) {
		t.Error("Should be equal or lower #2")
	}

	x = version.Parse("2.0.10")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrLowerThan(x) {
		t.Error("Should be equal or lower #2")
	}

	x = version.Parse("3.0.0")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrLowerThan(x) {
		t.Error("Should be equal or lower #2")
	}

	// new ver
	v = version.Parse("2.1.5")

	x = version.Parse("2.0.10")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrHigherThan(x) {
		t.Error("Should be equal or higher #2")
	}

	x = version.Parse("2.1.3")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrHigherThan(x) {
		t.Error("Should be equal or higher #2")
	}

	x = version.Parse("1.8.8")
	if v.Equal(x) {
		t.Error("Should not be equal")
	}
	if !v.EqualOrHigherThan(x) {
		t.Error("Should be equal or higher #2")
	}
	if !v.HigherThan(x) {
		t.Error("Should higher than")
	}

	x = version.Parse("2.1.5")
	if v.HigherThan(x) {
		t.Error("No higher")
	}
	if v.LowerThan(x) {
		t.Error("No lower")
	}

	x = version.Parse("2.1.6")
	if v.HigherThan(x) {
		t.Error("No higher")
	}
	if !v.LowerThan(x) {
		t.Error("Should be lower")
	}

	x = version.Parse("2.1.3")
	if !v.HigherThan(x) {
		t.Error("Should be higher")
	}
	if v.LowerThan(x) {
		t.Error("No lower")
	}

	x = version.Parse("2.0.6")
	if !v.HigherThan(x) {
		t.Error("Should be higher")
	}
	if v.LowerThan(x) {
		t.Error("No lower")
	}

	x = version.Parse("2.2.3")
	if v.HigherThan(x) {
		t.Error("No higher")
	}
	if !v.LowerThan(x) {
		t.Error("Should be lower")
	}
}
