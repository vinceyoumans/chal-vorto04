package strucs

type PMap300_ROUTE struct {
	RouteID           int
	Current_LoadID    int
	P1_StartLoaction  LatLong
	P2_PickUpLatLong  LatLong
	P3_DropOffLatLong LatLong
	Distance_P1P2     float64
	Distance_P2P3     float64
	Distance_P3PDepot float64
	Dist_P1P2P3PDepot float64

	Next_LoadID   int
	Next_P1PDepot float64
}
