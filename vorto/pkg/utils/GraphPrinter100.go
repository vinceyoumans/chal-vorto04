package utils

import (
	"fmt"
	"image/color"
	"log"

	"github.com/signintech/gopdf"
	"github.com/vinceyoumans/chal-vorto04/vorto/pkg/strucs"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// type Location struct {
// 	Latitude  float64
// 	Longitude float64
// }

// type Load struct {
// 	LoadNumber int
// 	Pickup     Location
// 	DropOff    Location
// 	BPickedUp  bool
// }

func Print_PNG_PDF_ByLoad(data []strucs.S230Load, FileDir, FileName string) {

	// Data
	// data := []S230Load{
	// 	{LoadNumber: 1, Pickup: Location{-81.31601096717336, 146.1861590972063}, DropOff: Location{-137.08683395394672, 47.813098996155006}},
	// 	{LoadNumber: 2, Pickup: Location{-161.65593927317377, -65.67481432724838}, DropOff: Location{-135.3054793021323, 42.51074152812761}},
	// 	{LoadNumber: 3, Pickup: Location{34.030733960074095, 22.813564486590252}, DropOff: Location{-28.166184315126266, 124.46635733053157}},
	// 	{LoadNumber: 4, Pickup: Location{-27.122557871110608, 5.57398329054593}, DropOff: Location{-134.03552084887804, 98.2264887616846}},
	// 	{LoadNumber: 5, Pickup: Location{-5.437650426430196, 46.79372987343775}, DropOff: Location{34.95686297226499, -77.86531858519476}},
	// }

	// Create a new plot
	p := plot.New()
	p.Title.Text = "Load Locations"
	p.X.Label.Text = "Longitude"
	p.Y.Label.Text = "Latitude"

	// Add depot point
	depot, err := plotter.NewScatter(plotter.XYs{{X: 0, Y: 0}})
	if err != nil {
		log.Fatalf("could not create depot scatter: %v", err)
	}
	depot.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	depot.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(depot)

	// Add load points
	for _, load := range data {
		pickup, err := plotter.NewScatter(plotter.XYs{{X: load.Pickup.Longitude, Y: load.Pickup.Latitude}})
		if err != nil {
			log.Fatalf("could not create pickup scatter: %v", err)
		}
		dropoff, err := plotter.NewScatter(plotter.XYs{{X: load.DropOff.Longitude, Y: load.DropOff.Latitude}})
		if err != nil {
			log.Fatalf("could not create dropoff scatter: %v", err)
		}
		pickup.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
		dropoff.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		pickup.GlyphStyle.Shape = draw.CircleGlyph{}
		dropoff.GlyphStyle.Shape = draw.CircleGlyph{}
		p.Add(pickup, dropoff)

		// Add line between pickup and dropoff
		line, err := plotter.NewLine(plotter.XYs{
			{X: load.Pickup.Longitude, Y: load.Pickup.Latitude},
			{X: load.DropOff.Longitude, Y: load.DropOff.Latitude},
		})
		if err != nil {
			log.Fatalf("could not create line: %v", err)
		}
		line.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		p.Add(line)
	}

	// Save plot to PNG
	err = p.Save(6*vg.Inch, 6*vg.Inch, "plot.png")
	if err != nil {
		log.Fatalf("could not save plot: %v", err)
	}

	// Create PDF and add the plot image
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err = pdf.Image("plot.png", 10, 10, nil)
	if err != nil {
		log.Fatalf("could not add image to PDF: %v", err)
	}

	// Save PDF
	err = pdf.WritePdf("output.pdf")
	if err != nil {
		log.Fatalf("could not save PDF: %v", err)
	}

	log.Println("PDF created successfully.")
}

func TestPDF200() {

	// Data
	data := []Load{
		{LoadNumber: 1, Pickup: Location{-81.31601096717336, 146.1861590972063}, DropOff: Location{-137.08683395394672, 47.813098996155006}},
		{LoadNumber: 2, Pickup: Location{-161.65593927317377, -65.67481432724838}, DropOff: Location{-135.3054793021323, 42.51074152812761}},
		{LoadNumber: 3, Pickup: Location{34.030733960074095, 22.813564486590252}, DropOff: Location{-28.166184315126266, 124.46635733053157}},
		{LoadNumber: 4, Pickup: Location{-27.122557871110608, 5.57398329054593}, DropOff: Location{-134.03552084887804, 98.2264887616846}},
		{LoadNumber: 5, Pickup: Location{-5.437650426430196, 46.79372987343775}, DropOff: Location{34.95686297226499, -77.86531858519476}},
	}

	// Create a new plot
	p := plot.New()
	p.Title.Text = "Load Locations"
	p.X.Label.Text = "Longitude"
	p.Y.Label.Text = "Latitude"

	// Calculate axis ranges to center the depot
	minX, maxX, minY, maxY := -180.0, 180.0, -180.0, 180.0
	for _, load := range data {
		if load.Pickup.Longitude < minX {
			minX = load.Pickup.Longitude
		}
		if load.Pickup.Longitude > maxX {
			maxX = load.Pickup.Longitude
		}
		if load.Pickup.Latitude < minY {
			minY = load.Pickup.Latitude
		}
		if load.Pickup.Latitude > maxY {
			maxY = load.Pickup.Latitude
		}
		if load.DropOff.Longitude < minX {
			minX = load.DropOff.Longitude
		}
		if load.DropOff.Longitude > maxX {
			maxX = load.DropOff.Longitude
		}
		if load.DropOff.Latitude < minY {
			minY = load.DropOff.Latitude
		}
		if load.DropOff.Latitude > maxY {
			maxY = load.DropOff.Latitude
		}
	}

	// Center the depot
	p.X.Min, p.X.Max = minX, maxX
	p.Y.Min, p.Y.Max = minY, maxY

	// Add a border to the plot
	p.Add(plotter.NewGrid())

	// Add depot point
	depot, err := plotter.NewScatter(plotter.XYs{{X: 0, Y: 0}})
	if err != nil {
		log.Fatalf("could not create depot scatter: %v", err)
	}
	depot.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	depot.GlyphStyle.Shape = draw.CircleGlyph{}
	p.Add(depot)

	// Add depot label
	depotLabel, err := plotter.NewLabels(plotter.XYLabels{
		XYs:    plotter.XYs{{X: 0, Y: 0}},
		Labels: []string{"DEPOT"},
	})
	if err != nil {
		log.Fatalf("could not create depot label: %v", err)
	}
	p.Add(depotLabel)

	// Add load points and labels
	for _, load := range data {
		pickup, err := plotter.NewScatter(plotter.XYs{{X: load.Pickup.Longitude, Y: load.Pickup.Latitude}})
		if err != nil {
			log.Fatalf("could not create pickup scatter: %v", err)
		}
		dropoff, err := plotter.NewScatter(plotter.XYs{{X: load.DropOff.Longitude, Y: load.DropOff.Latitude}})
		if err != nil {
			log.Fatalf("could not create dropoff scatter: %v", err)
		}
		pickup.GlyphStyle.Color = color.RGBA{R: 0, G: 255, B: 0, A: 255}
		dropoff.GlyphStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		pickup.GlyphStyle.Shape = draw.CircleGlyph{}
		dropoff.GlyphStyle.Shape = draw.CircleGlyph{}
		p.Add(pickup, dropoff)

		// Add line between pickup and dropoff
		line, err := plotter.NewLine(plotter.XYs{
			{X: load.Pickup.Longitude, Y: load.Pickup.Latitude},
			{X: load.DropOff.Longitude, Y: load.DropOff.Latitude},
		})
		if err != nil {
			log.Fatalf("could not create line: %v", err)
		}
		line.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
		p.Add(line)

		// Add load number label to the line
		midPoint := plotter.XY{
			X: (load.Pickup.Longitude + load.DropOff.Longitude) / 2,
			Y: (load.Pickup.Latitude + load.DropOff.Latitude) / 2,
		}
		lineLabel, err := plotter.NewLabels(plotter.XYLabels{
			XYs:    plotter.XYs{midPoint},
			Labels: []string{fmt.Sprintf("Load %d", load.LoadNumber)},
		})
		if err != nil {
			log.Fatalf("could not create line label: %v", err)
		}
		p.Add(lineLabel)
	}

	// File paths
	pngFilePath := "plot.png"
	pdfFilePath := "output.pdf"

	// Save plot to PNG
	err = p.Save(6*vg.Inch, 6*vg.Inch, pngFilePath)
	if err != nil {
		log.Fatalf("could not save plot: %v", err)
	}

	// Create PDF and add the plot image
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err = pdf.Image(pngFilePath, 10, 10, nil)
	if err != nil {
		log.Fatalf("could not add image to PDF: %v", err)
	}

	// Save PDF
	err = pdf.WritePdf(pdfFilePath)
	if err != nil {
		log.Fatalf("could not save PDF: %v", err)
	}

	log.Println("PDF created successfully.")

}
