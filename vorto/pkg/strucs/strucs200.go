package strucs

//	LoadID   int 	= delivered from Problem file
//	RouteID int  	= computed as the RouteID
//	RouteLeg int 	= The Leg of the Route
//	PickUp   []Problem100

type RemainingLoads230 struct {
	LoadID   int
	RouteID  int // 0
	RouteLeg int // 0
	PickUp   Problem100
}

// // Loads Sorted by PickUp
// type LOADByPickUpS230 struct {
// 	LoadID   int
// 	RouteID  int // 0
// 	RouteLeg int // 0
// 	PickUp   []Problem100
// }

// // Loads Sorted by DropOff
// type LOADByDropOffS230 struct {
// 	LoadID   int
// 	RouteID  int // 0
// 	RouteLeg int // 0
// 	PickUp   []Problem100
// }

type S230Load struct {
	LoadNumber int
	Pickup     LatLong
	DropOff    LatLong
	BPickedUp  bool
}
