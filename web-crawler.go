/*
Exercise: Web Crawler
In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
*/
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	urls map[string]string
	mux  sync.Mutex
}

func (cache *UrlCache) Contains(url string) bool {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	_, contains := cache.urls[url]
	return contains
}

func (cache *UrlCache) Add(url, body string) {
	cache.mux.Lock()
	cache.urls[url] = body
	cache.mux.Unlock()
}

var cache UrlCache = UrlCache{urls: make(map[string]string)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	contains := cache.Contains(url)
	if !contains {
		cache.Add(url, "fetching")
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		cache.Add(url, body)
		fmt.Printf("found: %s %q\n", url, body)
		done := make(chan bool)
		for i, u := range urls {
			fmt.Printf("-> Crawling child %v/%v of %v : %v\n", i+1, len(urls), url, u)
			go func(u string) {
				Crawl(u, depth-1, fetcher)
				done <- true
			}(u)
		}

		for i := range urls {
			fmt.Printf("<- [%v] %v/%v Waiting for child.\n", url, i+1, len(urls))
		    <-done
		}
		fmt.Printf("<- Done with %v\n", url)
	} else {
		fmt.Printf("<- Cache already contains %v", url)
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	fmt.Println(cache.urls)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
