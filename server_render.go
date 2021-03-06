package zzblog

import (
	"crypto/md5"
	"encoding/hex"
	hr "github.com/yang-zzhong/go-httprouter"
	"github.com/yang-zzhong/logf"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

type ServerRenderer struct {
	bots     []string
	hasPage  []string
	server   string
	cacheDir string
	zzblog   Zzblog
}

func NewServerRenderer(bots []string, server string, cachedir string, zzblog Zzblog) *ServerRenderer {
	_, err := os.Stat(cachedir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil
		}
		if err = os.MkdirAll(cachedir, 0755); err != nil {
			return nil
		}
	}
	return &ServerRenderer{bots, []string{}, server, cachedir, zzblog}
}

func (sr *ServerRenderer) Before(w *hr.Response, req *hr.Request) bool {
	for _, bot := range sr.bots {
		reg := regexp.MustCompile("(?i)" + bot)
		if reg.MatchString(req.Header.Get("User-Agent")) {
			logf.Printf("spider: %s %s\n", req.Header.Get("user-Agent"), req.URL.Path)
			return !sr.Render(w, req.Request)
		}
	}
	logf.Printf("user: %s %s %s\n", req.Header.Get("User-Agent"), req.URL.Path, req.RemoteAddr)
	return true
}

func (sr *ServerRenderer) After(w *hr.Response, _ *hr.Request) bool {
	return true
}

func (sr *ServerRenderer) is404(path string) bool {
	if path == "/" || path == "" {
		return false
	}
	// 是否是已存在的文章
	has := false
	sr.zzblog.Filter(func(l *LangGroup) *Blog {
		l.Each(func(b *Blog) bool {
			if "/"+b.URLID == path {
				has = true
			}
			return has
		})
		return nil
	})
	if has {
		return false
	}
	// 是否是请求标签
	for _, tag := range sr.zzblog.Tags("") {
		if path == "/tags/"+tag {
			return false
		}
	}
	// 是否是请求分类
	for _, cate := range sr.zzblog.Cates("") {
		if path == "/cates/"+cate {
			return false
		}
	}
	return true
}

func (sr *ServerRenderer) Render(w *hr.Response, req *http.Request) bool {
	if sr.is404(req.URL.Path) {
		w.WithStatus(404)
		w.WithString("Page Not Found")
		return true
	}
	target := sr.url(req)
	lang := req.FormValue("lang")
	if lang == "" {
		lang = "en"
	}
	write := func(r io.Reader, w *hr.Response) {
		if data, err := ioutil.ReadAll(r); err != nil {
			w.WithStatus(500)
			w.WithString("read body of server render error")
		} else {
			w.WithString(string(data))
		}
	}
	u := sr.server + "/render/" + url.QueryEscape(target)
	res, err := http.Get(u)
	if err != nil {
		logf.Printf("spider: server render error: %v\n", err)
		return false
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logf.Printf("spider: render status code %d\n", res.StatusCode)
		return false
	}
	for key, val := range res.Header {
		if len(val) > 0 {
			w.WithHeader(key, val[0])
		}
	}
	w.WithHeader("Content-Type", "text/html")
	write(res.Body, w)

	return true
}

func (sr *ServerRenderer) RenderWithCache(w *hr.Response, req *http.Request) bool {
	target := sr.url(req)
	lang := req.FormValue("lang")
	if lang == "" {
		lang = "en"
	}
	sum := md5.Sum([]byte(target))
	h := hex.EncodeToString(sum[:])
	cachefile := path.Join(sr.cacheDir, h+lang+".html")
	file, err := os.Open(cachefile)
	defer file.Close()
	write := func(r io.Reader, w *hr.Response) {
		if data, err := ioutil.ReadAll(r); err != nil {
			w.WithStatus(500)
			w.WithString("read body of server render error")
		} else {
			w.WithString(string(data))
		}
	}
	if err == nil {
		write(file, w)
		return true
	}
	if !os.IsNotExist(err) {
		return false
	}
	file, err = os.OpenFile(cachefile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return false
	}
	u := sr.server + "/render/" + url.QueryEscape(target)
	res, err := http.Get(u)
	if err != nil {
		logf.Printf("spider: render error: %v\n", err)
		return false
	}
	defer res.Body.Close()
	io.Copy(file, res.Body)
	file.Seek(0, os.SEEK_SET)
	write(file, w)
	return true
}

func (sr *ServerRenderer) url(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
}
