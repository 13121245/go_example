package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchItem string) ([]*Result, error)
}

func Match(matcher Matcher, feeds *Feed, searchItem string, results chan<- *Result) {
	searchResults, err := matcher.Search(feeds, searchItem)

	if err != nil {
		log.Println(err)
	}

	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	log.Printf("print results")
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
