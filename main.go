package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/logrusorgru/aurora"
)

const DEFAULT_USER_AGENT = "Go Link Health (github.com/nikoheikkila/go-link-health)"
const HTTP_MIN_STATUS = 200
const HTTP_MAX_STATUS = 299

func main() {
	target, urlError := getURL(os.Args)

	if urlError != nil {
		handleError(urlError)
		os.Exit(1)
	}

	collector := getCollector()
	collector.Visit(target.String())
	collector.Wait()
}

func getURL(args []string) (*url.URL, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Usage: %s <url>\n", args[0])
	}

	return url.Parse(args[1])
}

func getCollector() *colly.Collector {
	userAgent := flag.String("user-agent", DEFAULT_USER_AGENT, "User-Agent for scraping")
	depth := flag.Int("depth", 2, "Recursion depth for scraping")
	threads := flag.Int("threads", 4, "Number of threads to use for scraping")

	flag.Parse()

	collector := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent(*userAgent),
		colly.MaxDepth(*depth),
	)

	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: *threads,
		RandomDelay: 1 * time.Second,
	})

	if err != nil {
		handleError(err)
	}

	setHandlers(collector)

	return collector
}

func setHandlers(collector *colly.Collector) {
	collector.OnError(func(response *colly.Response, err error) {
		url := response.Request.URL
		reason := err.Error()

		handleError(fmt.Errorf("Request to %s failed. Reason: %s", url, reason))
	})

	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		link := element.Attr("href")
		element.Request.Visit(link)
	})

	collector.OnResponse(func(response *colly.Response) {
		url := response.Request.URL.String()
		status := response.StatusCode

		if !isHealthy(status) {
			printError(url, status)
			return
		}

		printSuccess(url)
	})
}

func isHealthy(status int) bool {
	return status >= HTTP_MIN_STATUS && status <= HTTP_MAX_STATUS
}

func handleError(error error) {
	fmt.Println(aurora.Red("Error:"), error)
}

func printSuccess(url string) {
	fmt.Printf("Link to %s is %s\n", url, aurora.Green("healthy"))
}

func printError(url string, status int) {
	fmt.Printf("Link to %s is %s with status %d\n", url, aurora.Red("down"), aurora.Bold(status))
}
