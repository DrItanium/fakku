//prints out the newest content from fakku
package main

import (
	"fmt"
	"github.com/DrItanium/fakku-toolchain"
)

func main() {
	//TODO: add support for grabbing beyond the first page
	//TODO: add support for spitting out CLIPS style facts (template assertions
	//or instances)
	posts, err := fakku.GetFrontPage()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Something bad happened! Perhaps Fakku is down?")
	}
	for i := 0; i < int(posts.Total); i++ {
		tmp := posts.Index[i]
		switch tmp.(type) {
		case fakku.Content:
			content := tmp.(fakku.Content)
			tags := content.Tags
			fmt.Printf("%s - ", content.Name)
			// print out the tags one after another in a form that can be easily
			// grepped through
			if len(tags) == 0 {
				fmt.Printf("No tags!")
			} else {
				fmt.Printf("{ %s", tags[0].Attribute)
				for j := 1; j < len(tags); j++ {
					fmt.Printf(", %s", tags[j].Attribute)
				}
				fmt.Printf(" }")
			}
			fmt.Printf(" - %s", content.Url)
			fmt.Println()
		}
	}
}
