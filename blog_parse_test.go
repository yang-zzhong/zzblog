package zzblog

import (
	"log"
	"os"
	"testing"
)

func TestParseBlog(t *testing.T) {
	f, e := os.Open("./test/root/blogs/blog.md")
	if e != nil {
		log.Print(e)
		return
	}
	blog := ParseBlog(f)
	log.Printf("title: %s\n", blog.Title)
	log.Printf("overview: %s\n", blog.Overview)
	log.Printf("urlid: %s\n", blog.URLID)
	log.Printf("content: %s\n", blog.Content)
	log.Printf("tags: %v\n", blog.Tags)
	log.Printf("lang: %s\n", blog.Lang)
	log.Printf("cate: %s\n", blog.Category)
}
