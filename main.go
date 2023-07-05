package main

import (
	"fmt"
	"time"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/rss-test", rssTest)

	e.Logger.Fatal(e.Start(":80"))
}

func rssTest(c echo.Context) error {
	feed, err := gofeed.NewParser().ParseURL("https://zenn.dev/topics/golang/feed")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error")
	}
	fmt.Println(feed.Title)
	fmt.Println(feed.FeedType, feed.FeedVersion)
	for _, item := range feed.Items {
		if item == nil {
			break
		}
		fmt.Println(item.Title)
		fmt.Println("\t->", item.Link)
		fmt.Println("\t->", item.PublishedParsed.Format(time.RFC3339))
	}
	return c.String(http.StatusOK, "RSS OK")
}
