package loader

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"

	pg "../postgres"
)

const (
	habrHref = "https://habr.com/ru/post/%d/"
	dateLayout = "2 Jan 2006 15:04"
)
var months = map[string]string{
	"января": "Jan",
	"февраля": "Feb",
	"марта": "Mar",
	"апреля": "Apr",
	"мая": "May",
	"июня": "Jun",
	"июля": "Jul",
	"августа": "Aug",
	"сентября": "Sep",
	"октября": "Oct",
	"ноября": "Nov",
	"декабря": "Dec",
}

func LoadPost(postID int) (pg.HabrPost, error)  {
	resp, err := soup.Get(fmt.Sprintf(habrHref, postID))
	if err != nil {
		return pg.HabrPost{}, err
	}
	doc := soup.HTMLParse(resp)
	if doc.Find("span", "class", "post__time").Error != nil {
		return pg.HabrPost{}, fmt.Errorf("post with %d post id not found", postID)
	}

	date, err := parseDateSpan(doc.Find("span", "class", "post__time").Text())
	if err != nil {
		return pg.HabrPost{}, err
	}
	viewsCount, err := parseViewsSpan(doc.Find("span", "class", "post-stats__views-count").Text())
	if err != nil {
		return pg.HabrPost{}, err
	}

	commentsSpan := doc.Find("span", "class", "comments-section__head-counter").Text()
	commentsCount, err := strconv.ParseInt(
		strings.TrimSpace(commentsSpan), 10, 64)
	if err != nil {
		return pg.HabrPost{}, err
	}

	bookmarksSpan := doc.Find("span", "class", "bookmark__counter").Text()
	bookmarksCount, err := strconv.ParseInt(
		strings.TrimSpace(bookmarksSpan), 10, 64)
	if err != nil {
		return pg.HabrPost{}, err
	}

	return pg.HabrPost{
		ID:             postID,
		Date:           date,
		Title:          doc.Find("span", "class", "post__title-text").Text(),
		Text:           doc.Find("div", "class", "post__text").Text(),
		ViewsCount:     viewsCount,
		CommentsCount:  int(commentsCount),
		BookmarksCount: int(bookmarksCount),
		Rating:         doc.Find("span", "class", "voting-wjt__counter").Text(),
	}, err
}

func parseViewsSpan(viewsSpan string) (int, error) {
	if strings.Contains(viewsSpan, "k") {
		viewsSpan = strings.Replace(viewsSpan, ",", ".", 1)
		viewsCount, err := strconv.ParseFloat(viewsSpan[:len(viewsSpan) - 1], 64)
		return int(viewsCount * 1000), err
	} else {
		viewsCount, err :=  strconv.ParseInt(viewsSpan, 10, 64)
		return int(viewsCount), err
	}
}

func parseDateSpan(postDate string) (time.Time, error) {
	month := strings.Split(postDate, " ")[1]
	postDate = strings.Replace(postDate, month, months[month], 1)
	postDate = strings.Replace(postDate, " в ", " ", 1)
	date, err := time.Parse(dateLayout, postDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, err
}