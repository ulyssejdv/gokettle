package main

import (
	"fmt"
	"os"

	gokettle "github.com/ulyssejdv/gokettle/internal"
)

func main() {
	fmt.Println("Thank's to use the kettle volume calculator")

	if len(os.Args) != 4 {
		fmt.Println("Usage : gokettle <kettle height> <kettle diameter> <empty height>")
		os.Exit(1)
	}

	kettle := gokettle.Kettle{
		TotalHeight: os.Args[1],
		Diameter:    os.Args[2],
		EmptyHeight: os.Args[3],
	}

	result := kettle.FilledVolume()
	resultString := fmt.Sprintf("%f", result)

	fmt.Println("Volume in the kettle : " + resultString)
}
