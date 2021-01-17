package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/jecrocker/crawl/core"
	"github.com/jecrocker/crawl/parsers"
	"github.com/jecrocker/crawl/utils"
)

var baseURL string
var t utils.Tracker = utils.NewTracker()

func init() {
	log.Println("Initialising Web Crawler")
	flag.StringVar(&baseURL, "baseURL", "", "The seed point for the URL")
	flag.Parse()

}

func main() {

	u, err := url.Parse(baseURL)

	if err != nil {
		log.Fatalf("Could not parse provided URL: %s, %v", baseURL, err)
		os.Exit(1)
	}

	log.Println("Starting Crawl")
	log.Printf("Starting at %v", u)

	urls := make(chan *url.URL, 10000)
	bodies := make(chan *core.URLMap, 1000)
	outputs := make(chan *core.URLMap, 1000)

	wg := &sync.WaitGroup{}

	t.AddURL(u.String())
	wg.Add(1)
	urls <- u

	go getPage(urls, bodies, wg)
	go parsePage(u, wg, urls, bodies, outputs)
	go generateOutput(outputs, wg)

	wg.Wait()
	log.Println("Done")
	close(urls)
}

// generateOutput outputs the content of the node and the links found
func generateOutput(outputs <-chan *core.URLMap, wg *sync.WaitGroup) {
	for {
		node, open := <-outputs
		if !open {
			break
		}
		fmt.Printf("- %s\n", node.URL.String())
		for _, link := range node.Links {
			fmt.Printf("|---- %s\n", link.String())
		}
		wg.Done()
	}
}

func parsePage(baseURL *url.URL, wg *sync.WaitGroup, urls chan *url.URL, processing <-chan *core.URLMap, outputs chan *core.URLMap) {
	for {
		node, open := <-processing
		if !open {
			break
		}

		// Define the new variables
		var newNode *core.URLMap
		var newURLs []*url.URL

		// Here we wish to switch based on the content type, this allows us to know
		// what type of content we're looking at
		switch {
		// If the content type contains html then parse the node as an HTML node
		case strings.Contains(node.ContentType, "text/html"):
			h := parsers.HTMLProvider{}
			newNode, newURLs = h.Parse(node, baseURL, t)
			// Add the node to be output
			wg.Add(1)
			outputs <- newNode

			for _, url := range newURLs {
				// Add the URL to the workgroup for getting
				wg.Add(1)
				urls <- url
			}
		}

		// Mark the node as parsed
		wg.Done()
	}
	close(outputs)
}

// getPage gets the body of a get request
func getPage(ch <-chan *url.URL, processing chan *core.URLMap, wg *sync.WaitGroup) {
	client := http.Client{}
	for {
		u, open := <-ch
		if !open {
			break
		}
		response, err := client.Get(u.String())
		if err != nil {
			// This enables the program to exit even if the page cannot be parsed
			wg.Done()
			continue
		}

		wg.Add(1)
		processing <- &core.URLMap{
			URL:         u,
			Body:        response.Body,
			ContentType: response.Header.Get("Content-Type"),
			Links:       []*url.URL{},
		}
		wg.Done()
	}
	close(processing)
}
