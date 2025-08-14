package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/iliamunaev/json-monitor/internal/builder"
	"github.com/iliamunaev/json-monitor/internal/comparator"
	"github.com/iliamunaev/json-monitor/internal/fetcher"
)

func main() {
	url := os.Getenv("URL")
	if strings.Compare(url, "") == 0 {
		log.Fatal("URL is not set up")
	}

	secondsStr := os.Getenv("SECONDS_TIMEOUT")
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil || seconds <= 30 {
		seconds = 59
	}

	interval := time.Duration(seconds) * time.Second
	log.Printf("Interval set to %v\n", interval)

	var prev any

	for {
		raw, err := fetcher.FetchJSON(url)
		if err != nil {
			log.Printf("Fetch error: %v", err)
			time.Sleep(interval)
			continue
		}

		cur, err := builder.NormalizeJSON(raw)
		if err != nil {
			log.Printf("Normalize error: %v", err)
			time.Sleep(interval)
			continue
		}

		if prev == nil {
			prev = cur
			log.Println("Initial fetch complete")
			comparator.CompareJSONs(cur, prev)

		} else {
			comparator.CompareJSONs(cur, prev)
		}

		prev = cur
		time.Sleep(interval)
	}
}
