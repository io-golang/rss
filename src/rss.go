package main

import (
	"html/template"
	"log"
	"time"
	"path/filepath"
	"encoding/xml"
	"bytes"
	"strings"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/html/charset"
)

// RSS struct contains multiple rss channels
type RSS struct {
	Channels []Channel `xml:"channel"`
}

// Channel struct contains multiple rss items
type Channel struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	Items []Item `xml:"item"`
	PubDate string `xml:"pubDate"`
}

// Item struct contains rss data definition
type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

// RssFeedResult struct for fetched feeds.
// It contains a pointer of type RSS and error
type RssFeedResult struct {
	feed *RSS
	err error
}

// LoadTemplate loads email templates
func LoadTemplate() *template.Template {
	templates, err := filepath.Glob("../email_templates/*")
	if err != nil {
		log.Println(err)
	}
	email := template.Must(template.New("layout.go.html").ParseFiles(templates...))
	return email
}

// FetchFeeds fetches feeds from urls by calling FetchFeed
func FetchFeeds(urls []string) ([]*RSS) {
	
	channel := make(chan RssFeedResult, len(urls))

	
	for _, url := range urls {
		go FetchFeed(strings.TrimSpace(url), channel)
	}

	
	feeds := []*RSS{}
	for i := 0; i < len(urls); i++ {
		res := <-channel
		
		if res.err != nil {
			log.Println(res.err)
			continue
		}
		feeds = append(feeds, res.feed)
	}

	return feeds
}

// FetchFeed fetches feed from url
func FetchFeed(url string, channel chan RssFeedResult) {

	net := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := net.Get(url)
	if err != nil {
		log.Println(err)
		channel <- RssFeedResult{nil, err}
		return
	}
	defer res.Body.Close()

	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		channel <- RssFeedResult{nil, err}
		return
	}

	feed, err := ParseFeed(body)
	if err != nil {
		log.Println(err)
		channel <- RssFeedResult{nil, err}
		return
	}

	channel <- RssFeedResult{feed, nil}
}

// ParseFeed parses feed from body
func ParseFeed(body []byte) (*RSS, error) {
	feed := RSS{}
	reader := bytes.NewReader(body)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&feed)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}


// TODO: Filter feeds
// func FilterFeeds(feeds []*RSS) ([]*RSS) {
// 	expired := time.Now()
// 	rfc1123 := "Mon, 02 Jan 2006 15:04:05 MST"
// 	rfc1123Z := "Mon, 02 Jan 2006 15:04:05 -0700"
// 	rfc3339 := "2006-01-02T15:04:05Z07:00"
// 	filtered := make([]*RSS, 0, len(feeds))
// 	var pub_date time.Time
	
// 	for _, feed := range feeds {
// 		for _, channel := range feed.Channels {
// 			// for _, item := range channel.Items {
// 			expired, err := time.Parse(rfc1123, expired.Format(rfc1123))
// 			if err != nil {
// 				log.Println(err)
// 				return []*RSS{}
// 			}
// 			pub_date, err = time.Parse(rfc1123, channel.PubDate)
// 			if err != nil {
// 				expired, err = time.Parse(rfc1123Z, expired.Format(rfc1123Z))
// 				if err != nil {
// 					log.Println(err)
// 					return []*RSS{}
// 				}
// 				pub_date, err = time.Parse(rfc1123Z, channel.PubDate)
// 				if err != nil {
// 					expired, err = time.Parse(rfc3339, expired.Format(rfc3339))
// 					if err != nil {
// 						log.Println(err)
// 						return []*RSS{}
// 					}
// 					pub_date, err = time.Parse(rfc3339, channel.PubDate)
// 					if err != nil {
// 						log.Println(err)
// 						// log.Println("Channel: ", channel)
// 						log.Println("RSS date format: " + channel.PubDate)
// 						return []*RSS{}
// 					}
// 				}
// 				// }
// 			}
// 		}
// 		expired = expired.Add(-24*time.Hour)
// 		if time.Time(pub_date).After(expired) {
// 			filtered = append(filtered, feed)
// 		}
// 	}
// 	return filtered
// }