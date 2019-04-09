package main

import (
	"flag"
	"fmt"
	"log"
)

var args1 = flag.String("c1", "", "The list of coordinates divived by comma")
var args2 = flag.String("c2", "", "The list of coordinates divived by comma")

// var operation = flag

func main() {
	flag.Parse()

	coordinates1, err := splitFlagCoordinates(*args1)
	if err != nil {
		log.Fatalln(err)
	}

	v1 := Vector{Coordinates: coordinates1}
	fmt.Println(v1.Str())

	coordinates2, err := splitFlagCoordinates(*args2)
	if err != nil {
		log.Fatalln(err)
	}

	v2 := Vector{Coordinates: coordinates2}
	fmt.Println(v2.Str())
}
