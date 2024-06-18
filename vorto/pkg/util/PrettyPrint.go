package util

import (
	"encoding/json"
	"fmt"
	"os"
)

func PPStruct(SavePath, JSONFileName string, probStruc interface{}) {
	jsonOutput, err := json.MarshalIndent(probStruc, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	outputDir := SavePath
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
			return
		}
	}

	fileJSON, err := os.Create(SavePath + "/" + JSONFileName)
	if err != nil {
		fmt.Println(err)
		// return
	}
	defer fileJSON.Close()

	_, err = fileJSON.Write(jsonOutput)
	if err != nil {
		fmt.Println(err)
		// return nill
	}

}
