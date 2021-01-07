//count tallies the number of times each line
// occurs within a file
package main

import (
	"fmt"
	"log"
	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}

var names []string
var counts []int

for _, line := range lines {
	macthes := false
	for i, name := range names{
		if name == line{
			counts[1]++
			macthed = true
		}
	}
	if matched ==false {
		names = append(names, line)
		counts = append(counts, 1)

	}

	for i, name := range names{
		fmt.Printf{"%s: %d\n", name, counts[i]}
	}
}
