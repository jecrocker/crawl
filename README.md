## Web Crawler

This is a basic web crawler, it is designed to crawl a given website and to create a sitemap from a given website.

### Assumptions

- The URL will have a **trusted** TLS certificate or no TLS certificate at all
- The application will respect the robots.txt if it is present
- If a sitemap.xml exists then the URIs from this will be added to the search list

### Tasks

#### Required

- [x] Print each of the URIs with a list of the links that are found
- [x] Respect the given domain, i.e. do not recurse sub-domains or external domains

#### Some helpful features

- [ ] Cache results to disk using the Cache-Control Header
- [ ] Add an option to ignore the cache forcing a get
- [ ] Allow the software to expose a RESTful API
- [ ] Add an option to scan sitemap.xml if available
- [ ] Respect Robots.txt

### Design

To build this application we'll break the application down into several components. We'll create some interfaces that will provide to the main application and will enable typed communication between the application.

#### Interfaces

| Name             | Purpose                                            |
| ---------------- | -------------------------------------------------- |
| PageProvider     | Provide a URL and a list of links from that page   |
| CachePersistence | An interface allowing for a the nodes to be cached |
