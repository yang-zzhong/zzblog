package zzblog

import (
	"time"
	"io"
)

// Blog a blog item
type Blog struct {
	URLID string			`json:"url_id"`
	Title string			`json:"title"`
	Tags []string			`json:"tags"`
	Category string			`json:"category"`
	Overview string			`json:"overview"`
	File string				`json:"file"`
	Lang string				`json:"lang"`
	Langs map[string]string	`json:"langs"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
}

const (
	ST_ASC = 0
	ST_DESC = 1
	SC_TIME = 3
	SC_TITLE = 4
	SC_ID = 5
)

type MBlogSet struct {
	blogs []*Blog
	sort map[int]int
	page int
	pageSize int
}

func NewMBlogSet(blogs []*Blog) *MBlogSet {
	bs := new(MBlogSet)
	bs.sort = make(map[int]int)
	bs.page = 0
	bs.pageSize = 10
	bs.blogs = blogs
	return bs
}

func (set *MBlogSet) Sort(key, t int) BlogSet {
	set.sort[key] = t
	return set
}

func (set *MBlogSet) Page(p, ps int) BlogSet {
	set.page = p
	set.pageSize = ps
	return set
}

func (set *MBlogSet) Get() []*Blog {
	if set.page == 0 {
		return set.blogs
	}
	b := (set.page - 1) * set.pageSize
	e := set.page * set.pageSize
	if e > len(set.blogs) {
		e = len(set.blogs)
	}
	return set.blogs[b:e]
}

type Zzblog interface {
	Has(id string) bool
	Get(id string) *Blog
	AddByReader(r io.Reader) (*Blog, error)
	Add(*Blog) error
	Filter(func(*Blog) bool) *BlogSet
}

type BlogSet interface {
	Sort(key, t int) BlogSet
	Page(p, ps int) BlogSet
	Get() []*Blog
}