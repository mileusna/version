package version

import (
	"fmt"
	"strconv"
	"strings"
)

// Version struct
type Version struct {
	Major  int
	Minor  int
	Patch  int
	Prefix string
	Suffix string
}

// Parse version string to struct
// For easier usage, it returns only one value, Version struct, without error
// If not parsable, empty version (0.0.0) will be returned
func Parse(s string) Version {
	var v Version
	for i, c := range s {
		if c >= '0' && c <= '9' {
			s = s[i:]
			break
		}
		v.Prefix += string(c)
	}

	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c >= '0' && c <= '9' {
			s = s[:i+1]
			break
		}
		v.Suffix = string(c) + v.Suffix
	}

	parts := strings.Split(s, ".")

	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return Version{}
		}
		switch i {
		case 0:
			v.Major = n
		case 1:
			v.Minor = n
		case 2:
			v.Patch = n
		}
	}

	if v.Prefix != "" {
		v.Prefix = strings.Trim(v.Prefix, " -.,")
	}
	if v.Suffix != "" {
		v.Suffix = strings.Trim(v.Suffix, " -.,")
	}

	return v
}

// String returns full version string with 3 components
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// ShortString returns full version string with 2 components
func (v Version) ShortString() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

// Equal returns true if versions are eaqual
func (v Version) Equal(compareTo Version) bool {
	return v.Major == compareTo.Major && v.Minor == compareTo.Minor && v.Patch == compareTo.Patch
}

// EqualString returns true if vresions are equal, accepts version string as param
func (v Version) EqualString(compareTo string) bool {
	compare := Parse(compareTo)
	return v.Equal(compare)
}

// EqualOrHigherThan returns true if version is equal or higher than compared one
func (v Version) EqualOrHigherThan(compareTo Version) bool {
	switch {
	case v.Major > compareTo.Major:
		return true
	case v.Major < compareTo.Major:
		return false
	default:
		switch {
		case v.Minor > compareTo.Minor:
			return true
		case v.Minor < compareTo.Minor:
			return false
		default:
			return v.Patch >= compareTo.Patch
		}
	}
}

// EqualOrHigherThanString returns true if version is equal or higher than compared one, accepts version as param
func (v Version) EqualOrHigherThanString(compareTo string) bool {
	compare := Parse(compareTo)
	return v.EqualOrHigherThan(compare)
}

//EqualOrLowerThan returns true if version is equal or lower than compared
func (v Version) EqualOrLowerThan(compareTo Version) bool {
	switch {
	case v.Major > compareTo.Major:
		return false
	case v.Major < compareTo.Major:
		return true
	default:
		switch {
		case v.Minor > compareTo.Minor:
			return false
		case v.Minor < compareTo.Minor:
			return true
		default:
			return v.Patch <= compareTo.Patch
		}
	}
}

//EqualOrLowerThanString returns true if version is equal or lower than compared, accepts version as param
func (v Version) EqualOrLowerThanString(compareTo string) bool {
	compare := Parse(compareTo)
	return v.EqualOrLowerThan(compare)
}

// HigherThan returns true if version is higher than compare
func (v Version) HigherThan(compareTo Version) bool {
	switch {
	case v.Major > compareTo.Major:
		return true
	case v.Major < compareTo.Major:
		return false
	default:
		switch {
		case v.Minor > compareTo.Minor:
			return true
		case v.Minor < compareTo.Minor:
			return false
		default:
			return v.Patch > compareTo.Patch
		}
	}
}

// HigherThanString returns true if version is higher than compare, accepts version as param
func (v Version) HigherThanString(compareTo string) bool {
	compare := Parse(compareTo)
	return v.HigherThan(compare)
}

// LowerThan returns true if version is lower
func (v Version) LowerThan(compareTo Version) bool {
	switch {
	case v.Major > compareTo.Major:
		return false
	case v.Major < compareTo.Major:
		return true
	default:
		switch {
		case v.Minor > compareTo.Minor:
			return false
		case v.Minor < compareTo.Minor:
			return true
		default:
			return v.Patch < compareTo.Patch
		}
	}
}

// LowerThanString returns true if version is lower, accepts version as param
func (v Version) LowerThanString(compareTo string) bool {
	compare := Parse(compareTo)
	return v.LowerThan(compare)
}
