package v100

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vinceyoumans/vorto04/vorto/pkg/strucs"
)

func StepV100(ProblemPathFile string) {

	slog := log.New(os.Stdout, "INFO: ", log.Ltime|log.Lshortfile)

	slog.Println("v100")

	// Step A100 - Consume ProblemFile
	// Seed Everything with an array 0
	// Theory is that Array Item 1 is Load also route 1
	var probStrucS []strucs.Problem100
	depotLocation := strucs.Problem100{
		LoadNumber: 0,
		Pickup: strucs.LatLong{
			Latitude:  0,
			Longitude: 0,
		},
		DropOff: strucs.LatLong{
			Latitude:  0,
			Longitude: 0,
		},
	}

	file, err := os.Open(ProblemPathFile)
	if err != nil {
		slog.Println("V100 - Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Skip the header line
	if scanner.Scan() {
		_ = scanner.Text()
		// header := scanner.Text()
		// fmt.Println("Header:", header)
		// slogpkg.LogVortoP100(header)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		parts[1] = strings.Trim(parts[1], "()")
		parts[2] = strings.Trim(parts[2], "()")

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
			LoadNumber: loadNumber,
			Pickup:     pickup,
			DropOff:    dropoff,
			BPickedUp:  false,
		}

		probStrucS = append(probStrucS, probStruc)

	}

	return probStrucS
}
