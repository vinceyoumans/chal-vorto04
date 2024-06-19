package v100

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
	util "github.com/vinceyoumans/chal-vorto04/vorto/pkg/utils"
)

func V100Start(ProblemPathFile string, bSaveOutput bool) {

	// These graphs work
	// util.TestPDF()
	// util.TestPDF200()

	P100S := StV100(ProblemPathFile)
	if bSaveOutput {
		util.PPStruct("../output/v100/p100", "p100s.json", P100S)
	}

	// fmt.Println("=============. P200")
	//. I see the Flaw in logic.
	// Flaw:  computing each Distance from DEPOT.
	// Solution:  Compute the Distance from Previos P1.
	// get sorted list of P200
	P200A_ByID, P200B_ByPickUp, P200C_ByDropOff := SStV200(P100S)
	if bSaveOutput {
		util.PPStruct("../output/v200/p200", "pP200A_ByID.json", P200A_ByID)
	}
	if bSaveOutput {
		util.PPStruct("../output/v200/p200", "P200B_ByPickUp.json", P200B_ByPickUp)
	}
	if bSaveOutput {
		util.PPStruct("../output/v200/p200", "P200C_ByDropOff.json", P200C_ByDropOff)
	}

	PP230A_ByID := SStV230(P200A_ByID, P100S)
	PP230B_ByPickUp := SStV230(P200B_ByPickUp, P100S)
	PP230C_ByDropOff := SStV230(P200C_ByDropOff, P100S)

	if bSaveOutput {
		util.PPStruct("../output/v230/p230", "PP230A_ByID.json", PP230A_ByID)
	}
	if bSaveOutput {
		util.PPStruct("../output/v230/p230", "PP230B_ByPickUp.json", PP230B_ByPickUp)
	}
	if bSaveOutput {
		util.PPStruct("../output/v230/p230", "PP230C_ByDropOff.json", PP230C_ByDropOff)
	}

	// Create Graphic
	gdir := "../output/graphic/v230"
	gfile := "v230"
	SStV230_Print_PNG_PDF_ByLoad(PP230A_ByID, gdir, gfile)

	PM300 := StV300(PP230B_ByPickUp, P100S.Problems)
	if bSaveOutput {
		util.PPStruct("../output/v300/p300", "PP300_ByPickUp.json", PM300)
	}

	PM400 := StV400(PM300)
	if bSaveOutput {
		util.PPStruct("../output/v400/p400", "PP400_ByPickUp.json", PM400)
	}

	PM500 := StV500(PM400)

	if bSaveOutput {
		util.PPStruct("../output/v500/p500", "PP500_ByPickUp.json", PM500)
	}

	// fmt.Println("============== StXX900")

	// PM900XX := StX900XX(PP230B_ByPickUp)

	// PM900X := Stxxx900(PP230B_ByPickUp)
	PM900X := Stxxx900(PP230B_ByPickUp)
	if bSaveOutput {
		util.PPStruct("../output/v900/p900XX", "PP900XX_random.json", PM900X)
	}

	// print out return
	RR900(PM900X)

}

func Graph900(PM900X strucs.RTS900SS, P200A_ByID []strucs.LOADS) {
	panic("unimplemented")
}

func RR900(PM900X strucs.RTS900SS) {
	RR := PM900X.Routes

	for _, v := range RR {

		strSlice := make([]string, len(v.LoadIDS))
		for i, num := range v.LoadIDS {
			strSlice[i] = strconv.Itoa(num)
		}

		// Join the slice of strings with commas
		result := "[" + strings.Join(strSlice, ",") + "]"

		// Print the result
		fmt.Println(result)

	}

}
