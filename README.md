## Web Crawler

This is a basic web crawler, it is designed to crawl a given website and to create a sitemap from a given website.

### Assumptions

- The URL will have a **trusted** TLS certificate or no TLS certificate at all
- The application will respect the robots.txt if it is present
- If a sitemap.xml exists then the URIs from this will be added to the search list

### Tasks

#### Required

- [ ] Print each of the URIs with a list of the links that are found
- [ ] Respect the given domain, i.e. do not recurse sub-domains or external domains

#### Some helpful features

- [ ] Cache results to disk using the Cache-Control Header
- [ ] Add an option to ignore the cache forcing a get
- [ ] Allow the software to expose a RESTful API
- [ ] Add an option to scan sitemap.xml if available
