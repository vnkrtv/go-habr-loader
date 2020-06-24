package loader

import (
	"os"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"

	pg "../postgres"
)

const habrHref = "https://habr.com/ru/post/%d/"

func LoadPost(postID int) (pg.HabrPost, error)  {
	resp, err := soup.Get(fmt.Sprintf(habrHref, postID))
	if err != nil {
		return pg.HabrPost{}, err
	}
	doc := soup.HTMLParse(resp)


	return pg.HabrPost{
		ID:             0,
		Date:           time.Time{},
		Title:          "",
		Text:           "",
		ViewsCount:     0,
		CommentsCount:  0,
		BookmarksCount: 0,
		Rating:         "",
	}, err
}

func GetViews(doc soup.Root) (int, error) {
	viewsSpan := doc.Find("span", "class", "post-stats__views-count").Text()
	if strings.Contains(viewsSpan, "k") {
		viewsCount, err := strconv.ParseFloat(viewsSpan[:len(viewsSpan) - 1], 64)
		return int(viewsCount * 1000), err
	} else {
		viewsCount, err :=  strconv.ParseInt(viewsSpan, 10, 64)
		return int(viewsCount), err
	}
}

func GetDate(doc soup.Root) (time.Time, error) {
	postDate := doc.Find("span", "class", "post__time").Text()
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, postDate)

	if err != nil {
		fmt.Println(err)
	}

}