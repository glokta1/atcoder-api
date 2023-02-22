package main

import (
	"fmt"

	// "github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	// "github.com/gofiber/fiber/v2"
)

type Contest struct {
	Name      string
	Duration  int
	StartTime int
	Url       string
}

func main() {c := colly.NewCollector()
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    // c.OnHTML("table.table", func(h *colly.HTMLElement) {
    //     h.DOM.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
    //         fmt.Println("Case #", i)
    //         // fmt.Println(s.Html())
    //         // startTime := s.Find(":nth-child(1)").Text()
    //         // name := s.Find(":nth-child(2)").Text()
    //         duration := s.Find(":nth-child(3)").Text()
    //         // fmt.Println("contest name:", name) 
    //         // fmt.Println("Time: ", startTime) 
    //         fmt.Println("Duration:", duration)
    //     })

    //     // if err != nil {
    //     //     fmt.Println(err)
    //     // }

    //     // fmt.Print(s)
    // })

    c.OnHTML("div[id=contest-table-upcoming] > div.panel > div.table-responsive > table.table > tbody", func(h *colly.HTMLElement) {
        h.ForEach("tr", func(i int, h *colly.HTMLElement) {
            name := h.DOM.Find(":nth-child(2) > a").Text()
            link, exists:= h.DOM.Find(":nth-child(2) > a").Attr("href")
            if !exists {
                fmt.Println("href attribute does not exist")
            }

            startTime := h.DOM.Find(":nth-child(1) > a > time.fixtime-full").Text()
            duration := h.DOM.ChildrenFiltered(":nth-child(3)").Text()
            fmt.Println(name)
            fmt.Println(link)
            fmt.Println(startTime)
            fmt.Println(duration)
            fmt.Println("*************")
        })
        // fmt.Println(name)
        // h.DOM.Each(func(i int, s *goquery.Selection) {
        //     name, err:= s.Find(":nth-child(1)").Children().Children().Html()
        //     if err != nil {
        //         fmt.Println(err)
        //     }
        //     fmt.Println(name)
        //     fmt.Println("*************")
        // })
    })

    c.Visit("https://atcoder.jp/contests/")

	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// app.Get("/lol", func(c *fiber.Ctx) error {
	// 	return c.SendString("LOL😂")
	// })

	// app.Listen(":3000")
}
