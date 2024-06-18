package strucs

type Problems100S struct {
	Problems []Problem100
}

// --------------------------------
type Problem100 struct {
	LoadNumber int
	Pickup     LatLong
	DropOff    LatLong
	BPickedUp  bool // If a Load has been picked up in previous Route
}
type LatLong struct {
	Latitude  float64
	Longitude float64
}
