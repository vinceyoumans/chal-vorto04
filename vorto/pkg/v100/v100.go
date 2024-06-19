package v100

import (
	util "github.com/vinceyoumans/chal-vorto04/vorto/pkg/utils"
)

func V100Start(ProblemPathFile string, bSaveOutput bool) {

	// These graphs work
	// util.TestPDF()
	// util.TestPDF200()

	P100S := StepV100(ProblemPathFile)
	if bSaveOutput {
		util.PPStruct("../output/v100/p100", "p100s.json", P100S)
	}

	// fmt.Println("=============. P200")
	//. I see the Flaw in logic.
	// Flaw:  computing each Distance from DEPOT.
	// Solution:  Compute the Distance from Previos P1.
	// get sorted list of P200
	P200A_ByID, P200B_ByPickUp, P200C_ByDropOff := StepV200(P100S)
	if bSaveOutput {
		util.PPStruct("../output/v200/p200", "pP200A_ByID.json", P200A_ByID)
	}
	if bSaveOutput {
		util.PPStruct("../output/v200/p200", "P200B_ByPickUp.json", P200B_ByPickUp)
	}
	if bSaveOutput {
		util.PPStruct("../output/v200/p200", "P200C_ByDropOff.json", P200C_ByDropOff)
	}

	PP230A_ByID := StepV230(P200A_ByID, P100S)
	PP230B_ByPickUp := StepV230(P200B_ByPickUp, P100S)
	PP230C_ByDropOff := StepV230(P200C_ByDropOff, P100S)

	if bSaveOutput {
		util.PPStruct("../output/v230/p230", "PP230A_ByID.json", PP230A_ByID)
	}
	if bSaveOutput {
		util.PPStruct("../output/v230/p230", "PP230B_ByPickUp.json", PP230B_ByPickUp)
	}
	if bSaveOutput {
		util.PPStruct("../output/v230/p230", "PP230C_ByDropOff.json", PP230C_ByDropOff)
	}

	// StepV230_Print_PNG_PDF_ByLoad(PP230A_ByID)

}
