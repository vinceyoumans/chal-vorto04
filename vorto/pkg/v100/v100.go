package v100

import (
	"github.com/vinceyoumans/chal-vorto04/pkg/util"
	v100 "github.com/vinceyoumans/chal-vorto04/vorto/pkg/V100"
)

func V100Start(ProblemPathFile string, bSaveOutput bool) {

	P100S := v100.StepV100(ProblemPathFile)
	if bSaveOutput {
		util.PPStruct("../output/v100/p100", "p100s.json", P100S)
	}

	return

}
