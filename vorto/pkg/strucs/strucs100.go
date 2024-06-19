package strucs

import (
	"math"
	"sort"
)

type Problems100S struct {
	Problems []Problem100
}

// --------------------------------
type Problem100 struct {
	LoadID    int
	Pickup    LatLong
	DropOff   LatLong
	BPickedUp bool // If a Load has been picked up in previous Route
}
type LatLong struct {
	Latitude  float64
	Longitude float64
}

type LOADS struct {
	LoadID             int
	Distance_P1toP2    float64
	Distance_P2toP3    float64
	Distance_P3toDepot float64
	Distance_P1toP3    float64
	Distance_P1toDepot float64
}

func (p Problems100S) GetSortOfRemainingLoadsByID() []LOADS {

	tLoad := LOADS{
		LoadID:             0,
		Distance_P1toP2:    0,
		Distance_P2toP3:    0,
		Distance_P3toDepot: 0,
	}
	PMapLoads := make(map[int]LOADS)

	P1 := Depot() // Assume starting point is Depot

	for i := 0; i <= len(p.Problems)-1; i++ {
		prob := p.Problems[i]
		// Check if Load has been picked up in previous Route
		if prob.BPickedUp {
			// skip this Load
			continue
		}

		tLoad.LoadID = i
		tLoad.Distance_P1toP2 = DistanceTo(P1, prob.Pickup)
		tLoad.Distance_P2toP3 = DistanceTo(prob.Pickup, prob.DropOff)
		tLoad.Distance_P3toDepot = DistanceTo(prob.DropOff, Depot())
		tLoad.Distance_P1toP3 = tLoad.Distance_P1toP2 + tLoad.Distance_P2toP3
		tLoad.Distance_P1toDepot = tLoad.Distance_P1toP2 + tLoad.Distance_P2toP3 + tLoad.Distance_P3toDepot
		// Add to Map
		PMapLoads[i] = tLoad
		P1 = prob.DropOff // update P1 to the next DropOff
	}

	// fmt.Println("==== len(PMapLoads)", len(PMapLoads))
	var sLoad []LOADS
	for i := 0; i <= len(PMapLoads)-1; i++ {
		sLoad = append(sLoad, PMapLoads[i])
	}
	// fmt.Println("==== len(sLoad)", len(sLoad))

	return sLoad

}

// GetSortOfRemainingLoadsByClosestPickupAndDropOff
//   - returns slice of integers of
//     the remaining loads in order of next PICKUP and DropOff
func (p Problems100S) GetSortOfRemainingLoadsByClosestPickup() []LOADS {
	Loads := p.GetSortOfRemainingLoadsByID()

	loadSlice := Loads

	// Sort the slice by Distance_P1toP2
	sort.Slice(loadSlice, func(i, j int) bool {
		return loadSlice[i].Distance_P1toP2 < loadSlice[j].Distance_P1toP2
	})

	// Print the sorted slice
	// for _, load := range loadSlice {
	// 	fmt.Printf("Load ID: %d, Distance P1 to P2: %f, Distance P2 to P3: %f, Distance P3 to Depot: %f, Distance P1 to P3: %f, Distance P1 to Depot: %f\n",
	// 		load.loadID, load.Distance_P1toP2, load.Distance_P2toP3, load.Distance_P3toDepot, load.Distance_P1toP3, load.Distance_P1toDepot)
	// }

	return loadSlice

}

// GetSortOfRemainingLoadsByClosestPickupAndDropOff
//   - returns slice of integers of
//     the remaining loads in order of next PICKUP and DropOff
func (p Problems100S) GetSortOfRemainingLoadsByClosestDropOff() []LOADS {
	Loads := p.GetSortOfRemainingLoadsByID()

	loadSlice := Loads

	// Sort the slice by Distance_P1toP2
	sort.Slice(loadSlice, func(i, j int) bool {
		return loadSlice[i].Distance_P1toP3 < loadSlice[j].Distance_P1toP3
	})

	// Print the sorted slice
	// for _, load := range loadSlice {
	// 	fmt.Printf("Load ID: %d, Distance P1 to P2: %f, Distance P2 to P3: %f, Distance P3 to Depot: %f, Distance P1 to P3: %f, Distance P1 to Depot: %f\n",
	// 		load.loadID, load.Distance_P1toP2, load.Distance_P2toP3, load.Distance_P3toDepot, load.Distance_P1toP3, load.Distance_P1toDepot)
	// }

	return loadSlice
}

// DistanceTo calculates the distance between Px and Py
func DistanceTo(p1, p2 LatLong) float64 {
	return math.Sqrt(math.Pow(p2.Latitude-p1.Latitude, 2) + math.Pow(p2.Longitude-p1.Longitude, 2))
}
