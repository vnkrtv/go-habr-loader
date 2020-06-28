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

	titleNode := doc.Find("span", "class", "post__title-text")
	textNode := doc.Find("div", "class", "post__text")
	ratingNode := doc.Find("span", "class", "voting-wjt__counter")

	if titleNode.Pointer == nil || textNode.Pointer == nil || ratingNode.Pointer == nil {
		return pg.HabrPost{}, fmt.Errorf("post with post_id %d not found", postID)
	}

	date, err := getDate(&doc)
	if err != nil {
		return pg.HabrPost{}, err
	}

	viewsCount, err := getViewsCount(&doc)
	if err != nil {
		return pg.HabrPost{}, err
	}

	commentsCount, err := getCommentsCount(&doc)
	if err != nil {
		return pg.HabrPost{}, err
	}

	bookmarksCount, err := getBookmarksCount(&doc)
	if err != nil {
		return pg.HabrPost{}, err
	}

	return pg.HabrPost{
		ID:             postID,
		Date:           date,
		Title:          titleNode.FullText(),
		Text:           textNode.FullText(),
		ViewsCount:     viewsCount,
		CommentsCount:  commentsCount,
		BookmarksCount: bookmarksCount,
		Rating:         ratingNode.Text(),
	}, err
}

func getDate(doc *soup.Root) (time.Time, error) {
	postTimeNode := doc.Find("span", "class", "post__time")
	if postTimeNode.Pointer == nil {
		return time.Time{}, fmt.Errorf("post time span not found")
	}
	return parseDateSpan(postTimeNode.Text())
}

func getViewsCount(doc *soup.Root) (int, error) {
	viewsSpanNode := doc.Find("span", "class", "post-stats__views-count")
	if viewsSpanNode.Pointer == nil {
		return 0, fmt.Errorf("views span not found")
	}
	viewsCount, err := parseViewsSpan(viewsSpanNode.Text())
	if err != nil {
		return 0, err
	}
	return viewsCount, err
}

func getCommentsCount(doc *soup.Root) (int, error) {
	commentsSpanNode := doc.Find("span", "class", "comments-section__head-counter")
	if commentsSpanNode.Pointer == nil {
		return 0, fmt.Errorf("comments span not found")
	}
	commentsCount, err := strconv.ParseInt(
		strings.TrimSpace(commentsSpanNode.Text()), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(commentsCount), err
}

func getBookmarksCount(doc *soup.Root) (int, error) {
	bookmarksSpanNode := doc.Find("span", "class", "bookmark__counter")
	if bookmarksSpanNode.Pointer == nil {
		return 0, fmt.Errorf("bookmarks span not found")
	}
	bookmarksCount, err := strconv.ParseInt(
		strings.TrimSpace(bookmarksSpanNode.Text()), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(bookmarksCount), err
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
