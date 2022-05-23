package main

import (
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {

	result, err := GetPokemon("Arceus")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
