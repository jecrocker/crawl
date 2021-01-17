package parsers

import (
	"net/url"

	"github.com/jecrocker/crawl/core"
	"github.com/jecrocker/crawl/utils"
	"golang.org/x/net/html"
)

// HTMLProvider is for providing an HTML Page to the system
type HTMLProvider struct {
}

// Parse parses an HTML node
func (h *HTMLProvider) Parse(node *core.URLMap, baseURL *url.URL, t utils.Tracker) (*core.URLMap, []*url.URL) {

	newURLs := []*url.URL{}
	processedNode := &core.URLMap{}

	// We need to parse the body of the HTML
	parsed, err := html.Parse(node.Body)
	if err != nil {
		// Again with preventing a deadlock
		return nil, nil
	}

	anchorTags := h.findAnchorTags(parsed)
	for _, tag := range anchorTags {
		valid, url := utils.ValidateURL(h.pullHref(tag), baseURL)
		if valid && !t.HasItem(url.String()) {
			t.AddURL(url.String())
			newURLs = append(newURLs, url)
		}
		if valid {
			node.Links = append(node.Links, url)
		}
	}
	processedNode = node

	return processedNode, newURLs
}

// findAnchorTags pulls the anchor tags from the base node
// this needs to be a private function as it's not something that
// the wider application needs to see.
func (h *HTMLProvider) findAnchorTags(baseNode *html.Node) []*html.Node {
	// Create a list of nodes
	nodes := []*html.Node{}

	// Create an anonymous function signature
	var anchorFinder func(*html.Node)

	// Define the anonymous function
	anchorFinder = func(node *html.Node) {
		// If the current search node is an anchor tag then append it to the list
		if node.Type == html.ElementNode && node.Data == "a" {
			nodes = append(nodes, node)
		}
		// recurse down the tree and to siblings
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			anchorFinder(child)
		}
	}

	// Start from the base node
	anchorFinder(baseNode)

	return nodes
}

// PullHref gets the href links from the anchor tags
func (h *HTMLProvider) pullHref(node *html.Node) string {
	// Go over all of the attributes to find href
	for _, value := range node.Attr {
		// If the attribute key is href, return the value
		if value.Key == "href" {
			return value.Val
		}
	}
	// If there is no url then return an empty string
	return ""
}
