//prints out the fakku front page poll
package main

import (
	"fmt"
	"github.com/DrItanium/fakku-toolchain"
)

func main() {
	poll, err := fakku.GetFrontPagePoll()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Something bad happened! Perhaps Fakku is down?")
	}
	fmt.Printf("Question: %s\n", poll.Question)
	fmt.Printf("URL: %s\n", poll.Url)
	for i := 0; i < len(poll.Options); i++ {
		currOption := poll.Options[i]
		fmt.Printf("\t- \"%s\"\t %d votes\n", currOption.Text, currOption.Votes)
	}
}
