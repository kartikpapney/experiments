package main

import (
	"flag"
	"net/http"

	"github.com/kartikovvy/golang/Documents/personal-files/experiments/link"
)

/*
	Steps:

	1. make a GET request to the webpage
	2. parse all the links from the webpage
	3. build urls from the links we get
	4. filter out any links with a different domain
	5. find all the pages (BFS)
	6. print out xml

*/

func main() {

	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")

	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	/*
		whenever the main function will exit the defer function will be called

		Advantages:
			easy to check memory leaks and if resources are not being closed properly
			say you returned an error in the middle of the function, the defer function will still be called unlike if you had written the code at the end of the function

	*/
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)
	for _, l := range links {
		println(l.Href)
	}
}

func bfs(url string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var queue []string = append([]string{}, url)

	for i := 0; i <= maxDepth; i++ {
		size := len(queue)
		for j := 0; j < size; j++ {
			currentUrl := queue[0]
			queue = queue[1:]

			if _, ok := seen[currentUrl]; ok {
				continue
			}
			seen[currentUrl] = struct{}{}

			children := get(currentUrl)
			queue = append(queue, children...)
		}
	}

	result := make([]string, 0, len(seen))
	for key := range seen {
		result = append(result, key)
	}
	return result
}

func get(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)
	urls := []string{}
	for _, l := range links {
		urls = append(urls, l.Href)
	}
	return urls
}
