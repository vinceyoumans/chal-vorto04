package v100

import (
	"os"

	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/utils"
)

func StepV200(prob100S strucs.Problems100S) ([]strucs.LOADS, []strucs.LOADS, []strucs.LOADS) {
	aLoadsByID := prob100S.GetSortOfRemainingLoadsByID()
	aLoadsByPickUpSort := prob100S.GetSortOfRemainingLoadsByClosestPickup()
	aLoadsByDropOffSort := prob100S.GetSortOfRemainingLoadsByClosestDropOff()

	// aaLoadsByID := []strucs.LOADS(aLoadsByID)
	var aaLoadsByID []strucs.LOADS
	aaLoadsByID = append(aaLoadsByID, aLoadsByID...)

	// aaLoadsByPickUpSort := []strucs.LOADS(aLoadsByPickUpSort)
	var aaLoadsByPickUpSort []strucs.LOADS
	aaLoadsByPickUpSort = append(aaLoadsByPickUpSort, aLoadsByPickUpSort...)

	// aaLoadsByDropOffSort := []strucs.LOADS(aLoadsByDropOffSort)
	var aaLoadsByDropOffSort []strucs.LOADS
	aaLoadsByDropOffSort = append(aaLoadsByDropOffSort, aLoadsByDropOffSort...)

	return aaLoadsByID, aaLoadsByPickUpSort, aaLoadsByDropOffSort
}

func StepV230(P200 []strucs.LOADS, P100S strucs.Problems100S) []strucs.RemainingLoads230 {

	var tempRL230 strucs.RemainingLoads230
	var tempRL230S []strucs.RemainingLoads230

	for i := 0; i < len(P200); i++ {
		tempRL230.LoadID = P200[i].LoadID
		tempRL230.RouteID = 0
		tempRL230.RouteLeg = 0
		tempP100 := P100S.Problems[tempRL230.LoadID]
		tempRL230.PickUp = tempP100
		tempRL230S = append(tempRL230S, tempRL230)
	}
	return tempRL230S
}

func StepV230_Print_PNG_PDF_ByLoad(ls []strucs.RemainingLoads230, GraphDir, FileName string) {
	var tLoads []strucs.S230Load
	var tload strucs.S230Load
	for _, v := range ls {
		tload.LoadNumber = v.LoadID
		tload.Pickup = v.PickUp.Pickup
		tload.DropOff = v.PickUp.DropOff
		tLoads = append(tLoads, tload)
	}

	filePDF, err := os.Create(GraphDir + "/" + FileName + "pdf")
	filePNG, err := os.Create(GraphDir + "/" + FileName + "png")

	if err != nil {
		// slog.Println(err)
		// return
	}
	defer filePDF.Close()
	defer filePNG.Close()

	// utils.Print_PNG_PDF_ByLoad(tLoads, "../output/graph/v230", "v230")
	utils.Print_PNG_PDF_ByLoad(tLoads, GraphDir, FileName)

}
