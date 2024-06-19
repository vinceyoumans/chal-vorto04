package strucs

type Route400S struct {
	RouteID         int
	RunningDistance float64
	// TotalLegs       int
	Route []Route400R
}

type Route400R struct {
	I                  int
	Leg                int
	LoadID             int
	Loc_P1             LatLong
	Loc_P2             LatLong
	Loc_P3             LatLong
	Loc_P4             LatLong
	Distance_P1toP3    float64
	Distance_P1toP4    float64
	Distance_P1toDepot float64
}
