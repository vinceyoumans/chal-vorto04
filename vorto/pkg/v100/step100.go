package v100

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
)

func StV100(ProblemPathFile string) strucs.Problems100S {
	// Define the log file path
	logFilePath := "../output/slog/V100.log"

	// Ensure the directory exists
	logDir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	// Open a file for writing log messages
	logFile, err := os.OpenFile("../output/slog/V100.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Create a new logger that writes to the file
	slog := log.New(logFile, "INFO: ", log.Ltime|log.Lshortfile)

	// Write a log message
	slog.Println("v100 - step100")

	// Step A100 - Consume ProblemFile
	// Seed Everything with an array 0
	// Theory is that Array Item 1 is Load also route 1
	var probStrucS []strucs.Problem100
	depotLocation := strucs.Problem100{
		LoadID: 0,
		Pickup: strucs.LatLong{
			Latitude:  0,
			Longitude: 0,
		},
		DropOff: strucs.LatLong{
			Latitude:  0,
			Longitude: 0,
		},
	}
	probStrucS = append(probStrucS, depotLocation)

	file, err := os.Open(ProblemPathFile)
	if err != nil {
		slog.Println("V100 - Error opening file:", err)
		// return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Skip the header line
	if scanner.Scan() {
		header := scanner.Text()
		slog.Println("V100 - Header:", header)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		// parts[1] = strings.Trim(parts[1], "()")
		// parts[2] = strings.Trim(parts[2], "()")

		if len(parts) != 3 {
			slog.Println("*****  Invalid line format:", line)
			continue
		}

		loadNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			slog.Println("Error parsing load number:", err)
			continue
		}
		pickup := parseLatLong(parts[1])
		dropoff := parseLatLong(parts[2])

		probStruc := strucs.Problem100{
			LoadID:    loadNumber,
			Pickup:    pickup,
			DropOff:   dropoff,
			BPickedUp: false,
		}

		probStrucS = append(probStrucS, probStruc)

	}
	var p strucs.Problems100S
	p.Problems = probStrucS

	return p
}

func parseLatLong(str string) strucs.LatLong {
	// Open a file for writing log messages
	logFile, err := os.OpenFile("../output/slog/V100B.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Create a new logger that writes to the file
	slog := log.New(logFile, "INFO: ", log.Ltime|log.Lshortfile)

	// Write a log message
	slog.Println("v100 - step100- parseLatLong")

	str = strings.Trim(str, "()")
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		slog.Println("Invalid LatLong format:", str)
		return strucs.LatLong{}
	}

	latitude, err1 := strconv.ParseFloat(parts[0], 64)
	longitude, err2 := strconv.ParseFloat(parts[1], 64)
	if err1 != nil || err2 != nil {
		slog.Println("Error parsing LatLong:", err1, err2)
		return strucs.LatLong{}
	}

	return strucs.LatLong{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
