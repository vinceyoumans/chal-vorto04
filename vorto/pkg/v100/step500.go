package v100

import (
	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
)

func StV500(S400 []strucs.Route400S) []strucs.RTS500 {

	var tRTS500 strucs.RTS500
	var tRTS500S []strucs.RTS500

	for _, v := range S400 {
		// fmt.Println("--------------------")
		// fmt.Println(i, v)
		// tRTS500. = v.I
		tRTS500.RouteID = v.RouteID
		// S500[i].Route = append(S500[i].Route, v.Routes...)
		tRTS500S = append(tRTS500S, tRTS500)
	}
	return tRTS500S
}
