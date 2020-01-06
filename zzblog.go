package zzblog

import (
	"io"
	"os"
	"sort"
	"time"
)

// Blog a blog item
type Blog struct {
	URLID     string    `json:"url_id"`
	Title     string    `json:"title"`
	Tags      []string  `json:"tags"`
	Category  string    `json:"category"`
	Overview  string    `json:"overview"`
	Lang      string    `json:"lang"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upadted_at"`
	file      string
}

type LangGroup struct {
	blogs []*Blog
	defaultLang string
}

func NewLangGroup(defaultLang string, blog *Blog) *LangGroup {
	return &LangGroup{[]*Blog{blog}, defaultLang}
}

func (lg *LangGroup) Select(lang string) *Blog {
	for _, blog := range lg.blogs {
		if blog.Lang == lang {
			return blog
		}
	}
	for _, blog := range lg.blogs {
		if blog.Lang == lg.defaultLang {
			return blog
		}
	}
	return lg.blogs[0]
}

type OneFilter struct {
	Lang string
	Cate string
	Tag string
}

func (lg *LangGroup) One(filter * OneFilter) *Blog {
	hasTag := func (tag string, blog *Blog) bool {
		for _, t := range blog.Tags {
			if t == tag {
				return true
			}
		}
		return false
	}
	if filter.Lang != "" {
		blog := lg.Select(filter.Lang)
		if filter.Cate != "" && filter.Cate != blog.Category {
			return nil
		}
		if filter.Tag != "" && !hasTag(filter.Tag, blog) {
			return nil
		}
		return blog
	}
	for _, blog := range lg.blogs {
		if filter.Cate != "" && filter.Cate != blog.Category {
			continue
		}
		if filter.Tag != "" && !hasTag(filter.Tag, blog) {
			continue
		}
		return blog
	}

	return nil
}

func (lg *LangGroup) Add(blog *Blog) {
	lg.blogs = append(lg.blogs, blog)
}

func (lg *LangGroup) Del(lang string) {
	for i, blog := range lg.blogs {
		if blog.Lang == lang {
			lg.blogs = append(lg.blogs[:i], lg.blogs[i + 1:]...)
			return
		}
	}
}

type Contact struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Href  string `json:"href"`
}

type Author struct {
	Name     string    `json:"name"`
	Bio      string    `json:"bio"`
	Avatar   string    `json:"avatar"`
	Bg       string    `json:"bg"`
	Contacts []Contact `json:"contacts"`
}

func (b *Blog) SetFile(file string) {
	b.file = file
}

func (b *Blog) Detail() (blog *ParsedBlog, err error) {
	file, e := os.Open(b.file)
	if e != nil {
		e = err
		return
	}
	defer file.Close()
	blog = ParseBlog(file)
	return
}

const (
	ST_ASC   = 0
	ST_DESC  = 1
	SC_TIME  = 3
	SC_TITLE = 4
	SC_ID    = 5
)

type MBlogSet struct {
	blogs    []*Blog
	sort     map[int]int
	page     int
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

func (set *MBlogSet) Len() int {
	return len(set.blogs)
}

func (set *MBlogSet) Less(i, j int) bool {
	for k, v := range set.sort {
		if k == SC_TIME {
			if v == ST_ASC {
				return set.blogs[i].UpdatedAt.Before(set.blogs[j].UpdatedAt)
			}
			return set.blogs[j].UpdatedAt.Before(set.blogs[i].UpdatedAt)
		} else if k == SC_TITLE {
			if v == ST_ASC {
				return set.blogs[i].Title > set.blogs[j].Title
			}
			return set.blogs[i].Title < set.blogs[j].Title
		}
	}

	return true
}

func (set *MBlogSet) Swap(i, j int) {
	blog := set.blogs[i]
	set.blogs[i] = set.blogs[j]
	set.blogs[j] = blog
}

func (set *MBlogSet) Get() []*Blog {
	sort.Sort(set)
	if set.page == 0 {
		return set.blogs
	}
	b := (set.page - 1) * set.pageSize
	if b > len(set.blogs) {
		return []*Blog{}
	}
	e := set.page * set.pageSize
	if e > len(set.blogs) {
		e = len(set.blogs)
	}
	return set.blogs[b:e]
}

type Zzblog interface {
	Init() error
	Has(id, lang string) bool
	Get(id, lang string) *Blog
	AddByReader(r io.Reader) (*Blog, error)
	Add(*Blog) error
	AddImage(r ImageReader) error
	GetImage(id string) *Image
	GetImageByFilename(filename string) *Image
	Cates(lang string) []string
	Tags(lang string) []string
	Author(lang string) *Author
	Theme() []map[string]string
	Filter(func(*LangGroup) *Blog) BlogSet
}

type BlogSet interface {
	Sort(key, t int) BlogSet
	Page(p, ps int) BlogSet
	Get() []*Blog
}
