package v100

import (
	"fmt"

	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
)

func StepV400(pm300S []strucs.PMap300_ROUTE) []strucs.Route400S {
	const max_Driving_Time = 12 * 60
	// const max_Driving_Time = 12 * 60 * 60
	var RouteS []strucs.Route400S
	var Route strucs.Route400S
	var RR strucs.Route400R

	// var R400R []strucs.Route400R
	RouteID := 1
	CurrentRouteRunningTime := float64(0.0)
	CurrentRouteRunningTimePlusP1PDepot := float64(0.0)
	leg := 1

	for i, v := range pm300S {
		fmt.Println("--- i", i)
		fmt.Println("--- v", v)
		CurrentRouteRunningTimePlusP1PDepot = CurrentRouteRunningTime + v.Distance_P1P2 + v.Distance_P2P3 + v.Next_P1PDepot
		if CurrentRouteRunningTimePlusP1PDepot > max_Driving_Time {
			// Return to Depot
			// The next Load would violate 12HR
			fmt.Println("12HR --- i", i)
			fmt.Println("--- v", v)
			fmt.Println("CurrentRouteRunningTime", CurrentRouteRunningTime)

			Route.RouteID = RouteID
			Route.RunningDistance = CurrentRouteRunningTime
			RouteS = append(RouteS, Route)

			CurrentRouteRunningTime = float64(0.0)
			RouteID++

		} else {
			fmt.Println("--- i", i)
			fmt.Println("--- v", v)
			fmt.Println("CurrentRouteRunningTime", CurrentRouteRunningTime)
			fmt.Println("CurrentRouteRunningTimePlusP1PDepot", CurrentRouteRunningTimePlusP1PDepot)
			RR.I = i
			RR.Leg = leg
			RR.LoadID = v.Current_LoadID
			RR.Loc_P1 = v.P1_StartLoaction
			RR.Loc_P2 = v.P2_PickUpLatLong
			RR.Loc_P3 = v.P3_DropOffLatLong
			// RR.RunningTime = CurrentRouteRunningTimePlusP1PDepot

			Route.Route = append(Route.Route, RR)
			leg++
		}

	}

	return RouteS
}

// func StepV400B(pm300S []strucs.PMap300_ROUTE) []strucs.Route400S {

// 	// vm500.MaxDriveTime_Min = 12 * 60 * 60 // 12 hours * 60 Minutes * 60 seconds
// 	const max_Driving_Time = 12 * 60
// 	// const max_Driving_Time = 12 * 60 * 60

// 	var R strucs.Route400S
// 	var RS []strucs.Route400S
// 	var route400 strucs.Route400
// 	var route400S []strucs.Route400
// 	// AllRoutesRunningDistance := float64(0.0)
// 	CurrentRouteRunningDistance := float64(0.0)
// 	RouteID := 1
// 	LegID := 1
// 	numberOfpm300S := len(pm300S)

// 	// 1st Loop
// 	// for _, p300 := range pm300S {
// 	for i := 0; i <= numberOfpm300S; i++ {
// 		fmt.Println("i = ", i)
// 		if i == numberOfpm300S {
// 			fmt.Println("i = ", i)
// 			break
// 		}

// 		p300 := pm300S[i]

// 		// 2nd Loop
// 		// If the Next Route violate 12HR, then return NOW
// 		v := (CurrentRouteRunningDistance + p300.Dist_P1P2P3PDepot)
// 		if v > max_Driving_Time {
// 			// RTD - Return To Depot
// 			// Next Load would violate 12HR
// 			fmt.Println("=====.  CurrentRouteRunningDistance ", CurrentRouteRunningDistance)
// 			R.RouteID = RouteID
// 			R.RunningDistance = CurrentRouteRunningDistance
// 			R.Route = route400S
// 			RS = append(RS, R)

// 			CurrentRouteRunningDistance = float64(0.0)
// 			RouteID += 1
// 		}
// 		route400.I = i
// 		route400.LoadID = p300.Current_LoadID
// 		route400.Leg = LegID
// 		route400.LoadID = p300.Current_LoadID
// 		route400.Loc_P1 = p300.P1_StartLoaction
// 		route400.Loc_P2 = p300.P2_PickUpLatLong
// 		route400.Loc_P3 = p300.P3_DropOffLatLong // 1st Delivery
// 		// route400.Loc_P4 = p300.	// 2nd Delivery
// 		// route400.Loc_P5 = p300.	// 3rd Delivery
// 		route400.Distance_P1toP3 = p300.Distance_P1P2 + p300.Distance_P2P3
// 		route400.Distance_P1toP4 = p300.Distance_P1P2 + p300.Distance_P2P3 + p300.Distance_P3PDepot
// 		route400.Distance_P1toDepot = p300.Distance_P1P2 + strucs.DistanceTo(route400.Loc_P3, strucs.Depot())

// 		route400S = append(route400S, route400)
// 		CurrentRouteRunningDistance = v
// 		LegID++
// 	}

// 	return RS
// }
