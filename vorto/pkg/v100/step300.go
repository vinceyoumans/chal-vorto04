package v100

import (
	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
)

func StV300(SRL230 []strucs.RemainingLoads230, pm100 []strucs.Problem100) []strucs.PMap300_ROUTE {

	var PMap300_ROUTES []strucs.PMap300_ROUTE
	var PMap300_ROUTE strucs.PMap300_ROUTE

	p1 := strucs.Depot()

	for i := 0; i <= len(SRL230)-1; i++ {

		PMap300_ROUTE.RouteID = i
		PMap300_ROUTE.Current_LoadID = pm100[SRL230[i].LoadID].LoadID
		PMap300_ROUTE.P1_StartLoaction = p1
		PMap300_ROUTE.P2_PickUpLatLong = pm100[SRL230[i].LoadID].Pickup
		PMap300_ROUTE.P3_DropOffLatLong = pm100[SRL230[i].LoadID].DropOff

		PMap300_ROUTE.Distance_P1P2 = strucs.DistanceTo(p1, PMap300_ROUTE.P2_PickUpLatLong)
		PMap300_ROUTE.Distance_P2P3 = strucs.DistanceTo(PMap300_ROUTE.P2_PickUpLatLong, PMap300_ROUTE.P3_DropOffLatLong)
		PMap300_ROUTE.Distance_P3PDepot = strucs.DistanceTo(PMap300_ROUTE.P3_DropOffLatLong, strucs.Depot())
		PMap300_ROUTE.Dist_P1P2P3PDepot = PMap300_ROUTE.Distance_P1P2 + PMap300_ROUTE.Distance_P2P3 + PMap300_ROUTE.Distance_P3PDepot

		PMap300_ROUTES = append(PMap300_ROUTES, PMap300_ROUTE)
		p1 = PMap300_ROUTE.P3_DropOffLatLong
	}
	LPM300S := len(PMap300_ROUTES)
	// fmt.Println("Len - LPM300S---", LPM300S)
	// fmt.Println("Len - PMap300_ROUTES---", len(PMap300_ROUTES))
	for i, v := range PMap300_ROUTES {
		if i == LPM300S-1 {
			continue
		}

		v.Next_LoadID = PMap300_ROUTES[i+1].Current_LoadID
		v.Next_P1PDepot = PMap300_ROUTES[i+1].Dist_P1P2P3PDepot
		// fmt.Println("i, v", i, v)
		PMap300_ROUTES[i] = v

	}

	return PMap300_ROUTES

}
