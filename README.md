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

### Design

To build this application we'll break the application down into several components. We'll create some interfaces that will provide to the main application and will enable typed communication between the application.

#### Interfaces

| Name             | Purpose                                                              |
| ---------------- | -------------------------------------------------------------------- |
| PageProvider     | Provide a URL and a list of links from that page                     |
| TypeProvider     | Provide the type of asset based on MIME type and Content-Type Header |
| CachePersistence | An interface allowing for a the nodes to be cached                   |
| Processor        | An interface for processing data                                     |

#### Global Data Structures

| Name            | Purpose                                                                   |
| --------------- | ------------------------------------------------------------------------- |
| ExploredNodes   | A slice of URLs that have already been explored                           |
| DiscoveredNodes | A slice of unprocessed URLs that have been found                          |
| DiscardedNodes  | A slice of nodes that have been found, but should not be explored further |
| Sitemap         | A Graph showing how the explored nodes are connected                      |
