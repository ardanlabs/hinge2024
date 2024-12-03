package name

import (
	"fmt"
	"regexp"
)

type Name struct {
	value string
}

func (n Name) String() string {
	return n.value
}

func (n Name) Equal(n2 Name) bool {
	return n.value == n2.value
}

func (n Name) MarshalText() ([]byte, error) {
	return []byte(n.value), nil
}

// =============================================================================

var nameRegEx = regexp.MustCompile("^[a-zA-Z0-9' -]{3,20}$")

func Parse(value string) (Name, error) {
	if !nameRegEx.MatchString(value) {
		return Name{}, fmt.Errorf("invalid name %q", value)
	}

	return Name{value}, nil
}

func MustParse(value string) Name {
	name, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return name
}

// =============================================================================

type Null struct {
	value string
	valid bool
}

func (n Null) String() string {
	if !n.valid {
		return "NULL"
	}

	return n.value
}

func (n Null) Valid() bool {
	return n.valid
}

func (n Null) Equal(n2 Null) bool {
	return n.value == n2.value
}

func ParseNull(value string) (Null, error) {
	if value == "" {
		return Null{}, nil
	}

	if !nameRegEx.MatchString(value) {
		return Null{}, fmt.Errorf("invalid name %q", value)
	}

	return Null{value, true}, nil
}

func MustParseNull(value string) Null {
	name, err := ParseNull(value)
	if err != nil {
		panic(err)
	}

	return name
}
