package utils

import (
	"encoding/json"
	"log"
	"os"
)

func PPStruct(SavePath, JSONFileName string, probStruc interface{}) {
	slog := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	jsonOutput, err := json.MarshalIndent(probStruc, "", "  ")
	if err != nil {
		slog.Println("Error marshalling JSON:", err)
	}

	outputDir := SavePath
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			slog.Println("Error creating output directory:", err)
			return
		}
	}

	fileJSON, err := os.Create(SavePath + "/" + JSONFileName)
	if err != nil {
		slog.Println(err)
		// return
	}
	defer fileJSON.Close()

	_, err = fileJSON.Write(jsonOutput)
	if err != nil {
		slog.Println(err)
		// return nill
	}

}
