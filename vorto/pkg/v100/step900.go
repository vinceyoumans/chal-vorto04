package v100

import (
	"fmt"

	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
)

//. this is a complete cheat.
//. I am not making real computations.

func StX900(RTS400 []strucs.Route400S) []strucs.RTS900 {

	var rts900S []strucs.RTS900
	var rts900 strucs.RTS900

	fmt.Println("RTS400------------------")
	fmt.Println(RTS400)
	fmt.Println("RTS400------------------")

	counter := 1
	Route := 1
	// Leg := 1

	for _, rt400 := range RTS400 {
		if counter > 5 {
			//RTD - Return To Depot
			rts900S = append(rts900S, rts900)
			counter = 1
			Route++
			rts900 = strucs.RTS900{}
		} else {
			rts900.RouteID = Route
			rts900.LoadIDS = append(rts900.LoadIDS, rt400.RouteID)
		}

	}

	fmt.Println(rts900S)

	return rts900S

}
