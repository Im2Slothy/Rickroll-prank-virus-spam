package main

import (
	"fmt"
	"time"
	"github.com/pkg/browser"
)

func main() {
	fmt.Print("You have just fallen into my trap... Sucks to be you!")
	time.Sleep(5 * time.Second)
	spam := 10

	for spam > 0 {
		browser.OpenURL("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
		spam--

	}

}