package core

import (
	"io"
	"net/url"
)

// URLMap contains data about a URL that has been found
type URLMap struct {
	URL         *url.URL
	Body        io.ReadCloser
	ContentType string
	Links       []*url.URL
}
