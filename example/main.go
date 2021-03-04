package main

import (
	"fmt"
	"log"

	"github.com/mileusna/version"
)

func main() {

	// parse version string and show its parts
	v := version.Parse("5.2.10")
	fmt.Println("Major release:", v.Major)
	fmt.Println("Minor release:", v.Minor)
	fmt.Println("Patch:", v.Patch)
	fmt.Println(v.String())      // 5.2.10
	fmt.Println(v.ShortString()) // 5.2

	// compare to second version
	v2 := version.Parse("5.2.1")
	if v.EqualOrHigherThan(v2) {
		fmt.Println("You have the lateste release!")
	}

	// compare directy to other version String
	if v.HigherThanString("5.2") {
		fmt.Println("You have the lateste release!")
	}

	// no errors are returned
	// if unable to parse, it will return empty struct (version 0.0.0)
	v3 := version.Parse("2.skfhaskjh.10")
	if v3.String() == "0.0.0" {
		log.Println("Wrong version string")
	}

	// having some prefix and/or suffix is OK
	// they will be stored in prefix/suffix fields
	v4 := version.Parse("iOS 14.2")
	fmt.Println(v4.Prefix) // iOS

	v5 := version.Parse("version 12.4b")
	fmt.Println(v5.Suffix) // b
}
