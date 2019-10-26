package zzblog

import (
	"testing"
	"os"
	"log"
)

func TestParseBlog(t * testing.T) {
	f, e := os.Open("./blog.md")
	if e != nil {
		log.Print(e)
		return
	}
	if blog, err := ParseBlog(f); err != nil {
		log.Print(err)
	} else {
		log.Printf("title: %s\n", blog.Title)
		log.Printf("overview: %s\n", blog.Overview)
		log.Printf("urlid: %s\n", blog.UrlID)
		log.Printf("content: %s\n", blog.Content)
		log.Printf("tags: %v\n", blog.Tags)
		log.Printf("cate: %s\n", blog.Category)
	}
}