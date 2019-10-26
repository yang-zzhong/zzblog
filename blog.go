package zzblog

import (
	"io"
	"fmt"
	"strings"
	"errors"
	"time"
	"bytes"
)

type Blog struct {
	UrlID 	string
	Title 	string
	Tags 	[]string
	Category string
	Overview string
	Content string
	images []string
	Lang string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	BLOG = 0
	IN_TAG = 1
	IN_TAG_BEGIN = 2
	IN_TAG_CONTENT = 3
	IN_TAG_END = 4
	HEAD_LEADING = 5
	HEAD_H = 6
	HEAD_HE = 7
	HEAD_HEA = 8
	HEAD_HEAD = 9

	HEAD_KEY = 12
	HEAD_VALUE = 13
	HEAD_END = 14

	TITLE = "title"
	URLID = "urlid"
	TAGS = "tags"
	CATE = "cate"
	CATEGORY = "category"
	OVERVIEW = "overview"
	LANG = "lang"
)

func ParseBlog(r io.Reader) (info * Blog, err error) {
	info = &Blog{}
	buffer := make([]byte, 1)
	state := BLOG
	var temp string
	var col, line int = 0, 1
	tmp := "unexpected '%s' at %d line %d"
	outBuf := make([]byte, len(tmp) + 16)
	out := bytes.NewBuffer(outBuf)
	var key, value string
	maps := map[string]string{}
	for {
		if _, e := r.Read(buffer); e != nil {
			if e == io.EOF {
				return
			}
			err = e
			return
		}
		switch state {
		case BLOG:
			if buffer[0] == '<' {
				temp = ""
				state = IN_TAG
				break
			}
			if buffer[0] == '`' {
				temp = "`"
				state = HEAD_LEADING
				break
			}
			if info.Content == "" && buffer[0] == '\n' {
				break
			}
			info.Content = info.Content + string(buffer[0])
		case HEAD_LEADING:
			if buffer[0] == '`' && len(temp) < 3 {
				temp = temp + string(buffer[0])
				break
			}
			if buffer[0] == 'h' || buffer[0] == 'H' {
				temp = temp + string(buffer[0])
				state = HEAD_H
				break
			}
			info.Content = info.Content + temp
			temp = ""
			if buffer[0] == '<' {
				state = IN_TAG
				break
			}
			state = BLOG
			info.Content = info.Content + string(buffer[0])
		case HEAD_H:
			if buffer[0] == 'E' || buffer[0] == 'e' {
				temp = temp + string(buffer[0])
				state = HEAD_HE
				break
			}
			info.Content = info.Content + temp
			temp = ""
			if buffer[0] == '<' {
				state = IN_TAG
				break
			}
			state = BLOG
			info.Content = info.Content + string(buffer[0])
		case HEAD_HE:
			if buffer[0] == 'A' || buffer[0] == 'a' {
				temp = temp + string(buffer[0])
				state = HEAD_HEA
				break
			}
			info.Content = info.Content + temp
			temp = ""
			if buffer[0] == '<' {
				state = IN_TAG
				break
			}
			state = BLOG
			info.Content = info.Content + string(buffer[0])
		case HEAD_HEA:
			if buffer[0] == 'D' || buffer[0] == 'd' {
				temp = temp + string(buffer[0])
				state = HEAD_HEAD
				break
			}
			info.Content = info.Content + temp
			temp = ""
			if buffer[0] == '<' {
				state = IN_TAG
				break
			}
			state = BLOG
		case HEAD_HEAD:
			if buffer[0] == '\n' {
				temp = ""
				state = HEAD_KEY
				break
			}
		case HEAD_KEY:
			if buffer[0] >= 'A' && buffer[0] <= 'Z' || buffer[0] >= 'a' && buffer[0] <= 'z' {
				temp = temp + string(buffer[0])
				break
			}
			if buffer[0] == ':' {
				key = strings.TrimSpace(temp)
				temp = ""
				state = HEAD_VALUE
				break
			}
			if buffer[0] == '`' && strings.TrimSpace(temp) == "" {
				temp = "`"
				state = HEAD_END
				break
			}
			fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
			err = errors.New(out.String())
			return
		case HEAD_VALUE:
			if buffer[0] == '\n' {
				if e := setProp(key, strings.TrimSpace(temp), info) ; e != nil {
					err = e
					return
				}
				temp = ""
				state = HEAD_KEY
				break
			}
			temp = temp + string(buffer[0])
		case HEAD_END:
			if buffer[0] == '`' {
				temp = temp + string(buffer[0])
				break
			}
			if strings.TrimSpace(temp) == "```" {
				temp = ""
				state = BLOG
				break
			}
			fmt.Printf("temp: %s\n", temp)
			fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
			err = errors.New(out.String())
			return
		case IN_TAG:
			if buffer[0] == '>' {
				fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
				err = errors.New(out.String())
				return
			}
			if buffer[0] == '/' {
				state = IN_TAG_END
				break
			}
			if buffer[0] < 'A' || buffer[0] > 'z' {
				fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
				err = errors.New(out.String())
				return
			}
			temp = string(buffer[0])
			state = IN_TAG_BEGIN
		case IN_TAG_BEGIN:
			if buffer[0] == '>' {
				key = temp
				temp = ""
				state = IN_TAG_CONTENT
				break
			}
			if buffer[0] < 'A' || buffer[0] > 'z' {
				fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
				err = errors.New(out.String())
				return
			}
			temp = temp + string(buffer[0])
		case IN_TAG_CONTENT:
			if buffer[0] == '<' {
				value = temp
				temp = ""
				state = IN_TAG
				break
			}
			temp = temp + string(buffer[0])
		case IN_TAG_END:
			if buffer[0] == '>' {
				if temp != key {
					fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
					err = errors.New(out.String())
					return
				}
				maps[key] = value
				if e := setProp(key, value, info); e != nil {
					err = e
					return
				}
				key = ""
				value = ""
				temp = ""
				state = BLOG
				break
			}
			if buffer[0] > 'z' || buffer[0] < 'A' {
				fmt.Fprintf(out, tmp, string(buffer[0]), col, line)
				err = errors.New(out.String())
				return
			}
			temp = temp + string(buffer[0])
		}
		if buffer[0] == '\n' {
			line = line + 1
			col = 0
		} else {
			col = col + 1
		}
	}
}

func is(k string, p string) bool {
	return strings.ToLower(k) == p;
}

func setProp(k string, v string, blog * Blog) error {
	if is(k, TITLE) {
		blog.Title = v
	} else if is(k, URLID) {
		return handleUrlid(v, blog)
	} else if is(k, TAGS) {
		return handleTags(v, blog)
	} else if is(k, CATE) || is(k, CATEGORY) {
		blog.Category = v
	} else if is(k, OVERVIEW) {
		blog.Overview = v
	}else if is(k, LANG) {
		blog.Lang = v
	} else {
		return errors.New("undefined tag " + k)
	}

	return nil
}

func handleTags(v string, blog * Blog) error {
	r := strings.NewReader(v)
	buf := make([]byte, 1)
	var s_space, s_tag int = 0, 1
	var temp string
	state := s_space
	for {
		if l, e := r.Read(buf); e != nil {
			if e == io.EOF {
				return nil
			}
			return e
		} else if l == 0 {
			continue
		}
		switch state {
		case s_space:
			if buf[0] == '[' {
				state = s_tag
			}
		case s_tag:
			if buf[0] == ']' {
				blog.Tags = append(blog.Tags, temp)
				temp = ""
				state = s_space
				break
			}
			if buf[0] == '[' {
				return errors.New("an error in <tags></tags>")
			}
			temp = temp + string(buf[0])
		}
	}
}

func handleUrlid(v string, blog * Blog) error {
	r := strings.NewReader(v)
	buf := make([]byte, 1)
	temp := ""
	for {
		if l, e := r.Read(buf); e != nil {
			if e == io.EOF {
				blog.UrlID = temp
				return nil
			}
			return e
		} else if l == 0 {
			continue
		}
		if buf[0] >= 'A' && buf[0] <= 'z' || buf[0] == '-' {
			temp = temp + string(buf[0])
		} else {
			return errors.New("urlid only support A-z_, found '" + string(buf[0]) + "'")
		}
	}
}