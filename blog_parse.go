package zzblog

import (
	"io"
	"fmt"
	"strings"
	"errors"
	"bytes"
)

// ParsedBlog a parsed result
type ParsedBlog struct {
	URLID 			string
	Title 			string
	Tags 			[]string
	Category 		string
	Overview 		string
	Content 		[]byte
	Images 			[]string
	Lang 			string
	Langs		   map[string]string
}

const (
	sBlog = iota
	sTag
	sTagContent
	sTagBegin
	sTagEnd
	sTagBeginMaybeEnd
	sTagAttrValueContent
	sTagAttrValueContentEnd
	sTagAttrKey
	sTagAttrValue

	tTitle = "title"
	tUrlid = "urlid"
	tTags = "tags"
	tCate = "cate"
	tCategory = "category"
	tOverview = "overview"
	tLang = "lang"
)

func unexcepted(b byte, col, line int) error {
	errTemplate := "unexpected '%s' at %d line %d"
	outBuf := make([]byte, len(errTemplate) + 16)
	out := bytes.NewBuffer(outBuf)
	fmt.Fprintf(out, errTemplate, b, col, line)
	return errors.New(out.String())
}

func tagNotMatch(a, b string) error {
	errTemplate := "tag <%s>, </%s> not match" 
	outBuf := make([]byte, len(errTemplate) + 32)
	out := bytes.NewBuffer(outBuf)
	fmt.Fprintf(out, errTemplate, a, b)
	return errors.New(out.String())
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

func ParseBlog(r io.Reader) (info *ParsedBlog, err error) {
	info = new(ParsedBlog)
	buf := make([]byte, 1)
	var tagName, tagContent string
	var col, line int = 0, 1
	cache := []byte{}
	temp := []byte{}
	state := sBlog
	attrkey := ""
	attr := make(map[string]string)
	for {
		if _, e := r.Read(buf); e != nil {
			if e == io.EOF {
				return
			}
			err = e
			return
		}
		char := buf[0]
		switch state {
		case sBlog:
			if char == '<' {
				temp = append(temp, char)
				cache = []byte{}
				state  = sTag
				break
			}
			if isWhiteSpace(char) && len(info.Content) == 0 {
				break
			}
			info.Content = append(info.Content, char)
		case sTag:
			temp = append(temp, char)
			if isAlpha(char) {
				cache = append(cache, char)
				state = sTagBegin
				break
			}
			state = sBlog
			cache = []byte{}
			info.Content = append(info.Content, temp...)
			temp = []byte{}
		case sTagBegin:
			temp = append(temp, char)
			if isW(char) {
				cache = append(cache, char)
				break
			}
			if isWhiteSpace(char) {
				tagName = string(cache)
				cache = []byte{}
				state = sTagAttrKey
				attrkey = ""
				break
			}
			if char == '/' {
				state = sTagBeginMaybeEnd
				break
			} 
			if char == '>' {
				tagName = string(cache)
				cache = []byte{}
				state = sTagContent
				break
			}
			err = unexcepted(char, col, line)
			return
		case sTagBeginMaybeEnd:
			temp = append(temp, char)
			if char != '>' {
				err = unexcepted(char, col, line)
				break
			}
			if ok, e := tagEnded(tagName, tagContent, attr, info); e != nil {
				err = e
				return
			} else if !ok {
				info.Content = append(info.Content, temp...)
			}
			temp = []byte{}
			tagName = ""
			tagContent = ""
			attr = make(map[string]string)
			cache = []byte{}
			state = sBlog
		case sTagAttrKey:
			temp = append(temp, char)
			if isWhiteSpace(char) {
				if len(cache) == 0 {
					break
				}
				attr[string(cache)] = "true"
				cache = []byte{}
				break
			}
			if isAlpha(char) || isNumber(char) && len(cache) > 0 {
				cache = append(cache, char)
				break
			}
			if char == '=' {
				attrkey = string(cache)
				cache = []byte{}
				state = sTagAttrValue
				break
			}
			if char == '/' {
				if len(attrkey) == 0 && len(cache) == 0 {
					err = unexcepted(char, col, line)
					return
				} else if len(cache) > 0 && len(attrkey) == 0 {
					attrkey = string(cache)
				}
				state = sTagBeginMaybeEnd
				cache = []byte{}
				break
			}
			err = unexcepted(char, col, line)
			return
		case sTagAttrValue:
			temp = append(temp, char)
			if isWhiteSpace(char) {
				if len(cache) == 0 {
					break
				}
				attr[attrkey] = string(cache)
				attrkey = ""
				state = sTagBegin
				cache = []byte{}
				break
			}
			if (char == '"' || char == '\'') && len(cache) == 0 {
				cache = []byte{}
				state = sTagAttrValueContent
				break
			}
			cache = append(cache, char)
		case sTagAttrValueContent:
			temp = append(temp, char)
			if char == '"' || char == '\'' {
				attr[attrkey] = string(cache)
				state = sTagAttrValueContentEnd
				break
			}
			cache = append(cache, char)
		case sTagAttrValueContentEnd:
			temp = append(temp, char)
			if isWhiteSpace(char) {
				state = sTagAttrKey
				cache = []byte{}
				break
			}
			if char == '>' {
				state = sTagContent
				break
			}
			err = unexcepted(char, col, line)
			return
		case sTagContent:
			temp = append(temp, char)
			if char == '<' {
				tagContent = string(cache)
				cache = []byte{}
				state = sTagEnd
				break
			}
			cache = append(cache, char)
		case sTagEnd:
			temp = append(temp, char)
			if char == '/' && len(cache) == 0 {
				break
			}
			if isAlpha(char) || isNumber(char) && len(cache) == 0 {
				cache = append(cache, char)
				break
			}
			if char == '>' {
				if !eq(string(cache), tagName) {
					err = tagNotMatch(tagName, string(cache))
					return
				}
				if ok, e := tagEnded(tagName, tagContent, attr, info); e != nil {
					err = e
					return
				} else if !ok {
					info.Content = append(info.Content, temp...)
				}
				temp = []byte{}
				cache = []byte{}
				state = sBlog
				tagName = ""
				tagContent = ""
				attr = make(map[string]string)
				break
			}
			err = unexcepted(char, col, line)
			return
		}
		if char == '\n' {
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

func setProp(k string, v string, blog * ParsedBlog) (ok bool, err error) {
	if is(k, tTitle) {
		ok = true
		err = nil
		blog.Title = v
	} else if is(k, tUrlid) {
		ok = true
		err = handleUrlid(v, blog)
	} else if is(k, tTags) {
		ok = true
		err = handleTags(v, blog)
	} else if is(k, tCate) || is(k, tCategory) {
		ok = true
		err = nil
		blog.Category = v
	} else if is(k, tOverview) {
		ok = true
		err = nil
		blog.Overview = v
	}else if is(k, tLang) {
		ok = true
		err = nil
		blog.Lang = v
	} else {
		ok = false
		err = nil
	}
	return
}

func handleTags(v string, blog * ParsedBlog) error {
	r := strings.NewReader(v)
	buf := make([]byte, 1)
	var sSpace, sTag int = 0, 1
	temp := []byte{}
	state := sSpace
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
		case sSpace:
			if buf[0] == '[' {
				state = sTag
			}
		case sTag:
			if buf[0] == ']' {
				blog.Tags = append(blog.Tags, string(temp))
				temp = []byte{}
				state = sSpace
				break
			}
			if buf[0] == '[' {
				return errors.New("an error in <tags></tags>")
			}
			temp = append(temp, buf[0])
		}
	}
}

func handleUrlid(v string, blog * ParsedBlog) error {
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
		if buf[0] >= 'A' && buf[0] <= 'z' || buf[0] == '-' {
			temp = append(temp, buf[0])
		} else {
			return errors.New("urlid only support A-z_, found '" + string(buf[0]) + "'")
		}
	}
}

func tagEnded(name, content string, attr map[string]string, info *ParsedBlog) (ok bool, err error) {
	k := name
	v := content
	if val, ok := attr["for"]; ok {
		k = val
	}
	if val, ok := attr["value"]; ok {
		v = val
	}
	return setProp(k, v, info)
}