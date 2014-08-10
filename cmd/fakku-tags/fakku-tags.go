// list all known tags
package main

import (
	"flag"
	"fmt"
	"github.com/DrItanium/fakku-toolchain"
)

var verboseFlag bool

func init() {
	const usage = "print descriptions of the tags"
	flag.BoolVar(&verboseFlag, "verbose", false, usage)
	flag.BoolVar(&verboseFlag, "v", false, usage+" (shorthand)")
}
func main() {
	flag.Parse()
	tags, err := fakku.GetTags()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Something bad happened! Perhaps Fakku is down?")
	}
	var i uint
	for i = 0; i < tags.Total; i++ {
		curr := tags.Tags[i]
		if verboseFlag {
			fmt.Printf("%s - %s\n", curr.Name, curr.Description)
		} else {
			fmt.Println(curr.Name)
		}
	}
}
