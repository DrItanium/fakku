//queries the forums and prints out the API changes post
package main

import (
	"fmt"
	"github.com/DrItanium/fakku-toolchain"
)

func main() {
	apiPosts, err := fakku.GetForumPosts("fakku-developers", "fakku-api-changes")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", apiPosts.Topic.Title)
		for i := 0; i < len(apiPosts.Posts); i++ {
			curr := apiPosts.Posts[i]
			fmt.Printf("\t%d:\t%s\n", i, curr.Text)
		}
	}
}
