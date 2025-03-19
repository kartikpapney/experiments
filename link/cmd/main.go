package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kartikpapney/experiments/link"
)

func main() {
	filename := flag.String("file", "ex1.html", "the HTML file to parse for links")
	flag.Parse()

	s, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(s)
	if err != nil {
		panic(err)
	}

	for _, l := range links {
		fmt.Printf("%+v, %+v\n", l.Href, l.Text)
	}
}
