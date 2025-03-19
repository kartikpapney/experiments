package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return getAllLinks(doc), nil
}

func getAllLinks(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				txt := extractText(n)
				links = append(links, Link{a.Val, txt})
			}
		}
	}

	// loop through all children nodes and recursively call getAllLinks
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		exLinks := getAllLinks(c)
		links = append(links, exLinks...)
	}
	return links
}

func extractText(n *html.Node) string {
	var text string
	if n.Type != html.ElementNode && n.Data != "a" && n.Type != html.CommentNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return strings.Trim(text, "\n")
}
