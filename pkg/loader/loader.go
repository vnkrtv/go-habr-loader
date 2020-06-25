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
	month := strings.Split(postDate, " ")[1]
	postDate = strings.Replace(postDate, month, months[month], 1)
	postDate = strings.Replace(postDate, " в ", " ", 1)
	date, err := time.Parse(dateLayout, postDate)
	if err != nil {
		return time.Time{}, err
	}
	return date, err
}