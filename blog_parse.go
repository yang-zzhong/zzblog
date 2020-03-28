package zzblog

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

const (
	sText = iota
	sHeaderBeginLeading
	sHeaderBegin
	sHeader
	sKey
	sValLeading
	sVal
	sHeaderEnd
	sHeaderEndLeading
	sEnd

	tTitle    = "title"
	tUrlid    = "urlid"
	tTags     = "tags"
	tCate     = "cate"
	tCategory = "category"
	tOverview = "overview"
	tImage    = "image"
	tImg      = "img"
	tLang     = "lang"
)

// ParsedBlog a parsed result
type ParsedBlog struct {
	URLID    string
	Title    string
	Tags     []string
	Category string
	Overview string
	Content  []byte
	Image    string
	Lang     string
}

func NewParsedBlog() *ParsedBlog {
	p := new(ParsedBlog)
	p.Tags = []string{}
	p.Content = []byte{}
	p.Lang = "en"
	return p
}

func (blog *ParsedBlog) UnmarshalJSON(b []byte) error {
	d := make(map[string]interface{})
	if err := json.Unmarshal(b, &d); err != nil {
		return err
	}
	blog.URLID = d["url_id"].(string)
	blog.Title = d["title"].(string)
	blog.Tags = d["tags"].([]string)
	blog.Category = d["category"].(string)
	blog.Content = []byte(d["content"].(string))
	blog.Lang = d["lang"].(string)

	return nil
}

func (blog *ParsedBlog) MarshalJSON() ([]byte, error) {
	d := make(map[string]interface{})

	d["url_id"] = blog.URLID
	d["title"] = blog.Title
	d["tags"] = blog.Tags
	d["category"] = blog.Category
	d["overview"] = blog.Overview
	d["content"] = string(blog.Content)
	d["lang"] = blog.Lang

	return json.Marshal(d)
}

func ParseBlog(r io.Reader) *ParsedBlog {
	parser := new(BlogParser)
	return parser.Parse(r)
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z'
}

func isW(b byte) bool {
	return isNumber(b) || isAlpha(b)
}

func isWhiteSpace(b byte) bool {
	return b == '\t' || b == ' ' || b == '\n'
}

func eq(a, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

type BlogParser struct {
	buf       []byte
	cache     []byte
	state     int
	col, line int
	key, val  []byte
}

func (p *BlogParser) Parse(r io.Reader) *ParsedBlog {
	blog := NewParsedBlog()
	p.buf = make([]byte, 1)
	p.cache = []byte{}
	p.state = sText
	for {
		if _, e := r.Read(p.buf); e != nil {
			if e == io.EOF {
				return blog
			}
			return nil
		}
		switch p.state {
		case sText:
			p.inText(blog)
		case sHeaderBeginLeading:
			p.inHeaderLeading(blog)
		case sHeaderBegin:
			p.inHeaderBegin(blog)
		case sHeader:
			p.inHeader(blog)
		case sKey:
			p.inKey(blog)
		case sValLeading:
			p.inValLeading(blog)
		case sVal:
			p.inVal(blog)
		case sHeaderEndLeading:
			p.inHeaderLeading(blog)
		case sHeaderEnd:
			p.inHeaderEnd(blog, r)
		case sEnd:
			return blog
		}
	}

	return blog
}

func (p *BlogParser) inText(blog *ParsedBlog) {
	char := p.buf[0]
	if char == '+' {
		p.state = sHeaderBeginLeading
		p.cache = append(p.cache, char)
		return
	}
	blog.Content = append(blog.Content, char)
}

func (p *BlogParser) inHeaderBegin(blog *ParsedBlog) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if char == '+' {
		p.state = sHeader
	} else if char != '-' {
		p.toText(blog)
	}
}

func (p *BlogParser) inHeader(blog *ParsedBlog) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if isAlpha(char) {
		p.state = sKey
		p.key = []byte{char}
	} else if char == '+' {
		p.state = sHeaderEnd
	}
}

func (p *BlogParser) inKey(blog *ParsedBlog) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if isW(char) || char == '-' {
		p.key = append(p.key, char)
		return
	} else if char == ':' {
		p.state = sValLeading
		return
	}
	p.toText(blog)
}

func (p *BlogParser) inValLeading(blog *ParsedBlog) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if !isWhiteSpace(char) {
		p.state = sVal
		p.val = []byte{char}
	}
}

func (p *BlogParser) inVal(blog *ParsedBlog) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if char != '\n' {
		p.val = append(p.val, char)
		return
	}
	p.setBlogAttr(blog)
	p.state = sHeader
}

func is(k string, p string) bool {
	return strings.ToLower(k) == p
}

func (p *BlogParser) setBlogAttr(blog *ParsedBlog) error {
	k := string(p.key)
	if is(k, tTitle) {
		blog.Title = string(p.val)
		return nil
	} else if is(k, tUrlid) {
		err := handleUrlid(string(p.val), blog)
		return err
	} else if is(k, tTags) {
		handleTags(string(p.val), blog)
		return nil
	} else if is(k, tCate) || is(k, tCategory) {
		blog.Category = string(p.val)
		return nil
	} else if is(k, tOverview) {
		blog.Overview = string(p.val)
		return nil
	} else if is(k, tLang) {
		blog.Lang = string(p.val)
		return nil
	} else if is(k, tImg) || is(k, tImage) {
		blog.Image = string(p.val)
	}
	return errors.New("undefined key '" + k + "'")
}

func handleUrlid(v string, blog *ParsedBlog) error {
	r := strings.NewReader(v)
	buf := make([]byte, 1)
	temp := []byte{}
	for {
		if l, e := r.Read(buf); e != nil {
			if e == io.EOF {
				blog.URLID = string(temp)
				return nil
			}
			return e
		} else if l == 0 {
			continue
		}
		if (buf[0] >= 'A') && (buf[0] <= 'Z') || (buf[0] >= 'a' && buf[0] <= 'z') || (buf[0] == '-') {
			temp = append(temp, buf[0])
		}
	}
}

func handleTags(v string, blog *ParsedBlog) {
	tags := strings.Split(v, "#")
	for _, tag := range tags {
		t := strings.Trim(tag, " ,")
		if t != "" {
			blog.Tags = append(blog.Tags, t)
		}
	}
}

func (p *BlogParser) inHeaderLeading(blog *ParsedBlog) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if char == '-' {
		if p.state == sHeaderBeginLeading {
			p.state = sHeaderBegin
		} else {
			p.state = sHeaderEnd
		}
		return
	}
	p.toText(blog)
}

func (p *BlogParser) inHeaderEnd(blog *ParsedBlog, r io.Reader) {
	char := p.buf[0]
	p.cache = append(p.cache, char)
	if char == '+' {
		buf := make([]byte, 2048)
		for {
			if _, e := r.Read(buf); e != nil {
				p.state = sEnd
				p.cache = []byte{}
				return
			}
			blog.Content = append(blog.Content, buf...)
		}
	} else if char != '-' {
		p.toText(blog)
	}
}

func (p *BlogParser) toText(blog *ParsedBlog) {
	p.state = sText
	blog.Content = append(blog.Content, p.cache...)
	p.cache = []byte{}
}
