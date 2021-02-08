package main

import (
	"log"
	"os"

	"rss/rss/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
