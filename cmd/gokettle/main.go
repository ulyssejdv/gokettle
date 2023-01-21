package main

import (
	"fmt"
	"os"

	gokettle "github.com/ulyssejdv/gokettle/internal"
)

func main() {
	fmt.Println("Thank's to use the kettle volume calculator")

	if len(os.Args) != 4 {
		fmt.Println("Usage : gokettle <height> <diameter> <empty height>")
		os.Exit(1)
	}

	run(os.Args[1], os.Args[2], os.Args[3])
}

func run(totalHeight string, diamter string, emptyHeight string) {
	kettle := gokettle.Kettle{
		TotalHeight: totalHeight,
		Diameter:    diamter,
		EmptyHeight: emptyHeight,
	}
	filledVolume := kettle.FilledVolume()
	fmt.Printf("Volume in the kettle : %.2f\n", filledVolume)
}
