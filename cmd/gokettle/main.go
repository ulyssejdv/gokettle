package main

import (
	"fmt"
	"os"
	"strconv"

	gokettle "github.com/ulyssejdv/gokettle/pkg"
)

func main() {
	fmt.Println("Thank's to use the kettle volume calculator")

	if len(os.Args) != 4 {
		fmt.Println("Usage : gokettle <kettle height> <kettle diameter> <empty height>")
		os.Exit(1)
	}

	kettleHeight, errArg1 := strconv.ParseFloat(os.Args[1], 64)
	kettleDiameter, errArg2 := strconv.ParseFloat(os.Args[2], 64)
	emptyHeight, errArg3 := strconv.ParseFloat(os.Args[3], 64)

	if errArg1 != nil || errArg2 != nil || errArg3 != nil {
		fmt.Println("Unable to parse args as float")
		os.Exit(1)
	}

	kettle := gokettle.Kettle{Height: kettleHeight, Diameter: kettleDiameter}

	result := kettle.ContentsVolumeByEmptyHeight(emptyHeight)

	resultString := fmt.Sprintf("%f", result)

	fmt.Println("Volume in the kettle : " + resultString)
}
