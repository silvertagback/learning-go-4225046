package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")
	states := make(map[string]string)
	states["NY"] = "New York"
	states["CA"] = "California"
	states["TX"] = "Texas"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println("CA stands for", california)

	fmt.Println("Number of states:", len(states))

	delete(states, "TX")
	fmt.Println("States after deletion:", states)

	_, exists := states["TX"]
	fmt.Println("Is TX present?", exists)

	for abbr, name := range states {
		fmt.Printf("%s stands for %s\n", abbr, name)
	}

	keys := make([]string, 0, len(states))
	for abbr := range states {
		keys = append(keys, abbr)
	}
	sort.Strings(keys)
	fmt.Println("Sorted state abbreviations:", keys)

	for index := range keys {
		fmt.Println(states[keys[index]])
	}
}
