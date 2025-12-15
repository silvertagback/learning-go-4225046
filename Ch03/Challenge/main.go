package main

import "fmt"


func main() {
	colorNames := []string{
		"Red",
		"Green",
		"Blue",
	}
	hexValues := []int{
		0xFF0000,
		0x00FF00,
		0x000FF,
	}

	fmt.Println("Color Names Slice:", colorNames)
	fmt.Println("Hex Values Slice:", hexValues)

	outSlice := slicesToObjects(colorNames, hexValues)
	fmt.Println("Color Objects Slice:", outSlice)
}

func slicesToObjects(colorNames []string, hexValues []int) []Color {
    var colorObjects []Color
    for i, name := range colorNames {
        hex := hexValues[i]
        color := Color{
            Name: name,
            Hex:  fmt.Sprintf("%06X", hex),
        }
        colorObjects = append(colorObjects, color)
    }
    return colorObjects
}

type Color struct {
	Name string
	Hex  string
}
