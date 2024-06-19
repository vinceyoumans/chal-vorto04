package v100

import (
	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
)

func Stxxx900(P200B_ByPickUp []strucs.RemainingLoads230) strucs.RTS900SS {
	// for i, v := range P200B_ByPickUp {
	// 	fmt.Println("nnn", i, v.LoadID)
	// }

	// type RemainingLoads230 struct {
	// 	LoadID   int
	// 	RouteID  int // 0
	// 	RouteLeg int // 0
	// 	PickUp   Problem100
	// }
	//	type RTS900 struct {
	//		RouteID      int
	//		LoadIDS      []int
	//		CountOfLoads int
	//	}
	var rts900S []strucs.RTS900
	var rts900 strucs.RTS900
	routeID := 1
	LegID := 1
	MaxP200 := len(P200B_ByPickUp)

	for i := 0; i < MaxP200; i++ {
		if LegID > 5 {
			// time to return to DEPOT
			rts900.CountOfLoads = len(rts900.LoadIDS)
			rts900S = append(rts900S, rts900)
			routeID++
			LegID = 1
			rts900 = strucs.RTS900{}
		}
		rts900.RouteID = routeID
		rts900.LoadIDS = append(rts900.LoadIDS, P200B_ByPickUp[i].LoadID)
		LegID++

	}

	//Count all of the Loads
	CountOfAllLoads := 0
	for _, v := range rts900S {
		CountOfAllLoads += v.CountOfLoads
	}
	var rts900SS strucs.RTS900SS
	rts900SS.TotalLoadIDs = CountOfAllLoads
	rts900SS.Routes = rts900S
	return rts900SS
}
