package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	// "github.com/gofiber/fiber/v2"
)

type Contest struct {
	Name      string `json:"name"` 
    Duration  int `json:"duration"`
	StartTime int64 `json:"startTime"`
	Link       string `json:"link"`
}

func main() {
    c := colly.NewCollector()
    contests := make([]Contest, 0)
        
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnHTML("div[id=contest-table-upcoming] > div.panel > div.table-responsive > table.table > tbody", func(h *colly.HTMLElement) {
        h.ForEach("tr", func(i int, h *colly.HTMLElement) {
            name := h.DOM.Find(":nth-child(2) > a").Text()

            baseURL := "https://atcoder.jp"
            contestPath, exists:= h.DOM.Find(":nth-child(2) > a").Attr("href")
            if !exists {
                fmt.Println("href attribute does not exist")
            }
            link := baseURL + contestPath

            datetime := h.DOM.Find(":nth-child(1) > a > time.fixtime-full").Text()
            // reference time written in datetime's format for later parsing
            datetimeLayout := "2006-01-02 15:04:05-0700"
            t, err := time.Parse(datetimeLayout, datetime)
            if err != nil {
                fmt.Println(err)
            }
            startTime := t.Unix()

            durationString := h.DOM.ChildrenFiltered(":nth-child(3)").Text()
            duration := parseDuration(durationString)

            // initialize Contest struct var with parsed values
            contest := Contest{
                Name: name,
                Duration: duration,
                StartTime: startTime,
                Link: link,
            }

            contests = append(contests, contest)
        })
    })

    c.Visit("https://atcoder.jp/contests/")

    fmt.Println(contests)
    writeJSON(contests)

 // app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// app.Get("/lol", func(c *fiber.Ctx) error {
	// 	return c.SendString("LOL😂")
	// })

	// app.Listen(":3000")
}

// converts duration specified in HH:MM to seconds
func parseDuration(durationString string) int {
    parts := strings.Split(durationString, ":")
    hours, err := strconv.Atoi(parts[0])
    if err != nil {
        fmt.Println("Couldn't parse hours")
    }

    minutes, err := strconv.Atoi(parts[1])
    if err != nil {
        fmt.Println("Couldn't parse minutes")
    }

    duration := hours*60*60 + minutes*60
    return duration
}

func writeJSON(contests []Contest) {
    f, err := json.MarshalIndent(contests, "", " ")
    if err != nil {
        fmt.Println("Failed to JSON encode contests slice")
    }

    os.WriteFile("contests.json", f, 0644)
}