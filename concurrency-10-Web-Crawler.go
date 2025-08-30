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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of
//depth and adds each crawled url to fetchedUrls.
//sync.WaitGroup is used to halt main thread execution
//until all Crawl routines have executed.
func Crawl(url string, depth int, fetcher Fetcher, fetchedUrls *safeFetchedUrls, wg *sync.WaitGroup) {
    // DONE: Fetch URLs in parallel.
    // DONE: Don't fetch the same URL twice.
    defer wg.Done()
    if depth <= 0 {
        return
    }
    fetchedUrls.add(url)
    fmt.Printf("Added \"%s\" to fetchedUrls at depth %d\n", url, depth)
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        if !fetchedUrls.read(u) {
            wg.Add(1)
            go Crawl(u, depth-1, fetcher, fetchedUrls, wg)
        }
    }
    return
}

func main() {
    var wg sync.WaitGroup
    fetchedUrls := safeFetchedUrls{urls: make(map[string]bool)}
    wg.Add(1)
    go Crawl("https://golang.org/", 4, fetcher, &fetchedUrls, &wg)
    wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

type safeFetchedUrls struct {
    mu sync.Mutex
    urls map[string]bool
}

func (fetchedUrls *safeFetchedUrls) add(url string) {
    fetchedUrls.mu.Lock()
    defer fetchedUrls.mu.Unlock()
    fetchedUrls.urls[url] = true
}

func (fetchedUrls *safeFetchedUrls) read(url string) bool {
    fetchedUrls.mu.Lock()
    defer fetchedUrls.mu.Unlock()
    return fetchedUrls.urls[url]
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