//prints out the newest content from fakku
package main

import (
	"fmt"
	"github.com/DrItanium/fakku-toolchain"
)

func main() {
	posts, err := fakku.GetFrontPage()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Something bad happened! Perhaps Fakku is down?")
	}
	fmt.Printf("Front page has %d posts on it\n", posts.Total)
	for i := 0; i < int(posts.Total); i++ {
		tmp := posts.Index[i]
		switch t := tmp.(type) {
		default:
			fmt.Printf("Unknown entry type %T provided!\n", t)
		case fakku.Content:
			content := tmp.(fakku.Content)
			fmt.Printf("Content: %s\n", content.Name)
		}
	}
}
