package parsers

import (
	"net/url"

	"github.com/jecrocker/crawl/core"
)

// PageProvider provides the URLs from a page type
type PageProvider interface {
	Parse(core.URLMap) (core.URLMap, []*url.URL)
}
