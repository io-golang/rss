package main

import (
	"os"
	// "github.com/joho/godotenv"
	"log"
	"strings"
)

func main() {
	log.Println("Running... RSS...")	
	// uncomment below to load envs vars from .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	
	parsed_urls := strings.Split(os.Getenv("RSS_URLS"), ",")
	log.Println("Fetching feeds from: ", parsed_urls)
	feeds := FetchFeeds(parsed_urls)
	log.Println("Found ", len(feeds), " feeds")
	// TODO filter feeds
	if len(feeds) == 0 {
		log.Println("No feeds found")
		return
	}

	// load templates for email content
	log.Println("Loading email templates")
	template := LoadTemplate()
	str_writer := &strings.Builder{}
	err := template.Execute(str_writer, feeds)
	if err != nil {
		log.Fatalln(err)
	}

	// Send email
	log.Println("Sending email")
	err = Emailer(str_writer.String())
	if err != nil {
		log.Fatalln(err)
	}
}